package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const defaultManifest = "corpus/corpus.json"

type manifest struct {
	Version      int          `json:"version"`
	Workspace    string       `json:"workspace"`
	Evidence     string       `json:"evidence"`
	Repositories []repository `json:"repositories"`
}

type repository struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Ref         string `json:"ref"`
	Commit      string `json:"commit"`
	Description string `json:"description"`
}

type sourceEvidence struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Ref         string `json:"ref"`
	Commit      string `json:"commit"`
	Description string `json:"description"`
	SourcePath  string `json:"sourcePath"`
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout io.Writer, stderr io.Writer) int {
	fs := flag.NewFlagSet("corpus", flag.ContinueOnError)
	fs.SetOutput(stderr)
	manifestPath := fs.String("manifest", defaultManifest, "corpus manifest path")
	repoFilter := fs.String("repo", "", "run only one repository id")
	skipClone := fs.Bool("skip-clone", false, "reuse existing .corpus clones without fetch/checkout")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	root, err := repoRoot()
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	cfg, err := loadManifest(filepath.Join(root, *manifestPath))
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	if err := validateManifest(cfg); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}

	selected := cfg.Repositories
	if *repoFilter != "" {
		selected = selected[:0]
		for _, repo := range cfg.Repositories {
			if repo.ID == *repoFilter {
				selected = append(selected, repo)
			}
		}
		if len(selected) == 0 {
			fmt.Fprintf(stderr, "unknown repository id %q\n", *repoFilter)
			return 2
		}
	}

	for _, repo := range selected {
		fmt.Fprintf(stdout, "==> %s\n", repo.ID)
		if err := runRepository(root, cfg, repo, *skipClone); err != nil {
			fmt.Fprintf(stderr, "%s: %v\n", repo.ID, err)
			return 1
		}
	}
	fmt.Fprintf(stdout, "Corpus evidence written to %s\n", filepath.ToSlash(cfg.Evidence))
	return 0
}

func repoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = io.Discard
	if err := cmd.Run(); err != nil {
		return os.Getwd()
	}
	return filepath.Clean(strings.TrimSpace(out.String())), nil
}

func loadManifest(path string) (manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return manifest{}, err
	}
	var cfg manifest
	if err := json.Unmarshal(data, &cfg); err != nil {
		return manifest{}, fmt.Errorf("%s: %w", path, err)
	}
	return cfg, nil
}

func validateManifest(cfg manifest) error {
	if cfg.Version != 1 {
		return fmt.Errorf("unsupported corpus manifest version %d", cfg.Version)
	}
	if cfg.Workspace == "" {
		return errors.New("manifest workspace is required")
	}
	if cfg.Evidence == "" {
		return errors.New("manifest evidence is required")
	}
	seen := map[string]bool{}
	for _, repo := range cfg.Repositories {
		if repo.ID == "" || repo.URL == "" || repo.Commit == "" {
			return fmt.Errorf("repository entries require id, url, and commit: %#v", repo)
		}
		if seen[repo.ID] {
			return fmt.Errorf("duplicate repository id %q", repo.ID)
		}
		seen[repo.ID] = true
	}
	if len(cfg.Repositories) == 0 {
		return errors.New("manifest must include at least one repository")
	}
	return nil
}

