# Maturity Analysis - 2026-06-16

This document evaluates the current maturity of `cervo-transpiler-java-to-go` as a modern Java-to-Go transpiler.

## Executive Summary

Current maturity: **2.3 / 5**

Stage: **serious technical MVP, not production migration tooling yet**.

The project is no longer just an academic compiler exercise. It now has the correct modern shape:

```text
Java source -> ANTLR4 parse tree -> parser facade -> semantic layer -> Cervo IR -> Go AST emitter -> gofmt/golden tests
```

The strongest parts are the architecture boundary, deterministic tests, and Go backend using `go/ast` plus `go/format`.

The weakest parts are Java semantic depth, modern-pipeline coverage, CI/release automation, and real-project migration support.

## Maturity Scale

| Level | Meaning |
| --- | --- |
| 0 | Idea only. No executable pipeline. |
| 1 | Spike/proof of concept. One happy path works. |
| 2 | MVP. Deterministic pipeline exists, but scope is narrow. |
| 3 | Controlled internal tool. Reliable for a documented subset and CI-verified. |
| 4 | Production migration assistant. Handles real projects with reports, review workflow, and known escape hatches. |
| 5 | Enterprise-grade transpiler platform. Handles large codebases, classpaths, semantic resolution, incremental migration, and governance. |

The repo is currently at **Level 2**, with parts of Level 3 already visible.

## Dimension Scores

| Dimension | Score | Current Evidence | Main Gap |
| --- | ---: | --- | --- |
| Architecture | 3.0 / 5 | Modern layers are split across `internal/parser`, `internal/semantic`, `internal/ir`, `internal/lower`, `internal/emitgo`, and `internal/pipeline`. | IR contracts are still informal and source-location preservation is limited. |
| Parser/frontend | 2.5 / 5 | ANTLR4 grammar exists in `grammar/JavaSubset.g4`; generated parser is committed under `internal/parser/gen`. | Grammar is a reduced Java subset, not full Java; no classpath/project model. |
| Semantic analysis | 1.5 / 5 | Symbol table and semantic diagnostics exist under `internal/semantic`. | No real type checking, overload resolution, method lookup, package/import resolution, or object model. |
| IR and lowering | 2.0 / 5 | Cervo IR exists and the lowerer maps classes, static methods, vars, assignments, `if`, `while`, `for`, calls, and returns. | Lowerer still relies on `GetText()` and string-prefix checks in several places; unsupported constructs are not uniformly modeled as diagnostics. |
| Go generation | 2.8 / 5 | Emitter uses `go/ast`, `go/token`, and `go/format`; `fmt` import is conditional. | Output is always `package main`; no package mapping, import mapping, receiver generation, idiomatic error conversion, or multi-file package assembly. |
| CLI and developer workflow | 2.0 / 5 | `cmd/j2go` accepts `-out` and file/directory inputs. | CLI reports only the first diagnostic per failed file; no JSON/Markdown report mode, no dry-run, no summary, no version flag. |
| Tests and verification | 2.5 / 5 | `go test ./...` passes; legacy suite verifies 30 cases; modern golden path exists. | Modern golden coverage has only one fixture; no GitHub Actions CI; no fuzzing, corpus tests, or real Java project fixtures. |
| Diagnostics and migration reporting | 2.0 / 5 | Legacy path has structured `JTG` diagnostics and recommendations; docs list diagnostic codes. | Modern path needs the same diagnostic richness and a deterministic report mode. |
| AI-assisted migration | 1.0 / 5 | Sidecar architecture is documented in `docs/ai-assisted-migration.md`; issue #3 tracks it. | No prompt files, offline provider, report-fed explainer, or tested AI workflow yet. |
| Release/operations | 1.0 / 5 | Repo is buildable locally. | No CI, release artifacts, versioning, install script, or compatibility matrix. |

## Verification Snapshot

Commands run successfully on 2026-06-16:

```bash
go test ./...
go test ./... -cover
go run ./tools/legacytest
```

Coverage snapshot from `go test ./... -cover`:

| Package | Coverage |
| --- | ---: |
| `internal/parser` | 93.8% |
| `internal/ir` | 61.5% |
| `internal/semantic` | 53.8% |
| `internal/lower` | 47.0% |
| `internal/pipeline` | 43.8% |
| `internal/emitgo` | 43.0% |
| `cmd/j2go` | 0.0% |

Legacy evidence suite:

- 20 supported Java-to-Go translation cases.
- 1 syntax-error case with line/column expectation.
- 9 unsupported-feature diagnostic cases.
- 30 total evidence cases.

## What Is Mature Already

1. **Correct compiler architecture**

   The project has moved away from direct grammar-rule text emission. The modern implementation goes through parse tree, lowering, IR, and Go emission.

2. **Deterministic default behavior**

   The default transpilation path does not depend on AI or probabilistic generation.

3. **Explicit unsupported-scope thinking**

   Unsupported Java features are documented and, in the legacy path, tested with stable diagnostic codes and recommendations.

4. **Go backend quality foundation**

   Using `go/ast` and `go/format` is the right foundation for maintainable Go output.

