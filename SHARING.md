# Sharing Rules

This project is open source under the Apache License, Version 2.0. These rules explain how to share the project, examples, reports, corpus evidence, and contributions safely.

This document is project policy, not legal advice.

## What You Can Share

You may use, copy, modify, distribute, and sublicense this project under the terms of the Apache License, Version 2.0.

When sharing this project or a modified version:

- Include the `LICENSE` file.
- Include the `NOTICE` file when distributing source, binaries, or substantial derivatives.
- Preserve copyright, patent, trademark, and attribution notices.
- Mark files that you changed.
- Do not imply the original maintainers endorse your fork, product, service, or migration results.

## Contributions

By submitting a pull request, issue comment, patch, fixture, prompt, report, or documentation change, you agree that the contribution is licensed under Apache-2.0 unless you clearly mark it as `Not a Contribution`.

Only contribute material you have the right to share.

## Java Inputs And Generated Output

The Apache-2.0 license applies to this transpiler and its repository content. It does not automatically grant rights to:

- Java projects used as migration input.
- Third-party libraries in external corpora.
- Proprietary code snippets in issues or reports.
- Generated Go that is derived from Java code you do not own or cannot relicense.

Before publishing generated Go, verify that you have rights to the original Java input and any dependency code that shaped the output.

## Corpus And Evidence Rules

Use public corpus projects only when their licenses allow the intended use.

For external corpus evidence:

- Pin source repositories and commits in `corpus/corpus.json`.
- Keep generated evidence reproducible.
- Prefer reports, summaries, and generated output over vendoring third-party source.
- Do not include private repositories, customer code, internal package names, credentials, or confidential migration findings.
- If a corpus license requires attribution, preserve that attribution in the evidence notes.

## AI Sidecar Rules

AI support is advisory only. Do not use AI output as deterministic generated code unless a future design explicitly adds review gates for that behavior.

When using external AI providers:

- Do not send secrets, credentials, customer code, or proprietary source unless you have explicit permission.
- Prefer deterministic reports, diagnostics, and minimal snippets over full repositories.
- Mark AI-generated analysis as advisory.
- Keep tests offline by using the canned provider.

## Security And Private Disclosure

Do not disclose vulnerabilities in public issues before maintainers have had a chance to review them.

Use the process in `SECURITY.md` for suspected vulnerabilities, credential leaks, unsafe generated code behavior, or supply-chain concerns.
