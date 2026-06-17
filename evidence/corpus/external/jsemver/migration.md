# Java-To-Go Migration Report

Project: `.corpus/jsemver/src/main/java`

## Summary

| Metric | Count |
| --- | ---: |
| Java files | 25 |
| Packages | 3 |
| Classes | 25 |
| Methods | 80 |
| Constructors | 43 |
| Fields | 32 |
| Unsupported features | 0 |
| Diagnostics | 741 |
| Internal imports | 44 |

## Build Files

No Maven or Gradle build file detected.

## Packages

- `com.github.zafarkhaja.semver`
- `com.github.zafarkhaja.semver.expr`
- `com.github.zafarkhaja.semver.util`

## Feature Counts

| Feature | Count |
| --- | ---: |
| class declarations | 25 |
| class fields | 32 |
| constructors | 43 |
| import declarations | 61 |
| instance methods | 62 |
| package declarations | 25 |
| static methods | 18 |

## Unsupported Feature Counts

None.

## Files

### `com/github/zafarkhaja/semver/ParseException.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Classes: `com.github.zafarkhaja.semver.ParseException`
- Fields in `com.github.zafarkhaja.semver.ParseException`: `long serialVersionUID`
- Constructors in `com.github.zafarkhaja.semver.ParseException`: `ParseException(0 params)`, `ParseException(1 params)`, `ParseException(2 params)`
- Methods in `com.github.zafarkhaja.semver.ParseException`: `instance String toString(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/Parser.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Diagnostics: `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/UnexpectedCharacterException.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Imports: `com.github.zafarkhaja.semver.VersionParser.CharType`, `com.github.zafarkhaja.semver.util.UnexpectedElementException`, `java.util.Arrays`
- Classes: `com.github.zafarkhaja.semver.UnexpectedCharacterException`
- Fields in `com.github.zafarkhaja.semver.UnexpectedCharacterException`: `long serialVersionUID`, `Character unexpected`, `int position`, `CharType[] expected`
- Constructors in `com.github.zafarkhaja.semver.UnexpectedCharacterException`: `UnexpectedCharacterException(1 params)`, `UnexpectedCharacterException(2 params)`
- Methods in `com.github.zafarkhaja.semver.UnexpectedCharacterException`: `instance Character getUnexpectedCharacter(0 params)`, `instance int getPosition(0 params)`, `instance CharType[] getExpectedCharTypes(0 params)`, `instance String toString(0 params)`, `static String createMessage(2 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/Version.java`

- Source root: `.`
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

