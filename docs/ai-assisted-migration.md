# AI-Assisted Migration Layer

This repo keeps the compiler deterministic and puts AI around it as an optional migration assistant.

Default pipeline:

```text
Java source -> Flex lexer -> Bison parser -> C++ AST -> Go generator -> gofmt/tests
```

AI sidecar pipeline:

```text
Java source + structured diagnostics + generated Go
  -> AI migration assistant
  -> report, explanation, review notes, optional draft suggestions
```

## Rule

AI must not be the default source of truth for generated Go.

The compiler decides what is supported and what gets emitted. AI can explain diagnostics, summarize migration risk, propose reviewed golden tests, and draft advisory code for unsupported constructs. Those outputs must be clearly marked as suggestions.

## First-Phase Foundation

The current first phase already provides the deterministic foundation needed by an AI sidecar:

- structured `JTG` diagnostics in `src/diagnostics.hpp` and `src/diagnostics.cpp`
- unsupported-feature fixtures in `tests/fixtures/unsupported_*.java`
- reproducible diagnostics documented in `docs/diagnostics.md`
- a 30-case test suite in `test.ps1`

## Useful AI Features

### 1. Migration Report

Add a report mode that scans Java files and summarizes risk.

Example future command:

```powershell
.\build\javago.exe --report src build\migration-report.md
```

Report contents:

- files analyzed
- supported features found
- unsupported feature counts
- risk level per file or method
- recommended implementation order

### 2. Diagnostic Explainer

Turn a deterministic compiler diagnostic into a human migration note.

Compiler diagnostic:

```text
OrderService.java:42:13: JTG1007: unsupported feature: try/catch exceptions
```

AI explanation:

```text
This method uses Java exceptions. Go does not have try/catch, so the migration needs explicit error returns.
```

### 3. Controlled Fallback Suggestions

Add an opt-in flag for advisory draft code.

Example future command:

```powershell
.\build\javago.exe --ai-suggest examples\Main.java build\Main.go
```

Suggested code must include a warning:

```go
// AI suggestion: review before using.
// Original Java construct was not supported by the deterministic transpiler.
```

### 4. Golden Test Assistance

AI can propose expected Go for new fixtures, but the reviewed golden file remains the truth. The test runner must never call AI.

### 5. Idiomatic Go Review

After deterministic generation, AI can suggest cleaner Go patterns such as explicit `error` returns, simpler package-level functions, or `context.Context` for cancellation-heavy APIs.

## Proposed Files

- `src/report_generator.hpp`
- `src/report_generator.cpp`
- `docs/ai-prompts.md`
- `scripts/ai_explain.ps1`
- `tests/fixtures/report_*.java`
- `tests/expected/*.stderr.txt`

## Implementation Order

1. Add deterministic Markdown/JSON report output.
2. Add `docs/ai-prompts.md` with versioned prompts.
3. Add an offline AI provider or canned-response script for tests.
4. Add an optional real-provider script that reads diagnostics JSON.
5. Add `--ai-suggest` only after report output is stable.

This keeps Cervo's migration workflow practical without sacrificing reproducibility.
