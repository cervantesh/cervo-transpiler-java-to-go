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
| Unsupported features | 0 |
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
| class fields | 1 |
| constructors | 1 |
| import declarations | 1 |
| instance methods | 1 |
| package declarations | 3 |
| static methods | 3 |

## Unsupported Feature Counts

None.

## Files

### `src/main/java/com/example/Calculator.java`

- Source root: `src/main/java`
- Risk: `medium`
- Package: `com.example`
- Classes: `com.example.Calculator`
- Fields in `com.example.Calculator`: `int last`
- Constructors in `com.example.Calculator`: `Calculator(0 params)`
- Methods in `com.example.Calculator`: `instance int add(2 params)`, `static int twice(1 params)`

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

- `src/main/java/com/example/Calculator.java` (medium risk)
- `src/main/java/com/example/MathUtil.java` (medium risk)
- `src/test/java/com/example/CalculatorTest.java` (medium risk)
