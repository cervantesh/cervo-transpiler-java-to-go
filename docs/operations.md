# Operations Guide

This guide documents stable CLI behavior for running `j2go` as an internal migration tool.

## Install

For tagged releases, download the binary for your OS and architecture from GitHub Releases and verify `checksums.txt`.

Local build:

```powershell
go build -o build\j2go.exe ./cmd/j2go
.\build\j2go.exe --version
```

## Configuration

Use `cervo-migration.yaml` for repeatable migration jobs. Start from `cervo-migration.example.yaml`.

Validate config:

```powershell
j2go config validate --config cervo-migration.yaml
```

Supported config fields:

```yaml
version: 1
project:
  source: ./java-lib
  module: example.com/migrated
migration:
  out: ./go-lib
  report: build/migration-report.md
  dryRun: true
logs:
  file: build/j2go.log
ai:
  provider: canned
```

## Scan

```powershell
j2go scan ./java-lib
```

Exit code `0` means scan completed. Exit code `1` means runtime failure, such as unreadable files. Exit code `2` means CLI usage error.

## Report

```powershell
j2go report ./java-lib --format json --out build/migration-report.json
j2go report ./java-lib --format markdown --out build/migration-report.md
```

Reports are deterministic for the same input tree and tool version.

## Migrate

Dry run from config:

```powershell
j2go migrate --config cervo-migration.yaml
```

Explicit command:

```powershell
j2go migrate ./java-lib --out ./go-lib --report build/migration-report.md --log-file build/j2go.log --dry-run
j2go migrate ./java-lib --out ./go-lib --report build/migration-report.md --log-file build/j2go.log
```

The migrate command may generate supported files and skip unsupported files in the same run. It maps supported internal Java imports to deterministic Go imports, writes supported Java test fixtures as `_test.go`, and reports Go symbol collisions before emitting ambiguous output. Skipped work is reported through stderr, the Markdown report, and the operation log.

## Review

Generate advisory review notes from deterministic reports:

```powershell
j2go ai explain --report build/migration-report.json --out build/ai-review.md --provider canned
```

Use `--provider external` only when `J2GO_AI_COMMAND` is configured. The external command receives the bounded JSON payload on stdin and must return Markdown on stdout.

## Exit Codes

| Code | Meaning |
| ---: | --- |
| 0 | Command succeeded. Partial migration can still report skipped files. |
| 1 | Runtime or processing error. |
| 2 | CLI usage error. |

## Logs

`j2go migrate --log-file <path>` writes a deterministic key-value log:

```text
command=migrate
out=./go-lib
report=build/migration-report.md
dryRun=true
javaFiles=2
generated=1
skipped=1
diagnostics=1
```

Logs intentionally omit timestamps so CI comparisons remain stable.
