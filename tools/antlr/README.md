# ANTLR Tooling

The modern parser is generated from `grammar/JavaSubset.g4`.

Requirements:

- JDK 17 or newer for running the ANTLR tool jar.
- Go module dependency `github.com/antlr4-go/antlr/v4 v4.13.1`.

Generate parser files:

```bash
make generate-parser
```

Equivalent direct command:

```bash
go run ./tools/antlrgen -jar tools/antlr/antlr-4.13.1-complete.jar -grammar grammar/JavaSubset.g4 -out internal/parser/gen
```

`tools/antlrgen` downloads the ANTLR jar when it is missing. Set `JTG_JAVA` if the desired JDK is not first on `PATH`.
