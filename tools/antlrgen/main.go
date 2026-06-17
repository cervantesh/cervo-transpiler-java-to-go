package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const defaultANTLRVersion = "4.13.1"

type options struct {
	jar     string
	grammar string
	out     string
	version string
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	opts, err := parseArgs(args)
	if err != nil {
		return err
	}
	root, err := repoRoot()
	if err != nil {
		return err
	}
	jarPath := filepath.Join(root, filepath.FromSlash(opts.jar))
	if err := ensureANTLRJar(jarPath, opts.version); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(root, filepath.FromSlash(opts.out)), 0755); err != nil {
		return err
	}
	java, err := javaCommand()
	if err != nil {
		return err
	}
	cmd := exec.Command(
		java,
		"-jar", jarPath,
		"-Dlanguage=Go",
		"-visitor",
		"-package", "gen",
		"-Xexact-output-dir",
		"-o", filepath.FromSlash(opts.out),
		filepath.FromSlash(opts.grammar),
	)
	cmd.Dir = root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ANTLR parser generation failed: %w", err)
	}
	return nil
}

func parseArgs(args []string) (options, error) {
	opts := options{
		jar:     "tools/antlr/antlr-" + defaultANTLRVersion + "-complete.jar",
		grammar: "grammar/JavaSubset.g4",
		out:     "internal/parser/gen",
		version: defaultANTLRVersion,
	}
	for i := 0; i < len(args); i++ {
		if i+1 >= len(args) {
			return options{}, fmt.Errorf("%s requires a value", args[i])
		}
		value := args[i+1]
		switch args[i] {
		case "-jar":
			opts.jar = value
		case "-grammar":
			opts.grammar = value
		case "-out":
			opts.out = value
		case "-version":
			opts.version = value
		default:
			return options{}, fmt.Errorf("unknown option %s", args[i])
		}
		i++
	}
	return opts, nil
}

func repoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		wd, wdErr := os.Getwd()
		if wdErr != nil {
			return "", wdErr
		}
		return wd, nil
	}
	return filepath.Clean(stringTrimSpace(out)), nil
}

func ensureANTLRJar(path string, version string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	url := "https://www.antlr.org/download/antlr-" + version + "-complete.jar"
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download ANTLR from %s: %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("download ANTLR from %s: HTTP %s", url, resp.Status)
	}
	tmp := path + ".tmp"
	file, err := os.Create(tmp)
	if err != nil {
		return err
	}
	_, copyErr := io.Copy(file, resp.Body)
	closeErr := file.Close()
	if copyErr != nil {
		_ = os.Remove(tmp)
		return copyErr
	}
	if closeErr != nil {
		_ = os.Remove(tmp)
		return closeErr
	}
	return os.Rename(tmp, path)
}

func javaCommand() (string, error) {
	candidates := []string{}
	if value := os.Getenv("JTG_JAVA"); value != "" {
		candidates = append(candidates, value)
	}
	if value := os.Getenv("JAVA_HOME"); value != "" {
		candidates = append(candidates, filepath.Join(value, "bin", executableName("java")))
	}
	if runtime.GOOS == "windows" {
		candidates = append(candidates, `C:\dev\jdk-24.0.1\bin\java.exe`)
	}
	candidates = append(candidates, "java")
	for _, candidate := range candidates {
		if filepath.IsAbs(candidate) {
			if _, err := os.Stat(candidate); err == nil {
				return candidate, nil
			}
			continue
		}
		if path, err := exec.LookPath(candidate); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("missing java runtime; install JDK 17+ or set JTG_JAVA")
}

func executableName(name string) string {
	if runtime.GOOS == "windows" {
		return name + ".exe"
	}
	return name
}

func stringTrimSpace(data []byte) string {
	start := 0
	end := len(data)
	for start < end && (data[start] == ' ' || data[start] == '\t' || data[start] == '\n' || data[start] == '\r') {
		start++
	}
	for end > start && (data[end-1] == ' ' || data[end-1] == '\t' || data[end-1] == '\n' || data[end-1] == '\r') {
		end--
	}
	return string(data[start:end])
}
