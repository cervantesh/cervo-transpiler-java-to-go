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
