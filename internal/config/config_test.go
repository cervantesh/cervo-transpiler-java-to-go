package config

import (
	"strings"
	"testing"
)

func TestParseCervoMigrationConfig(t *testing.T) {
	cfg, err := Parse([]byte(`version: 1
project:
  source: ./java-lib
  module: example.com/migrated
migration:
  out: ./go-lib
  report: build/migration-report.md
  dryRun: true
logs:
  file: build/j2go.log
ai:
  provider: canned
`))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if cfg.Version != "1" || cfg.Project.Source != "./java-lib" || cfg.Project.Module != "example.com/migrated" {
		t.Fatalf("unexpected project config: %#v", cfg)
	}
	if cfg.Migration.Out != "./go-lib" || cfg.Migration.Report != "build/migration-report.md" || !cfg.Migration.DryRun {
		t.Fatalf("unexpected migration config: %#v", cfg.Migration)
	}
	if cfg.Logs.File != "build/j2go.log" || cfg.AI.Provider != "canned" {
		t.Fatalf("unexpected ops config: %#v", cfg)
	}
}

func TestParseRejectsUnknownKeys(t *testing.T) {
	_, err := Parse([]byte("migration:\n  surprise: true\n"))
	if err == nil || !strings.Contains(err.Error(), "unknown migration key") {
		t.Fatalf("expected unknown key error, got %v", err)
	}
}
