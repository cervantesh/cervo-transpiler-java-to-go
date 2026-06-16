package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/pipeline"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/version"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout io.Writer, stderr io.Writer) int {
	fs := flag.NewFlagSet("j2go", flag.ContinueOnError)
	fs.SetOutput(stderr)
	outDir := fs.String("out", "", "directory for generated Go files")
	showVersion := fs.Bool("version", false, "print version")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *showVersion {
		fmt.Fprintf(stdout, "j2go %s\n", version.Version)
		return 0
	}
	if *outDir == "" {
		fmt.Fprintln(stderr, "-out is required")
		return 2
	}
	inputs := fs.Args()
	if len(inputs) == 0 {
		fmt.Fprintln(stderr, "at least one Java file or directory is required")
		return 2
	}
	if err := os.MkdirAll(*outDir, 0755); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	for _, input := range inputs {
		if err := transpilePath(input, *outDir); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
	}
	return 0
}

func transpilePath(input string, outDir string) error {
	info, err := os.Stat(input)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return filepath.WalkDir(input, func(path string, entry os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() || !strings.HasSuffix(path, ".java") {
				return nil
			}
			return transpileFile(path, outDir)
		})
	}
	return transpileFile(input, outDir)
}

func transpileFile(input string, outDir string) error {
	source, err := os.ReadFile(input)
	if err != nil {
		return err
	}
	result := pipeline.TranspileSource(input, string(source))
	if len(result.Diagnostics) > 0 {
		return result.Diagnostics[0]
	}
	base := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input)) + ".go"
	return os.WriteFile(filepath.Join(outDir, base), result.Code, 0644)
}
