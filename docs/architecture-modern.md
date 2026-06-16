# Modern Architecture

The target architecture for `cervo-transpiler-java-to-go` is:

```text
Java source -> ANTLR4 parse tree -> semantic analysis -> Cervo IR -> Go generator -> gofmt/golden tests
```

The current Flex/Bison/C++ implementation is a legacy reference. It proves a small Java-to-Go translation surface, but the Cervo implementation should move to a modern parser, an explicit intermediate representation, and a Go backend that formats output through Go tooling.

## Technology Choices

| Layer | Choice | Notes |
| --- | --- | --- |
| Parser | ANTLR4 | Modern parser generator with visitors/listeners and a clear grammar boundary. |
| Java grammar | Reduced ANTLR Java subset first | Keeps the MVP testable while preserving a path to `antlr/grammars-v4`. |
| Editor tooling | Tree-sitter | Useful for fast scans and editor integrations, not the semantic source of truth. |
| Java AST alternatives | JavaParser or Spoon | Good candidates once real Java projects require richer semantics. |
| Refactoring framework | OpenRewrite | Useful for pre-normalizing enterprise Java before transpilation. |
| IR | `internal/ir` | Stable boundary between Java-specific lowering and Go-specific emission. |
| Go output | `go/ast` + `go/format` | Ensures generated Go is syntactically valid and consistently formatted. |

## Rule

The parser must not emit Go directly. All code generation must go through Cervo IR and the Go emitter.

## Implementation Plan

The detailed task plan is:

[docs/superpowers/plans/2026-06-16-modern-java-to-go-transpiler.md](superpowers/plans/2026-06-16-modern-java-to-go-transpiler.md)
