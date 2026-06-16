# Compatibility Matrix

## Toolchain

| Component | Supported version | Notes |
| --- | --- | --- |
| Go | 1.25.x | Used for CI, local tests, generated module defaults, and release builds. |
| Java runtime | 21 | Used in CI to run ANTLR parser generation. |
| ANTLR | 4.13.1 | Parser generation source of truth for `internal/parser/gen`. |
| Maven | Metadata scan only | `pom.xml` is detected, but dependency resolution is not implemented yet. |
| Gradle | Metadata scan only | `build.gradle` and `build.gradle.kts` are detected, but dependency resolution is not implemented yet. |
| Spring | Exploratory only | Deferred until pure Java library migration is stronger. |

## Operating Systems

Release workflow builds `j2go` binaries for:

- Windows amd64 / arm64
- Linux amd64 / arm64
- macOS amd64 / arm64

## Java Language Coverage

The productive target is pure Java libraries first. The current supported subset includes static methods, primitive types, simple control flow, fields, constructors, instance methods, direct `new Class(...)`, and simple interface declarations.

Unsupported or partial areas include inheritance, `implements`, overloads, generics, exceptions, annotations, broad arrays, reflection, broad standard-library mapping, and framework behavior.
