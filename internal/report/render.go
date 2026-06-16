package report

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/javaproject"
)

func JSON(project javaproject.Project) ([]byte, error) {
	out, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(out, '\n'), nil
}

func Markdown(project javaproject.Project) []byte {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "# Java-To-Go Migration Report\n\n")
	fmt.Fprintf(&buffer, "Project: `%s`\n\n", project.Root)

	fmt.Fprintf(&buffer, "## Summary\n\n")
	fmt.Fprintf(&buffer, "| Metric | Count |\n")
	fmt.Fprintf(&buffer, "| --- | ---: |\n")
	fmt.Fprintf(&buffer, "| Java files | %d |\n", project.Summary.JavaFiles)
	fmt.Fprintf(&buffer, "| Packages | %d |\n", project.Summary.Packages)
	fmt.Fprintf(&buffer, "| Classes | %d |\n", project.Summary.Classes)
	fmt.Fprintf(&buffer, "| Methods | %d |\n", project.Summary.Methods)
	fmt.Fprintf(&buffer, "| Constructors | %d |\n", project.Summary.Constructors)
	fmt.Fprintf(&buffer, "| Fields | %d |\n", project.Summary.Fields)
	fmt.Fprintf(&buffer, "| Unsupported features | %d |\n", project.Summary.Unsupported)
	fmt.Fprintf(&buffer, "| Diagnostics | %d |\n", project.Summary.Diagnostics)
	fmt.Fprintf(&buffer, "| Internal imports | %d |\n\n", project.Summary.InternalImports)

	fmt.Fprintf(&buffer, "## Build Files\n\n")
	if len(project.BuildFiles) == 0 {
		fmt.Fprintf(&buffer, "No Maven or Gradle build file detected.\n\n")
	} else {
		for _, buildFile := range project.BuildFiles {
			fmt.Fprintf(&buffer, "- `%s` (%s)\n", buildFile.Path, buildFile.Kind)
		}
		fmt.Fprintf(&buffer, "\n")
	}

	fmt.Fprintf(&buffer, "## Packages\n\n")
	if len(project.Packages) == 0 {
		fmt.Fprintf(&buffer, "No package declarations detected.\n\n")
	} else {
		for _, packageName := range project.Packages {
			fmt.Fprintf(&buffer, "- `%s`\n", packageName)
		}
		fmt.Fprintf(&buffer, "\n")
	}

	fmt.Fprintf(&buffer, "## Feature Counts\n\n")
	writeCounts(&buffer, project.FeatureCounts)

	fmt.Fprintf(&buffer, "## Unsupported Feature Counts\n\n")
	writeCounts(&buffer, project.UnsupportedCounts)

	fmt.Fprintf(&buffer, "## Files\n\n")
	for _, file := range project.Files {
		fmt.Fprintf(&buffer, "### `%s`\n\n", file.RelativePath)
		fmt.Fprintf(&buffer, "- Source root: `%s`\n", file.SourceRoot)
		fmt.Fprintf(&buffer, "- Risk: `%s`\n", file.Risk)
		if file.PackageName != "" {
			fmt.Fprintf(&buffer, "- Package: `%s`\n", file.PackageName)
		}
		if len(file.Imports) > 0 {
			fmt.Fprintf(&buffer, "- Imports: ")
			for i, importDecl := range file.Imports {
				if i > 0 {
					fmt.Fprintf(&buffer, ", ")
				}
				name := importDecl.Name
				if importDecl.Wildcard {
					name += ".*"
				}
				fmt.Fprintf(&buffer, "`%s`", name)
			}
			fmt.Fprintf(&buffer, "\n")
		}
		if len(file.Classes) > 0 {
			fmt.Fprintf(&buffer, "- Classes: ")
			for i, class := range file.Classes {
				if i > 0 {
					fmt.Fprintf(&buffer, ", ")
				}
				fmt.Fprintf(&buffer, "`%s`", class.Qualified)
			}
			fmt.Fprintf(&buffer, "\n")
		}
		for _, class := range file.Classes {
			if len(class.Fields) > 0 {
				fmt.Fprintf(&buffer, "- Fields in `%s`: ", class.Qualified)
				for i, field := range class.Fields {
					if i > 0 {
						fmt.Fprintf(&buffer, ", ")
					}
					fmt.Fprintf(&buffer, "`%s %s`", field.Type, field.Name)
				}
				fmt.Fprintf(&buffer, "\n")
			}
			if len(class.Constructors) > 0 {
				fmt.Fprintf(&buffer, "- Constructors in `%s`: ", class.Qualified)
				for i, constructor := range class.Constructors {
					if i > 0 {
						fmt.Fprintf(&buffer, ", ")
					}
					fmt.Fprintf(&buffer, "`%s(%d params)`", constructor.Name, len(constructor.Params))
				}
				fmt.Fprintf(&buffer, "\n")
			}
			if len(class.Methods) > 0 {
				fmt.Fprintf(&buffer, "- Methods in `%s`: ", class.Qualified)
				for i, method := range class.Methods {
					if i > 0 {
						fmt.Fprintf(&buffer, ", ")
					}
					scope := "instance"
					if method.Static {
						scope = "static"
					}
					fmt.Fprintf(&buffer, "`%s %s %s(%d params)`", scope, method.ReturnType, method.Name, len(method.Params))
				}
				fmt.Fprintf(&buffer, "\n")
			}
		}
		if len(file.Unsupported) > 0 {
			fmt.Fprintf(&buffer, "- Unsupported:\n")
			for _, unsupported := range file.Unsupported {
				fmt.Fprintf(&buffer, "  - `%s`", unsupported.Name)
				if unsupported.Code != "" {
					fmt.Fprintf(&buffer, " (%s)", unsupported.Code)
				}
				if unsupported.Recommendation != "" {
					fmt.Fprintf(&buffer, ": %s", unsupported.Recommendation)
				}
				fmt.Fprintf(&buffer, "\n")
			}
		}
		if len(file.Diagnostics) > 0 {
			fmt.Fprintf(&buffer, "- Diagnostics: ")
			for i, diagnostic := range file.Diagnostics {
				if i > 0 {
					fmt.Fprintf(&buffer, ", ")
				}
				fmt.Fprintf(&buffer, "`%s`", diagnostic.Code)
			}
			fmt.Fprintf(&buffer, "\n")
		}
		fmt.Fprintf(&buffer, "\n")
	}

	fmt.Fprintf(&buffer, "## Recommended Migration Order\n\n")
	for _, file := range migrationOrder(project.Files) {
		fmt.Fprintf(&buffer, "- `%s` (%s risk)\n", file.RelativePath, file.Risk)
	}
	return buffer.Bytes()
}