func runRepository(root string, cfg manifest, repo repository, skipClone bool) error {
	sourceDir := filepath.Join(root, cfg.Workspace, repo.ID)
	evidenceDir := filepath.Join(root, cfg.Evidence, repo.ID)
	sourceArg := filepath.ToSlash(filepath.Join(cfg.Workspace, repo.ID))
	evidenceArg := filepath.ToSlash(filepath.Join(cfg.Evidence, repo.ID))
	goArg := filepath.ToSlash(filepath.Join(cfg.Evidence, repo.ID, "go"))

	if !skipClone {
		if err := checkoutRepository(sourceDir, repo); err != nil {
			return err
		}
	} else if _, err := os.Stat(sourceDir); err != nil {
		return fmt.Errorf("skip-clone requested but source directory is unavailable: %w", err)
	}

	if err := os.RemoveAll(evidenceDir); err != nil {
		return err
	}
	if err := os.MkdirAll(evidenceDir, 0755); err != nil {
		return err
	}
	if err := writeSourceEvidence(evidenceDir, repo, sourceArg); err != nil {
		return err
	}

	if err := runJ2GO(root, filepath.Join(evidenceDir, "scan.txt"), "scan", sourceArg); err != nil {
		return err
	}
	if err := runJ2GO(root, filepath.Join(evidenceDir, "report-json.txt"), "report", sourceArg, "--format", "json", "--out", evidenceArg+"/report.json"); err != nil {
		return err
	}
	if err := runJ2GO(root, filepath.Join(evidenceDir, "report-md.txt"), "report", sourceArg, "--format", "markdown", "--out", evidenceArg+"/report.md"); err != nil {
		return err
	}
	if err := runJ2GO(root, filepath.Join(evidenceDir, "migrate.txt"), "migrate", sourceArg, "--out", goArg, "--report", evidenceArg+"/migration.md", "--log-file", evidenceArg+"/migration.log", "--module", "github.com/cervantesh/cervo-transpiler-java-to-go/evidence/corpus/"+repo.ID+"/go"); err != nil {
		return err
	}
	if err := runTranscript(filepath.Join(root, goArg), filepath.Join(evidenceDir, "go-test.txt"), "go", "test", "./..."); err != nil {
		return err
	}
	return nil
}

func checkoutRepository(sourceDir string, repo repository) error {
	if _, err := os.Stat(sourceDir); errors.Is(err, os.ErrNotExist) {
		if err := runShell("git", "clone", repo.URL, sourceDir); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		if err := runIn(sourceDir, "git", "fetch", "--all", "--tags", "--prune"); err != nil {
			return err
		}
	}
	if err := runIn(sourceDir, "git", "checkout", "--detach", repo.Commit); err != nil {
		return err
	}
	if err := runIn(sourceDir, "git", "clean", "-fdx"); err != nil {
		return err
	}
	return nil
}

func writeSourceEvidence(evidenceDir string, repo repository, sourceDir string) error {
	payload := sourceEvidence{
		ID:          repo.ID,
		Name:        repo.Name,
		URL:         repo.URL,
		Ref:         repo.Ref,
		Commit:      repo.Commit,
		Description: repo.Description,
		SourcePath:  filepath.ToSlash(sourceDir),
	}
	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(filepath.Join(evidenceDir, "source.json"), data, 0644)
}

func runJ2GO(root string, transcript string, args ...string) error {
	fullArgs := append([]string{"run", "./cmd/j2go"}, args...)
	return runTranscript(root, transcript, "go", fullArgs...)
}

func runTranscript(dir string, transcript string, name string, args ...string) error {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	var combined bytes.Buffer
	combined.WriteString("$ " + commandLine(name, args) + "\n")
	if stdout.Len() > 0 {
		combined.WriteString("\n[stdout]\n")
		combined.Write(stdout.Bytes())
	}
	if stderr.Len() > 0 {
		combined.WriteString("\n[stderr]\n")
		combined.Write(stderr.Bytes())
	}
	if err != nil {
		combined.WriteString("\n[exit]\n")
		combined.WriteString(err.Error() + "\n")
	}
	if writeErr := os.WriteFile(transcript, combined.Bytes(), 0644); writeErr != nil {
		return writeErr
	}
	return err
}

func runShell(name string, args ...string) error {
	return runIn("", name, args...)
}

func runIn(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func commandLine(name string, args []string) string {
	quoted := []string{name}
	for _, arg := range args {
		if strings.ContainsAny(arg, " \t\n\"") {
			quoted = append(quoted, fmt.Sprintf("%q", arg))
			continue
		}
		quoted = append(quoted, arg)
	}
	return strings.Join(quoted, " ")
}
