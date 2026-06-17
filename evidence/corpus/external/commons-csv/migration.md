# Java-To-Go Migration Report

Project: `.corpus/commons-csv/src/main/java`

## Summary

| Metric | Count |
| --- | ---: |
| Java files | 12 |
| Packages | 1 |
| Classes | 13 |
| Methods | 125 |
| Constructors | 28 |
| Fields | 101 |
| Unsupported features | 0 |
| Diagnostics | 1300 |
| Internal imports | 0 |

## Build Files

No Maven or Gradle build file detected.

## Packages

- `org.apache.commons.csv`

## Feature Counts

| Feature | Count |
| --- | ---: |
| class declarations | 13 |
| class fields | 101 |
| constructors | 28 |
| import declarations | 103 |
| instance methods | 116 |
| package declarations | 12 |
| static methods | 9 |

## Unsupported Feature Counts

None.

## Files

### `org/apache/commons/csv/CSVException.java`

- Source root: `.`
- Risk: `high`
- Package: `org.apache.commons.csv`
- Imports: `java.io.IOException`, `java.util.Formatter`, `java.util.IllegalFormatException`
- Classes: `org.apache.commons.csv.CSVException`
- Fields in `org.apache.commons.csv.CSVException`: `long serialVersionUID`
- Constructors in `org.apache.commons.csv.CSVException`: `CSVException(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `org/apache/commons/csv/CSVFormat.java`

- Source root: `.`
- Risk: `high`
- Package: `org.apache.commons.csv`
- Imports: `staticorg.apache.commons.io.IOUtils.EOF`, `java.io.File`, `java.io.IOException`, `java.io.InputStream`, `java.io.OutputStream`, `java.io.Reader`, `java.io.Serializable`, `java.io.StringWriter`, `java.io.Writer`, `java.nio.charset.Charset`, `java.nio.file.Files`, `java.nio.file.Path`, `java.sql.ResultSet`, `java.sql.ResultSetMetaData`, `java.sql.SQLException`, `java.util.Arrays`, `java.util.HashSet`, `java.util.Objects`, `java.util.Set`, `java.util.function.Supplier`, `org.apache.commons.codec.binary.Base64OutputStream`, `org.apache.commons.io.IOUtils`, `org.apache.commons.io.function.IOStream`, `org.apache.commons.io.function.Uncheck`, `org.apache.commons.io.output.AppendableOutputStream`
- Classes: `org.apache.commons.csv.CSVFormat`, `org.apache.commons.csv.Builder`
- Fields in `org.apache.commons.csv.Builder`: `boolean allowMissingColumnNames`, `boolean autoFlush`, `Character commentMarker`, `String delimiter`, `DuplicateHeaderMode duplicateHeaderMode`, `Character escapeCharacter`, `String[] headerComments`, `String[] headers`, `boolean ignoreEmptyLines`, `boolean ignoreHeaderCase`, `boolean ignoreSurroundingSpaces`, `String nullString`, `Character quoteCharacter`, `String quotedNullString`, `QuoteMode quoteMode`, `String recordSeparator`, `boolean skipHeaderRecord`, `boolean lenientEof`, `boolean trailingData`, `boolean trailingDelimiter`, `boolean trim`, `long maxRows`
- Constructors in `org.apache.commons.csv.Builder`: `Builder(0 params)`, `Builder(1 params)`
- Methods in `org.apache.commons.csv.Builder`: `static Builder create(0 params)`, `static Builder create(1 params)`, `instance CSVFormat build(0 params)`, `instance CSVFormat get(0 params)`, `instance Builder setAllowDuplicateHeaderNames(1 params)`, `instance Builder setAllowMissingColumnNames(1 params)`, `instance Builder setAutoFlush(1 params)`, `instance Builder setCommentMarker(1 params)`, `instance Builder setCommentMarker(1 params)`, `instance Builder setDelimiter(1 params)`, `instance Builder setDelimiter(1 params)`, `instance Builder setDuplicateHeaderMode(1 params)`, `instance Builder setEscape(1 params)`, `instance Builder setEscape(1 params)`, `instance Builder setHeader(1 params)`, `instance Builder setHeader(1 params)`, `instance Builder setHeader(1 params)`, `instance Builder setHeader(0 params)`, `instance Builder setHeaderComments(0 params)`, `instance Builder setHeaderComments(0 params)`, `instance Builder setIgnoreEmptyLines(1 params)`, `instance Builder setIgnoreHeaderCase(1 params)`, `instance Builder setIgnoreSurroundingSpaces(1 params)`, `instance Builder setLenientEof(1 params)`, `instance Builder setMaxRows(1 params)`, `instance Builder setNullString(1 params)`, `instance Builder setQuote(1 params)`, `instance Builder setQuote(1 params)`, `instance Builder setQuoteMode(1 params)`, `instance Builder setRecordSeparator(1 params)`, `instance Builder setRecordSeparator(1 params)`, `instance Builder setSkipHeaderRecord(1 params)`, `instance Builder setTrailingData(1 params)`, `instance Builder setTrailingDelimiter(1 params)`, `instance Builder setTrim(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `org/apache/commons/csv/CSVParser.java`

- Source root: `.`
- Risk: `high`
- Package: `org.apache.commons.csv`
- Imports: `staticorg.apache.commons.csv.Token.Type.TOKEN`, `java.io.Closeable`, `java.io.File`, `java.io.IOException`, `java.io.InputStream`, `java.io.InputStreamReader`, `java.io.Reader`, `java.io.StringReader`, `java.io.UncheckedIOException`, `java.net.URL`, `java.nio.charset.Charset`, `java.nio.file.Files`, `java.nio.file.Path`, `java.util.ArrayList`, `java.util.Arrays`, `java.util.Collections`, `java.util.Iterator`, `java.util.LinkedHashMap`, `java.util.List`, `java.util.Map`, `java.util.NoSuchElementException`, `java.util.Objects`, `java.util.Spliterator`, `java.util.Spliterators`, `java.util.TreeMap`, `java.util.stream.Collectors`, `java.util.stream.Stream`, `java.util.stream.StreamSupport`, `org.apache.commons.io.Charsets`, `org.apache.commons.io.build.AbstractStreamBuilder`, `org.apache.commons.io.function.Uncheck`
- Classes: `org.apache.commons.csv.CSVParser`, `org.apache.commons.csv.Builder`, `org.apache.commons.csv.CSVRecordIterator`, `org.apache.commons.csv.Headers`
- Fields in `org.apache.commons.csv.Builder`: `CSVFormat format`, `long byteOffset`, `long characterOffset`, `long recordNumber`, `boolean trackBytes`
- Constructors in `org.apache.commons.csv.Builder`: `Builder(0 params)`
- Methods in `org.apache.commons.csv.Builder`: `instance CSVParser get(0 params)`, `instance Builder setByteOffset(1 params)`, `instance Builder setCharacterOffset(1 params)`, `instance Builder setFormat(1 params)`, `instance Builder setRecordNumber(1 params)`, `instance Builder setTrackBytes(1 params)`
- Fields in `org.apache.commons.csv.CSVRecordIterator`: `CSVRecord current`, `long recordCount`
- Methods in `org.apache.commons.csv.CSVRecordIterator`: `instance CSVRecord getNextRecord(0 params)`, `instance boolean hasNext(0 params)`, `instance CSVRecord next(0 params)`, `instance void remove(0 params)`
- Fields in `org.apache.commons.csv.Headers`: `String headerComment`, `String trailerComment`, `CSVFormat format`, `Headers headers`, `Lexer lexer`, `CSVRecordIterator csvRecordIterator`, `long recordNumber`, `long byteOffset`, `long characterOffset`, `Token reusableToken`
- Constructors in `org.apache.commons.csv.Headers`: `Headers(0 params)`, `SuppressWarnings(0 params)`, `CSVParser(1 params)`, `CSVParser(2 params)`, `CSVParser(4 params)`, `createEmptyHeaderMap(0 params)`, `getHeaderMap(0 params)`, `getHeaderMapRaw(0 params)`, `getHeaderNames(0 params)`, `getRecords(0 params)`, `iterator(0 params)`, `stream(0 params)`
- Methods in `org.apache.commons.csv.Headers`: `static Builder builder(0 params)`, `static CSVParser parse(3 params)`, `static CSVParser parse(3 params)`, `static CSVParser parse(3 params)`, `static CSVParser parse(2 params)`, `static CSVParser parse(2 params)`, `static CSVParser parse(3 params)`, `instance void addRecordValue(1 params)`, `instance void close(0 params)`, `instance Headers createHeaders(0 params)`, `instance long getCurrentLineNumber(0 params)`, `instance String getFirstEndOfLine(0 params)`, `instance String getHeaderComment(0 params)`, `instance long getRecordNumber(0 params)`, `instance String getTrailerComment(0 params)`, `instance String handleNull(1 params)`, `instance boolean hasHeaderComment(0 params)`, `instance boolean hasTrailerComment(0 params)`, `instance boolean isClosed(0 params)`, `instance boolean isStrictQuoteMode(0 params)`, `instance CSVRecord nextRecord(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `org/apache/commons/csv/CSVPrinter.java`

- Source root: `.`
- Risk: `high`
- Package: `org.apache.commons.csv`
- Imports: `staticorg.apache.commons.csv.Constants.CR`, `staticorg.apache.commons.csv.Constants.LF`, `staticorg.apache.commons.csv.Constants.SP`, `java.io.Closeable`, `java.io.Flushable`, `java.io.IOException`, `java.io.InputStream`, `java.io.Reader`, `java.sql.Blob`, `java.sql.Clob`, `java.sql.ResultSet`, `java.sql.SQLException`, `java.sql.Statement`, `java.util.Arrays`, `java.util.Objects`, `java.util.concurrent.locks.ReentrantLock`, `java.util.stream.Stream`, `org.apache.commons.io.function.IOStream`
- Classes: `org.apache.commons.csv.CSVPrinter`
- Fields in `org.apache.commons.csv.CSVPrinter`: `implements Flushable`, `implements Closeable`, `Appendable appendable`, `CSVFormat format`, `boolean newRecord`, `long recordCount`, `ReentrantLock lock`
- Constructors in `org.apache.commons.csv.CSVPrinter`: `CSVPrinter(2 params)`
- Methods in `org.apache.commons.csv.CSVPrinter`: `instance void close(0 params)`, `instance void close(1 params)`, `instance void endOfRecord(0 params)`, `instance void flush(0 params)`, `instance Appendable getOut(0 params)`, `instance long getRecordCount(0 params)`, `instance void print(1 params)`, `instance void printComment(1 params)`, `instance void printHeaders(1 params)`, `instance void println(0 params)`, `instance void printRaw(1 params)`, `instance void printRecord(0 params)`, `instance void printRecord(0 params)`, `instance void printRecord(0 params)`, `instance void printRecordObject(1 params)`, `instance void printRecords(0 params)`, `instance void printRecords(0 params)`, `instance void printRecords(0 params)`, `instance void printRecords(1 params)`, `instance void printRecords(2 params)`, `instance void printRecords(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `org/apache/commons/csv/CSVRecord.java`

- Source root: `.`
- Risk: `high`
- Package: `org.apache.commons.csv`
- Imports: `java.io.Serializable`, `java.util.Arrays`, `java.util.Iterator`, `java.util.LinkedHashMap`, `java.util.List`, `java.util.Map`, `java.util.stream.Collectors`, `java.util.stream.Stream`
- Classes: `org.apache.commons.csv.CSVRecord`
- Fields in `org.apache.commons.csv.CSVRecord`: `implements Serializable`, `implements Iterable`, `long serialVersionUID`, `long characterPosition`, `long bytePosition`, `String comment`, `long recordNumber`, `String[] values`, `CSVParser parser`
- Constructors in `org.apache.commons.csv.CSVRecord`: `CSVRecord(6 params)`, `getHeaderMapRaw(0 params)`, `iterator(0 params)`, `stream(0 params)`, `toList(0 params)`, `toMap(0 params)`
- Methods in `org.apache.commons.csv.CSVRecord`: `instance String get(0 params)`, `instance String get(1 params)`, `instance String get(1 params)`, `instance long getBytePosition(0 params)`, `instance long getCharacterPosition(0 params)`, `instance String getComment(0 params)`, `instance CSVParser getParser(0 params)`, `instance long getRecordNumber(0 params)`, `instance boolean hasComment(0 params)`, `instance boolean isConsistent(0 params)`, `instance boolean isMapped(1 params)`, `instance boolean isSet(1 params)`, `instance boolean isSet(1 params)`, `instance M putIn(1 params)`, `instance int size(0 params)`, `instance String toString(0 params)`, `instance String[] values(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `org/apache/commons/csv/Constants.java`

- Source root: `.`
- Risk: `high`
- Package: `org.apache.commons.csv`
- Classes: `org.apache.commons.csv.Constants`
- Fields in `org.apache.commons.csv.Constants`: `char BACKSLASH`, `char BACKSPACE`, `String COMMA`, `char COMMENT`, `char CR`, `String CRLF`, `Character DOUBLE_QUOTE_CHAR`, `String EMPTY`, `String[] EMPTY_STRING_ARRAY`, `char FF`, `char LF`, `String LINE_SEPARATOR`, `String NEXT_LINE`, `String PARAGRAPH_SEPARATOR`, `char PIPE`, `char RS`, `char SP`, `String SQL_NULL_STRING`, `char TAB`, `int UNDEFINED`, `char US`
- Constructors in `org.apache.commons.csv.Constants`: `Constants(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `org/apache/commons/csv/DuplicateHeaderMode.java`

- Source root: `.`
- Risk: `high`
- Package: `org.apache.commons.csv`
- Diagnostics: `JTG0001`

### `org/apache/commons/csv/ExtendedBufferedReader.java`

- Source root: `.`
- Risk: `high`
- Package: `org.apache.commons.csv`
- Imports: `staticorg.apache.commons.csv.Constants.CR`, `staticorg.apache.commons.csv.Constants.LF`, `staticorg.apache.commons.csv.Constants.UNDEFINED`, `staticorg.apache.commons.io.IOUtils.EOF`, `java.io.IOException`, `java.io.Reader`, `java.nio.CharBuffer`, `java.nio.charset.CharacterCodingException`, `java.nio.charset.Charset`, `java.nio.charset.CharsetEncoder`, `org.apache.commons.io.IOUtils`, `org.apache.commons.io.input.UnsynchronizedBufferedReader`
- Classes: `org.apache.commons.csv.ExtendedBufferedReader`
- Fields in `org.apache.commons.csv.ExtendedBufferedReader`: `int lastChar`, `int lastCharMark`, `long lineNumber`, `long lineNumberMark`, `long position`, `long positionMark`, `long bytesRead`, `long bytesReadMark`, `CharsetEncoder encoder`
- Constructors in `org.apache.commons.csv.ExtendedBufferedReader`: `ExtendedBufferedReader(1 params)`, `ExtendedBufferedReader(3 params)`, `CharacterCodingException(0 params)`
- Methods in `org.apache.commons.csv.ExtendedBufferedReader`: `instance void close(0 params)`, `instance long getEncodedCharLength(3 params)`, `instance int getEncodedCharLength(1 params)`, `instance int getLastChar(0 params)`, `instance long getLineNumber(0 params)`, `instance long getPosition(0 params)`, `instance void mark(1 params)`, `instance int read(0 params)`, `instance int read(3 params)`, `instance String readLine(0 params)`, `instance void reset(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `org/apache/commons/csv/Lexer.java`

- Source root: `.`
- Risk: `high`
- Package: `org.apache.commons.csv`
- Imports: `staticorg.apache.commons.io.IOUtils.EOF`, `java.io.Closeable`, `java.io.IOException`, `java.util.Arrays`, `org.apache.commons.io.IOUtils`
- Classes: `org.apache.commons.csv.Lexer`
- Fields in `org.apache.commons.csv.Lexer`: `String CR_STRING`, `String LF_STRING`, `char[] delimiter`, `char[] delimiterBuf`, `char[] escapeDelimiterBuf`, `int escape`, `int quoteChar`, `int commentStart`, `boolean ignoreSurroundingSpaces`, `boolean ignoreEmptyLines`, `boolean lenientEof`, `boolean trailingData`, `ExtendedBufferedReader reader`, `String firstEol`, `boolean isLastTokenDelimiter`
- Constructors in `org.apache.commons.csv.Lexer`: `Lexer(2 params)`
- Methods in `org.apache.commons.csv.Lexer`: `instance void appendNextEscapedCharacterToToken(1 params)`, `instance void close(0 params)`, `instance boolean isMetaChar(1 params)`, `instance boolean isQuoteChar(1 params)`, `instance boolean isStartOfLine(1 params)`, `instance Token nextToken(1 params)`, `instance int nullToDisabled(1 params)`, `instance Token parseEncapsulatedToken(1 params)`, `instance Token parseSimpleToken(2 params)`, `instance void trimTrailingSpaces(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `org/apache/commons/csv/QuoteMode.java`

- Source root: `.`
- Risk: `high`
- Package: `org.apache.commons.csv`
- Diagnostics: `JTG0001`

### `org/apache/commons/csv/Token.java`

- Source root: `.`
- Risk: `high`
- Package: `org.apache.commons.csv`
- Imports: `staticorg.apache.commons.csv.Token.Type.INVALID`
- Classes: `org.apache.commons.csv.Token`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `org/apache/commons/csv/package-info.java`

- Source root: `.`
- Risk: `medium`
- Package: `org.apache.commons.csv`

## Recommended Migration Order

- `org/apache/commons/csv/package-info.java` (medium risk)
- `org/apache/commons/csv/CSVException.java` (high risk)
- `org/apache/commons/csv/CSVFormat.java` (high risk)
- `org/apache/commons/csv/CSVParser.java` (high risk)
- `org/apache/commons/csv/CSVPrinter.java` (high risk)
- `org/apache/commons/csv/CSVRecord.java` (high risk)
- `org/apache/commons/csv/Constants.java` (high risk)
- `org/apache/commons/csv/DuplicateHeaderMode.java` (high risk)
- `org/apache/commons/csv/ExtendedBufferedReader.java` (high risk)
- `org/apache/commons/csv/Lexer.java` (high risk)
- `org/apache/commons/csv/QuoteMode.java` (high risk)
- `org/apache/commons/csv/Token.java` (high risk)

## Migration Execution Summary

| Metric | Count |
| --- | ---: |
| Java files | 12 |
| Generated files | 1 |
| Skipped files | 11 |
| Diagnostics | 316 |
| Dry run | false |

### Migration Diagnostics

- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVException.java:31:26: missing '{' at 'extends'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVException.java:33:50: extraneous input 'L' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVException.java:42:57: mismatched input '.' expecting Identifier`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVException.java:43:14: mismatched input 'String' expecting {'(', ')', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVException.java:43:20: extraneous input '.' expecting Identifier`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVException.java:43:27: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVException.java:43:34: mismatched input ',' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:228:8: JTG2004: overload not supported: method org.apache.commons.csv.Builder.create`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:438:8: JTG2004: overload not supported: method org.apache.commons.csv.Builder.setCommentMarker`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:462:8: JTG2004: overload not supported: method org.apache.commons.csv.Builder.setDelimiter`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:504:8: JTG2004: overload not supported: method org.apache.commons.csv.Builder.setEscape`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:563:8: JTG2004: overload not supported: method org.apache.commons.csv.Builder.setHeader`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:587:8: JTG2004: overload not supported: method org.apache.commons.csv.Builder.setHeader`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:587:8: JTG2004: overload not supported: method org.apache.commons.csv.Builder.setHeader`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:621:8: JTG2004: overload not supported: method org.apache.commons.csv.Builder.setHeader`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:621:8: JTG2004: overload not supported: method org.apache.commons.csv.Builder.setHeader`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:621:8: JTG2004: overload not supported: method org.apache.commons.csv.Builder.setHeader`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:703:8: JTG2001: duplicate symbol: method org.apache.commons.csv.Builder.setHeaderComments`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:804:8: JTG2004: overload not supported: method org.apache.commons.csv.Builder.setQuote`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:852:8: JTG2004: overload not supported: method org.apache.commons.csv.Builder.setRecordSeparator`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:189:18: JTG2001: duplicate Go symbol after migration: Builder`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:277:8: JTG2001: duplicate Go symbol after migration: NewBuilder`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:281:8: JTG2001: duplicate Go symbol after migration: NewBuilder`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:211:8: JTG2001: duplicate Go symbol after migration: Create`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVFormat.java:228:8: JTG2001: duplicate Go symbol after migration: Create`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:154:18: JTG2001: duplicate symbol: class org.apache.commons.csv.Builder`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:369:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:393:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:393:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:418:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:418:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:418:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:437:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:437:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:437:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:437:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:465:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:465:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:465:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:465:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:465:4: JTG2004: overload not supported: method org.apache.commons.csv.Headers.parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:309:8: JTG2001: duplicate Go symbol after migration: NewHeaders`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:510:5: JTG2001: duplicate Go symbol after migration: NewHeaders`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:511:4: JTG2001: duplicate Go symbol after migration: NewHeaders`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:541:4: JTG2001: duplicate Go symbol after migration: NewHeaders`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:570:4: JTG2001: duplicate Go symbol after migration: NewHeaders`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:600:33: JTG2001: duplicate Go symbol after migration: NewHeaders`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:717:32: JTG2001: duplicate Go symbol after migration: NewHeaders`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:731:25: JTG2001: duplicate Go symbol after migration: NewHeaders`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:747:24: JTG2001: duplicate Go symbol after migration: NewHeaders`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:780:27: JTG2001: duplicate Go symbol after migration: NewHeaders`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:886:31: JTG2001: duplicate Go symbol after migration: NewHeaders`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:956:29: JTG2001: duplicate Go symbol after migration: NewHeaders`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:154:18: JTG2001: duplicate Go symbol after migration: Builder`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:321:4: JTG2001: duplicate Go symbol after migration: Builder`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:342:4: JTG2001: duplicate Go symbol after migration: Parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:369:4: JTG2001: duplicate Go symbol after migration: Parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:393:4: JTG2001: duplicate Go symbol after migration: Parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:418:4: JTG2001: duplicate Go symbol after migration: Parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:437:4: JTG2001: duplicate Go symbol after migration: Parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:465:4: JTG2001: duplicate Go symbol after migration: Parse`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVParser.java:165:8: JTG2001: duplicate Go symbol after migration: NewBuilder`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:139:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVPrinter.close`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:350:4: JTG2001: duplicate symbol: method org.apache.commons.csv.CSVPrinter.printRecord`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:369:4: JTG2001: duplicate symbol: method org.apache.commons.csv.CSVPrinter.printRecord`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:369:4: JTG2001: duplicate symbol: method org.apache.commons.csv.CSVPrinter.printRecord`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:433:4: JTG2001: duplicate symbol: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:474:4: JTG2001: duplicate symbol: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:474:4: JTG2001: duplicate symbol: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:489:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:489:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:489:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:528:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:528:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:528:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:528:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:576:4: JTG2001: duplicate symbol: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:576:4: JTG2001: duplicate symbol: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:576:4: JTG2001: duplicate symbol: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:576:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVPrinter.java:576:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVPrinter.printRecords`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVRecord.java:98:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVRecord.get`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVRecord.java:125:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVRecord.get`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVRecord.java:125:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVRecord.get`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVRecord.java:267:4: JTG2004: overload not supported: method org.apache.commons.csv.CSVRecord.isSet`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVRecord.java:70:4: JTG2001: duplicate Go symbol after migration: NewCSVRecord`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVRecord.java:175:33: JTG2001: duplicate Go symbol after migration: NewCSVRecord`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVRecord.java:277:28: JTG2001: duplicate Go symbol after migration: NewCSVRecord`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVRecord.java:316:26: JTG2001: duplicate Go symbol after migration: NewCSVRecord`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVRecord.java:329:24: JTG2001: duplicate Go symbol after migration: NewCSVRecord`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/CSVRecord.java:343:31: JTG2001: duplicate Go symbol after migration: NewCSVRecord`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:27:34: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:27:35: token recognition error at: '\'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:27:36: token recognition error at: '\'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:27:37: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:27:38: mismatched input ';' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:29:34: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:29:35: token recognition error at: '\'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:29:37: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:36:32: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:36:33: token recognition error at: '#'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:36:34: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:36:35: mismatched input ';' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:38:27: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:38:28: token recognition error at: '\'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:38:30: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:43:65: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:43:66: token recognition error at: '"');  // Explicit boxing is intentional.'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:45:4: mismatched input 'static' expecting {'(', ')', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:47:47: mismatched input '{' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:49:27: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:49:28: token recognition error at: '\'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:49:30: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:51:27: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:51:28: token recognition error at: '\'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:51:30: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:68:29: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:68:30: token recognition error at: '|''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:68:32: mismatched input ';' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:73:27: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:73:29: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:73:30: mismatched input ';' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:77:28: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:77:29: token recognition error at: '\'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Constants.java:77:31: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/DuplicateHeaderMode.java:28:7: no viable alternative at input 'publicenum'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/ExtendedBufferedReader.java:143:4: JTG2004: overload not supported: method org.apache.commons.csv.ExtendedBufferedReader.getEncodedCharLength`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/ExtendedBufferedReader.java:216:4: JTG2004: overload not supported: method org.apache.commons.csv.ExtendedBufferedReader.read`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/ExtendedBufferedReader.java:72:4: JTG2001: duplicate Go symbol after migration: NewExtendedBufferedReader`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/ExtendedBufferedReader.java:84:4: JTG2001: duplicate Go symbol after migration: NewExtendedBufferedReader`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/ExtendedBufferedReader.java:156:18: JTG2001: duplicate Go symbol after migration: NewExtendedBufferedReader`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:33:18: missing '{' at 'implements'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:56:20: no viable alternative at input 'this.reader='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:57:23: no viable alternative at input 'this.delimiter='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:58:20: no viable alternative at input 'this.escape='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:59:23: no viable alternative at input 'this.quoteChar='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:60:26: no viable alternative at input 'this.commentStart='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:61:37: no viable alternative at input 'this.ignoreSurroundingSpaces='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:62:30: no viable alternative at input 'this.ignoreEmptyLines='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:63:24: no viable alternative at input 'this.lenientEof='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:64:26: no viable alternative at input 'this.trailingData='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:65:26: no viable alternative at input 'this.delimiterBuf='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:65:36: mismatched input '[' expecting '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:65:57: extraneous input ']' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:66:32: no viable alternative at input 'this.escapeDelimiterBuf='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:66:42: mismatched input '[' expecting '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:66:67: extraneous input ']' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:76:70: missing '{' at 'throws'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:76:89: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:82:44: extraneous input 'escape' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:82:66: mismatched input 'reader' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:82:86: extraneous input ')' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:84:44: extraneous input 'unescaped' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:95:4: token recognition error at: '@'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:96:4: no viable alternative at input 'Overridepublic'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:96:24: missing '{' at 'throws'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:96:43: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:105:21: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:105:22: mismatched input ')' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:105:24: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:114:29: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:114:30: mismatched input ')' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:114:32: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:123:29: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:123:30: mismatched input ')' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:123:32: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:127:22: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:127:23: mismatched input ')' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:127:25: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:131:20: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:131:21: mismatched input ')' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:131:23: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:135:26: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:135:27: mismatched input 'final' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:135:39: mismatched input ')' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:147:23: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:147:24: mismatched input 'final' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:147:36: mismatched input ')' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:147:57: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:149:27: mismatched input '[' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:149:29: mismatched input ']' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:157:51: mismatched input '+' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:157:52: mismatched input ')' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:158:28: mismatched input '[' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:158:30: mismatched input ']' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:158:45: missing ']' at 'i'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:158:47: mismatched input '+' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:158:50: mismatched input ']' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:172:23: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:172:24: mismatched input 'final' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:172:36: mismatched input ')' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:181:20: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:181:21: mismatched input 'final' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:181:33: mismatched input ')' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:193:29: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:193:30: extraneous input ')' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:193:39: mismatched input 'IOException' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:194:40: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:194:41: token recognition error at: '\'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:194:43: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:196:30: mismatched input '[' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:196:32: mismatched input ']' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:196:47: extraneous input '0' expecting ']'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:196:49: mismatched input ')' expecting Identifier`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:199:48: mismatched input '+' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:199:49: mismatched input ')' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:200:34: mismatched input '[' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:200:40: mismatched input ']' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:200:55: extraneous input 'i' expecting ']'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:200:58: extraneous input '||' expecting Identifier`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:200:79: mismatched input '[' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:200:89: mismatched input ']' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:208:4: extraneous input 'private' expecting {';', '{', '}', 'final', '(', 'boolean', 'int', 'double', 'String', 'if', 'while', 'for', 'return', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:237:39: missing '{' at 'throws'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:237:58: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:253:31: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:261:23: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:268:27: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:274:23: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:277:34: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:277:35: token recognition error at: '\'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:277:37: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:282:53: extraneous input 'c' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:290:27: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:294:27: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:301:27: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:302:30: no viable alternative at input 'token.isReady='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:312:4: extraneous input 'private' expecting {';', '{', '}', 'final', '(', 'boolean', 'int', 'double', 'String', 'if', 'while', 'for', 'return', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:313:25: token recognition error at: '?'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:313:27: missing ';' at 'Constants'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:313:47: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:313:50: mismatched input '.' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:339:60: missing '{' at 'throws'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:339:79: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:340:23: no viable alternative at input 'token.isQuoted='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:350:48: extraneous input 'c' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:356:39: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:360:39: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:361:42: no viable alternative at input 'token.isReady='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:365:39: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:369:56: extraneous input 'c' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:370:66: extraneous input 'c' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:372:34: no viable alternative at input 'thrownew'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:381:31: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:382:34: no viable alternative at input 'token.isReady='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:386:22: no viable alternative at input 'thrownew'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:389:44: extraneous input 'c' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:412:4: extraneous input 'private' expecting {';', '{', '}', 'final', '(', 'boolean', 'int', 'double', 'String', 'if', 'while', 'for', 'return', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:412:68: missing '{' at 'throws'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:412:87: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:417:27: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:421:27: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:422:30: no viable alternative at input 'token.isReady='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:426:27: no viable alternative at input 'token.type='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:433:44: extraneous input 'cur' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:450:25: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:450:26: mismatched input 'final' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:450:38: mismatched input ')' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:450:59: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:458:30: no viable alternative at input 'this.firstEol='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:464:30: no viable alternative at input 'this.firstEol='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:466:30: no viable alternative at input 'this.firstEol='`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:482:18: missing ';' at '('`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:482:19: extraneous input ')' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:482:28: mismatched input 'IOException' expecting ')'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:485:20: missing ';' at '{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:486:13: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:486:15: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:486:16: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:487:12: missing ';' at 'return'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:488:13: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:488:15: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:488:16: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:489:12: missing ';' at 'return'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:490:13: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:490:15: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:490:16: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:491:12: missing ';' at 'return'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:492:13: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:492:15: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:492:16: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:493:12: missing ';' at 'return'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:494:13: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:494:15: token recognition error at: '''`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:494:16: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:495:12: missing ';' at 'return'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:496:22: mismatched input '.' expecting ';'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:496:25: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:497:13: missing ';' at 'Constants'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:497:25: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:498:13: missing ';' at 'Constants'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:498:25: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:499:13: missing ';' at 'Constants'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:499:26: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:500:13: missing ';' at 'Constants'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:500:32: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:502:16: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:503:12: missing ';' at 'throw'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:504:15: token recognition error at: ':'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:506:12: no viable alternative at input 'defaultif'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:514:4: extraneous input 'void' expecting {';', '{', '}', 'final', '(', 'boolean', 'int', 'double', 'String', 'if', 'while', 'for', 'return', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Lexer.java:517:20: mismatched input ';' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/QuoteMode.java:26:7: no viable alternative at input 'publicenum'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:32:14: no viable alternative at input 'enumType{'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:51:25: no viable alternative at input 'privatestaticfinalint'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:54:4: extraneous input 'Token' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:57:10: no viable alternative at input 'finalStringBuilder'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:60:4: extraneous input 'boolean' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:62:4: extraneous input 'boolean' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:64:4: extraneous input 'void' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:66:8: extraneous input 'type' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:67:8: extraneous input 'isReady' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:68:8: extraneous input 'isQuoted' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:69:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:76:4: token recognition error at: '@'`
- `.corpus/commons-csv/src/main/java/org/apache/commons/csv/Token.java:79:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
