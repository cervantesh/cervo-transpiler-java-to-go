# Migration Corpus

This directory defines the external Java repositories used as reproducible evidence for Java-to-Go migration runs.

The checked-in manifest is [`corpus.json`](corpus.json). It pins each repository by URL and exact commit so a run can be reproduced without vendoring third-party source code into this repository.

## Layout

```text
corpus/
  corpus.json
tools/
  corpus/
    main.go
.corpus/             # local clones, ignored by git
evidence/corpus/     # generated migration evidence, versionable
```

## Run

Use the Go runner instead of PowerShell for better cross-platform behavior:

```bash
go run ./tools/corpus
```

Useful options:

```bash
go run ./tools/corpus -repo jsemver
go run ./tools/corpus -skip-clone
go run ./tools/corpus -manifest corpus/corpus.json
```

The runner clones or updates each repository under `.corpus/`, checks out the pinned commit, then runs:

```bash
go run ./cmd/j2go scan <repo>
go run ./cmd/j2go report <repo> --format json --out <evidence>/report.json
go run ./cmd/j2go report <repo> --format markdown --out <evidence>/report.md
go run ./cmd/j2go migrate <repo> --out <evidence>/go --report <evidence>/migration.md --log-file <evidence>/migration.log
```

Each corpus output directory also receives `source.json`, `scan.txt`, and `migrate.txt` so the evidence can be reviewed without rerunning the tool.
The runner also executes `go test ./...` inside the generated Go module and saves the transcript as `go-test.txt`.

## Current Corpus

| ID | Repository | Why it is useful |
| --- | --- | --- |
| `jsemver` | <https://github.com/zafarkhaja/jsemver> | Small Maven library with domain objects, parsing, comparison, and tests. |
| `commons-csv` | <https://github.com/apache/commons-csv> | Real Apache library with a compact API and CSV-oriented object model. |
| `java-diff-utils` | <https://github.com/java-diff-utils/java-diff-utils> | Medium-sized library with algorithms, patches, modules, and package structure. |
