package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAllProjectsSeparatesCuratedAndExternalDefaults(t *testing.T) {
	cfg := manifest{
		LocalProjects: []localProject{{
			ID:   "curated-lib",
			Path: "corpus/projects/curated-lib",
		}},
		Repositories: []repository{{
			ID:     "external-lib",
			URL:    "https://example.com/lib.git",
			Commit: "abc123",
		}},
	}

	projects := allProjects(cfg)
	if len(projects) != 2 {
		t.Fatalf("expected two projects, got %#v", projects)
	}
	if projects[0].Kind != "local" || projects[0].Category != "curated" {
		t.Fatalf("expected local curated project, got %#v", projects[0])
	}
	if projects[1].Kind != "git" || projects[1].Category != "external" {
		t.Fatalf("expected git external project, got %#v", projects[1])
	}
}

func TestCheckExpectationsRejectsCuratedRegression(t *testing.T) {
	maxSkipped := 0
	maxDiagnostics := 0
	project := corpusProject{
		ID: "curated-lib",
		Expectation: expectation{
			MinGenerated:   5,
			MaxSkipped:     &maxSkipped,
			MaxDiagnostics: &maxDiagnostics,
		},
	}
	logPath := filepath.Join(t.TempDir(), "migration.log")
	if err := os.WriteFile(logPath, []byte("javaFiles=5\ngenerated=4\nskipped=1\ndiagnostics=0\n"), 0644); err != nil {
		t.Fatal(err)
	}

	err := checkExpectations(logPath, project)
	if err == nil || !strings.Contains(err.Error(), "expected at least 5") {
		t.Fatalf("expected min generated failure, got %v", err)
	}
}

func TestScrubEvidencePathsLeavesGeneratedGoUntouched(t *testing.T) {
	root := filepath.Join(t.TempDir(), "repo")
	evidenceDir := filepath.Join(root, "evidence")
	if err := os.MkdirAll(filepath.Join(evidenceDir, "go"), 0755); err != nil {
		t.Fatal(err)
	}
	reportPath := filepath.Join(evidenceDir, "report.md")
	goPath := filepath.Join(evidenceDir, "go", "Example.go")
	if err := os.WriteFile(reportPath, []byte(filepath.ToSlash(root)+"/corpus/project\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(goPath, []byte("package main\n"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := scrubEvidencePaths(evidenceDir, root); err != nil {
		t.Fatal(err)
	}
	report, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(report), filepath.ToSlash(root)) || !strings.Contains(string(report), "corpus/project") {
		t.Fatalf("expected scrubbed report, got %q", string(report))
	}
	generated, err := os.ReadFile(goPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(generated) != "package main\n" {
		t.Fatalf("expected generated Go untouched, got %q", string(generated))
	}
}
