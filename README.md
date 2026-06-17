# cervo-transpiler-java-to-go

Modern Java-to-Go migration toolkit for the Cervo migration effort.

The project uses a deterministic compiler-style pipeline:

```text
Java source -> ANTLR4 parse tree -> semantic analysis -> Cervo IR -> Go generator -> gofmt/tests
```

The current tool is a controlled-subset transpiler plus a project migration/reporting framework. AI support is advisory only and does not change deterministic generated code by default.

## Quick Start

From the repo root:

```bash
go test ./...
go run ./tools/antlrgen -jar tools/antlr/antlr-4.13.1-complete.jar -grammar grammar/JavaSubset.g4 -out internal/parser/gen
go build -o build/j2go ./cmd/j2go
```

Run the CLI through Go when you want a platform-neutral command:

```bash
go run ./cmd/j2go --version
go run ./cmd/j2go scan ./internal/javaproject/testdata/pure-java-lib
go run ./cmd/j2go report ./internal/javaproject/testdata/pure-java-lib --format markdown --out build/migration-report.md
go run ./cmd/j2go migrate ./internal/javaproject/testdata/pure-java-lib --out build/go-lib --report build/migration-report.md --dry-run
```

Run the evidence suites:

```bash
go run ./tools/legacytest
go run ./tools/corpus
```

## Main Commands

- `j2go scan <path>`: inspect a Java project.
- `j2go report <path> --format json|markdown --out <file>`: create deterministic migration reports.
- `j2go transpile <file-or-dir> -out <dir>`: transpile supported Java files.
- `j2go migrate <project> --out <dir> --report <file> [--dry-run]`: generate a partial Go migration.
- `j2go ai explain --report <file> --out <file>`: optional advisory AI review.

## Current Scope

The transpiler supports a focused Java subset: simple classes, static methods, primitive types, variables, assignments, expressions, `if`, `while`, `System.out.println`, basic fields, constructors, instance methods, and simple interfaces.

Unsupported Java features should produce structured `JTG` diagnostics instead of misleading Go output. Broad Java framework migration, including Spring, is intentionally outside the current deterministic path.

The legacy Flex/Bison implementation under `src/` remains as reference material while the modern Go/ANTLR pipeline continues to mature.

## Documentation

Start with the GitHub Wiki:

- [Wiki Home](https://github.com/cervantesh/cervo-transpiler-java-to-go/wiki)
- [Project Guide](https://github.com/cervantesh/cervo-transpiler-java-to-go/wiki/Project-Guide)

Core docs:

- [Operations](docs/operations.md)
- [Modern Architecture](docs/architecture-modern.md)
- [Diagnostics](docs/diagnostics.md)
- [Compatibility](docs/compatibility.md)
- [Maturity Analysis](docs/maturity-analysis-2026-06-16.md)
- [Productive Migration Roadmap](docs/productive-migration-roadmap.md)
- [AI-Assisted Migration](docs/ai-assisted-migration.md)
- [Spring Adapter Decision](docs/spring-adapter-decision.md)
