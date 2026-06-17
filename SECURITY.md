# Security Policy

## Supported Versions

Security review focuses on the default branch and the latest published release.

| Version | Supported |
| --- | --- |
| `main` | Yes |
| latest release | Yes |
| older releases | Best effort |

## Reporting A Vulnerability

Do not open a public issue for vulnerabilities, credential leaks, or private migration data exposure.

Use GitHub private vulnerability reporting for this repository when available. If private reporting is not available, contact the maintainer through a non-public channel and include only the minimum information needed to reproduce the issue.

Include:

- Affected commit, release, or command.
- Minimal reproduction steps.
- Expected security boundary.
- Actual behavior.
- Whether generated Go, migration reports, logs, AI sidecar output, or corpus evidence are involved.

## Security Boundaries

The project should not:

- Leak secrets into generated Go, reports, logs, corpus evidence, or AI prompts.
- Execute untrusted generated code as part of normal scan/report workflows.
- Call external AI providers unless the user explicitly opts in.
- Treat advisory AI output as deterministic transpiler output.
- Vendor private or proprietary Java corpus material.

## Disclosure Process

Maintainers will triage reports, request reproduction details when needed, and coordinate a fix before public disclosure when appropriate.
