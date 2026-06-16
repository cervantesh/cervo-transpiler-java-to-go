package javaproject

import (
	"path/filepath"
	"testing"
)

func TestScanPureJavaLibrary(t *testing.T) {
	project, err := Scan(filepath.Join("testdata", "pure-java-lib"))
	if err != nil {
		t.Fatalf("Scan failed: %v", err)
	}

	if project.Summary.JavaFiles != 3 {
		t.Fatalf("expected 3 Java files, got %#v", project.Summary)
	}
	if len(project.BuildFiles) != 1 || project.BuildFiles[0].Kind != "maven" {
		t.Fatalf("expected Maven build file, got %#v", project.BuildFiles)
	}
	if len(project.Packages) != 2 {
		t.Fatalf("expected two packages, got %#v", project.Packages)
	}
	if project.Summary.Classes != 3 {
		t.Fatalf("expected three classes, got %#v", project.Summary)
	}
	if project.Summary.InternalImports != 1 {
		t.Fatalf("expected one internal dependency, got %#v", project.InternalDependencies)
	}
	if project.Summary.Unsupported != 0 {
		t.Fatalf("expected object-model features to be supported, got %#v", project.Summary)
	}
	for _, feature := range []string{"class fields", "constructors", "instance methods"} {
		if project.FeatureCounts[feature] == 0 {
			t.Fatalf("expected supported feature count for %q, got %#v", feature, project.FeatureCounts)
		}
	}
}

func TestScanSingleJavaFile(t *testing.T) {
	project, err := Scan(filepath.Join("testdata", "single-file", "Hello.java"))
	if err != nil {
		t.Fatalf("Scan failed: %v", err)
	}
	if project.Summary.JavaFiles != 1 {
		t.Fatalf("expected one Java file, got %#v", project.Summary)
	}
	if project.Files[0].SourceRoot != "." {
		t.Fatalf("expected root source root, got %q", project.Files[0].SourceRoot)
	}
}
