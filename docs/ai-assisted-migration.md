# AI-Assisted Migration Layer

This repo keeps the compiler deterministic and puts AI around it as an optional migration assistant.

Default pipeline:

```text
Java source -> ANTLR4 parser -> Cervo IR -> semantic analysis -> Go generator -> gofmt/tests
```

AI sidecar pipeline:

```text
deterministic report + structured diagnostics + explicit generated Go snippets
  -> optional AI migration assistant
  -> advisory explanation, risk summary, review notes, proposed tests
```

## Rule

AI must not be the default source of truth for generated Go.

The compiler decides what is supported and what gets emitted. AI can explain diagnostics, summarize migration risk, propose reviewed golden tests, and draft advisory code for unsupported constructs. Those outputs must be clearly marked as suggestions.

## Current Sidecar

The current implementation exposes the sidecar through:

```bash
j2go ai explain --report build/migration-report.json --out build/ai-review.md --provider canned
j2go ai explain --report build/migration-report.json --out build/ai-review.md --provider external
```

The `canned` provider is deterministic and runs without network access or API keys. The `external` provider is opt-in and reads its executable from `J2GO_AI_COMMAND` plus optional JSON-string-array arguments from `J2GO_AI_ARGS`; it receives a bounded JSON payload on stdin and returns advisory Markdown on stdout.

Versioned prompts are documented in [ai-prompts.md](ai-prompts.md).

## Useful AI Features

### 1. Migration Report

Generate a deterministic report first:

```bash
j2go report ./my-java-lib --format json --out build/migration-report.json
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

Add explicit generated snippets only when review of generated Go is desired:

```bash
j2go ai explain --report build/migration-report.json --out build/ai-review.md --snippet ./go-lib/com/example/Service.go
```

### 4. Golden Test Assistance

AI can propose expected Go for new fixtures, but the reviewed golden file remains the truth. The test runner must never call AI.

### 5. Idiomatic Go Review

After deterministic generation, AI can suggest cleaner Go patterns such as explicit `error` returns, simpler package-level functions, or `context.Context` for cancellation-heavy APIs.

## Implementation Order

1. Add deterministic Markdown/JSON report output. Done.
2. Add `docs/ai-prompts.md` with versioned prompts. Done.
3. Add an offline AI provider for tests. Done.
4. Add an optional real-provider hook that reads JSON on stdin. Done.
5. Keep AI out of default `transpile` and `migrate` paths. Enforced by CLI tests.

This keeps Cervo's migration workflow practical without sacrificing reproducibility.