### `com/github/zafarkhaja/semver/VersionParser.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver`
- Imports: `com.github.zafarkhaja.semver.util.Stream`, `com.github.zafarkhaja.semver.util.UnexpectedElementException`, `java.util.ArrayList`, `java.util.EnumSet`, `java.util.List`, `staticcom.github.zafarkhaja.semver.VersionParser.CharType.*`
- Classes: `com.github.zafarkhaja.semver.VersionParser`
- Methods in `com.github.zafarkhaja.semver.VersionParser`: `instance boolean isMatchedBy(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/And.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.And`
- Fields in `com.github.zafarkhaja.semver.expr.And`: `Expression left`, `Expression right`
- Constructors in `com.github.zafarkhaja.semver.expr.And`: `And(2 params)`
- Methods in `com.github.zafarkhaja.semver.expr.And`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/CompositeExpression.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.ParseException`, `com.github.zafarkhaja.semver.UnexpectedCharacterException`, `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.CompositeExpression`, `com.github.zafarkhaja.semver.expr.Helper`
- Methods in `com.github.zafarkhaja.semver.expr.Helper`: `static CompositeExpression not(1 params)`, `static CompositeExpression eq(1 params)`, `static CompositeExpression eq(1 params)`, `static CompositeExpression neq(1 params)`, `static CompositeExpression neq(1 params)`, `static CompositeExpression gt(1 params)`, `static CompositeExpression gt(1 params)`, `static CompositeExpression gte(1 params)`, `static CompositeExpression gte(1 params)`, `static CompositeExpression lt(1 params)`, `static CompositeExpression lt(1 params)`, `static CompositeExpression lte(1 params)`, `static CompositeExpression lte(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/Equal.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.Equal`
- Fields in `com.github.zafarkhaja.semver.expr.Equal`: `Version parsedVersion`
- Constructors in `com.github.zafarkhaja.semver.expr.Equal`: `Equal(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.Equal`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/Expression.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`, `java.util.function.Predicate`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/ExpressionParser.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Parser`, `com.github.zafarkhaja.semver.Version`, `com.github.zafarkhaja.semver.expr.Lexer.Token`, `com.github.zafarkhaja.semver.util.Stream`, `com.github.zafarkhaja.semver.util.Stream.ElementType`, `com.github.zafarkhaja.semver.util.UnexpectedElementException`, `java.util.EnumSet`, `java.util.Iterator`, `staticcom.github.zafarkhaja.semver.expr.CompositeExpression.Helper.*`, `staticcom.github.zafarkhaja.semver.expr.Lexer.Token.Type.*`
- Classes: `com.github.zafarkhaja.semver.expr.ExpressionParser`
- Fields in `com.github.zafarkhaja.semver.expr.ExpressionParser`: `Lexer lexer`
- Constructors in `com.github.zafarkhaja.semver.expr.ExpressionParser`: `ExpressionParser(1 params)`, `newInstance(0 params)`
- Methods in `com.github.zafarkhaja.semver.expr.ExpressionParser`: `instance Expression parse(1 params)`, `instance CompositeExpression parseSemVerExpression(0 params)`, `instance CompositeExpression parseMoreExpressions(1 params)`, `instance CompositeExpression parseRange(0 params)`, `instance CompositeExpression parseComparisonRange(0 params)`, `instance CompositeExpression parseTildeRange(0 params)`, `instance CompositeExpression parseCaretRange(0 params)`, `instance boolean isWildcardRange(0 params)`, `instance CompositeExpression parseWildcardRange(0 params)`, `instance boolean isHyphenRange(0 params)`, `instance CompositeExpression parseHyphenRange(0 params)`, `instance boolean isPartialVersionRange(0 params)`, `instance CompositeExpression parsePartialVersionRange(0 params)`, `instance Version parseVersion(0 params)`, `instance boolean isVersionFollowedBy(1 params)`, `instance Token consumeNextToken(0 params)`, `instance long consumeNextNumeric(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/Greater.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.Greater`
- Fields in `com.github.zafarkhaja.semver.expr.Greater`: `Version parsedVersion`
- Constructors in `com.github.zafarkhaja.semver.expr.Greater`: `Greater(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.Greater`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/GreaterOrEqual.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.GreaterOrEqual`
- Fields in `com.github.zafarkhaja.semver.expr.GreaterOrEqual`: `Version parsedVersion`
- Constructors in `com.github.zafarkhaja.semver.expr.GreaterOrEqual`: `GreaterOrEqual(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.GreaterOrEqual`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/Less.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.Less`
- Fields in `com.github.zafarkhaja.semver.expr.Less`: `Version parsedVersion`
- Constructors in `com.github.zafarkhaja.semver.expr.Less`: `Less(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.Less`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/LessOrEqual.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.LessOrEqual`
- Fields in `com.github.zafarkhaja.semver.expr.LessOrEqual`: `Version parsedVersion`
- Constructors in `com.github.zafarkhaja.semver.expr.LessOrEqual`: `LessOrEqual(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.LessOrEqual`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/Lexer.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.util.Stream`, `java.util.ArrayList`, `java.util.List`, `java.util.regex.Matcher`, `java.util.regex.Pattern`
- Classes: `com.github.zafarkhaja.semver.expr.Lexer`, `com.github.zafarkhaja.semver.expr.Token`
- Constructors in `com.github.zafarkhaja.semver.expr.Token`: `NUMERIC(0 params)`, `DOT(0 params)`, `HYPHEN(0 params)`, `EQUAL(0 params)`, `NOT_EQUAL(0 params)`, `GREATER(0 params)`, `GREATER_EQUAL(0 params)`, `LESS(0 params)`, `LESS_EQUAL(0 params)`, `TILDE(0 params)`, `WILDCARD(0 params)`, `CARET(0 params)`, `AND(0 params)`, `OR(0 params)`, `NOT(0 params)`, `LEFT_PAREN(0 params)`, `RIGHT_PAREN(0 params)`, `WHITESPACE(0 params)`, `EOI(0 params)`
- Methods in `com.github.zafarkhaja.semver.expr.Token`: `instance String toString(0 params)`, `instance boolean isMatchedBy(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/LexerException.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.ParseException`
- Classes: `com.github.zafarkhaja.semver.expr.LexerException`
- Fields in `com.github.zafarkhaja.semver.expr.LexerException`: `long serialVersionUID`, `String expr`
- Constructors in `com.github.zafarkhaja.semver.expr.LexerException`: `LexerException(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.LexerException`: `instance String toString(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/Not.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.Not`
- Fields in `com.github.zafarkhaja.semver.expr.Not`: `Expression expr`
- Constructors in `com.github.zafarkhaja.semver.expr.Not`: `Not(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.Not`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/NotEqual.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.NotEqual`
- Fields in `com.github.zafarkhaja.semver.expr.NotEqual`: `Version parsedVersion`
- Constructors in `com.github.zafarkhaja.semver.expr.NotEqual`: `NotEqual(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.NotEqual`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/Or.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.Version`
- Classes: `com.github.zafarkhaja.semver.expr.Or`
- Fields in `com.github.zafarkhaja.semver.expr.Or`: `Expression left`, `Expression right`
- Constructors in `com.github.zafarkhaja.semver.expr.Or`: `Or(2 params)`
- Methods in `com.github.zafarkhaja.semver.expr.Or`: `instance boolean interpret(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/UnexpectedTokenException.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.expr`
- Imports: `com.github.zafarkhaja.semver.ParseException`, `com.github.zafarkhaja.semver.expr.Lexer.Token`, `com.github.zafarkhaja.semver.util.UnexpectedElementException`, `java.util.Arrays`
- Classes: `com.github.zafarkhaja.semver.expr.UnexpectedTokenException`
- Fields in `com.github.zafarkhaja.semver.expr.UnexpectedTokenException`: `long serialVersionUID`, `Token unexpected`, `Token.Type[] expected`
- Constructors in `com.github.zafarkhaja.semver.expr.UnexpectedTokenException`: `UnexpectedTokenException(1 params)`, `UnexpectedTokenException(1 params)`
- Methods in `com.github.zafarkhaja.semver.expr.UnexpectedTokenException`: `instance Token getUnexpectedToken(0 params)`, `instance Token.Type[] getExpectedTokenTypes(0 params)`, `instance String toString(0 params)`, `static String createMessage(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/expr/package-info.java`

- Source root: `.`
- Risk: `medium`
- Package: `com.github.zafarkhaja.semver.expr`

### `com/github/zafarkhaja/semver/package-info.java`

- Source root: `.`
- Risk: `medium`
- Package: `com.github.zafarkhaja.semver`

### `com/github/zafarkhaja/semver/util/Stream.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.util`
- Imports: `java.util.Arrays`, `java.util.Iterator`, `java.util.NoSuchElementException`
- Classes: `com.github.zafarkhaja.semver.util.Stream`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/util/UnexpectedElementException.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.zafarkhaja.semver.util`
- Imports: `com.github.zafarkhaja.semver.util.Stream.ElementType`, `java.util.Arrays`
- Classes: `com.github.zafarkhaja.semver.util.UnexpectedElementException`
- Fields in `com.github.zafarkhaja.semver.util.UnexpectedElementException`: `long serialVersionUID`, `Object unexpected`, `int position`
- Constructors in `com.github.zafarkhaja.semver.util.UnexpectedElementException`: `UnexpectedElementException(2 params)`, `getExpectedElementTypes(0 params)`
- Methods in `com.github.zafarkhaja.semver.util.UnexpectedElementException`: `instance Object getUnexpectedElement(0 params)`, `instance int getPosition(0 params)`, `instance String toString(0 params)`, `static String createMessage(2 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/zafarkhaja/semver/util/package-info.java`

- Source root: `.`
- Risk: `medium`
- Package: `com.github.zafarkhaja.semver.util`

## Recommended Migration Order

- `com/github/zafarkhaja/semver/expr/package-info.java` (medium risk)
- `com/github/zafarkhaja/semver/package-info.java` (medium risk)
- `com/github/zafarkhaja/semver/util/package-info.java` (medium risk)
- `com/github/zafarkhaja/semver/ParseException.java` (high risk)
- `com/github/zafarkhaja/semver/Parser.java` (high risk)
- `com/github/zafarkhaja/semver/UnexpectedCharacterException.java` (high risk)
- `com/github/zafarkhaja/semver/Version.java` (high risk)
- `com/github/zafarkhaja/semver/VersionParser.java` (high risk)
- `com/github/zafarkhaja/semver/expr/And.java` (high risk)
- `com/github/zafarkhaja/semver/expr/CompositeExpression.java` (high risk)
- `com/github/zafarkhaja/semver/expr/Equal.java` (high risk)
- `com/github/zafarkhaja/semver/expr/Expression.java` (high risk)
- `com/github/zafarkhaja/semver/expr/ExpressionParser.java` (high risk)
- `com/github/zafarkhaja/semver/expr/Greater.java` (high risk)
- `com/github/zafarkhaja/semver/expr/GreaterOrEqual.java` (high risk)
- `com/github/zafarkhaja/semver/expr/Less.java` (high risk)
- `com/github/zafarkhaja/semver/expr/LessOrEqual.java` (high risk)
- `com/github/zafarkhaja/semver/expr/Lexer.java` (high risk)
- `com/github/zafarkhaja/semver/expr/LexerException.java` (high risk)
- `com/github/zafarkhaja/semver/expr/Not.java` (high risk)
- `com/github/zafarkhaja/semver/expr/NotEqual.java` (high risk)
- `com/github/zafarkhaja/semver/expr/Or.java` (high risk)
- `com/github/zafarkhaja/semver/expr/UnexpectedTokenException.java` (high risk)
- `com/github/zafarkhaja/semver/util/Stream.java` (high risk)
- `com/github/zafarkhaja/semver/util/UnexpectedElementException.java` (high risk)

## Migration Execution Summary

| Metric | Count |
| --- | ---: |
| Java files | 25 |
| Generated files | 3 |
| Skipped files | 22 |
| Diagnostics | 296 |
| Dry run | false |

### Migration Diagnostics

- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/ParseException.java:39:4: JTG2001: duplicate Go symbol after migration: NewParseException`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/ParseException.java:48:4: JTG2001: duplicate Go symbol after migration: NewParseException`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/ParseException.java:59:4: JTG2001: duplicate Go symbol after migration: NewParseException`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/Parser.java:34:23: mismatched input '<' expecting '{'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/Parser.java:43:0: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/UnexpectedCharacterException.java:0:0: JTG2005: import unresolved: com.github.zafarkhaja.semver.VersionParser.CharType`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/UnexpectedCharacterException.java:63:4: JTG2001: duplicate Go symbol after migration: NewUnexpectedCharacterException`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/UnexpectedCharacterException.java:79:4: JTG2001: duplicate Go symbol after migration: NewUnexpectedCharacterException`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/Version.java:92:8: JTG2004: overload not supported: method com.github.zafarkhaja.semver.Builder.setVersionCore`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/Version.java:106:8: JTG2004: overload not supported: method com.github.zafarkhaja.semver.Builder.setVersionCore`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/Version.java:106:8: JTG2004: overload not supported: method com.github.zafarkhaja.semver.Builder.setVersionCore`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/Version.java:0:0: JTG2005: import unresolved: com.github.zafarkhaja.semver.expr.Expression`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/Version.java:69:8: JTG2001: duplicate Go symbol after migration: NewBuilder`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/Version.java:284:8: JTG2001: duplicate Go symbol after migration: NewBuilder`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:39:20: missing '{' at 'implements'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:50:12: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:55:30: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:55:32: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:55:44: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:55:46: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:57:9: extraneous input ',' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:62:12: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:66:16: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:67:31: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:67:33: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:67:45: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:67:47: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:68:31: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:68:33: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:68:45: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:68:47: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:69:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:75:12: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:79:16: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:80:30: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:80:32: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:81:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:87:12: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:91:16: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:92:30: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:92:32: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:93:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:99:12: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:103:16: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:104:30: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:104:32: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:105:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:111:12: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:114:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:120:12: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:123:16: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:123:35: token recognition error at: ':'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:126:20: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:129:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:138:15: no viable alternative at input 'staticCharType'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:139:31: token recognition error at: ':'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:142:16: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:145:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:156:18: no viable alternative at input 'privatefinalStream'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:165:4: extraneous input 'VersionParser' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:167:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:178:8: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:180:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:182:8: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:182:24: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:182:44: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:184:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:186:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:196:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:199:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:212:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:226:8: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:227:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:240:8: extraneous input 'long' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:241:8: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:242:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:255:8: extraneous input 'String' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:256:8: extraneous input 'parser' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:257:8: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:258:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:271:8: extraneous input 'String' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:272:8: extraneous input 'parser' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:273:8: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:274:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:292:8: extraneous input 'String' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:293:8: extraneous input 'String' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:295:8: extraneous input 'Character' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:296:8: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:298:12: extraneous input 'next' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:299:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:302:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:304:8: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:305:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:321:8: extraneous input 'long' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:322:8: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:324:12: extraneous input 'minor' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:325:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:328:8: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:330:12: extraneous input 'patch' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:331:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:334:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:352:8: extraneous input 'List' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:353:8: extraneous input 'do' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:355:12: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:357:16: extraneous input 'continue' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:358:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:360:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:361:8: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:362:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:378:8: extraneous input 'CharType' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:379:8: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:381:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:383:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:402:8: extraneous input 'List' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:403:8: extraneous input 'do' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:405:12: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:407:16: extraneous input 'continue' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:408:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:410:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:411:8: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:412:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:428:8: extraneous input 'CharType' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:429:8: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:431:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:433:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:451:8: extraneous input 'try' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:453:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:455:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:474:8: extraneous input 'do' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:476:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:477:8: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:478:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:494:8: extraneous input 'do' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:496:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:497:8: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:498:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:507:27: token recognition error at: ':'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:508:31: token recognition error at: ':'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:511:16: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:515:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:524:8: extraneous input 'Character' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:525:8: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:525:34: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:525:36: token recognition error at: '''`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:529:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:540:8: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:549:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:562:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:564:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/VersionParser.java:580:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/And.java:34:10: missing '{' at 'implements'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/And.java:54:18: no viable alternative at input 'this.left='`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/And.java:55:19: no viable alternative at input 'this.right='`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/And.java:65:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/And.java:66:4: no viable alternative at input 'Overridepublic'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:76:8: JTG2004: overload not supported: method com.github.zafarkhaja.semver.expr.Helper.eq`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:101:8: JTG2004: overload not supported: method com.github.zafarkhaja.semver.expr.Helper.neq`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:126:8: JTG2004: overload not supported: method com.github.zafarkhaja.semver.expr.Helper.gt`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:151:8: JTG2004: overload not supported: method com.github.zafarkhaja.semver.expr.Helper.gte`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:176:8: JTG2004: overload not supported: method com.github.zafarkhaja.semver.expr.Helper.lt`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:201:8: JTG2004: overload not supported: method com.github.zafarkhaja.semver.expr.Helper.lte`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:62:8: JTG2001: duplicate Go symbol after migration: Eq`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:76:8: JTG2001: duplicate Go symbol after migration: Eq`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:137:8: JTG2001: duplicate Go symbol after migration: Gte`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:151:8: JTG2001: duplicate Go symbol after migration: Gte`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:112:8: JTG2001: duplicate Go symbol after migration: Gt`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:126:8: JTG2001: duplicate Go symbol after migration: Gt`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:87:8: JTG2001: duplicate Go symbol after migration: Neq`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:101:8: JTG2001: duplicate Go symbol after migration: Neq`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:51:8: JTG2001: duplicate Go symbol after migration: Not`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:162:8: JTG2001: duplicate Go symbol after migration: Lt`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:176:8: JTG2001: duplicate Go symbol after migration: Lt`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:187:8: JTG2001: duplicate Go symbol after migration: Lte`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/CompositeExpression.java:201:8: JTG2001: duplicate Go symbol after migration: Lte`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Equal.java:34:12: missing '{' at 'implements'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Equal.java:47:27: no viable alternative at input 'this.parsedVersion='`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Equal.java:58:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Equal.java:59:4: no viable alternative at input 'Overridepublic'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Expression.java:37:28: missing '{' at 'extends'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Expression.java:37:45: mismatched input '<' expecting '('`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Expression.java:37:53: mismatched input '>' expecting Identifier`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Expression.java:47:12: extraneous input 'boolean' expecting Identifier`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Expression.java:47:42: mismatched input '{' expecting ';'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Expression.java:48:24: extraneous input '(' expecting Identifier`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Expression.java:48:32: missing '(' at ')'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Expression.java:50:0: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/ExpressionParser.java:0:0: JTG2005: import unresolved: com.github.zafarkhaja.semver.Parser`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/ExpressionParser.java:0:0: JTG2005: import unresolved: com.github.zafarkhaja.semver.expr.Lexer.Token`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/ExpressionParser.java:0:0: JTG2005: import unresolved: com.github.zafarkhaja.semver.util.Stream.ElementType`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/ExpressionParser.java:61:4: JTG2001: duplicate Go symbol after migration: NewExpressionParser`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/ExpressionParser.java:72:37: JTG2001: duplicate Go symbol after migration: NewExpressionParser`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Greater.java:34:14: missing '{' at 'implements'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Greater.java:48:27: no viable alternative at input 'this.parsedVersion='`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Greater.java:59:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Greater.java:60:4: no viable alternative at input 'Overridepublic'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/GreaterOrEqual.java:34:21: missing '{' at 'implements'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/GreaterOrEqual.java:48:27: no viable alternative at input 'this.parsedVersion='`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/GreaterOrEqual.java:60:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/GreaterOrEqual.java:61:4: no viable alternative at input 'Overridepublic'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Less.java:34:11: missing '{' at 'implements'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Less.java:48:27: no viable alternative at input 'this.parsedVersion='`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Less.java:59:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Less.java:60:4: no viable alternative at input 'Overridepublic'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/LessOrEqual.java:34:18: missing '{' at 'implements'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/LessOrEqual.java:48:27: no viable alternative at input 'this.parsedVersion='`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/LessOrEqual.java:60:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/LessOrEqual.java:61:4: no viable alternative at input 'Overridepublic'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:50:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:51:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:52:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:53:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:54:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:55:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:56:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:57:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:58:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:59:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:60:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:61:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:62:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:63:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:64:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:65:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:66:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:67:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Lexer.java:68:12: JTG2001: duplicate Go symbol after migration: NewToken`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/LexerException.java:35:28: missing '{' at 'extends'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/LexerException.java:37:50: extraneous input 'L' expecting ';'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/LexerException.java:53:18: no viable alternative at input 'this.expr='`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/LexerException.java:61:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/LexerException.java:62:4: no viable alternative at input 'Overridepublic'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Not.java:34:0: JTG2001: duplicate Go symbol after migration: Not`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/NotEqual.java:34:15: missing '{' at 'implements'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/NotEqual.java:47:27: no viable alternative at input 'this.parsedVersion='`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/NotEqual.java:58:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/NotEqual.java:59:4: no viable alternative at input 'Overridepublic'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Or.java:34:9: missing '{' at 'implements'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Or.java:54:18: no viable alternative at input 'this.left='`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Or.java:55:19: no viable alternative at input 'this.right='`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Or.java:65:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/Or.java:66:4: no viable alternative at input 'Overridepublic'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/UnexpectedTokenException.java:0:0: JTG2005: import unresolved: com.github.zafarkhaja.semver.expr.Lexer.Token`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/UnexpectedTokenException.java:57:4: JTG2001: duplicate Go symbol after migration: NewUnexpectedTokenException`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/expr/UnexpectedTokenException.java:71:4: JTG2001: duplicate Go symbol after migration: NewUnexpectedTokenException`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:38:19: mismatched input '<' expecting '{'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:46:32: mismatched input '<' expecting '{'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:56:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:68:12: no viable alternative at input 'privateint'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:77:11: no viable alternative at input 'publicStream'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:79:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:90:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:92:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:105:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:108:8: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:108:33: token recognition error at: ':'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:111:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:114:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:122:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:132:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:144:8: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:146:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:148:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:157:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:169:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:171:33: token recognition error at: ':'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:174:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:177:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:191:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:197:8: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:197:24: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:197:46: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:199:12: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:201:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:202:37: token recognition error at: ':'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:205:16: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:209:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:223:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:228:24: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:228:36: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:229:37: token recognition error at: ':'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:232:16: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:236:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:243:4: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:260:12: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:260:13: extraneous input 'Override' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:263:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:268:12: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:272:16: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:274:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:279:12: token recognition error at: '@'`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:282:12: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:284:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/Stream.java:296:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/UnexpectedElementException.java:0:0: JTG2005: import unresolved: com.github.zafarkhaja.semver.util.Stream.ElementType`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/UnexpectedElementException.java:63:4: JTG2001: duplicate Go symbol after migration: NewUnexpectedElementException`
- `.corpus/jsemver/src/main/java/com/github/zafarkhaja/semver/util/UnexpectedElementException.java:94:28: JTG2001: duplicate Go symbol after migration: NewUnexpectedElementException`
