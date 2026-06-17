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
	"strconv"
	"strings"
)

const defaultManifest = "corpus/corpus.json"

type manifest struct {
	Version       int            `json:"version"`
	Workspace     string         `json:"workspace"`
	Evidence      string         `json:"evidence"`
	LocalProjects []localProject `json:"localProjects"`
	Repositories  []repository   `json:"repositories"`
}

type localProject struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Category    string      `json:"category"`
	Description string      `json:"description"`
	Expectation expectation `json:"expectation"`
}

type repository struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	URL         string      `json:"url"`
	Ref         string      `json:"ref"`
	Commit      string      `json:"commit"`
	Category    string      `json:"category"`
	Description string      `json:"description"`
	Expectation expectation `json:"expectation"`
}

type expectation struct {
	MinGenerated   int  `json:"minGenerated,omitempty"`
	MaxSkipped     *int `json:"maxSkipped,omitempty"`
	MaxDiagnostics *int `json:"maxDiagnostics,omitempty"`
}

type corpusProject struct {
	ID          string
	Name        string
	Kind        string
	Category    string
	Description string
	Path        string
	URL         string
	Ref         string
	Commit      string
	Expectation expectation
}

type sourceEvidence struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Kind        string      `json:"kind"`
	Category    string      `json:"category"`
	Path        string      `json:"path,omitempty"`
	URL         string      `json:"url"`
	Ref         string      `json:"ref"`
	Commit      string      `json:"commit"`
	Description string      `json:"description"`
	SourcePath  string      `json:"sourcePath"`
	Expectation expectation `json:"expectation,omitempty"`
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout io.Writer, stderr io.Writer) int {
	fs := flag.NewFlagSet("corpus", flag.ContinueOnError)
	fs.SetOutput(stderr)
	manifestPath := fs.String("manifest", defaultManifest, "corpus manifest path")
	repoFilter := fs.String("repo", "", "run only one corpus project id")
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

	selected := allProjects(cfg)
	if *repoFilter != "" {
		selected = selected[:0]
		for _, project := range allProjects(cfg) {
			if project.ID == *repoFilter {
				selected = append(selected, project)
			}
		}
		if len(selected) == 0 {
			fmt.Fprintf(stderr, "unknown corpus project id %q\n", *repoFilter)
			return 2
		}
	}

	for _, project := range selected {
		fmt.Fprintf(stdout, "==> %s/%s\n", project.Category, project.ID)
		if err := runProject(root, cfg, project, *skipClone); err != nil {
			fmt.Fprintf(stderr, "%s: %v\n", project.ID, err)
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
	for _, project := range cfg.LocalProjects {
		if project.ID == "" || project.Path == "" {
			return fmt.Errorf("local project entries require id and path: %#v", project)
		}
		if seen[project.ID] {
			return fmt.Errorf("duplicate corpus project id %q", project.ID)
		}
		seen[project.ID] = true
	}
	for _, repo := range cfg.Repositories {
		if repo.ID == "" || repo.URL == "" || repo.Commit == "" {
			return fmt.Errorf("repository entries require id, url, and commit: %#v", repo)
		}
		if seen[repo.ID] {
			return fmt.Errorf("duplicate repository id %q", repo.ID)
		}
		seen[repo.ID] = true
	}
	if len(seen) == 0 {
		return errors.New("manifest must include at least one corpus project")
	}
	return nil
}

func allProjects(cfg manifest) []corpusProject {
	projects := make([]corpusProject, 0, len(cfg.LocalProjects)+len(cfg.Repositories))
	for _, local := range cfg.LocalProjects {
		category := local.Category
		if category == "" {
			category = "curated"
		}
		projects = append(projects, corpusProject{
			ID:          local.ID,
			Name:        local.Name,
			Kind:        "local",
			Category:    category,
			Description: local.Description,
			Path:        local.Path,
			Expectation: local.Expectation,
		})
	}
	for _, repo := range cfg.Repositories {
		category := repo.Category
		if category == "" {
			category = "external"
		}
		projects = append(projects, corpusProject{
			ID:          repo.ID,
			Name:        repo.Name,
			Kind:        "git",
			Category:    category,
			Description: repo.Description,
			URL:         repo.URL,
			Ref:         repo.Ref,
			Commit:      repo.Commit,
			Expectation: repo.Expectation,
		})
	}
	return projects
}

func runProject(root string, cfg manifest, project corpusProject, skipClone bool) error {
	sourceArg := project.Path
	sourceDir := filepath.Join(root, project.Path)
	if project.Kind == "git" {
		sourceArg = filepath.ToSlash(filepath.Join(cfg.Workspace, project.ID))
		sourceDir = filepath.Join(root, cfg.Workspace, project.ID)
	}
	evidenceRel := filepath.Join(cfg.Evidence, project.Category, project.ID)
	evidenceDir := filepath.Join(root, evidenceRel)
	evidenceArg := filepath.ToSlash(evidenceRel)
	goArg := filepath.ToSlash(filepath.Join(evidenceRel, "go"))

	if project.Kind == "git" && !skipClone {
		if err := checkoutRepository(sourceDir, project); err != nil {
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
	if err := writeSourceEvidence(evidenceDir, project, sourceArg); err != nil {
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
	if err := runJ2GO(root, filepath.Join(evidenceDir, "migrate.txt"), "migrate", sourceArg, "--out", goArg, "--report", evidenceArg+"/migration.md", "--log-file", evidenceArg+"/migration.log", "--module", "github.com/cervantesh/cervo-transpiler-java-to-go/"+evidenceArg+"/go"); err != nil {
		return err
	}
	if err := checkExpectations(filepath.Join(evidenceDir, "migration.log"), project); err != nil {
		return err
	}
	if err := runTranscript(filepath.Join(root, goArg), filepath.Join(evidenceDir, "go-test.txt"), "go", "test", "./..."); err != nil {
		return err
	}
	if err := scrubEvidencePaths(evidenceDir, root); err != nil {
		return err
	}
	return nil
}

func checkoutRepository(sourceDir string, repo corpusProject) error {
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

func writeSourceEvidence(evidenceDir string, repo corpusProject, sourceDir string) error {
	payload := sourceEvidence{
		ID:          repo.ID,
		Name:        repo.Name,
		Kind:        repo.Kind,
		Category:    repo.Category,
		Path:        repo.Path,
		URL:         repo.URL,
		Ref:         repo.Ref,
		Commit:      repo.Commit,
		Description: repo.Description,
		SourcePath:  filepath.ToSlash(sourceDir),
		Expectation: repo.Expectation,
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

type migrationStats struct {
	JavaFiles   int
	Generated   int
	Skipped     int
	Diagnostics int
}

func checkExpectations(logPath string, project corpusProject) error {
	if project.Expectation.MinGenerated == 0 && project.Expectation.MaxSkipped == nil && project.Expectation.MaxDiagnostics == nil {
		return nil
	}
	stats, err := readMigrationStats(logPath)
	if err != nil {
		return err
	}
	if stats.Generated < project.Expectation.MinGenerated {
		return fmt.Errorf("%s generated %d files, expected at least %d", project.ID, stats.Generated, project.Expectation.MinGenerated)
	}
	if project.Expectation.MaxSkipped != nil && stats.Skipped > *project.Expectation.MaxSkipped {
		return fmt.Errorf("%s skipped %d files, expected at most %d", project.ID, stats.Skipped, *project.Expectation.MaxSkipped)
	}
	if project.Expectation.MaxDiagnostics != nil && stats.Diagnostics > *project.Expectation.MaxDiagnostics {
		return fmt.Errorf("%s produced %d diagnostics, expected at most %d", project.ID, stats.Diagnostics, *project.Expectation.MaxDiagnostics)
	}
	return nil
}

func readMigrationStats(logPath string) (migrationStats, error) {
	data, err := os.ReadFile(logPath)
	if err != nil {
		return migrationStats{}, err
	}
	stats := migrationStats{}
	for _, line := range strings.Split(string(data), "\n") {
		key, value, ok := strings.Cut(strings.TrimSpace(line), "=")
		if !ok {
			continue
		}
		parsed, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		switch key {
		case "javaFiles":
			stats.JavaFiles = parsed
		case "generated":
			stats.Generated = parsed
		case "skipped":
			stats.Skipped = parsed
		case "diagnostics":
			stats.Diagnostics = parsed
		}
	}
	return stats, nil
}

func scrubEvidencePaths(evidenceDir string, root string) error {
	root = filepath.Clean(root)
	replacements := []string{
		filepath.ToSlash(root) + "/",
		filepath.ToSlash(root),
		root + string(os.PathSeparator),
		root,
	}
	return filepath.WalkDir(evidenceDir, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || filepath.Ext(path) == ".go" || filepath.Base(path) == "go.mod" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		text := string(data)
		for _, replacement := range replacements {
			text = strings.ReplaceAll(text, replacement, "")
		}
		return os.WriteFile(path, []byte(text), 0644)
	})
}
