# Structured Diagnostics

The legacy and modern transpilers run a deterministic unsupported-feature scan before invoking their parsers.

This scan is intentionally not a full Java parser. It exists to turn common out-of-scope Java constructs into stable diagnostics that can be consumed by tests, migration reports, and the AI-assisted sidecar.

Diagnostic format:

```text
error: <file>:<line>:<column>: <code>: unsupported feature: <feature>
  recommendation: <next step>
```

Example:

```text
error: tests\fixtures\unsupported_inheritance.java:1:37: JTG1004: unsupported feature: inheritance
  recommendation: Rewrite inheritance as composition or add an object model lowering pass.
```

## Codes

| Code | Feature | Recommendation |
| --- | --- | --- |
| `JTG1001` | package declarations | Remove the package declaration for the current subset or add package-to-module mapping. |
| `JTG1002` | import declarations | Inline supported code or add import mapping before transpiling. |
| `JTG1003` | legacy interfaces | Add interface lowering before using Java interfaces in the legacy path. |
| `JTG1004` | inheritance | Rewrite inheritance as composition or add an object model lowering pass. |
| `JTG1005` | interface implementation | Add Java interface mapping before transpiling implements clauses. |
| `JTG1006` | unsupported object construction | Add or simplify class/constructor lowering before using non-direct new expressions. |
| `JTG1007` | try/catch exceptions | Design explicit Go error returns before transpiling exceptions. |
| `JTG1008` | catch blocks | Design explicit Go error returns before transpiling exceptions. |
| `JTG1009` | finally blocks | Map cleanup to defer before transpiling finally blocks. |
| `JTG1010` | throw statements | Map thrown exceptions to Go error values before transpiling throw. |
| `JTG1011` | throws declarations | Map checked exceptions to explicit Go error returns. |
| `JTG1012` | lambdas | Rewrite the lambda as a named method or add functional-interface lowering. |
| `JTG1013` | annotations | Remove annotations or add an annotation policy for migration metadata. |
| `JTG1014` | generics | Specialize the generic type manually or add a Java-to-Go type-parameter mapping pass. |
| `JTG1015` | arrays and indexing | Add array declarations, indexing, and length lowering before using arrays. |
| `JTG1016` | legacy class fields | Add struct field lowering before transpiling Java fields in the legacy path. |
| `JTG1017` | legacy instance methods | Add Go receiver generation before transpiling instance methods in the legacy path. |
| `JTG1018` | method overloading | Rename overloaded methods or add overload name mangling before generating Go. |
| `JTG1019` | assignment expressions | Move the assignment to a standalone statement before transpiling. |
| `JTG1020` | standalone blocks | Inline the block or add block-scope lowering before transpiling nested standalone blocks. |
| `JTG2001` | duplicate symbols | Rename the duplicate package, class, field, method, parameter, or local variable before migration. |
| `JTG2002` | unknown symbols | Declare the symbol, fix the name, or add the missing import/resolution rule before migration. |
| `JTG2003` | incompatible types | Rewrite the expression or add an explicit conversion/mapping before migration. |
| `JTG2004` | unsupported overloads | Rename overloaded Java methods or add a deterministic overload name-mangling policy. |
| `JTG2005` | unresolved internal imports | Add the missing source file, fix the package/import name, or mark the dependency as external. |

## Test Coverage

`go run ./tools/legacytest` validates the legacy structured diagnostics with these fixtures. The modern Go pipeline keeps structured diagnostics for still-unsupported features and adds positive object-model/interface coverage in `internal/pipeline`.

- `tests\fixtures\unsupported_field.java`
- `tests\fixtures\unsupported_package_import.java`
- `tests\fixtures\unsupported_exception.java`
- `tests\fixtures\unsupported_instance_method.java`
- `tests\fixtures\unsupported_overload.java`
- `tests\fixtures\unsupported_interface.java`
- `tests\fixtures\unsupported_inheritance.java`
- `tests\fixtures\unsupported_generics.java`
- `tests\fixtures\unsupported_array_indexing.java`

Each fixture must fail before Go output is written and must include:

- a stable `JTG` diagnostic code
- the unsupported feature name
- a migration recommendation

These diagnostics make unsupported cases reproducible for migration planning and automated review.
