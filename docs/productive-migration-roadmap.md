# Productive Migration Roadmap

This roadmap moves `cervo-transpiler-java-to-go` from a technical MVP to a productive Java-to-Go migration tool.

The product direction combines two goals:

- increasingly automatic Java-to-Go transpilation
- enterprise migration support with reports, CI, traceability, releases, and review workflows

The first real corpus is **pure Java libraries**. Spring and framework-specific migration work is intentionally deferred until the Java language and project model are reliable.

## Phase Issues

| Phase | Issue | Result |
| --- | --- | --- |
| Phase 3 | [#8](https://github.com/cervantesh/cervo-transpiler-java-to-go/issues/8) `modern pipeline hardening and CI gate` | Raise the modern ANTLR/IR pipeline to a reliable internal-tool baseline. |
| Phase 4 | [#9](https://github.com/cervantesh/cervo-transpiler-java-to-go/issues/9) `pure Java library project model` | Discover and scan real Java library projects with packages, imports, source roots, and build files. |
| Phase 5 | [#2](https://github.com/cervantesh/cervo-transpiler-java-to-go/issues/2) `deterministic migration report output` | Emit deterministic JSON/Markdown migration reports that work even when transpilation fails. |
| Phase 6 | [#10](https://github.com/cervantesh/cervo-transpiler-java-to-go/issues/10) `Java semantic engine and typed IR` | Add real semantic resolution and typed IR. |
| Phase 7 | [#11](https://github.com/cervantesh/cervo-transpiler-java-to-go/issues/11) `object model lowering to Go` | Lower Java classes, fields, constructors, instance methods, and interfaces to Go. |
| Phase 8 | [#12](https://github.com/cervantesh/cervo-transpiler-java-to-go/issues/12) `automatic migration pipeline` | Generate complete Go package trees with partial-migration reporting. |
| Phase 9 | [#3](https://github.com/cervantesh/cervo-transpiler-java-to-go/issues/3) `AI sidecar for review and explanation` | Add optional advisory AI explanations and review notes. |
| Phase 10 | [#13](https://github.com/cervantesh/cervo-transpiler-java-to-go/issues/13) `enterprise release and governance` | Add releases, checksums, config, compatibility docs, and stable CLI operations. |
| Phase 11 | [#14](https://github.com/cervantesh/cervo-transpiler-java-to-go/issues/14) `Spring adapter exploration` | Evaluate Spring-specific migration after pure Java support is stable. |

## Product Gates

### Gate 1: Controlled Internal Tool

Tracked by Phase 3.

The modern Go pipeline must reach parity with the existing 30-case legacy evidence suite, run in CI, and fail unsupported Java deterministically with structured diagnostics.

### Gate 2: Useful Migration Scanner

Tracked by Phases 4 and 5.

The tool must scan a real pure Java library, identify packages/classes/methods/features, and produce stable reports even when no Go output is generated.

### Gate 3: Semantically Reliable Transpiler

Tracked by Phases 6 and 7.

The tool must resolve Java symbols and types well enough to lower object-oriented Java constructs into maintainable Go.

### Gate 4: Productive Migration Tool

Tracked by Phases 8 through 10.

The tool must migrate a small Java library into a Go package tree, report manual gaps, support advisory AI review, and ship as a versioned CLI.

### Gate 5: Framework Adapters

Tracked by Phase 11.

Framework behavior such as Spring annotations, dependency injection, controllers, repositories, and exception handlers should be handled only after the core Java migration foundation is stable.

## Defaults

- Pure Java libraries are the first real target.
- Spring support is exploratory and later-stage.
- AI output is advisory and opt-in.
- Deterministic compiler/report output remains the source of truth.
- Partial migration is acceptable only when the report clearly identifies manual work.
