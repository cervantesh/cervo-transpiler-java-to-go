# ANTLR Tooling

The modern parser is generated from `grammar/JavaSubset.g4`.

Requirements:

- JDK 17 or newer for running the ANTLR tool jar.
- Go module dependency `github.com/antlr4-go/antlr/v4 v4.13.1`.

Generate parser files:

```powershell
make generate-parser
```

On this workstation `java` on `PATH` currently points to Java 8. `tools/antlr/generate-parser.ps1` prefers `C:\dev\jdk-24.0.1\bin\java.exe` when it is available, then falls back to `java`.