5. **Legacy reference remains useful**

   The Flex/Bison/C++ reference still provides a broader evidence suite while the modern Go implementation catches up.

## What Is Still Immature

1. **Modern feature coverage is thin**

   The modern pipeline has a single golden fixture. It proves the pipeline shape, not broad translation reliability.

2. **Semantic analysis is shallow**

   Real Java migration requires type resolution, method lookup, static vs instance semantics, overload resolution, constructor modeling, imports, and package/classpath understanding.

3. **Lowering has brittle spots**

   Some lowering decisions use raw parse text, string prefixes, and sentinel calls such as `__unsupported_block` / `__assign_...`. These should become typed IR nodes or structured diagnostics.

4. **No CI gate**

   A repo at maturity Level 3 needs GitHub Actions for parser generation checks, Go tests, legacy tests, and golden verification.

5. **No migration report yet**

   Issue #2 is important. Before AI, the tool needs deterministic JSON/Markdown reporting: files scanned, features found, unsupported constructs, risk levels, and recommended migration order.

6. **AI is architectural only**

   The AI sidecar is correctly positioned, but not implemented. There is no tested prompt library or offline test provider.

## Production Readiness Assessment

| Use Case | Ready? | Notes |
| --- | --- | --- |
| Demonstrate modern transpiler architecture | Yes | The architecture is implemented and verified. |
| Teach compiler pipeline concepts | Yes | Good separation of parser, IR, lowerer, emitter, and tests. |
| Translate controlled toy Java snippets | Partially | Modern path supports a narrow subset; legacy path has broader evidence. |
| Run on a small real Java project | No | Missing classpath, package/import mapping, object model, and broad diagnostics. |
| Use as an enterprise migration assistant | No | Needs deterministic reports, review workflow, CI, and AI sidecar. |
| Fully automated Java-to-Go migration | No | This should not be a near-term goal without a much larger semantic engine. |

## Recommended Next Maturity Gates

### Gate 1: Reach Level 3 - controlled internal tool

Goal: make the modern Go pipeline as trustworthy as the legacy reference suite.

Required work:

- Port the 30 legacy evidence cases to the modern pipeline.
- Add modern unsupported-feature diagnostics with stable `JTG` codes.
- Replace string-prefix lowering checks with typed parse-context handling.
- Add GitHub Actions for:
  - `go test ./...`
  - legacy `go run ./tools/legacytest`
  - modern CLI golden check
  - parser generation drift check
- Add CLI tests for `cmd/j2go`.
- Add a `--version` flag.

Exit criteria:

- Modern pipeline passes at least 30 golden/diagnostic cases.
- CI blocks regressions.
- Unsupported constructs fail deterministically before Go output is written.

### Gate 2: Add deterministic migration intelligence

Goal: make the tool useful before full transpilation coverage exists.

Required work:

- Implement issue #2: deterministic Markdown/JSON migration reports.
- Report per-file unsupported features.
- Report feature counts and risk categories.
- Preserve source spans in diagnostics.
- Emit all diagnostics per file, not only the first.
- Add fixtures for report output.

Exit criteria:

- A Java codebase can be scanned even when it cannot be transpiled.
- The user gets a prioritized migration plan with stable evidence.

### Gate 3: Add AI sidecar safely

Goal: use AI for explanation and review, not as the compiler.

Required work:

- Implement issue #3 with versioned prompt templates.
- Add an offline/canned provider for tests.
- Feed AI only structured diagnostics and generated Go.
- Mark AI output as advisory.
- Add tests proving the default transpilation path never calls AI.

Exit criteria:

- AI can explain diagnostics and suggest review notes.
- Builds remain deterministic without AI credentials.

### Gate 4: Start real Java project support

Goal: move from language-subset translation to project-aware migration.

Required work:

- Decide whether to stay with ANTLR-only semantics or integrate JavaParser/Spoon for richer Java analysis.
- Add package/import mapping.
- Add class fields -> Go structs.
- Add instance methods -> Go receivers.
- Add constructors -> factory functions or initializer patterns.
- Add arrays and indexing.
- Add method call resolution.
- Add exception strategy: explicit `error` returns, `defer`, and cleanup translation policy.

Exit criteria:

- The tool can scan and partially migrate a small real Java project with documented manual follow-up.

## Immediate Priority Order

1. **Modern diagnostics parity**

   The modern path should reject unsupported features as clearly as the legacy path.

2. **Modern golden suite expansion**

   Move the 30-case evidence suite toward the ANTLR/IR pipeline.

3. **CI**

   Add GitHub Actions before the codebase grows further.

4. **Migration report**

   Implement deterministic reports before AI sidecar work.

5. **Lowerer hardening**

   Remove `GetText()`-driven structural decisions where ANTLR contexts already provide typed access.

## Bottom Line

The project has the right architecture and enough implementation to be considered a real MVP. It is mature enough to guide development and demonstrate the Cervo approach.

It is not yet mature enough to migrate real Java applications automatically. The next useful milestone is not "more AI"; it is deterministic migration reporting, modern diagnostic parity, broader golden coverage, and CI.
