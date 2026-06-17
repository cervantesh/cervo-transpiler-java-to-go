# Migration Corpus

This directory defines reproducible Java corpus projects for Java-to-Go migration runs.

The checked-in manifest is [`corpus.json`](corpus.json). It includes:

- Curated local projects that demonstrate high-confidence multi-file transpilation.
- External Git repositories pinned by URL and exact commit for stress/reporting evidence.

## Layout

```text
corpus/
  corpus.json
  projects/          # curated local projects
tools/
  corpus/
    main.go
.corpus/             # local clones, ignored by git
evidence/corpus/     # generated migration evidence, versionable
```

## Run

Use the Go runner for cross-platform behavior:

```bash
go run ./tools/corpus
```

Useful options:

```bash
go run ./tools/corpus -repo jsemver
go run ./tools/corpus -repo curated-store-lib
go run ./tools/corpus -skip-clone
go run ./tools/corpus -manifest corpus/corpus.json
```

For external Git repositories, the runner clones or updates each repository under `.corpus/` and checks out the pinned commit. For curated local projects, it reads the source directly from `corpus/projects/`.
Corpus entries may set `sourceRoot` to analyze a deterministic subdirectory such as `src/main/java`. This keeps third-party test suites and build fixtures from diluting the library migration signal while still preserving the pinned repository metadata in `source.json`.

For every corpus project it then runs:

```bash
go run ./cmd/j2go scan <source-root>
go run ./cmd/j2go report <source-root> --format json --out <evidence>/report.json
go run ./cmd/j2go report <source-root> --format markdown --out <evidence>/report.md
go run ./cmd/j2go migrate <source-root> --out <evidence>/go --report <evidence>/migration.md --log-file <evidence>/migration.log
```

Each corpus output directory also receives `source.json`, `scan.txt`, and `migrate.txt` so the evidence can be reviewed without rerunning the tool.
The runner also executes `go test ./...` inside the generated Go module and saves the transcript as `go-test.txt`.

## Evidence Categories

| Category | Purpose | Expected signal |
| --- | --- | --- |
| `curated` | Prove deterministic multi-file transpilation inside the supported subset. | High generated/skipped ratio, ideally 100% generated. |
| `external` | Stress the scanner, reports, diagnostics, partial migration, and unsupported-feature visibility against real Java libraries. | Partial generation is acceptable; skipped files are expected until Java coverage improves. |

## Current Corpus

| ID | Category | Source | Source root | Why it is useful |
| --- | --- | --- | --- | --- |
| `curated-store-lib` | `curated` | `corpus/projects/curated-store-lib` | Full curated project | Controlled multi-package library with internal imports, object model, static methods, fields, and a generated Go test. |
| `jsemver` | `external` | <https://github.com/zafarkhaja/jsemver> | `src/main/java` | Small Maven library with domain objects, parsing, and comparison logic. |
| `commons-csv` | `external` | <https://github.com/apache/commons-csv> | `src/main/java` | Real Apache library with a compact API and CSV-oriented object model. |
| `java-diff-utils` | `external` | <https://github.com/java-diff-utils/java-diff-utils> | `java-diff-utils/src/main/java` | Medium-sized core library module with algorithms, patches, and package structure. |