func ScanText(project javaproject.Project) []byte {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "Project: %s\n", project.Root)
	fmt.Fprintf(&buffer, "Java files: %d\n", project.Summary.JavaFiles)
	fmt.Fprintf(&buffer, "Packages: %d\n", project.Summary.Packages)
	fmt.Fprintf(&buffer, "Classes: %d\n", project.Summary.Classes)
	fmt.Fprintf(&buffer, "Unsupported features: %d\n", project.Summary.Unsupported)
	if len(project.Packages) > 0 {
		fmt.Fprintf(&buffer, "Package list:\n")
		for _, packageName := range project.Packages {
			fmt.Fprintf(&buffer, "- %s\n", packageName)
		}
	}
	if len(project.InternalDependencies) > 0 {
		fmt.Fprintf(&buffer, "Internal dependencies:\n")
		for _, dependency := range project.InternalDependencies {
			fmt.Fprintf(&buffer, "- %s imports %s\n", dependency.FromFile, dependency.Import)
		}
	}
	if len(project.Files) > 0 {
		fmt.Fprintf(&buffer, "Files:\n")
		for _, file := range project.Files {
			fmt.Fprintf(&buffer, "- %s (%s risk)\n", file.RelativePath, file.Risk)
			for _, class := range file.Classes {
				fmt.Fprintf(&buffer, "  - class %s\n", class.Qualified)
			}
		}
	}
	return buffer.Bytes()
}

func writeCounts(buffer *bytes.Buffer, counts map[string]int) {
	if len(counts) == 0 {
		fmt.Fprintf(buffer, "None.\n\n")
		return
	}
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Fprintf(buffer, "| Feature | Count |\n")
	fmt.Fprintf(buffer, "| --- | ---: |\n")
	for _, key := range keys {
		fmt.Fprintf(buffer, "| %s | %d |\n", key, counts[key])
	}
	fmt.Fprintf(buffer, "\n")
}

func migrationOrder(files []javaproject.File) []javaproject.File {
	out := append([]javaproject.File(nil), files...)
	rank := map[string]int{"low": 0, "medium": 1, "high": 2}
	sort.Slice(out, func(i, j int) bool {
		if rank[out[i].Risk] != rank[out[j].Risk] {
			return rank[out[i].Risk] < rank[out[j].Risk]
		}
		return out[i].RelativePath < out[j].RelativePath
	})
	return out
}
