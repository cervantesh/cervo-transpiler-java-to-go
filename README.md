# cervo-transpiler-java-to-go

Java-to-Go transpiler project for the Cervo migration effort.

The target architecture is modern compiler-style translation:

```text
Java source -> ANTLR4 parse tree -> semantic analysis -> Cervo IR -> Go generator -> gofmt/golden tests
```

The detailed implementation plan is in [docs/superpowers/plans/2026-06-16-modern-java-to-go-transpiler.md](docs/superpowers/plans/2026-06-16-modern-java-to-go-transpiler.md), and the architecture summary is in [docs/architecture-modern.md](docs/architecture-modern.md).

Current maturity is documented in [docs/maturity-analysis-2026-06-16.md](docs/maturity-analysis-2026-06-16.md).

The productive migration roadmap is tracked in [docs/productive-migration-roadmap.md](docs/productive-migration-roadmap.md).

## Project Objective

The project goal is to translate a controlled subset of Java into valid Go code through an explicit parser, semantic-analysis, IR, and backend pipeline.

The existing checked-in compiler core is a legacy reference pipeline:

```text
Java source
  -> Flex lexer
  -> Bison parser
  -> C++ AST
  -> deterministic Go generator
  -> gofmt
  -> go run / golden tests
```

The parser builds an AST first, and the Go generator emits code from that AST. AI can assist around the compiler through reports and suggestions, but the default transpilation path remains deterministic.

## Modern MVP

The modern MVP now implements:

- ANTLR4 parser generation from `grammar/JavaSubset.g4`.
- Parser facade with syntax diagnostics in `internal/parser`.
- Cervo IR in `internal/ir`.
- Semantic diagnostics and scoped symbols in `internal/semantic`.
- Parse-tree lowering in `internal/lower`.
- Go emission through `go/ast` and `go/format` in `internal/emitgo`.
- End-to-end pipeline and CLI in `internal/pipeline` and `cmd/j2go`.
- Golden tests under `modern-tests`.

## Legacy Reference

The original reference implementation remains available while the modern pipeline catches up in coverage. It implements:

- Tokenization with Flex.
- Grammar-based parsing with Bison.
- Use of modern Bison features such as location tracking and detailed parse errors.
- AST construction in C++ using ownership-oriented data structures.
- Separation between parsing and code generation.
- Translation from Java syntax into Go syntax.
- Golden-file testing for compiler output.
- Validation of generated Go with `gofmt` and `go run`.
- Error reporting with line and column information for unsupported or invalid input.

## Legacy Frontend

The current frontend uses Flex and Bison because they make the compiler stages explicit and keep the first phase compact:

- Flex defines how raw Java text becomes tokens.
- Bison defines which token sequences form valid Java-subset programs.
- The semantic actions create AST nodes.
- A separate generator walks the AST and emits Go.

A later production migration tool should move toward a richer Java frontend, type resolution, classpath analysis, project-wide semantic understanding, and an AI-assisted migration sidecar.

## Project Structure

```text
.
|-- build.ps1                  Windows build script
|-- Makefile                   Make-style build for compatible environments
|-- go.mod                     Modern Go module
|-- README.md                  Project explanation and usage
|-- cmd/j2go/                  Modern CLI
|-- docs/
|   |-- ai-assisted-migration.md
|   |-- architecture-modern.md
|   |-- diagnostics.md
|   `-- evidence/
|       `-- verification-2026-06-16.md
|-- grammar/
|   `-- JavaSubset.g4          ANTLR grammar for the modern MVP
|-- internal/                  Modern parser, IR, lowerer, emitter, pipeline
|-- modern-tests/              Modern golden fixtures
|-- examples/
|   `-- Main.java              Demo Java input
|-- src/
|   |-- lexer.l                Flex lexer
|   |-- parser.y               Bison parser
|   |-- ast.hpp / ast.cpp      AST node model
|   |-- diagnostics.hpp / .cpp Unsupported-feature diagnostics
|   |-- generator.hpp / .cpp   AST-to-Go code generator
|   |-- parser_driver.hpp      Parser facade
|   `-- main.cpp               CLI entry point
|-- test.ps1                   Golden-test runner
`-- tests/
    |-- fixtures/              Java test inputs
    `-- expected/              Expected formatted Go outputs
```

Generated binaries and test outputs are written to `build/`, `tests/generated/`, and `modern-tests/generated/`, which are intentionally ignored by Git.

## Supported Java Subset

The current version supports:

- One public class.
- Static methods.
- `public static void main(String[] args)`.
- Types: `int`, `double`, `boolean`, `String`, `void`.
- Method parameters and return values.
- Variable declarations:
  - `int x;`
  - `int x = 5;`
- Assignments:
  - `x = x + 1;`
- Expressions:
  - identifiers and literals
  - `+`, `-`, `*`, `/`
  - `<`, `>`, `<=`, `>=`, `==`, `!=`
  - `&&`, `||`, `!`
  - parenthesized expressions
- Control flow:
  - `if (...) { ... } else { ... }`
  - `while (...) { ... }`
- Output:
  - `System.out.println(expr);`
- Comments:
  - `// ...`
  - `/* ... */`
- Object model subset:
  - class fields -> Go struct fields
  - constructors and default constructors -> `NewClass(...) *Class`
  - instance methods -> Go receiver methods
  - direct `new Class(...)` calls -> `NewClass(...)`

