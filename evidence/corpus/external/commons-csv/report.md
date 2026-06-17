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
