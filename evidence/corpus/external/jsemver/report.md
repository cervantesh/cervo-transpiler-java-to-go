# Java-To-Go Migration Report

Project: `.corpus/jsemver`

## Summary

| Metric | Count |
| --- | ---: |
| Java files | 44 |
| Packages | 3 |
| Classes | 92 |
| Methods | 213 |
| Constructors | 49 |
| Fields | 36 |
| Unsupported features | 0 |
| Diagnostics | 2432 |
| Internal imports | 61 |

## Build Files

- `pom.xml` (maven)

## Packages

- `com.github.zafarkhaja.semver`
- `com.github.zafarkhaja.semver.expr`
- `com.github.zafarkhaja.semver.util`

## Feature Counts

| Feature | Count |
| --- | ---: |
| class declarations | 92 |
| class fields | 36 |
| constructors | 49 |
| import declarations | 139 |
| instance methods | 194 |
| package declarations | 44 |
| static methods | 19 |

## Unsupported Feature Counts

None.

## Files

### `src/main/java/com/github/zafarkhaja/semver/ParseException.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Classes: `com.github.zafarkhaja.semver.ParseException`
- Fields in `com.github.zafarkhaja.semver.ParseException`: `long serialVersionUID`
- Constructors in `com.github.zafarkhaja.semver.ParseException`: `ParseException(0 params)`, `ParseException(1 params)`, `ParseException(2 params)`
- Methods in `com.github.zafarkhaja.semver.ParseException`: `instance String toString(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/Parser.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Diagnostics: `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/UnexpectedCharacterException.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Imports: `com.github.zafarkhaja.semver.VersionParser.CharType`, `com.github.zafarkhaja.semver.util.UnexpectedElementException`, `java.util.Arrays`
- Classes: `com.github.zafarkhaja.semver.UnexpectedCharacterException`
- Fields in `com.github.zafarkhaja.semver.UnexpectedCharacterException`: `long serialVersionUID`, `Character unexpected`, `int position`, `CharType[] expected`
- Constructors in `com.github.zafarkhaja.semver.UnexpectedCharacterException`: `UnexpectedCharacterException(1 params)`, `UnexpectedCharacterException(2 params)`
- Methods in `com.github.zafarkhaja.semver.UnexpectedCharacterException`: `instance Character getUnexpectedCharacter(0 params)`, `instance int getPosition(0 params)`, `instance CharType[] getExpectedCharTypes(0 params)`, `instance String toString(0 params)`, `static String createMessage(2 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/Version.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Imports: `com.github.zafarkhaja.semver.expr.Expression`, `com.github.zafarkhaja.semver.expr.ExpressionParser`, `java.io.InvalidObjectException`, `java.io.ObjectInputStream`, `java.io.Serializable`, `java.util.Arrays`, `java.util.Comparator`, `java.util.Locale`, `java.util.Optional`, `java.util.function.Predicate`, `staticcom.github.zafarkhaja.semver.Version.Validators.*`, `staticcom.github.zafarkhaja.semver.VersionParser.parseBuild`, `staticcom.github.zafarkhaja.semver.VersionParser.parsePreRelease`
- Classes: `com.github.zafarkhaja.semver.Version`, `com.github.zafarkhaja.semver.Builder`, `com.github.zafarkhaja.semver.Validators`, `com.github.zafarkhaja.semver.SerializationProxy`
- Fields in `com.github.zafarkhaja.semver.Builder`: `long major`, `long minor`, `long patch`, `String[] preReleaseIds`, `String[] buildIds`
- Constructors in `com.github.zafarkhaja.semver.Builder`: `Builder(0 params)`, `Builder(1 params)`
- Methods in `com.github.zafarkhaja.semver.Builder`: `instance Builder setVersionCore(1 params)`, `instance Builder setVersionCore(2 params)`, `instance Builder setVersionCore(3 params)`, `instance Builder setMajorVersion(1 params)`, `instance Builder setMinorVersion(1 params)`, `instance Builder setPatchVersion(1 params)`, `instance Builder setPreReleaseVersion(0 params)`, `instance Builder addPreReleaseIdentifiers(0 params)`, `instance Builder unsetPreReleaseVersion(0 params)`, `instance Builder setBuildMetadata(0 params)`, `instance Builder addBuildIdentifiers(0 params)`, `instance Builder unsetBuildMetadata(0 params)`, `instance Version build(0 params)`, `static String[] concatArrays(2 params)`, `instance Deprecated SuppressWarnings(0 params)`, `instance Builder setNormalVersion(1 params)`
- Methods in `com.github.zafarkhaja.semver.Validators`: `static long nonNegative(2 params)`, `instance T nonNull(2 params)`, `instance T[] nonEmpty(2 params)`, `instance T[] zeroOrMoreNonNulls(2 params)`, `instance T[] oneOrMoreNonNulls(2 params)`, `instance T nonNullOrThrow(2 params)`
- Fields in `com.github.zafarkhaja.semver.SerializationProxy`: `long serialVersionUID`, `String version`
- Constructors in `com.github.zafarkhaja.semver.SerializationProxy`: `SerializationProxy(1 params)`
- Methods in `com.github.zafarkhaja.semver.SerializationProxy`: `instance Object readResolve(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/VersionParser.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Imports: `com.github.zafarkhaja.semver.util.Stream`, `com.github.zafarkhaja.semver.util.UnexpectedElementException`, `java.util.ArrayList`, `java.util.EnumSet`, `java.util.List`, `staticcom.github.zafarkhaja.semver.VersionParser.CharType.*`
- Classes: `com.github.zafarkhaja.semver.VersionParser`
- Methods in `com.github.zafarkhaja.semver.VersionParser`: `instance boolean isMatchedBy(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/And.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.And`
- Fields in `com.github.zafarkhaja.semver.expr.And`: `Expression left`, `Expression right`
- Constructors in `com.github.zafarkhaja.semver.expr.And`: `And(2 params)`
- Methods in `com.github.zafarkhaja.semver.expr.And`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.ParseException`, `com.github.zafarkhaja.semver.UnexpectedCharacterException`, `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.CompositeExpression`, `com.github.zafarkhaja.semver.expr.Helper`
- Methods in `com.github.zafarkhaja.semver.expr.Helper`: `static CompositeExpression not(1 params)`, `static CompositeExpression eq(1 params)`, `static CompositeExpression eq(1 params)`, `static CompositeExpression neq(1 params)`, `static CompositeExpression neq(1 params)`, `static CompositeExpression gt(1 params)`, `static CompositeExpression gt(1 params)`, `static CompositeExpression gte(1 params)`, `static CompositeExpression gte(1 params)`, `static CompositeExpression lt(1 params)`, `static CompositeExpression lt(1 params)`, `static CompositeExpression lte(1 params)`, `static CompositeExpression lte(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/Equal.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.Equal`
- Fields in `com.github.zafarkhaja.semver.expr.Equal`: `Version parsedVersion`
- Constructors in `com.github.zafarkhaja.semver.expr.Equal`: `Equal(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.Equal`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/Expression.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`, `java.util.function.Predicate`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/ExpressionParser.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Parser`, `com.github.zafarkhaja.semver.Version`, `com.github.zafarkhaja.semver.expr.Lexer.Token`, `com.github.zafarkhaja.semver.util.Stream`, `com.github.zafarkhaja.semver.util.Stream.ElementType`, `com.github.zafarkhaja.semver.util.UnexpectedElementException`, `java.util.EnumSet`, `java.util.Iterator`, `staticcom.github.zafarkhaja.semver.expr.CompositeExpression.Helper.*`, `staticcom.github.zafarkhaja.semver.expr.Lexer.Token.Type.*`
- Classes: `com.github.zafarkhaja.semver.expr.ExpressionParser`
- Fields in `com.github.zafarkhaja.semver.expr.ExpressionParser`: `Lexer lexer`
- Constructors in `com.github.zafarkhaja.semver.expr.ExpressionParser`: `ExpressionParser(1 params)`, `newInstance(0 params)`
- Methods in `com.github.zafarkhaja.semver.expr.ExpressionParser`: `instance Expression parse(1 params)`, `instance CompositeExpression parseSemVerExpression(0 params)`, `instance CompositeExpression parseMoreExpressions(1 params)`, `instance CompositeExpression parseRange(0 params)`, `instance CompositeExpression parseComparisonRange(0 params)`, `instance CompositeExpression parseTildeRange(0 params)`, `instance CompositeExpression parseCaretRange(0 params)`, `instance boolean isWildcardRange(0 params)`, `instance CompositeExpression parseWildcardRange(0 params)`, `instance boolean isHyphenRange(0 params)`, `instance CompositeExpression parseHyphenRange(0 params)`, `instance boolean isPartialVersionRange(0 params)`, `instance CompositeExpression parsePartialVersionRange(0 params)`, `instance Version parseVersion(0 params)`, `instance boolean isVersionFollowedBy(1 params)`, `instance Token consumeNextToken(0 params)`, `instance long consumeNextNumeric(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/Greater.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.Greater`
- Fields in `com.github.zafarkhaja.semver.expr.Greater`: `Version parsedVersion`
- Constructors in `com.github.zafarkhaja.semver.expr.Greater`: `Greater(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.Greater`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/GreaterOrEqual.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.GreaterOrEqual`
- Fields in `com.github.zafarkhaja.semver.expr.GreaterOrEqual`: `Version parsedVersion`
- Constructors in `com.github.zafarkhaja.semver.expr.GreaterOrEqual`: `GreaterOrEqual(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.GreaterOrEqual`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/Less.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.Less`
- Fields in `com.github.zafarkhaja.semver.expr.Less`: `Version parsedVersion`
- Constructors in `com.github.zafarkhaja.semver.expr.Less`: `Less(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.Less`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/LessOrEqual.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.LessOrEqual`
- Fields in `com.github.zafarkhaja.semver.expr.LessOrEqual`: `Version parsedVersion`
- Constructors in `com.github.zafarkhaja.semver.expr.LessOrEqual`: `LessOrEqual(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.LessOrEqual`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.util.Stream`, `java.util.ArrayList`, `java.util.List`, `java.util.regex.Matcher`, `java.util.regex.Pattern`
- Classes: `com.github.zafarkhaja.semver.expr.Lexer`, `com.github.zafarkhaja.semver.expr.Token`
- Fields in `com.github.zafarkhaja.semver.expr.Token`: `Pattern pattern`
- Constructors in `com.github.zafarkhaja.semver.expr.Token`: `NUMERIC(0 params)`, `DOT(0 params)`, `HYPHEN(0 params)`, `EQUAL(0 params)`, `NOT_EQUAL(0 params)`, `GREATER(0 params)`, `GREATER_EQUAL(0 params)`, `LESS(0 params)`, `LESS_EQUAL(0 params)`, `TILDE(0 params)`, `WILDCARD(0 params)`, `CARET(0 params)`, `AND(0 params)`, `OR(0 params)`, `NOT(0 params)`, `LEFT_PAREN(0 params)`, `RIGHT_PAREN(0 params)`, `WHITESPACE(0 params)`, `EOI(0 params)`, `Type(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.Token`: `instance String toString(0 params)`, `instance boolean isMatchedBy(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/LexerException.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.ParseException`
- Classes: `com.github.zafarkhaja.semver.expr.LexerException`
- Fields in `com.github.zafarkhaja.semver.expr.LexerException`: `long serialVersionUID`, `String expr`
- Constructors in `com.github.zafarkhaja.semver.expr.LexerException`: `LexerException(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.LexerException`: `instance String toString(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/Not.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.Not`
- Fields in `com.github.zafarkhaja.semver.expr.Not`: `Expression expr`
- Constructors in `com.github.zafarkhaja.semver.expr.Not`: `Not(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.Not`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/NotEqual.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.NotEqual`
- Fields in `com.github.zafarkhaja.semver.expr.NotEqual`: `Version parsedVersion`
- Constructors in `com.github.zafarkhaja.semver.expr.NotEqual`: `NotEqual(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.NotEqual`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/Or.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.Or`
- Fields in `com.github.zafarkhaja.semver.expr.Or`: `Expression left`, `Expression right`
- Constructors in `com.github.zafarkhaja.semver.expr.Or`: `Or(2 params)`
- Methods in `com.github.zafarkhaja.semver.expr.Or`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/UnexpectedTokenException.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.ParseException`, `com.github.zafarkhaja.semver.expr.Lexer.Token`, `com.github.zafarkhaja.semver.util.UnexpectedElementException`, `java.util.Arrays`
- Classes: `com.github.zafarkhaja.semver.expr.UnexpectedTokenException`
- Fields in `com.github.zafarkhaja.semver.expr.UnexpectedTokenException`: `long serialVersionUID`, `Token unexpected`, `Token.Type[] expected`
- Constructors in `com.github.zafarkhaja.semver.expr.UnexpectedTokenException`: `UnexpectedTokenException(1 params)`, `UnexpectedTokenException(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.UnexpectedTokenException`: `instance Token getUnexpectedToken(0 params)`, `instance Token.Type[] getExpectedTokenTypes(0 params)`, `instance String toString(0 params)`, `static String createMessage(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/expr/package-info.java`

- Source root: `src/main/java`
- Risk: `medium`
- Package: `com.github.zafarkhaja.semver.expr`

### `src/main/java/com/github/zafarkhaja/semver/package-info.java`

- Source root: `src/main/java`
- Risk: `medium`
- Package: `com.github.zafarkhaja.semver`

### `src/main/java/com/github/zafarkhaja/semver/util/Stream.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.util`
- Imports: `java.util.Arrays`, `java.util.Iterator`, `java.util.NoSuchElementException`
- Classes: `com.github.zafarkhaja.semver.util.Stream`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/util/UnexpectedElementException.java`

- Source root: `src/main/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.util`
- Imports: `com.github.zafarkhaja.semver.util.Stream.ElementType`, `java.util.Arrays`
- Classes: `com.github.zafarkhaja.semver.util.UnexpectedElementException`
- Fields in `com.github.zafarkhaja.semver.util.UnexpectedElementException`: `long serialVersionUID`, `Object unexpected`, `int position`
- Constructors in `com.github.zafarkhaja.semver.util.UnexpectedElementException`: `UnexpectedElementException(2 params)`, `getExpectedElementTypes(0 params)`
- Methods in `com.github.zafarkhaja.semver.util.UnexpectedElementException`: `instance Object getUnexpectedElement(0 params)`, `instance int getPosition(0 params)`, `instance String toString(0 params)`, `static String createMessage(2 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/main/java/com/github/zafarkhaja/semver/util/package-info.java`

- Source root: `src/main/java`
- Risk: `medium`
- Package: `com.github.zafarkhaja.semver.util`

### `src/test/java/com/github/zafarkhaja/semver/ParserErrorHandlingTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Imports: `com.github.zafarkhaja.semver.VersionParser.CharType`, `java.util.Arrays`, `java.util.Collection`, `org.junit.jupiter.params.ParameterizedTest`, `org.junit.jupiter.params.provider.MethodSource`, `staticcom.github.zafarkhaja.semver.VersionParser.CharType.*`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.ParserErrorHandlingTest`
- Constructors in `com.github.zafarkhaja.semver.ParserErrorHandlingTest`: `parameters(0 params)`
- Methods in `com.github.zafarkhaja.semver.ParserErrorHandlingTest`: `instance ParameterizedTest MethodSource(0 params)`, `instance void shouldCorrectlyHandleParseErrors(4 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/VersionParserCharTypeTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Imports: `com.github.zafarkhaja.semver.VersionParser.CharType`, `org.junit.jupiter.api.Test`, `staticcom.github.zafarkhaja.semver.VersionParser.CharType.*`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.VersionParserCharTypeTest`
- Methods in `com.github.zafarkhaja.semver.VersionParserCharTypeTest`: `instance void shouldBeMatchedByDigit(0 params)`, `instance void shouldBeMatchedByLetter(0 params)`, `instance void shouldBeMatchedByDot(0 params)`, `instance void shouldBeMatchedByHyphen(0 params)`, `instance void shouldBeMatchedByPlus(0 params)`, `instance void shouldBeMatchedByEol(0 params)`, `instance void shouldBeMatchedByIllegal(0 params)`, `instance void shouldReturnCharTypeForCharacter(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/VersionParserTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Imports: `org.junit.jupiter.api.Test`, `org.junit.jupiter.params.ParameterizedTest`, `org.junit.jupiter.params.provider.ValueSource`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.VersionParserTest`, `com.github.zafarkhaja.semver.VersionParser`, `com.github.zafarkhaja.semver.VersionParser`, `com.github.zafarkhaja.semver.VersionParser`, `com.github.zafarkhaja.semver.VersionParser`, `com.github.zafarkhaja.semver.parser`, `com.github.zafarkhaja.semver.VersionParser`, `com.github.zafarkhaja.semver.Version`
- Methods in `com.github.zafarkhaja.semver.VersionParserTest`: `instance void shouldParseNormalVersion(0 params)`, `instance void shouldRaiseErrorIfNumericIdentifierHasLeadingZeroes(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/VersionTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Imports: `java.io.ByteArrayInputStream`, `java.io.ByteArrayOutputStream`, `java.io.ObjectInputStream`, `java.io.ObjectOutputStream`, `java.util.Locale`, `java.util.function.Predicate`, `org.junit.jupiter.api.Nested`, `org.junit.jupiter.api.Test`, `org.junit.jupiter.api.function.Executable`, `org.junit.jupiter.params.ParameterizedTest`, `org.junit.jupiter.params.provider.ValueSource`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.VersionTest`, `com.github.zafarkhaja.semver.Builder`, `com.github.zafarkhaja.semver.CoreFunctionality`, `com.github.zafarkhaja.semver.Version`, `com.github.zafarkhaja.semver.Version`, `com.github.zafarkhaja.semver.Version`, `com.github.zafarkhaja.semver.Version`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.v`, `com.github.zafarkhaja.semver.CompareToMethod`, `com.github.zafarkhaja.semver.IncrementOrderComparator`, `com.github.zafarkhaja.semver.PrecedenceOrderComparator`, `com.github.zafarkhaja.semver.EqualsMethod`, `com.github.zafarkhaja.semver.HashCodeMethod`, `com.github.zafarkhaja.semver.ToStringMethod`, `com.github.zafarkhaja.semver.Serialization`, `com.github.zafarkhaja.semver.exec`
- Methods in `com.github.zafarkhaja.semver.Builder`: `instance void shouldSetVersionCore(0 params)`, `instance void shouldNotAcceptNegativeNumbersForVersionCore(0 params)`, `instance void shouldSetVersionCoreWithDefaultMinorAndPatchValues(0 params)`, `instance void shouldNotAcceptNegativeNumbersForVersionCoreWithDefaultMinorAndPatchValues(0 params)`, `instance void shouldSetVersionCoreWithDefaultPatchValue(0 params)`, `instance void shouldNotAcceptNegativeNumbersForVersionCoreWithDefaultPatchValue(0 params)`, `instance void shouldSetMajorVersion(0 params)`, `instance void shouldNotAcceptNegativeNumbersForMajorVersion(0 params)`, `instance void shouldSetMinorVersion(0 params)`, `instance void shouldNotAcceptNegativeNumbersForMinorVersion(0 params)`, `instance void shouldSetPatchVersion(0 params)`, `instance void shouldNotAcceptNegativeNumbersForPatchVersion(0 params)`, `instance void shouldSetPreReleaseVersion(0 params)`, `instance void shouldMakeDefensiveCopyOfArgumentsWhenSettingPreReleaseVersion(0 params)`, `instance void shouldNotAcceptNullsForPreReleaseVersion(0 params)`, `instance void shouldNotAcceptEmptyArraysForPreReleaseVersion(0 params)`, `instance void shouldAddPreReleaseIdentifiers(0 params)`, `instance void shouldMakeDefensiveCopyOfArgumentsWhenAddingPreReleaseIdentifiers(0 params)`, `instance void shouldNotAcceptNullsWhenAddingPreReleaseIdentifiers(0 params)`, `instance void shouldNotAcceptEmptyArraysWhenAddingPreReleaseIdentifiers(0 params)`, `instance void shouldUnsetPreReleaseVersion(0 params)`, `instance void shouldSetBuildMetadata(0 params)`, `instance void shouldMakeDefensiveCopyOfArgumentsWhenSettingBuildMetadata(0 params)`, `instance void shouldNotAcceptNullsForBuildMetadata(0 params)`, `instance void shouldNotAcceptEmptyArraysForBuildMetadata(0 params)`, `instance void shouldAddBuildIdentifiers(0 params)`, `instance void shouldMakeDefensiveCopyOfArgumentsWhenAddingBuildIdentifiers(0 params)`, `instance void shouldNotAcceptNullsWhenAddingBuildIdentifiers(0 params)`, `instance void shouldNotAcceptEmptyArraysWhenAddingBuildIdentifiers(0 params)`, `instance void shouldUnsetBuildVersion(0 params)`, `instance Test SuppressWarnings(0 params)`, `instance void shouldSetNormalVersion(0 params)`, `instance Test SuppressWarnings(0 params)`, `instance void shouldNotAcceptNullsWhenSettingNormalVersion(0 params)`
- Methods in `com.github.zafarkhaja.semver.CoreFunctionality`: `instance void shouldNormallyTakeTheFormXDotYDotZWhereXYZAreNonNegativeIntegers(0 params)`, `instance void mayHavePreReleaseVersionFollowingPatchVersionPrependedWithHyphen(0 params)`, `instance void mayHaveBuildMetadataFollowingPatchOrPreReleaseVersionPrependedWithPlus(0 params)`, `instance void shouldParseFullSemVerCompliantVersionStringsInStrictMode(0 params)`, `instance void shouldCheckInputStringForNullBeforeParsing(0 params)`
- Methods in `com.github.zafarkhaja.semver.CompareToMethod`: `instance void shouldCompareMajorVersionNumerically(0 params)`, `instance void shouldCompareMinorVersionNumerically(0 params)`, `instance void shouldComparePatchVersionNumerically(0 params)`, `instance void shouldCompareAlphanumericIdentifiersLexicallyInAsciiOrder(0 params)`, `instance void shouldGiveHigherPrecedenceToNonNumericIdentifierThanNumeric(0 params)`, `instance void shouldCompareNumericIdentifiersNumerically(0 params)`, `instance void shouldGiveHigherPrecedenceToVersionWithLargerSetOfIdentifiers(0 params)`, `instance void shouldGiveHigherPrecedenceToStableVersionThanPreReleaseVersion(0 params)`, `instance void shouldGiveHigherPrecedenceToVersionWithBuildMetadata(0 params)`, `instance void shouldBeConsistentWithEquals(0 params)`, `instance void shouldCorrectlyCompareVersionsWithBuildMetadata(0 params)`
- Methods in `com.github.zafarkhaja.semver.IncrementOrderComparator`: `instance void shouldSortInIncrementOrder(0 params)`, `instance void shouldIgnoreBuildMetadata(0 params)`
- Methods in `com.github.zafarkhaja.semver.PrecedenceOrderComparator`: `instance void shouldSortInPrecedenceOrder(0 params)`, `instance void shouldIgnoreBuildMetadata(0 params)`
- Methods in `com.github.zafarkhaja.semver.EqualsMethod`: `instance void shouldBeReflexive(0 params)`, `instance void shouldBeSymmetric(0 params)`, `instance void shouldBeTransitive(0 params)`, `instance void shouldBeConsistent(0 params)`, `instance void shouldReturnFalseIfOtherVersionIsOfDifferentType(0 params)`, `instance void shouldReturnFalseIfOtherVersionIsNull(0 params)`, `instance void shouldNotIgnoreBuildMetadataWhenCheckingForExactEquality(0 params)`
- Methods in `com.github.zafarkhaja.semver.HashCodeMethod`: `instance void shouldReturnSameHashCodeIfVersionsAreEqual(0 params)`, `instance void shouldReturnDifferentHashCodesIfVersionsAreNotEqual(0 params)`
- Methods in `com.github.zafarkhaja.semver.ToStringMethod`: `instance void shouldReturnStringRepresentation(0 params)`, `instance void shouldUseRootLocale(0 params)`
- Methods in `com.github.zafarkhaja.semver.Serialization`: `instance void shouldBeSerializable(0 params)`, `instance byte[] serialize(1 params)`, `instance Version deserialize(1 params)`, `static void assertThrowsIllegalArgumentException(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/AndTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `org.junit.jupiter.api.Test`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.AndTest`
- Methods in `com.github.zafarkhaja.semver.expr.AndTest`: `instance void shouldCheckIfBothExpressionsEvaluateToTrue(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/CompositeExpressionTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `org.junit.jupiter.api.Test`, `staticcom.github.zafarkhaja.semver.expr.CompositeExpression.Helper.*`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.CompositeExpressionTest`
- Methods in `com.github.zafarkhaja.semver.expr.CompositeExpressionTest`: `instance void shouldSupportEqualExpression(0 params)`, `instance void shouldSupportNotEqualExpression(0 params)`, `instance void shouldSupportGreaterExpression(0 params)`, `instance void shouldSupportGreaterOrEqualExpression(0 params)`, `instance void shouldSupportLessExpression(0 params)`, `instance void shouldSupportLessOrEqualExpression(0 params)`, `instance void shouldSupportNotExpression(0 params)`, `instance void shouldSupportAndExpression(0 params)`, `instance void shouldSupportOrExpression(0 params)`, `instance void shouldSupportComplexExpressions(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/EqualTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`, `org.junit.jupiter.api.Test`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.EqualTest`
- Methods in `com.github.zafarkhaja.semver.expr.EqualTest`: `instance void shouldCheckIfVersionIsEqualToParsedVersion(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/ExpressionParserTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`, `java.util.Arrays`, `org.junit.jupiter.api.Test`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.ExpressionParserTest`, `com.github.zafarkhaja.semver.expr.parser`, `com.github.zafarkhaja.semver.expr.parser`, `com.github.zafarkhaja.semver.expr.parser`
- Constructors in `com.github.zafarkhaja.semver.expr.ExpressionParserTest`: `parse(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.ExpressionParserTest`: `instance void shouldParseEqualComparisonRange(0 params)`, `instance void shouldParseEqualComparisonRangeIfOnlyFullVersionGiven(0 params)`, `instance void shouldParseNotEqualComparisonRange(0 params)`, `instance void shouldParseGreaterComparisonRange(0 params)`, `instance void shouldParseGreaterOrEqualComparisonRange(0 params)`, `instance void shouldParseLessComparisonRange(0 params)`, `instance void shouldParseLessOrEqualComparisonRange(0 params)`, `instance void shouldSupportLongNumericIdentifiersInComparisonRanges(0 params)`, `instance void shouldParseTildeRange(0 params)`, `instance void shouldSupportLongNumericIdentifiersInTildeRanges(0 params)`, `instance void shouldRaiseErrorIfIncrementCausesOverflowInTildeRanges(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/GreaterOrEqualTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`, `org.junit.jupiter.api.Test`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.GreaterOrEqualTest`
- Methods in `com.github.zafarkhaja.semver.expr.GreaterOrEqualTest`: `instance void shouldCheckIfVersionIsGreaterThanOrEqualToParsedVersion(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/GreaterTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`, `org.junit.jupiter.api.Test`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.GreaterTest`
- Methods in `com.github.zafarkhaja.semver.expr.GreaterTest`: `instance void shouldCheckIfVersionIsGreaterThanParsedVersion(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/LessOrEqualTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`, `org.junit.jupiter.api.Test`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.LessOrEqualTest`
- Methods in `com.github.zafarkhaja.semver.expr.LessOrEqualTest`: `instance void shouldCheckIfVersionIsLessThanOrEqualToParsedVersion(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/LessTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`, `org.junit.jupiter.api.Test`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.LessTest`
- Methods in `com.github.zafarkhaja.semver.expr.LessTest`: `instance void shouldCheckIfVersionIsLessThanParsedVersion(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/LexerTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.expr.Lexer.*`, `com.github.zafarkhaja.semver.util.Stream`, `org.junit.jupiter.api.Test`, `staticcom.github.zafarkhaja.semver.expr.Lexer.Token.Type.*`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.LexerTest`, `com.github.zafarkhaja.semver.expr.lexer`
- Fields in `com.github.zafarkhaja.semver.expr.LexerTest`: `Lexer lexer`, `Lexer lexer`, `Lexer lexer`
- Constructors in `com.github.zafarkhaja.semver.expr.LexerTest`: `tokenize(0 params)`, `tokenize(0 params)`, `tokenize(0 params)`
- Methods in `com.github.zafarkhaja.semver.expr.LexerTest`: `instance void shouldTokenizeVersionString(0 params)`, `instance void shouldSkipWhitespaces(0 params)`, `instance void shouldEndWithEol(0 params)`, `instance void shouldRaiseErrorOnIllegalCharacter(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/LexerTokenTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.expr.Lexer.Token`, `org.junit.jupiter.api.Nested`, `org.junit.jupiter.api.Test`, `staticcom.github.zafarkhaja.semver.expr.Lexer.Token.Type.*`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.LexerTokenTest`, `com.github.zafarkhaja.semver.expr.EqualsMethod`, `com.github.zafarkhaja.semver.expr.HashCodeMethod`
- Methods in `com.github.zafarkhaja.semver.expr.EqualsMethod`: `instance void shouldBeReflexive(0 params)`, `instance void shouldBeSymmetric(0 params)`, `instance void shouldBeTransitive(0 params)`, `instance void shouldBeConsistent(0 params)`, `instance void shouldReturnFalseIfOtherVersionIsOfDifferentType(0 params)`, `instance void shouldReturnFalseIfOtherVersionIsNull(0 params)`, `instance void shouldReturnFalseIfTypesAreDifferent(0 params)`, `instance void shouldReturnFalseIfLexemesAreDifferent(0 params)`, `instance void shouldReturnFalseIfPositionsAreDifferent(0 params)`
- Methods in `com.github.zafarkhaja.semver.expr.HashCodeMethod`: `instance void shouldReturnSameHashCodeIfTokensAreEqual(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/NotEqualTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`, `org.junit.jupiter.api.Test`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.NotEqualTest`
- Methods in `com.github.zafarkhaja.semver.expr.NotEqualTest`: `instance void shouldCheckIfVersionIsNotEqualToParsedVersion(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/NotTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `org.junit.jupiter.api.Test`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.NotTest`
- Methods in `com.github.zafarkhaja.semver.expr.NotTest`: `instance void shouldRevertBooleanResultOfExpression(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/OrTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `org.junit.jupiter.api.Test`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.OrTest`
- Methods in `com.github.zafarkhaja.semver.expr.OrTest`: `instance void shouldCheckIfOneOfTwoExpressionsEvaluateToTrue(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/expr/ParserErrorHandlingTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.expr.Lexer.Token`, `java.util.Arrays`, `java.util.Collection`, `org.junit.jupiter.params.ParameterizedTest`, `org.junit.jupiter.params.provider.MethodSource`, `staticcom.github.zafarkhaja.semver.expr.Lexer.Token.Type.*`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.expr.ParserErrorHandlingTest`, `com.github.zafarkhaja.semver.expr.ExpressionParser`
- Methods in `com.github.zafarkhaja.semver.expr.ParserErrorHandlingTest`: `instance ParameterizedTest MethodSource(0 params)`, `instance void shouldCorrectlyHandleParseErrors(3 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `src/test/java/com/github/zafarkhaja/semver/util/StreamTest.java`

- Source root: `src/test/java`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.util`
- Imports: `java.util.Iterator`, `org.junit.jupiter.api.Test`, `staticorg.junit.jupiter.api.Assertions.*`
- Classes: `com.github.zafarkhaja.semver.util.StreamTest`, `com.github.zafarkhaja.semver.util.stream`
- Methods in `com.github.zafarkhaja.semver.util.StreamTest`: `instance void shouldBeBackedByArray(0 params)`, `instance void shouldImplementIterable(0 params)`, `instance void shouldNotReturnRealElementsArray(0 params)`, `instance void shouldReturnArrayOfElementsThatAreLeftInStream(0 params)`, `instance void shouldConsumeElementsOneByOne(0 params)`, `instance void shouldRaiseErrorWhenUnexpectedElementConsumed(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

## Recommended Migration Order

- `src/main/java/com/github/zafarkhaja/semver/expr/package-info.java` (medium risk)
- `src/main/java/com/github/zafarkhaja/semver/package-info.java` (medium risk)
- `src/main/java/com/github/zafarkhaja/semver/util/package-info.java` (medium risk)
- `src/main/java/com/github/zafarkhaja/semver/ParseException.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/Parser.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/UnexpectedCharacterException.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/Version.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/VersionParser.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/And.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/Equal.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/Expression.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/ExpressionParser.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/Greater.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/GreaterOrEqual.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/Less.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/LessOrEqual.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/LexerException.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/Not.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/NotEqual.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/Or.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/expr/UnexpectedTokenException.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/util/Stream.java` (high risk)
- `src/main/java/com/github/zafarkhaja/semver/util/UnexpectedElementException.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/ParserErrorHandlingTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/VersionParserCharTypeTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/VersionParserTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/VersionTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/AndTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/CompositeExpressionTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/EqualTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/ExpressionParserTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/GreaterOrEqualTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/GreaterTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/LessOrEqualTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/LessTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/LexerTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/LexerTokenTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/NotEqualTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/NotTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/OrTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/expr/ParserErrorHandlingTest.java` (high risk)
- `src/test/java/com/github/zafarkhaja/semver/util/StreamTest.java` (high risk)