Unsupported features still include inheritance, interfaces, packages/imports in direct transpile mode, general arrays, exceptions, generics, lambdas, annotations, overloading, reflection, and broad Java standard-library translation.

Unsupported features are detected before parsing when possible and reported as structured `JTG` diagnostics with line, column, feature name, and recommendation. See [docs/diagnostics.md](docs/diagnostics.md).

## Translation Rules

The transpiler uses deterministic mappings:

| Java | Go |
| --- | --- |
| `public class Main { ... }` | class wrapper is not emitted |
| `public static void main(String[] args)` | `func main()` |
| static Java method | top-level Go function |
| `int` | `int` |
| `double` | `float64` |
| `boolean` | `bool` |
| `String` | `string` |
| `void` | no return type |
| `while (condition)` | `for condition` |
| `System.out.println(x)` | `fmt.Println(x)` |

The generator only imports `fmt` when the AST contains a print statement.

## Example

Input Java:

```java
public class Main {
    public static void main(String[] args) {
        int x = 3;
        while (x < 5) {
            x = x + 1;
        }

        if (x >= 5) {
            System.out.println(add(x, 2));
        } else {
            System.out.println(0);
        }
    }

    public static int add(int a, int b) {
        return a + b;
    }
}
```

Generated Go:

```go
package main

import "fmt"

func main() {
	var x int = 3
	for x < 5 {
		x = x + 1
	}
	if x >= 5 {
		fmt.Println(add(x, 2))
	} else {
		fmt.Println(0)
	}
}

func add(a int, b int) int {
	return a + b
}
```

Program output:

```text
7
```

## Modern Build

Generate the ANTLR parser and run modern tests:

```powershell
New-Item -ItemType Directory -Force tools\antlr | Out-Null
Invoke-WebRequest -Uri https://www.antlr.org/download/antlr-4.13.1-complete.jar -OutFile tools\antlr\antlr-4.13.1-complete.jar
powershell -NoProfile -ExecutionPolicy Bypass -File tools\antlr\generate-parser.ps1 -Jar tools/antlr/antlr-4.13.1-complete.jar -Grammar grammar/JavaSubset.g4 -OutputDir internal/parser/gen
go test ./...
go build -o build\j2go.exe ./cmd/j2go
```

Run the modern CLI:

```powershell
.\build\j2go.exe -out modern-tests\generated modern-tests\fixtures\hello\Hello.java
gofmt -w modern-tests\generated\Hello.go
```

Scan a pure Java library project:

```powershell
.\build\j2go.exe scan .\internal\javaproject\testdata\pure-java-lib
```

Generate deterministic migration reports:

```powershell
.\build\j2go.exe report .\internal\javaproject\testdata\pure-java-lib --format json --out build\migration-report.json
.\build\j2go.exe report .\internal\javaproject\testdata\pure-java-lib --format markdown --out build\migration-report.md
```

The report mode scans packages, imports, classes, methods, fields, constructors, unsupported features, diagnostics, internal dependencies, per-file risk, and recommended migration order. It does not call AI.

## Legacy Build

Windows PowerShell:

```powershell
.\build.ps1
```

Manual equivalent:

```powershell
bison -d -o build\parser.cpp src\parser.y
flex -o build\lexer.cpp src\lexer.l
g++ -std=c++17 -Isrc -Ibuild build\parser.cpp build\lexer.cpp src\ast.cpp src\diagnostics.cpp src\generator.cpp src\main.cpp -o build\javago.exe
```

The full build script also compiles `src\diagnostics.cpp`; use `.\build.ps1` as the source of truth for local builds.

If `flex` or `bison` is missing, install WinFlexBison:

```powershell
scoop install winflexbison
```

or:

```powershell
choco install winflexbison3
```

## Run

```powershell
.\build\javago.exe examples\Main.java build\Main.go
gofmt -w build\Main.go
go run build\Main.go
```

Expected output:

```text
7
```

## Test

Run the golden-test suite:

```powershell
.\test.ps1
```

The test runner:

- Rebuilds the transpiler.
- Executes exactly 30 evidence cases.
- Transpiles Java fixtures from `tests/fixtures`.
- Formats generated Go with `gofmt`.
- Compares generated Go against `tests/expected`.
- Runs generated Go programs with `go run`.
- Verifies syntax errors include line and column details.
- Verifies unsupported Java features fail with structured `JTG` diagnostics and recommendations.

The 30-case suite includes 20 supported Java-to-Go translations, one syntax-error case, and nine unsupported-feature diagnostics.

## Verification Evidence

Verification evidence is stored in:

[docs/evidence/verification-2026-06-16.md](docs/evidence/verification-2026-06-16.md)

That file records the tool versions, build output, golden-test output, and negative syntax/unsupported-feature tests.

## Known Limitations

This phase is still intentionally small. It does not perform full classpath resolution, method overload resolution, interface lowering, inheritance lowering, or broad library migration.

Examples of unsupported input should fail with a structured diagnostic or parser error instead of producing misleading Go code. For example, interfaces are rejected before parsing with `JTG1003`.

## AI-Assisted Migration

AI belongs beside the deterministic transpiler, not inside the default code-generation path. See [docs/ai-assisted-migration.md](docs/ai-assisted-migration.md) for the sidecar architecture and implementation order.
