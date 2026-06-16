# Modern MVP Verification - 2026-06-16

This file records verification for the ANTLR4 + Cervo IR Java-to-Go MVP.

## Environment

```text
go version go1.25.6 windows/amd64
openjdk version "24.0.1" 2025-04-15
ANTLR tool jar 4.13.1
github.com/antlr4-go/antlr/v4 v4.13.1
```

## Verified Pipeline

```text
Java source -> ANTLR4 parse tree -> parser facade -> Cervo IR -> Go AST -> go/format -> golden tests
```

## Commands

Parser generation:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File tools\antlr\generate-parser.ps1 -Jar tools/antlr/antlr-4.13.1-complete.jar -Grammar grammar/JavaSubset.g4 -OutputDir internal/parser/gen
```

Go test suite:

```powershell
go test ./...
```

CLI verification:

```powershell
go build -o build\j2go.exe ./cmd/j2go
.\build\j2go.exe -out modern-tests\generated modern-tests\fixtures\hello\Hello.java
gofmt -w modern-tests\generated\Hello.go
```

Result:

```text
PASS modern CLI golden
```

## Acceptance Criteria Evidence

| Criterion | Evidence |
| --- | --- |
| ANTLR4 parser facade compiles and reports syntax diagnostics | `internal/parser` tests pass. |
| Cervo IR and semantic diagnostics are implemented | `internal/ir` and `internal/semantic` tests pass. |
| Lowerer maps MVP Java subset into IR | `internal/lower` tests pass. |
| Go emitter uses `go/ast` and `go/format` | `internal/emitgo` tests pass and emitter implementation uses Go AST nodes. |
| CLI `j2go` passes golden end-to-end tests | `internal/pipeline` golden test and manual CLI golden check pass. |
