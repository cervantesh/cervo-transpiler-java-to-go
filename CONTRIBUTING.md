# Contributing

Thank you for contributing to `cervo-transpiler-java-to-go`.

This project is a deterministic Java-to-Go transpiler and migration toolkit. Contributions should preserve reproducible output, structured diagnostics, and cross-platform operation.

## Contribution License

Unless you explicitly mark a submission as `Not a Contribution`, any contribution you intentionally submit to this project is provided under the Apache License, Version 2.0.

Do not submit code, tests, corpus files, logs, prompts, reports, or snippets unless you have the right to share them under terms compatible with this project.

## Development Rules

- Keep generated Go deterministic and reproducible.
- Prefer structured diagnostics over generated placeholder code.
- Do not make AI part of the default deterministic generation path.
- Keep command examples portable. Prefer `go run ...` and `go test ./...`.
- Do not add PowerShell-only project workflows.
- Do not commit generated build outputs from `build/`, `tests/generated/`, `modern-tests/generated/`, or `.corpus/`.
- Add or update golden fixtures when changing generated output.
- Keep unsupported Java features explicit through `JTG` diagnostics.

## Local Verification

Run the core suite before opening a pull request:

```bash
go test ./...
go run ./tools/legacytest
go run ./tools/corpus
```

When changing the grammar, also regenerate the parser:

```bash
go run ./tools/antlrgen -jar tools/antlr/antlr-4.13.1-complete.jar -grammar grammar/JavaSubset.g4 -out internal/parser/gen
```

## Pull Request Checklist

- The change is scoped to one clear behavior or documentation update.
- Tests or fixtures were added for behavior changes.
- Documentation was updated when commands, flags, or migration behavior changed.
- New diagnostics use stable `JTG` IDs.
- External Java corpus evidence does not include proprietary or confidential code.
- AI output, if any, is clearly marked advisory.

## Reporting Issues

Use the GitHub issue templates when possible. Include:

- Operating system and Go version.
- Exact command.
- Minimal Java input or fixture.
- Expected result.
- Actual result.
- Diagnostic output, if available.

Do not post secrets, private repository URLs, customer code, proprietary source files, API keys, tokens, or confidential migration reports in public issues.
