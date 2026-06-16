package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/pipeline"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	fs := flag.NewFlagSet("j2go", flag.ContinueOnError)
	outDir := fs.String("out", "", "directory for generated Go files")
	if err := fs.Parse(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	if *outDir == "" {
		fmt.Fprintln(os.Stderr, "-out is required")
		return 2
	}
	inputs := fs.Args()
	if len(inputs) == 0 {
		fmt.Fprintln(os.Stderr, "at least one Java file or directory is required")
		return 2
	}
	if err := os.MkdirAll(*outDir, 0755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	for _, input := range inputs {
		if err := transpilePath(input, *outDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
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
