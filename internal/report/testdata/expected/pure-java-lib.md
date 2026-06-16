# Java-To-Go Migration Report

Project: `<PROJECT_ROOT>`

## Summary

| Metric | Count |
| --- | ---: |
| Java files | 3 |
| Packages | 2 |
| Classes | 3 |
| Methods | 4 |
| Constructors | 1 |
| Fields | 1 |
| Unsupported features | 3 |
| Diagnostics | 0 |
| Internal imports | 1 |

## Build Files

- `pom.xml` (maven)

## Packages

- `com.example`
- `com.example.test`

## Feature Counts

| Feature | Count |
| --- | ---: |
| class declarations | 3 |
| import declarations | 1 |
| package declarations | 3 |
| static methods | 3 |

## Unsupported Feature Counts

| Feature | Count |
| --- | ---: |
| class fields | 1 |
| constructors | 1 |
| instance methods | 1 |

## Files

### `src/main/java/com/example/Calculator.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.example`
- Classes: `com.example.Calculator`
- Fields in `com.example.Calculator`: `int last`
- Constructors in `com.example.Calculator`: `Calculator(0 params)`
- Methods in `com.example.Calculator`: `instance int add(2 params)`, `static int twice(1 params)`
- Unsupported:
  - `class fields` (JTG1016): Add struct field lowering before transpiling Java fields.
  - `constructors` (JTG1006): Add constructor lowering before transpiling Java constructors.
  - `instance methods` (JTG1017): Add Go receiver generation before transpiling instance methods.

### `src/main/java/com/example/MathUtil.java`

- Source root: `src/main/java`
- Risk: `medium`
- Package: `com.example`
- Classes: `com.example.MathUtil`
- Methods in `com.example.MathUtil`: `static int inc(1 params)`

### `src/test/java/com/example/CalculatorTest.java`

- Source root: `src/test/java`
- Risk: `medium`
- Package: `com.example.test`
- Imports: `com.example.Calculator`
- Classes: `com.example.test.CalculatorTest`
- Methods in `com.example.test.CalculatorTest`: `static int smoke(0 params)`

## Recommended Migration Order

- `src/main/java/com/example/MathUtil.java` (medium risk)
- `src/test/java/com/example/CalculatorTest.java` (medium risk)
- `src/main/java/com/example/Calculator.java` (high risk)
