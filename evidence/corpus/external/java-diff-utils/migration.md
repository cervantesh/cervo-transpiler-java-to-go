# Java-To-Go Migration Report

Project: `.corpus/java-diff-utils`

## Summary

| Metric | Count |
| --- | ---: |
| Java files | 55 |
| Packages | 9 |
| Classes | 62 |
| Methods | 254 |
| Constructors | 252 |
| Fields | 106 |
| Unsupported features | 0 |
| Diagnostics | 3700 |
| Internal imports | 207 |

## Build Files

- `pom.xml` (maven)

## Packages

- `com.github.difflib`
- `com.github.difflib.algorithm`
- `com.github.difflib.algorithm.jgit`
- `com.github.difflib.algorithm.myers`
- `com.github.difflib.examples`
- `com.github.difflib.patch`
- `com.github.difflib.text`
- `com.github.difflib.text.deltamerge`
- `com.github.difflib.unifieddiff`

## Feature Counts

| Feature | Count |
| --- | ---: |
| class declarations | 62 |
| class fields | 106 |
| constructors | 252 |
| import declarations | 403 |
| instance methods | 236 |
| package declarations | 55 |
| static methods | 18 |

## Unsupported Feature Counts

None.

## Files

### `java-diff-utils-jgit/src/main/java/com/github/difflib/algorithm/jgit/HistogramDiff.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm.jgit`
- Imports: `com.github.difflib.algorithm.Change`, `com.github.difflib.algorithm.DiffAlgorithmI`, `com.github.difflib.algorithm.DiffAlgorithmListener`, `com.github.difflib.patch.DeltaType`, `java.util.ArrayList`, `java.util.List`, `java.util.Objects`, `org.eclipse.jgit.diff.Edit`, `org.eclipse.jgit.diff.EditList`, `org.eclipse.jgit.diff.Sequence`, `org.eclipse.jgit.diff.SequenceComparator`
- Classes: `com.github.difflib.algorithm.jgit.HistogramDiff`, `com.github.difflib.algorithm.jgit.DataListComparator`, `com.github.difflib.algorithm.jgit.DataList`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils-jgit/src/test/java/com/github/difflib/algorithm/jgit/HistogramDiffTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm.jgit`
- Imports: `staticorg.junit.jupiter.api.Assertions.assertEquals`, `staticorg.junit.jupiter.api.Assertions.assertNotNull`, `com.github.difflib.algorithm.DiffAlgorithmListener`, `com.github.difflib.patch.Patch`, `com.github.difflib.patch.PatchFailedException`, `java.util.ArrayList`, `java.util.Arrays`, `java.util.List`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.algorithm.jgit.HistogramDiffTest`
- Constructors in `com.github.difflib.algorithm.jgit.HistogramDiffTest`: `HistogramDiffTest(0 params)`, `generate(1 params)`, `generate(1 params)`
- Methods in `com.github.difflib.algorithm.jgit.HistogramDiffTest`: `instance void testDiff(0 params)`, `instance void testDiffWithListener(0 params)`, `instance void diffStart(0 params)`, `instance void diffStep(2 params)`, `instance void diffEnd(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils-jgit/src/test/java/com/github/difflib/algorithm/jgit/LRHistogramDiffTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm.jgit`
- Imports: `staticjava.util.stream.Collectors.toList`, `staticorg.junit.jupiter.api.Assertions.assertArrayEquals`, `staticorg.junit.jupiter.api.Assertions.assertEquals`, `com.github.difflib.algorithm.DiffAlgorithmListener`, `com.github.difflib.patch.Patch`, `com.github.difflib.patch.PatchFailedException`, `java.io.BufferedReader`, `java.io.IOException`, `java.io.InputStream`, `java.io.InputStreamReader`, `java.nio.charset.Charset`, `java.nio.charset.StandardCharsets`, `java.util.ArrayList`, `java.util.List`, `java.util.zip.ZipFile`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.algorithm.jgit.LRHistogramDiffTest`
- Methods in `com.github.difflib.algorithm.jgit.LRHistogramDiffTest`: `instance void testPossibleDiffHangOnLargeDatasetDnaumenkoIssue26(0 params)`, `instance void diffStart(0 params)`, `instance void diffStep(2 params)`, `instance void diffEnd(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib`
- Imports: `com.github.difflib.algorithm.DiffAlgorithmFactory`, `com.github.difflib.algorithm.DiffAlgorithmI`, `com.github.difflib.algorithm.DiffAlgorithmListener`, `com.github.difflib.algorithm.myers.MyersDiff`, `com.github.difflib.patch.AbstractDelta`, `com.github.difflib.patch.Patch`, `com.github.difflib.patch.PatchFailedException`, `java.util.ArrayList`, `java.util.Arrays`, `java.util.Collections`, `java.util.List`, `java.util.Objects`, `java.util.function.BiPredicate`
- Classes: `com.github.difflib.DiffUtils`
- Fields in `com.github.difflib.DiffUtils`: `DiffAlgorithmFactory DEFAULT_DIFF`
- Constructors in `com.github.difflib.DiffUtils`: `diff(1 params)`, `diff(1 params)`, `diff(1 params)`, `diff(3 params)`, `diff(1 params)`, `diff(1 params)`, `diff(1 params)`, `diff(1 params)`, `diffInline(2 params)`, `patch(1 params)`, `unpatch(1 params)`, `compressLines(0 params)`, `DiffUtils(0 params)`
- Methods in `com.github.difflib.DiffUtils`: `static void withDefaultDiffAlgorithmFactory(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/UnifiedDiffUtils.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib`
- Imports: `com.github.difflib.patch.AbstractDelta`, `com.github.difflib.patch.ChangeDelta`, `com.github.difflib.patch.Chunk`, `com.github.difflib.patch.Patch`, `java.util.ArrayList`, `java.util.HashMap`, `java.util.List`, `java.util.Map`, `java.util.Optional`, `java.util.regex.Matcher`, `java.util.regex.Pattern`, `java.util.stream.Collectors`
- Classes: `com.github.difflib.UnifiedDiffUtils`
- Fields in `com.github.difflib.UnifiedDiffUtils`: `Pattern UNIFIED_DIFF_CHUNK_REGEXP`, `String NULL_FILE_INDICATOR`
- Constructors in `com.github.difflib.UnifiedDiffUtils`: `parseUnifiedDiff(0 params)`, `generateUnifiedDiff(2 params)`, `processDeltas(0 params)`, `getDeltaText(0 params)`, `UnifiedDiffUtils(0 params)`, `generateOriginalAndDiff(0 params)`, `generateOriginalAndDiff(0 params)`, `insertOrig(0 params)`, `getRowMap(1 params)`, `getOrigList(0 params)`
- Methods in `com.github.difflib.UnifiedDiffUtils`: `static void processLinesInPrevChunk(0 params)`, `static void insertOrig(1 params)`, `static void insert(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/algorithm/Change.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm`
- Imports: `com.github.difflib.patch.DeltaType`
- Classes: `com.github.difflib.algorithm.Change`
- Fields in `com.github.difflib.algorithm.Change`: `DeltaType deltaType`, `int startOriginal`, `int endOriginal`, `int startRevised`, `int endRevised`
- Constructors in `com.github.difflib.algorithm.Change`: `Change(5 params)`
- Methods in `com.github.difflib.algorithm.Change`: `instance Change withEndOriginal(1 params)`, `instance Change withEndRevised(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm`
- Imports: `java.util.function.BiPredicate`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmI.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm`
- Imports: `java.util.Arrays`, `java.util.List`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmListener.java`

- Source root: `.`
- Risk: `medium`
- Package: `com.github.difflib.algorithm`

### `java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/MyersDiff.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm.myers`
- Imports: `com.github.difflib.algorithm.Change`, `com.github.difflib.algorithm.DiffAlgorithmFactory`, `com.github.difflib.algorithm.DiffAlgorithmI`, `com.github.difflib.algorithm.DiffAlgorithmListener`, `com.github.difflib.patch.DeltaType`, `com.github.difflib.patch.Patch`, `java.util.ArrayList`, `java.util.List`, `java.util.Objects`, `java.util.function.BiPredicate`
- Classes: `com.github.difflib.algorithm.myers.MyersDiff`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/MyersDiffWithLinearSpace.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm.myers`
- Imports: `com.github.difflib.algorithm.Change`, `com.github.difflib.algorithm.DiffAlgorithmFactory`, `com.github.difflib.algorithm.DiffAlgorithmI`, `com.github.difflib.algorithm.DiffAlgorithmListener`, `com.github.difflib.patch.DeltaType`, `java.util.ArrayList`, `java.util.List`, `java.util.Objects`, `java.util.function.BiPredicate`, `java.util.function.Consumer`
- Classes: `com.github.difflib.algorithm.myers.MyersDiffWithLinearSpace`, `com.github.difflib.algorithm.myers.DiffData`, `com.github.difflib.algorithm.myers.Snake`
- Fields in `com.github.difflib.algorithm.myers.DiffData`: `int size`, `int[] vDown`, `int[] vUp`
- Constructors in `com.github.difflib.algorithm.myers.DiffData`: `DiffData(1 params)`
- Fields in `com.github.difflib.algorithm.myers.Snake`: `int start`, `int end`, `int diag`
- Constructors in `com.github.difflib.algorithm.myers.Snake`: `Snake(3 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm.myers`
- Classes: `com.github.difflib.algorithm.myers.PathNode`
- Fields in `com.github.difflib.algorithm.myers.PathNode`: `int i`, `int j`, `PathNode prev`, `boolean snake`, `boolean bootstrap`
- Constructors in `com.github.difflib.algorithm.myers.PathNode`: `PathNode(5 params)`
- Methods in `com.github.difflib.algorithm.myers.PathNode`: `instance boolean isSnake(0 params)`, `instance boolean isBootstrap(0 params)`, `instance PathNode previousSnake(0 params)`, `instance String toString(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.io.Serializable`, `java.util.List`, `java.util.Objects`
- Classes: `com.github.difflib.patch.AbstractDelta`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.util.List`, `java.util.Objects`
- Classes: `com.github.difflib.patch.ChangeDelta`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.io.Serializable`, `java.util.ArrayList`, `java.util.Arrays`, `java.util.List`, `java.util.Objects`
- Classes: `com.github.difflib.patch.Chunk`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/patch/ConflictOutput.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.io.Serializable`, `java.util.List`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.util.List`
- Classes: `com.github.difflib.patch.DeleteDelta`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/patch/DeltaType.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Diagnostics: `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/patch/DiffException.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Classes: `com.github.difflib.patch.DiffException`
- Fields in `com.github.difflib.patch.DiffException`: `long serialVersionUID`
- Constructors in `com.github.difflib.patch.DiffException`: `DiffException(0 params)`, `DiffException(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/patch/EqualDelta.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.util.List`
- Classes: `com.github.difflib.patch.EqualDelta`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.util.List`
- Classes: `com.github.difflib.patch.InsertDelta`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `staticjava.util.Comparator.comparing`, `com.github.difflib.algorithm.Change`, `java.io.Serializable`, `java.util.ArrayList`, `java.util.Collections`, `java.util.List`, `java.util.ListIterator`
- Classes: `com.github.difflib.patch.Patch`, `com.github.difflib.patch.PatchApplyingContext`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/patch/PatchFailedException.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Classes: `com.github.difflib.patch.PatchFailedException`
- Fields in `com.github.difflib.patch.PatchFailedException`: `long serialVersionUID`
- Constructors in `com.github.difflib.patch.PatchFailedException`: `PatchFailedException(0 params)`, `PatchFailedException(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/patch/VerifyChunk.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Diagnostics: `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.text`
- Imports: `java.io.Serializable`, `java.util.Objects`
- Classes: `com.github.difflib.text.DiffRow`
- Fields in `com.github.difflib.text.DiffRow`: `Tag tag`, `String oldLine`, `String newLine`
- Constructors in `com.github.difflib.text.DiffRow`: `DiffRow(3 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/text/DiffRowGenerator.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.text`
- Imports: `staticjava.util.stream.Collectors.toList`, `com.github.difflib.DiffUtils`, `com.github.difflib.patch.AbstractDelta`, `com.github.difflib.patch.ChangeDelta`, `com.github.difflib.patch.Chunk`, `com.github.difflib.patch.DeleteDelta`, `com.github.difflib.patch.DeltaType`, `com.github.difflib.patch.InsertDelta`, `com.github.difflib.patch.Patch`, `com.github.difflib.text.DiffRow.Tag`, `com.github.difflib.text.deltamerge.DeltaMergeUtils`, `com.github.difflib.text.deltamerge.InlineDeltaMergeInfo`, `java.util.*`, `java.util.function.BiFunction`, `java.util.function.BiPredicate`, `java.util.function.Function`, `java.util.regex.Matcher`, `java.util.regex.Pattern`
- Classes: `com.github.difflib.text.DiffRowGenerator`, `com.github.difflib.text.Builder`
- Fields in `com.github.difflib.text.DiffRowGenerator`: `Object equals`, `StringUtils normalize`
- Constructors in `com.github.difflib.text.DiffRowGenerator`: `adjustWhitespace(1 params)`, `equals(1 params)`, `length(0 params)`, `toCharArray(0 params)`
- Fields in `com.github.difflib.text.Builder`: `boolean showInlineDiffs`, `boolean ignoreWhiteSpaces`, `boolean decompressDeltas`, `int columnWidth`, `boolean mergeOriginalRevised`, `boolean reportLinesUnchanged`, `boolean replaceOriginalLinefeedInChangesWithSpaces`
- Constructors in `com.github.difflib.text.Builder`: `Builder(0 params)`
- Methods in `com.github.difflib.text.Builder`: `instance Builder showInlineDiffs(1 params)`, `instance Builder ignoreWhiteSpaces(1 params)`, `instance Builder reportLinesUnchanged(1 params)`, `instance Builder oldTag(3 params)`, `instance Builder oldTag(2 params)`, `instance Builder newTag(3 params)`, `instance Builder newTag(2 params)`, `instance Builder processDiffs(0 params)`, `instance Builder processEqualities(0 params)`, `instance Builder columnWidth(1 params)`, `instance DiffRowGenerator build(0 params)`, `instance Builder mergeOriginalRevised(1 params)`, `instance Builder decompressDeltas(1 params)`, `instance Builder inlineDiffByWord(1 params)`, `instance Builder inlineDiffBySplitter(0 params)`, `instance Builder lineNormalizer(0 params)`, `instance Builder equalizer(0 params)`, `instance Builder replaceOriginalLinefeedInChangesWithSpaces(1 params)`, `instance Builder inlineDeltaMerger(2 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/text/StringUtils.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.text`
- Imports: `staticjava.util.stream.Collectors.toList`, `java.util.List`
- Classes: `com.github.difflib.text.StringUtils`
- Constructors in `com.github.difflib.text.StringUtils`: `wrapText(0 params)`, `StringUtils(0 params)`
- Methods in `com.github.difflib.text.StringUtils`: `static String htmlEntites(1 params)`, `static String normalize(1 params)`, `static String wrapText(2 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.text.deltamerge`
- Imports: `com.github.difflib.patch.AbstractDelta`, `com.github.difflib.patch.ChangeDelta`, `com.github.difflib.patch.Chunk`, `java.util.ArrayList`, `java.util.List`, `java.util.function.Predicate`
- Classes: `com.github.difflib.text.deltamerge.DeltaMergeUtils`
- Fields in `com.github.difflib.text.deltamerge.DeltaMergeUtils`: `int i`
- Constructors in `com.github.difflib.text.deltamerge.DeltaMergeUtils`: `mergeInlineDeltas(2 params)`, `add(0 params)`, `size(0 params)`, `get(1 params)`, `getOrigList(0 params)`, `subList(0 params)`, `getPosition(0 params)`, `getSource(0 params)`, `size(0 params)`, `getSource(0 params)`, `getPosition(0 params)`, `test(1 params)`, `addAll(0 params)`, `getLines(0 params)`, `addAll(1 params)`, `getSource(0 params)`, `getPosition(0 params)`, `getTarget(0 params)`, `getPosition(0 params)`, `remove(0 params)`, `add(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/InlineDeltaMergeInfo.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.text.deltamerge`
- Imports: `com.github.difflib.patch.AbstractDelta`, `java.util.List`
- Classes: `com.github.difflib.text.deltamerge.InlineDeltaMergeInfo`
- Constructors in `com.github.difflib.text.deltamerge.InlineDeltaMergeInfo`: `InlineDeltaMergeInfo(1 params)`, `getDeltas(0 params)`, `getOrigList(0 params)`, `getRevList(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiff.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Imports: `com.github.difflib.patch.PatchFailedException`, `java.util.ArrayList`, `java.util.Collections`, `java.util.List`, `java.util.function.Predicate`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiff`
- Fields in `com.github.difflib.unifieddiff.UnifiedDiff`: `String header`, `String tail`
- Constructors in `com.github.difflib.unifieddiff.UnifiedDiff`: `getFiles(0 params)`, `applyPatchTo(0 params)`
- Methods in `com.github.difflib.unifieddiff.UnifiedDiff`: `instance String getHeader(0 params)`, `instance void setHeader(1 params)`, `instance void addFile(1 params)`, `instance void setTailTxt(1 params)`, `instance String getTail(0 params)`, `static UnifiedDiff from(2 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffFile.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Imports: `com.github.difflib.patch.Patch`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiffFile`
- Fields in `com.github.difflib.unifieddiff.UnifiedDiffFile`: `String diffCommand`, `String fromFile`, `String fromTimestamp`, `String toFile`, `String renameFrom`, `String renameTo`, `String copyFrom`, `String copyTo`, `String toTimestamp`, `String index`, `String newFileMode`, `String oldMode`, `String newMode`, `String deletedFileMode`, `String binaryAdded`, `String binaryDeleted`, `String binaryEdited`, `boolean noNewLineAtTheEndOfTheFile`, `Integer similarityIndex`
- Constructors in `com.github.difflib.unifieddiff.UnifiedDiffFile`: `getPatch(0 params)`
- Methods in `com.github.difflib.unifieddiff.UnifiedDiffFile`: `instance String getDiffCommand(0 params)`, `instance void setDiffCommand(1 params)`, `instance String getFromFile(0 params)`, `instance void setFromFile(1 params)`, `instance String getToFile(0 params)`, `instance void setToFile(1 params)`, `instance void setIndex(1 params)`, `instance String getIndex(0 params)`, `instance String getFromTimestamp(0 params)`, `instance void setFromTimestamp(1 params)`, `instance String getToTimestamp(0 params)`, `instance void setToTimestamp(1 params)`, `instance Integer getSimilarityIndex(0 params)`, `instance void setSimilarityIndex(1 params)`, `instance String getRenameFrom(0 params)`, `instance void setRenameFrom(1 params)`, `instance String getRenameTo(0 params)`, `instance void setRenameTo(1 params)`, `instance String getCopyFrom(0 params)`, `instance void setCopyFrom(1 params)`, `instance String getCopyTo(0 params)`, `instance void setCopyTo(1 params)`, `static UnifiedDiffFile from(2 params)`, `instance void setNewFileMode(1 params)`, `instance String getNewFileMode(0 params)`, `instance String getDeletedFileMode(0 params)`, `instance void setDeletedFileMode(1 params)`, `instance String getOldMode(0 params)`, `instance void setOldMode(1 params)`, `instance String getNewMode(0 params)`, `instance void setNewMode(1 params)`, `instance String getBinaryAdded(0 params)`, `instance void setBinaryAdded(1 params)`, `instance String getBinaryDeleted(0 params)`, `instance void setBinaryDeleted(1 params)`, `instance String getBinaryEdited(0 params)`, `instance void setBinaryEdited(1 params)`, `instance boolean isNoNewLineAtTheEndOfTheFile(0 params)`, `instance void setNoNewLineAtTheEndOfTheFile(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffParserException.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiffParserException`
- Constructors in `com.github.difflib.unifieddiff.UnifiedDiffParserException`: `UnifiedDiffParserException(0 params)`, `UnifiedDiffParserException(1 params)`, `UnifiedDiffParserException(2 params)`, `UnifiedDiffParserException(1 params)`, `UnifiedDiffParserException(4 params)`
- Diagnostics: `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffReader.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Imports: `com.github.difflib.patch.AbstractDelta`, `com.github.difflib.patch.ChangeDelta`, `com.github.difflib.patch.Chunk`, `com.github.difflib.patch.DeleteDelta`, `com.github.difflib.patch.EqualDelta`, `com.github.difflib.patch.InsertDelta`, `java.io.BufferedReader`, `java.io.IOException`, `java.io.InputStream`, `java.io.InputStreamReader`, `java.io.Reader`, `java.util.ArrayList`, `java.util.List`, `java.util.Objects`, `java.util.function.BiConsumer`, `java.util.logging.Level`, `java.util.logging.Logger`, `java.util.regex.MatchResult`, `java.util.regex.Matcher`, `java.util.regex.Pattern`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiffReader`, `com.github.difflib.unifieddiff.getName`, `com.github.difflib.unifieddiff.UnifiedDiffLine`, `com.github.difflib.unifieddiff.InternalUnifiedDiffReader`
- Fields in `com.github.difflib.unifieddiff.UnifiedDiffReader`: `Pattern UNIFIED_DIFF_CHUNK_REGEXP`, `Pattern TIMESTAMP_REGEXP`, `InternalUnifiedDiffReader READER`, `UnifiedDiff data`, `UnifiedDiffLine DIFF_COMMAND`, `UnifiedDiffLine SIMILARITY_INDEX`, `UnifiedDiffLine INDEX`, `UnifiedDiffLine FROM_FILE`, `UnifiedDiffLine TO_FILE`, `UnifiedDiffLine RENAME_FROM`, `UnifiedDiffLine RENAME_TO`, `UnifiedDiffLine COPY_FROM`, `UnifiedDiffLine COPY_TO`, `UnifiedDiffLine NEW_FILE_MODE`, `UnifiedDiffLine DELETED_FILE_MODE`, `UnifiedDiffLine OLD_MODE`, `UnifiedDiffLine NEW_MODE`, `UnifiedDiffLine BINARY_ADDED`, `UnifiedDiffLine BINARY_DELETED`, `UnifiedDiffLine BINARY_EDITED`, `UnifiedDiffLine CHUNK`, `UnifiedDiffLine LINE_NORMAL`, `UnifiedDiffLine LINE_DEL`, `UnifiedDiffLine LINE_ADD`, `UnifiedDiffFile actualFile`
- Constructors in `com.github.difflib.unifieddiff.UnifiedDiffReader`: `UnifiedDiffReader(1 params)`
- Methods in `com.github.difflib.unifieddiff.UnifiedDiffReader`: `instance UnifiedDiff parse(0 params)`, `instance String checkForNoNewLineAtTheEndOfTheFile(1 params)`, `static String[] parseFileNames(1 params)`
- Fields in `com.github.difflib.unifieddiff.UnifiedDiffLine`: `Pattern pattern`, `boolean stopsHeaderParsing`
- Constructors in `com.github.difflib.unifieddiff.UnifiedDiffLine`: `UnifiedDiffLine(3 params)`, `UnifiedDiffLine(4 params)`, `UnifiedDiffLine(4 params)`
- Methods in `com.github.difflib.unifieddiff.UnifiedDiffLine`: `instance boolean validLine(1 params)`, `instance boolean processLine(1 params)`, `instance boolean isStopsHeaderParsing(0 params)`, `instance String toString(0 params)`
- Fields in `com.github.difflib.unifieddiff.InternalUnifiedDiffReader`: `String lastLine`
- Constructors in `com.github.difflib.unifieddiff.InternalUnifiedDiffReader`: `InternalUnifiedDiffReader(1 params)`
- Methods in `com.github.difflib.unifieddiff.InternalUnifiedDiffReader`: `instance String readLine(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffWriter.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Imports: `com.github.difflib.patch.AbstractDelta`, `java.io.IOException`, `java.io.Writer`, `java.util.ArrayList`, `java.util.List`, `java.util.Objects`, `java.util.function.Consumer`, `java.util.function.Function`, `java.util.logging.Level`, `java.util.logging.Logger`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiffWriter`, `com.github.difflib.unifieddiff.getName`
- Fields in `com.github.difflib.unifieddiff.UnifiedDiffWriter`: `Logger LOG`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/main/java/com/github/difflib/unifieddiff/package-info.java`

- Source root: `.`
- Risk: `medium`
- Package: `com.github.difflib.unifieddiff`

### `java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib`
- Imports: `staticjava.util.stream.Collectors.toList`, `staticorg.assertj.core.api.Assertions.assertThat`, `staticorg.junit.jupiter.api.Assertions.assertEquals`, `staticorg.junit.jupiter.api.Assertions.assertNotNull`, `staticorg.junit.jupiter.api.Assertions.assertTrue`, `com.github.difflib.patch.AbstractDelta`, `com.github.difflib.patch.ChangeDelta`, `com.github.difflib.patch.Chunk`, `com.github.difflib.patch.DeleteDelta`, `com.github.difflib.patch.EqualDelta`, `com.github.difflib.patch.InsertDelta`, `com.github.difflib.patch.Patch`, `java.io.BufferedReader`, `java.io.IOException`, `java.io.InputStream`, `java.io.InputStreamReader`, `java.nio.charset.Charset`, `java.nio.charset.StandardCharsets`, `java.nio.file.Files`, `java.nio.file.Paths`, `java.util.ArrayList`, `java.util.Arrays`, `java.util.Collections`, `java.util.List`, `java.util.zip.ZipFile`, `org.junit.jupiter.api.Disabled`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.DiffUtilsTest`
- Constructors in `com.github.difflib.DiffUtilsTest`: `getDeltas(0 params)`, `get(0 params)`, `getDeltas(0 params)`, `get(0 params)`, `asList(0 params)`, `diff(1 params)`, `getDeltas(0 params)`, `get(0 params)`, `getDeltas(0 params)`, `get(0 params)`, `diff(1 params)`, `readStringListFromInputStream(1 params)`, `getDeltas(0 params)`, `get(0 params)`
- Methods in `com.github.difflib.DiffUtilsTest`: `instance void testDiff_Insert(0 params)`, `instance void testDiff_Delete(0 params)`, `instance void testDiff_Change(0 params)`, `instance void testDiff_EmptyList(0 params)`, `instance void testDiff_EmptyListWithNonEmpty(0 params)`, `instance void testDiffInline(0 params)`, `instance void testDiffInline2(0 params)`, `instance void testDiffIntegerList(0 params)`, `instance void testDiffMissesChangeForkDnaumenkoIssue31(0 params)`, `instance void testPossibleDiffHangOnLargeDatasetDnaumenkoIssue26(0 params)`, `instance void testDiffMyersExample1(0 params)`, `instance void testDiff_Equal(0 params)`, `instance void testDiff_InsertWithEqual(0 params)`, `instance void testDiff_ProblemIssue42(0 params)`, `instance void testDiffPatchIssue189Problem(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib`
- Imports: `staticjava.util.stream.Collectors.joining`, `staticorg.assertj.core.api.Assertions.assertThat`, `staticorg.junit.jupiter.api.Assertions.assertEquals`, `staticorg.junit.jupiter.api.Assertions.assertThrows`, `staticorg.junit.jupiter.api.Assertions.assertTrue`, `staticorg.junit.jupiter.api.Assertions.fail`, `com.github.difflib.patch.Chunk`, `com.github.difflib.patch.Patch`, `com.github.difflib.patch.PatchFailedException`, `java.io.BufferedReader`, `java.io.FileNotFoundException`, `java.io.FileReader`, `java.io.IOException`, `java.util.ArrayList`, `java.util.Arrays`, `java.util.Collections`, `java.util.List`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.GenerateUnifiedDiffTest`, `com.github.difflib.DiffUtils`
- Constructors in `com.github.difflib.GenerateUnifiedDiffTest`: `fileToLines(1 params)`, `fileToLines(0 params)`, `parseUnifiedDiff(1 params)`, `fileToLines(0 params)`, `fileToLines(0 params)`, `parseUnifiedDiff(1 params)`, `getDeltas(0 params)`, `get(1 params)`, `getTarget(0 params)`, `fileToLines(0 params)`, `parseUnifiedDiff(1 params)`
- Methods in `com.github.difflib.GenerateUnifiedDiffTest`: `instance void testGenerateUnified(0 params)`, `instance void testGenerateUnifiedWithOneDelta(0 params)`, `instance void testGenerateUnifiedDiffWithoutAnyDeltas(0 params)`, `instance void testDiff_Issue10(0 params)`, `instance void testPatchWithNoDeltas(0 params)`, `instance void testDiff5(0 params)`, `instance void testDiffWithHeaderLineInText(0 params)`, `instance void testNewFileCreation(0 params)`, `instance void testChangePosition(0 params)`, `instance void validateChangePosition(0 params)`, `instance void verify(0 params)`, `instance void testFailingPatchByException(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/TestConstants.java`

- Source root: `.`
- Risk: `medium`
- Package: `com.github.difflib`
- Classes: `com.github.difflib.TestConstants`
- Fields in `com.github.difflib.TestConstants`: `String BASE_FOLDER_RESOURCES`, `String MOCK_FOLDER`
- Constructors in `com.github.difflib.TestConstants`: `TestConstants(0 params)`

### `java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/MyersDiffTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm.myers`
- Imports: `staticorg.junit.jupiter.api.Assertions.assertEquals`, `staticorg.junit.jupiter.api.Assertions.assertNotNull`, `com.github.difflib.algorithm.DiffAlgorithmListener`, `com.github.difflib.patch.Patch`, `java.util.ArrayList`, `java.util.Arrays`, `java.util.List`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.algorithm.myers.MyersDiffTest`
- Constructors in `com.github.difflib.algorithm.myers.MyersDiffTest`: `generate(1 params)`, `generate(1 params)`
- Methods in `com.github.difflib.algorithm.myers.MyersDiffTest`: `instance void testDiffMyersExample1Forward(0 params)`, `instance void testDiffMyersExample1ForwardWithListener(0 params)`, `instance void diffStart(0 params)`, `instance void diffStep(2 params)`, `instance void diffEnd(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/MyersDiffWithLinearSpaceTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm.myers`
- Imports: `staticjava.util.stream.Collectors.toList`, `staticorg.junit.jupiter.api.Assertions.*`, `com.github.difflib.DiffUtils`, `com.github.difflib.algorithm.DiffAlgorithmListener`, `com.github.difflib.patch.Patch`, `java.util.ArrayList`, `java.util.Arrays`, `java.util.List`, `java.util.stream.IntStream`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.algorithm.myers.MyersDiffWithLinearSpaceTest`
- Constructors in `com.github.difflib.algorithm.myers.MyersDiffWithLinearSpaceTest`: `generate(1 params)`, `generate(1 params)`
- Methods in `com.github.difflib.algorithm.myers.MyersDiffWithLinearSpaceTest`: `instance void testDiffMyersExample1Forward(0 params)`, `instance void testDiffMyersExample1ForwardWithListener(0 params)`, `instance void diffStart(0 params)`, `instance void diffStep(2 params)`, `instance void diffEnd(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/WithMyersDiffWithLinearSpacePatchTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm.myers`
- Imports: `staticorg.junit.jupiter.api.Assertions.assertEquals`, `staticorg.junit.jupiter.api.Assertions.assertThrows`, `staticorg.junit.jupiter.api.Assertions.fail`, `com.github.difflib.DiffUtils`, `com.github.difflib.patch.*`, `java.io.ByteArrayInputStream`, `java.io.ByteArrayOutputStream`, `java.io.IOException`, `java.io.ObjectInputStream`, `java.io.ObjectOutputStream`, `java.util.ArrayList`, `java.util.Arrays`, `java.util.Collection`, `java.util.Comparator`, `java.util.List`, `java.util.stream.Collectors`, `java.util.stream.IntStream`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.algorithm.myers.WithMyersDiffWithLinearSpacePatchTest`, `com.github.difflib.algorithm.myers.FuzzyApplyTestPair`, `com.github.difflib.algorithm.myers.FuzzyApplyTestDataGenerator`
- Constructors in `com.github.difflib.algorithm.myers.WithMyersDiffWithLinearSpacePatchTest`: `asList(0 params)`, `diff(1 params)`, `asList(0 params)`, `diff(1 params)`, `asList(0 params)`, `diff(1 params)`, `intRange(1 params)`, `join(0 params)`
- Methods in `com.github.difflib.algorithm.myers.WithMyersDiffWithLinearSpacePatchTest`: `instance void testPatch_Insert(0 params)`, `instance void testPatch_Delete(0 params)`, `instance void testPatch_Change(0 params)`
- Fields in `com.github.difflib.algorithm.myers.FuzzyApplyTestPair`: `int requiredFuzz`
- Constructors in `com.github.difflib.algorithm.myers.FuzzyApplyTestPair`: `FuzzyApplyTestPair(0 params)`
- Methods in `com.github.difflib.algorithm.myers.FuzzyApplyTestPair`: `instance void fuzzyApply(0 params)`
- Fields in `com.github.difflib.algorithm.myers.FuzzyApplyTestDataGenerator`: `FuzzyApplyTestPair[] FUZZY_APPLY_TEST_PAIRS`
- Constructors in `com.github.difflib.algorithm.myers.FuzzyApplyTestDataGenerator`: `FuzzyApplyTestPair(5 params)`, `asList(0 params)`
- Methods in `com.github.difflib.algorithm.myers.FuzzyApplyTestDataGenerator`: `static String createList(0 params)`, `static void main(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/examples/ApplyPatch.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.examples`
- Imports: `com.github.difflib.DiffUtils`, `com.github.difflib.TestConstants`, `com.github.difflib.UnifiedDiffUtils`, `com.github.difflib.patch.Patch`, `com.github.difflib.patch.PatchFailedException`, `java.io.File`, `java.io.IOException`, `java.nio.file.Files`, `java.util.List`
- Classes: `com.github.difflib.examples.ApplyPatch`
- Fields in `com.github.difflib.examples.ApplyPatch`: `String ORIGINAL`, `String PATCH`
- Methods in `com.github.difflib.examples.ApplyPatch`: `static void main(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/examples/ComputeDifference.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.examples`
- Imports: `com.github.difflib.DiffUtils`, `com.github.difflib.TestConstants`, `com.github.difflib.patch.AbstractDelta`, `com.github.difflib.patch.Patch`, `java.io.File`, `java.io.IOException`, `java.nio.file.Files`, `java.util.List`
- Classes: `com.github.difflib.examples.ComputeDifference`
- Fields in `com.github.difflib.examples.ComputeDifference`: `String ORIGINAL`, `String REVISED`
- Methods in `com.github.difflib.examples.ComputeDifference`: `static void main(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.examples`
- Imports: `staticjava.util.stream.Collectors.joining`, `staticorg.junit.jupiter.api.Assertions.fail`, `com.github.difflib.TestConstants`, `com.github.difflib.UnifiedDiffUtils`, `java.io.BufferedReader`, `java.io.FileNotFoundException`, `java.io.FileReader`, `java.io.IOException`, `java.util.ArrayList`, `java.util.List`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.examples.OriginalAndDiffTest`
- Constructors in `com.github.difflib.examples.OriginalAndDiffTest`: `fileToLines(1 params)`
- Methods in `com.github.difflib.examples.OriginalAndDiffTest`: `instance void testGenerateOriginalAndDiff(0 params)`, `instance void testGenerateOriginalAndDiffFirstLineChange(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `staticorg.junit.jupiter.api.Assertions.assertEquals`, `java.util.List`, `java.util.stream.Collectors`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.patch.ChunkTest`
- Constructors in `com.github.difflib.patch.ChunkTest`: `toCharList(1 params)`
- Methods in `com.github.difflib.patch.ChunkTest`: `instance void verifyChunk(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `staticorg.junit.jupiter.api.Assertions.assertEquals`, `staticorg.junit.jupiter.api.Assertions.fail`, `com.github.difflib.DiffUtils`, `com.github.difflib.algorithm.DiffAlgorithmFactory`, `com.github.difflib.algorithm.myers.MyersDiff`, `com.github.difflib.algorithm.myers.MyersDiffWithLinearSpace`, `java.io.ByteArrayInputStream`, `java.io.ByteArrayOutputStream`, `java.io.IOException`, `java.io.ObjectInputStream`, `java.io.ObjectOutputStream`, `java.util.Arrays`, `java.util.List`, `java.util.stream.Stream`, `org.junit.jupiter.api.AfterAll`, `org.junit.jupiter.params.ParameterizedTest`, `org.junit.jupiter.params.provider.Arguments`, `org.junit.jupiter.params.provider.MethodSource`
- Classes: `com.github.difflib.patch.PatchWithAllDiffAlgorithmsTest`
- Constructors in `com.github.difflib.patch.PatchWithAllDiffAlgorithmsTest`: `provideAlgorithms(0 params)`, `asList(0 params)`, `asList(0 params)`, `diff(1 params)`, `asList(0 params)`, `asList(0 params)`, `diff(1 params)`, `asList(0 params)`, `asList(0 params)`, `diff(1 params)`, `asList(0 params)`, `asList(0 params)`, `diff(1 params)`
- Methods in `com.github.difflib.patch.PatchWithAllDiffAlgorithmsTest`: `static void afterAll(0 params)`, `instance ParameterizedTest MethodSource(0 params)`, `instance void testPatch_Insert(1 params)`, `instance ParameterizedTest MethodSource(0 params)`, `instance void testPatch_Delete(1 params)`, `instance ParameterizedTest MethodSource(0 params)`, `instance void testPatch_Change(1 params)`, `instance ParameterizedTest MethodSource(0 params)`, `instance void testPatch_Serializable(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithMyerDiffTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `staticcom.github.difflib.patch.Patch.CONFLICT_PRODUCES_MERGE_CONFLICT`, `staticjava.util.stream.Collectors.joining`, `staticorg.junit.jupiter.api.Assertions.assertEquals`, `staticorg.junit.jupiter.api.Assertions.fail`, `com.github.difflib.DiffUtils`, `java.util.Arrays`, `java.util.List`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.patch.PatchWithMyerDiffTest`
- Constructors in `com.github.difflib.patch.PatchWithMyerDiffTest`: `asList(0 params)`, `diff(1 params)`
- Methods in `com.github.difflib.patch.PatchWithMyerDiffTest`: `instance void testPatch_Change_withExceptionProcessor(0 params)`, `instance void testPatchThreeWayIssue138(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithMyerDiffWithLinearSpaceTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `staticorg.junit.jupiter.api.Assertions.assertEquals`, `staticorg.junit.jupiter.api.Assertions.fail`, `com.github.difflib.DiffUtils`, `com.github.difflib.algorithm.myers.MyersDiff`, `com.github.difflib.algorithm.myers.MyersDiffWithLinearSpace`, `java.util.Arrays`, `java.util.List`, `org.junit.jupiter.api.AfterAll`, `org.junit.jupiter.api.BeforeAll`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.patch.PatchWithMyerDiffWithLinearSpaceTest`
- Constructors in `com.github.difflib.patch.PatchWithMyerDiffWithLinearSpaceTest`: `asList(0 params)`, `diff(1 params)`
- Methods in `com.github.difflib.patch.PatchWithMyerDiffWithLinearSpaceTest`: `static void setupClass(0 params)`, `static void resetClass(0 params)`, `instance void testPatch_Change_withExceptionProcessor(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.text`
- Imports: `staticorg.junit.jupiter.api.Assertions.assertEquals`, `staticorg.junit.jupiter.api.Assertions.assertTrue`, `java.util.Arrays`, `java.util.List`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.text.DiffRowGeneratorEqualitiesTest`
- Methods in `com.github.difflib.text.DiffRowGeneratorEqualitiesTest`: `instance void testDefaultEqualityProcessingLeavesTextUnchanged(0 params)`, `instance void testCustomEqualityProcessingIsApplied(0 params)`, `instance void testHtmlEscapingEqualitiesWorksWithDefaultNormalizer(0 params)`, `instance void testEqualitiesProcessedButInlineDiffStillPresent(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.text`
- Imports: `staticjava.util.stream.Collectors.joining`, `staticjava.util.stream.Collectors.toList`, `staticorg.assertj.core.api.Assertions.assertThat`, `staticorg.junit.jupiter.api.Assertions.assertEquals`, `staticorg.junit.jupiter.api.Assertions.assertTrue`, `com.github.difflib.DiffUtils`, `com.github.difflib.algorithm.myers.MyersDiffWithLinearSpace`, `com.github.difflib.patch.AbstractDelta`, `com.github.difflib.text.deltamerge.DeltaMergeUtils`, `com.github.difflib.text.deltamerge.InlineDeltaMergeInfo`, `java.io.File`, `java.io.IOException`, `java.net.URISyntaxException`, `java.nio.file.FileSystem`, `java.nio.file.FileSystems`, `java.nio.file.Files`, `java.nio.file.Paths`, `java.util.Arrays`, `java.util.Collections`, `java.util.List`, `java.util.function.Function`, `java.util.regex.Pattern`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.text.DiffRowGeneratorTest`
- Fields in `com.github.difflib.text.DiffRowGeneratorTest`: `DiffRowGenerator generator`
- Constructors in `com.github.difflib.text.DiffRowGeneratorTest`: `split(1 params)`, `assertInlineDiffResult(4 params)`
- Methods in `com.github.difflib.text.DiffRowGeneratorTest`: `instance void testGenerator_Default(0 params)`, `instance void testNormalize_List(0 params)`, `instance void testGenerator_Default2(0 params)`, `instance void testGenerator_InlineDiff(0 params)`, `instance void testGenerator_IgnoreWhitespaces(0 params)`, `instance void print(1 params)`, `instance void testGeneratorWithWordWrap(0 params)`, `instance void testGeneratorWithMerge(0 params)`, `instance void testGeneratorWithMerge2(0 params)`, `instance void testGeneratorWithMerge3(0 params)`, `instance void testGeneratorWithMergeByWord4(0 params)`, `instance void testGeneratorWithMergeByWord5(0 params)`, `instance void testSplitString(0 params)`, `instance void testSplitString2(0 params)`, `instance void testSplitString3(0 params)`, `instance void testGeneratorExample1(0 params)`, `instance void testGeneratorExample2(0 params)`, `instance void testGeneratorUnchanged(0 params)`, `instance void testGeneratorIssue14(0 params)`, `instance void testGeneratorIssue15(0 params)`, `instance void testGeneratorIssue22(0 params)`, `instance void testGeneratorIssue22_2(0 params)`, `instance void testGeneratorIssue22_3(0 params)`, `instance void testGeneratorIssue41DefaultNormalizer(0 params)`, `instance void testGeneratorIssue41UserNormalizer(0 params)`, `instance void testGenerationIssue44reportLinesUnchangedProblem(0 params)`, `instance void testIgnoreWhitespaceIssue66(0 params)`, `instance void testIgnoreWhitespaceIssue66_2(0 params)`, `instance void testIgnoreWhitespaceIssue64(0 params)`, `instance void testReplaceDiffsIssue63(0 params)`, `instance void testProblemTooManyDiffRowsIssue65(0 params)`, `instance void testProblemTooManyDiffRowsIssue65_NoMerge(0 params)`, `instance void testProblemTooManyDiffRowsIssue65_DiffByWord(0 params)`, `instance void testProblemTooManyDiffRowsIssue65_NoInlineDiff(0 params)`, `instance void testLinefeedInStandardTagsWithLineWidthIssue81(0 params)`, `instance void testIssue86WrongInlineDiff(0 params)`, `instance void testCorrectChangeIssue114(0 params)`, `instance void testCorrectChangeIssue114_2(0 params)`, `instance void testIssue119WrongContextLength(0 params)`, `instance void testIssue129WithDeltaDecompression(0 params)`, `instance void testIssue129SkipDeltaDecompression(0 params)`, `instance void testIssue129SkipWhitespaceChanges(0 params)`, `instance void testGeneratorWithWhitespaceDeltaMerge(0 params)`, `instance void testGeneratorWithMergingDeltasForShortEqualities(0 params)`, `instance void assertInlineDiffResult(4 params)`, `instance void testIssue188HangOnExamples(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/text/StringUtilsTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.text`
- Imports: `staticorg.junit.jupiter.api.Assertions.assertEquals`, `org.junit.jupiter.api.Assertions`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.text.StringUtilsTest`, `com.github.difflib.text.StringUtils`
- Methods in `com.github.difflib.text.StringUtilsTest`: `instance void testHtmlEntites(0 params)`, `instance void testNormalize_String(0 params)`, `instance void testWrapText_String_int(0 params)`, `instance void testWrapText_String_int_zero(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Imports: `staticorg.assertj.core.api.Assertions.assertThat`, `staticorg.junit.jupiter.api.Assertions.assertEquals`, `staticorg.junit.jupiter.api.Assertions.assertNull`, `staticorg.junit.jupiter.api.Assertions.assertTrue`, `com.github.difflib.patch.AbstractDelta`, `java.io.IOException`, `java.util.regex.Matcher`, `java.util.regex.Pattern`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiffReaderTest`
- Fields in `com.github.difflib.unifieddiff.UnifiedDiffReaderTest`: `UnifiedDiffFile file1`, `UnifiedDiffFile file1`, `UnifiedDiffFile file`, `UnifiedDiffFile file`, `UnifiedDiffFile file`, `UnifiedDiffFile file1`, `UnifiedDiffFile file1`, `UnifiedDiffFile file1`, `UnifiedDiffFile file1`, `UnifiedDiffFile file1`, `UnifiedDiffFile file1`
- Constructors in `com.github.difflib.unifieddiff.UnifiedDiffReaderTest`: `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `println(1 params)`, `getResourceAsStream(0 params)`, `println(1 params)`, `getResourceAsStream(0 params)`, `println(1 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `println(1 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `assertEquals(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `assertEquals(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `assertThat(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `assertThat(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `assertThat(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `size(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `isEqualTo(0 params)`, `getResourceAsStream(0 params)`, `assertThat(0 params)`, `isEqualTo(0 params)`
- Methods in `com.github.difflib.unifieddiff.UnifiedDiffReaderTest`: `instance void testSimpleParse(0 params)`, `instance void testParseDiffBlock(0 params)`, `instance void testChunkHeaderParsing(0 params)`, `instance void testChunkHeaderParsing2(0 params)`, `instance void testChunkHeaderParsing3(0 params)`, `instance void testSimpleParse2(0 params)`, `instance void testParseIssue201(0 params)`, `instance void testSimplePattern(0 params)`, `instance void testParseIssue46(0 params)`, `instance void testParseIssue33(0 params)`, `instance void testParseIssue51(0 params)`, `instance void testParseIssue79(0 params)`, `instance void testParseIssue84(0 params)`, `instance void testParseIssue85(0 params)`, `instance void testTimeStampRegexp(0 params)`, `instance void testParseIssue98(0 params)`, `instance void testParseIssue104(0 params)`, `instance void testParseIssue107BazelDiff(0 params)`, `instance void testParseIssue107_2(0 params)`, `instance void testParseIssue107_3(0 params)`, `instance void testParseIssue107_4(0 params)`, `instance void testParseIssue107_5(0 params)`, `instance void testParseIssue110(0 params)`, `instance void testParseIssue117(0 params)`, `instance void testParseIssue122(0 params)`, `instance void testParseIssue123(0 params)`, `instance void testParseIssue141(0 params)`, `instance void testParseIssue182add(0 params)`, `instance void testParseIssue182delete(0 params)`, `instance void testParseIssue182edit(0 params)`, `instance void testParseIssue182mode(0 params)`, `instance void testParseIssue193Copy(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripNewLineTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Imports: `staticjava.util.stream.Collectors.joining`, `staticorg.junit.jupiter.api.Assertions.assertEquals`, `com.github.difflib.patch.PatchFailedException`, `java.io.ByteArrayInputStream`, `java.io.IOException`, `java.util.Arrays`, `org.junit.jupiter.api.Disabled`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiffRoundTripNewLineTest`
- Methods in `com.github.difflib.unifieddiff.UnifiedDiffRoundTripNewLineTest`: `instance void testIssue135MissingNoNewLineInPatched(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Imports: `staticjava.util.stream.Collectors.joining`, `staticorg.junit.jupiter.api.Assertions.assertEquals`, `staticorg.junit.jupiter.api.Assertions.fail`, `com.github.difflib.DiffUtils`, `com.github.difflib.TestConstants`, `com.github.difflib.patch.Patch`, `com.github.difflib.patch.PatchFailedException`, `java.io.BufferedReader`, `java.io.ByteArrayInputStream`, `java.io.FileNotFoundException`, `java.io.FileReader`, `java.io.IOException`, `java.io.StringWriter`, `java.util.ArrayList`, `java.util.Arrays`, `java.util.List`, `org.junit.jupiter.api.Disabled`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiffRoundTripTest`
- Constructors in `com.github.difflib.unifieddiff.UnifiedDiffRoundTripTest`: `fileToLines(1 params)`, `fileToLines(0 params)`, `getFiles(0 params)`, `get(0 params)`, `getPatch(0 params)`, `fileToLines(0 params)`, `fileToLines(0 params)`
- Methods in `com.github.difflib.unifieddiff.UnifiedDiffRoundTripTest`: `instance void testGenerateUnified(0 params)`, `instance void testGenerateUnifiedWithOneDelta(0 params)`, `instance void testGenerateUnifiedDiffWithoutAnyDeltas(0 params)`, `instance void testDiff_Issue10(0 params)`, `instance void testPatchWithNoDeltas(0 params)`, `instance void testDiff5(0 params)`, `instance void testDiffWithHeaderLineInText(0 params)`, `instance void verify(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffWriterTest.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Imports: `staticorg.junit.jupiter.api.Assertions.assertEquals`, `com.github.difflib.DiffUtils`, `com.github.difflib.patch.Patch`, `java.io.ByteArrayInputStream`, `java.io.IOException`, `java.io.StringWriter`, `java.net.URI`, `java.net.URISyntaxException`, `java.nio.charset.Charset`, `java.nio.file.Files`, `java.nio.file.Paths`, `java.util.ArrayList`, `java.util.Collections`, `java.util.List`, `org.junit.jupiter.api.Test`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiffWriterTest`
- Fields in `com.github.difflib.unifieddiff.UnifiedDiffWriterTest`: `UnifiedDiff diff`, `StringWriter writer`
- Constructors in `com.github.difflib.unifieddiff.UnifiedDiffWriterTest`: `UnifiedDiffWriterTest(0 params)`, `getResource(0 params)`, `toURI(0 params)`, `defaultCharset(0 params)`, `write(1 params)`
- Methods in `com.github.difflib.unifieddiff.UnifiedDiffWriterTest`: `instance void testWrite(0 params)`, `instance void testWriteWithNewFile(0 params)`, `static String readFile(2 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

## Recommended Migration Order

- `java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmListener.java` (medium risk)
- `java-diff-utils/src/main/java/com/github/difflib/unifieddiff/package-info.java` (medium risk)
- `java-diff-utils/src/test/java/com/github/difflib/TestConstants.java` (medium risk)
- `java-diff-utils-jgit/src/main/java/com/github/difflib/algorithm/jgit/HistogramDiff.java` (high risk)
- `java-diff-utils-jgit/src/test/java/com/github/difflib/algorithm/jgit/HistogramDiffTest.java` (high risk)
- `java-diff-utils-jgit/src/test/java/com/github/difflib/algorithm/jgit/LRHistogramDiffTest.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/UnifiedDiffUtils.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/algorithm/Change.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmI.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/MyersDiff.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/MyersDiffWithLinearSpace.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/patch/ConflictOutput.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/patch/DeltaType.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/patch/DiffException.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/patch/EqualDelta.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/patch/PatchFailedException.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/patch/VerifyChunk.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/text/DiffRowGenerator.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/text/StringUtils.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/InlineDeltaMergeInfo.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiff.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffFile.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffParserException.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffReader.java` (high risk)
- `java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffWriter.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/MyersDiffTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/MyersDiffWithLinearSpaceTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/WithMyersDiffWithLinearSpacePatchTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/examples/ApplyPatch.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/examples/ComputeDifference.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithMyerDiffTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithMyerDiffWithLinearSpaceTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/text/StringUtilsTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripNewLineTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripTest.java` (high risk)
- `java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffWriterTest.java` (high risk)

## Migration Execution Summary

| Metric | Count |
| --- | ---: |
| Java files | 55 |
| Generated files | 3 |
| Skipped files | 52 |
| Diagnostics | 690 |
| Dry run | false |

### Migration Diagnostics

- `.corpus/java-diff-utils/java-diff-utils-jgit/src/main/java/com/github/difflib/algorithm/jgit/HistogramDiff.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmI`
- `.corpus/java-diff-utils/java-diff-utils-jgit/src/main/java/com/github/difflib/algorithm/jgit/HistogramDiff.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmListener`
- `.corpus/java-diff-utils/java-diff-utils-jgit/src/main/java/com/github/difflib/algorithm/jgit/HistogramDiff.java:0:0: JTG2005: import unresolved: com.github.difflib.patch.DeltaType`
- `.corpus/java-diff-utils/java-diff-utils-jgit/src/test/java/com/github/difflib/algorithm/jgit/HistogramDiffTest.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmListener`
- `.corpus/java-diff-utils/java-diff-utils-jgit/src/test/java/com/github/difflib/algorithm/jgit/HistogramDiffTest.java:35:2: JTG2001: duplicate Go symbol after migration: NewHistogramDiffTest`
- `.corpus/java-diff-utils/java-diff-utils-jgit/src/test/java/com/github/difflib/algorithm/jgit/HistogramDiffTest.java:45:14: JTG2001: duplicate Go symbol after migration: NewHistogramDiffTest`
- `.corpus/java-diff-utils/java-diff-utils-jgit/src/test/java/com/github/difflib/algorithm/jgit/HistogramDiffTest.java:63:38: JTG2001: duplicate Go symbol after migration: NewHistogramDiffTest`
- `.corpus/java-diff-utils/java-diff-utils-jgit/src/test/java/com/github/difflib/algorithm/jgit/LRHistogramDiffTest.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmListener`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmFactory`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmI`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmListener`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:60:29: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:73:29: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:86:29: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:98:30: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:114:29: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:122:29: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:142:29: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:165:29: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:180:30: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:205:28: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:217:28: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:221:30: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:228:2: JTG2001: duplicate Go symbol after migration: NewDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/DiffUtils.java:35:0: JTG2001: duplicate Go symbol after migration: DiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/UnifiedDiffUtils.java:47:30: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/UnifiedDiffUtils.java:144:29: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/UnifiedDiffUtils.java:212:30: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/UnifiedDiffUtils.java:304:30: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/UnifiedDiffUtils.java:315:2: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/UnifiedDiffUtils.java:330:29: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/UnifiedDiffUtils.java:348:29: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/UnifiedDiffUtils.java:373:30: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/UnifiedDiffUtils.java:437:38: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/UnifiedDiffUtils.java:453:30: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/Change.java:0:0: JTG2005: import unresolved: com.github.difflib.patch.DeltaType`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:26:2: extraneous input '<' expecting {';', '}', 'public', 'private', 'protected', 'static', 'final', 'void', 'boolean', 'int', 'double', 'String', Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:26:4: extraneous input '>' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:26:20: mismatched input '<' expecting '('`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:26:22: extraneous input '>' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:28:2: extraneous input '<' expecting {';', '}', 'public', 'private', 'protected', 'static', 'final', 'void', 'boolean', 'int', 'double', 'String', Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:28:4: extraneous input '>' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:28:20: mismatched input '<' expecting '('`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:28:22: extraneous input '>' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:28:43: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:28:42: extraneous input '<' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:28:51: mismatched input 'T' expecting ')'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:28:54: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:28:52: extraneous input ',' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:28:62: missing '(' at 'T'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmFactory.java:28:63: extraneous input '>' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmI.java:27:31: mismatched input '<' expecting '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmI.java:37:32: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmI.java:37:58: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmI.java:47:2: extraneous input 'default' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/DiffAlgorithmI.java:49:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/MyersDiff.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmFactory`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/MyersDiff.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmI`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/MyersDiff.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmListener`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/MyersDiff.java:0:0: JTG2005: import unresolved: com.github.difflib.patch.DeltaType`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/MyersDiffWithLinearSpace.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmFactory`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/MyersDiffWithLinearSpace.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmI`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/MyersDiffWithLinearSpace.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmListener`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/MyersDiffWithLinearSpace.java:0:0: JTG2005: import unresolved: com.github.difflib.patch.DeltaType`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java:50:11: no viable alternative at input 'this.i='`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java:51:11: no viable alternative at input 'this.j='`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java:52:19: no viable alternative at input 'this.bootstrap='`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java:54:16: no viable alternative at input 'this.prev='`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java:56:16: no viable alternative at input 'this.prev='`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java:56:31: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java:56:38: token recognition error at: ':'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java:56:33: missing ';' at 'null'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java:56:40: missing ';' at 'prev'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java:58:15: no viable alternative at input 'this.snake='`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java:95:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/algorithm/myers/PathNode.java:96:2: no viable alternative at input 'Overridepublic'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:26:7: no viable alternative at input 'publicabstract'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:26:35: mismatched input '<' expecting '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:28:16: no viable alternative at input 'privatefinalChunk'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:29:16: no viable alternative at input 'privatefinalDeltaType'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:31:9: no viable alternative at input 'publicAbstractDelta'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:33:4: extraneous input 'Objects' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:34:4: extraneous input 'Objects' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:35:4: extraneous input 'this' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:36:4: extraneous input 'this' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:37:4: extraneous input 'this' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:38:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:42:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:46:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:50:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:59:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:63:4: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:65:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:67:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:71:12: no viable alternative at input 'protectedabstract'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:81:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:81:3: extraneous input 'SuppressWarnings' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:85:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:92:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:92:3: extraneous input 'Override' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:95:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:97:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:101:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:104:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:107:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:108:24: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:108:50: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:109:4: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:111:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:114:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/AbstractDelta.java:116:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:27:30: mismatched input '<' expecting '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:37:4: extraneous input 'Objects' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:38:4: extraneous input 'Objects' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:39:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:41:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:44:4: extraneous input 'int' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:45:4: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:45:20: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:45:30: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:47:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:49:4: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:49:16: token recognition error at: ':'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:51:6: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:52:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:55:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:58:4: extraneous input 'int' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:59:4: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:59:20: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:59:30: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:61:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:63:4: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:63:16: token recognition error at: ':'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:65:6: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:66:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:71:4: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:71:23: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:71:40: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:73:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:76:4: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:76:16: token recognition error at: ':'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:78:6: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:79:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:82:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:86:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:88:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ChangeDelta.java:91:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:38:24: mismatched input '<' expecting '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:41:10: no viable alternative at input 'privateList'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:42:16: no viable alternative at input 'privatefinalList'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:51:9: no viable alternative at input 'publicChunk'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:53:4: extraneous input 'this' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:54:4: extraneous input 'this' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:54:49: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:54:83: token recognition error at: ':'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:55:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:65:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:76:4: extraneous input 'this' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:77:4: extraneous input 'this' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:77:49: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:77:83: token recognition error at: ':'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:78:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:88:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:99:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:113:4: extraneous input 'int' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:114:4: extraneous input 'int' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:116:4: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:118:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:119:29: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:119:44: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:122:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:125:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:132:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:136:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:143:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:150:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:154:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:161:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:163:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:166:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:168:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:172:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:175:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:178:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:179:10: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:179:28: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:180:4: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:183:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:186:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:188:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:190:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Chunk.java:193:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ConflictOutput.java:29:0: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ConflictOutput.java:29:1: extraneous input 'FunctionalInterface' expecting {<EOF>, 'package', ';', 'import', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ConflictOutput.java:30:31: mismatched input '<' expecting '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/ConflictOutput.java:34:0: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:26:30: mismatched input '<' expecting '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:36:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:38:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:41:4: extraneous input 'int' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:42:4: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:42:20: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:42:30: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:44:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:47:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:50:4: extraneous input 'List' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:51:4: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:51:20: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:51:38: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:53:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:56:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:60:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:62:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeleteDelta.java:65:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DeltaType.java:33:7: no viable alternative at input 'publicenum'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DiffException.java:27:2: JTG2001: duplicate Go symbol after migration: NewDiffException`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/DiffException.java:29:2: JTG2001: duplicate Go symbol after migration: NewDiffException`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/EqualDelta.java:24:23: mismatched input '<' expecting '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/EqualDelta.java:28:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/EqualDelta.java:30:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/EqualDelta.java:33:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/EqualDelta.java:39:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/EqualDelta.java:44:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/EqualDelta.java:48:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/EqualDelta.java:50:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/EqualDelta.java:53:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:26:30: mismatched input '<' expecting '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:36:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:38:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:41:4: extraneous input 'List' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:42:4: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:42:20: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:42:38: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:44:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:47:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:50:4: extraneous input 'int' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:51:4: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:51:20: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:51:30: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:53:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:56:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:60:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:62:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/InsertDelta.java:65:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:38:24: mismatched input '<' expecting '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:42:9: no viable alternative at input 'publicPatch'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:44:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:48:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:57:30: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:59:4: extraneous input 'applyToExisting' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:60:4: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:61:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:73:4: extraneous input 'while' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:75:6: extraneous input 'VerifyChunk' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:76:6: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:78:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:82:43: mismatched input '<' expecting '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:84:17: no viable alternative at input 'publicfinalint'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:87:11: no viable alternative at input 'publicint'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:90:11: no viable alternative at input 'publicint'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:92:11: no viable alternative at input 'publicint'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:93:11: no viable alternative at input 'publicboolean'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:94:11: no viable alternative at input 'publicboolean'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:96:12: no viable alternative at input 'privatePatchApplyingContext'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:98:6: extraneous input 'this' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:99:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:106:4: extraneous input 'int' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:108:4: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:108:32: token recognition error at: ':'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:110:6: extraneous input 'int' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:111:6: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:113:8: extraneous input 'lastPatchDelta' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:114:8: extraneous input 'ctx' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:115:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:117:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:121:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:125:23: extraneous input 'fuzz' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:125:44: extraneous input 'fuzz' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:127:6: extraneous input 'int' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:128:6: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:130:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:133:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:140:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:143:4: extraneous input 'ctx' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:147:4: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:147:28: extraneous input 'moreDelta' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:147:44: extraneous input 'moreDelta' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:149:6: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:151:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:154:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:158:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:167:6: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:169:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:175:6: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:177:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:182:6: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:184:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:188:6: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:190:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:193:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:201:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:206:22: no viable alternative at input 'publicstaticfinalConflictOutput'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:211:10: extraneous input 'for' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:211:26: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:211:56: extraneous input 'i' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:213:12: extraneous input 'result' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:214:10: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:217:10: extraneous input 'orgData' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:218:10: extraneous input 'orgData' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:219:10: extraneous input 'orgData' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:221:10: extraneous input 'result' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:223:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:226:8: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:229:10: no viable alternative at input 'privateConflictOutput'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:235:9: no viable alternative at input 'publicPatch'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:237:4: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:238:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:247:30: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:249:4: extraneous input 'restoreToExisting' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:250:4: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:251:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:263:4: extraneous input 'while' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:265:6: extraneous input 'delta' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:266:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:276:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:285:4: extraneous input 'return' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:286:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:288:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:290:43: token recognition error at: '''`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:290:45: token recognition error at: '''`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:291:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:295:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:297:66: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:299:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:302:11: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:302:39: token recognition error at: '?'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:304:4: extraneous input 'int' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:305:4: extraneous input 'int' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:307:4: extraneous input 'List' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:309:4: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:311:6: extraneous input 'Collections' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:312:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:314:23: token recognition error at: ':'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:320:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:323:6: extraneous input 'Chunk' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:324:6: extraneous input 'switch' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:325:19: token recognition error at: ':'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:327:10: extraneous input 'break' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:328:8: extraneous input 'case' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:328:19: token recognition error at: ':'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:330:10: extraneous input 'break' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:331:8: extraneous input 'case' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:331:19: token recognition error at: ':'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:333:10: extraneous input 'break' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:334:8: extraneous input 'default' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:334:15: token recognition error at: ':'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:338:6: extraneous input 'startRevised' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:339:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:345:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/Patch.java:348:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/PatchFailedException.java:27:2: JTG2001: duplicate Go symbol after migration: NewPatchFailedException`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/PatchFailedException.java:29:2: JTG2001: duplicate Go symbol after migration: NewPatchFailedException`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/patch/VerifyChunk.java:22:7: no viable alternative at input 'publicenum'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:26:27: missing '{' at 'implements'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:33:13: no viable alternative at input 'this.tag='`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:34:17: no viable alternative at input 'this.oldLine='`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:35:17: no viable alternative at input 'this.newLine='`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:38:18: no viable alternative at input 'publicenumTag{'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:48:9: no viable alternative at input 'publicTag'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:50:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:57:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:64:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:71:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:73:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:76:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:78:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:82:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:85:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:88:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:90:4: extraneous input 'if' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:93:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:96:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:100:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:103:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:107:6: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:110:4: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:112:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:114:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRow.java:117:2: extraneous input '}' expecting {<EOF>, ';', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRowGenerator.java:597:4: JTG2004: overload not supported: method com.github.difflib.text.Builder.oldTag`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRowGenerator.java:619:4: JTG2004: overload not supported: method com.github.difflib.text.Builder.newTag`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRowGenerator.java:0:0: JTG2005: import unresolved: com.github.difflib.patch.DeltaType`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRowGenerator.java:0:0: JTG2005: import unresolved: com.github.difflib.text.DiffRow.Tag`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRowGenerator.java:58:29: JTG2001: duplicate Go symbol after migration: NewDiffRowGenerator`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRowGenerator.java:58:56: JTG2001: duplicate Go symbol after migration: NewDiffRowGenerator`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRowGenerator.java:66:45: JTG2001: duplicate Go symbol after migration: NewDiffRowGenerator`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/DiffRowGenerator.java:67:36: JTG2001: duplicate Go symbol after migration: NewDiffRowGenerator`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/StringUtils.java:22:0: JTG2001: duplicate Go symbol after migration: StringUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/StringUtils.java:38:29: JTG2001: duplicate Go symbol after migration: NewStringUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/StringUtils.java:79:2: JTG2001: duplicate Go symbol after migration: NewStringUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:32:44: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:40:14: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:41:39: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:43:64: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:46:11: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:47:11: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:48:40: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:49:34: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:49:46: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:50:27: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:50:39: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:52:26: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:61:23: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:61:56: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:62:23: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:66:38: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:66:50: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:67:38: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:67:50: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:69:18: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/DeltaMergeUtils.java:72:18: JTG2001: duplicate Go symbol after migration: NewDeltaMergeUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/InlineDeltaMergeInfo.java:33:2: JTG2001: duplicate Go symbol after migration: NewInlineDeltaMergeInfo`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/InlineDeltaMergeInfo.java:39:37: JTG2001: duplicate Go symbol after migration: NewInlineDeltaMergeInfo`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/InlineDeltaMergeInfo.java:43:22: JTG2001: duplicate Go symbol after migration: NewInlineDeltaMergeInfo`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/text/deltamerge/InlineDeltaMergeInfo.java:47:22: JTG2001: duplicate Go symbol after migration: NewInlineDeltaMergeInfo`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiff.java:71:2: JTG2001: duplicate Go symbol after migration: From`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiff.java:46:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiff`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiff.java:58:22: JTG2001: duplicate Go symbol after migration: NewUnifiedDiff`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffFile.java:140:2: JTG2001: duplicate Go symbol after migration: From`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffParserException.java:24:2: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffParserException`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffParserException.java:26:2: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffParserException`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffParserException.java:30:2: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffParserException`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffParserException.java:34:2: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffParserException`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffParserException.java:38:2: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffParserException`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffReader.java:469:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffLine`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffReader.java:473:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffLine`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffReader.java:479:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffLine`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffReader.java:227:71: JTG2001: duplicate Go symbol after migration: GetName`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffWriter.java:35:71: JTG2001: duplicate symbol: class com.github.difflib.unifieddiff.getName`
- `.corpus/java-diff-utils/java-diff-utils/src/main/java/com/github/difflib/unifieddiff/UnifiedDiffWriter.java:35:71: JTG2001: duplicate Go symbol after migration: GetName`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:39:46: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:39:58: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:50:46: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:50:58: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:59:46: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:61:42: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:64:46: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:64:58: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:82:46: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:82:58: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:114:43: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:156:29: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:181:46: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/DiffUtilsTest.java:181:58: JTG2001: duplicate Go symbol after migration: NewDiffUtilsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:211:38: JTG2001: duplicate symbol: class com.github.difflib.DiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:25:29: JTG2001: duplicate Go symbol after migration: NewGenerateUnifiedDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:70:36: JTG2001: duplicate Go symbol after migration: NewGenerateUnifiedDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:71:45: JTG2001: duplicate Go symbol after migration: NewGenerateUnifiedDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:85:32: JTG2001: duplicate Go symbol after migration: NewGenerateUnifiedDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:92:32: JTG2001: duplicate Go symbol after migration: NewGenerateUnifiedDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:147:49: JTG2001: duplicate Go symbol after migration: NewGenerateUnifiedDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:167:44: JTG2001: duplicate Go symbol after migration: NewGenerateUnifiedDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:167:56: JTG2001: duplicate Go symbol after migration: NewGenerateUnifiedDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:167:67: JTG2001: duplicate Go symbol after migration: NewGenerateUnifiedDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:205:36: JTG2001: duplicate Go symbol after migration: NewGenerateUnifiedDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:206:45: JTG2001: duplicate Go symbol after migration: NewGenerateUnifiedDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/GenerateUnifiedDiffTest.java:211:38: JTG2001: duplicate Go symbol after migration: DiffUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/MyersDiffTest.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmListener`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/MyersDiffTest.java:39:14: JTG2001: duplicate Go symbol after migration: NewMyersDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/MyersDiffTest.java:53:38: JTG2001: duplicate Go symbol after migration: NewMyersDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/MyersDiffWithLinearSpaceTest.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmListener`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/MyersDiffWithLinearSpaceTest.java:40:38: JTG2001: duplicate Go symbol after migration: NewMyersDiffWithLinearSpaceTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/MyersDiffWithLinearSpaceTest.java:56:38: JTG2001: duplicate Go symbol after migration: NewMyersDiffWithLinearSpaceTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/WithMyersDiffWithLinearSpacePatchTest.java:266:8: JTG2001: duplicate Go symbol after migration: NewFuzzyApplyTestDataGenerator`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/WithMyersDiffWithLinearSpacePatchTest.java:268:15: JTG2001: duplicate Go symbol after migration: NewFuzzyApplyTestDataGenerator`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/WithMyersDiffWithLinearSpacePatchTest.java:28:46: JTG2001: duplicate Go symbol after migration: NewWithMyersDiffWithLinearSpacePatchTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/WithMyersDiffWithLinearSpacePatchTest.java:31:18: JTG2001: duplicate Go symbol after migration: NewWithMyersDiffWithLinearSpacePatchTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/WithMyersDiffWithLinearSpacePatchTest.java:42:46: JTG2001: duplicate Go symbol after migration: NewWithMyersDiffWithLinearSpacePatchTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/WithMyersDiffWithLinearSpacePatchTest.java:45:18: JTG2001: duplicate Go symbol after migration: NewWithMyersDiffWithLinearSpacePatchTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/WithMyersDiffWithLinearSpacePatchTest.java:56:46: JTG2001: duplicate Go symbol after migration: NewWithMyersDiffWithLinearSpacePatchTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/WithMyersDiffWithLinearSpacePatchTest.java:59:18: JTG2001: duplicate Go symbol after migration: NewWithMyersDiffWithLinearSpacePatchTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/WithMyersDiffWithLinearSpacePatchTest.java:69:23: JTG2001: duplicate Go symbol after migration: NewWithMyersDiffWithLinearSpacePatchTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/algorithm/myers/WithMyersDiffWithLinearSpacePatchTest.java:74:29: JTG2001: duplicate Go symbol after migration: NewWithMyersDiffWithLinearSpacePatchTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/ApplyPatch.java:18:2: JTG2001: duplicate Go symbol after migration: main`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/ComputeDifference.java:17:2: JTG2001: duplicate Go symbol after migration: main`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:18:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:19:2: no viable alternative at input 'Testpublic'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:20:9: mismatched input 'String' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:20:15: extraneous input '>' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:21:9: mismatched input 'String' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:21:15: extraneous input '>' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:22:8: no viable alternative at input 'try{'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:25:25: extraneous input 'e' expecting ')'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:25:28: missing ';' at '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:29:9: mismatched input 'String' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:29:15: extraneous input '>' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:33:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:34:2: no viable alternative at input 'Testpublic'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:35:9: mismatched input 'String' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:35:15: extraneous input '>' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:36:9: mismatched input 'String' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:36:15: extraneous input '>' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:37:8: no viable alternative at input 'try{'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:40:25: extraneous input 'e' expecting ')'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:40:28: missing ';' at '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:44:9: mismatched input 'String' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:44:15: extraneous input '>' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:48:20: no viable alternative at input 'publicstaticList<'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:48:58: missing '{' at 'throws'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:48:100: missing ';' at '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:49:9: mismatched input 'String' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:49:15: extraneous input '>' expecting Identifier`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:49:38: mismatched input '<' expecting '('`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:49:39: extraneous input '>' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:49:41: mismatched input ')' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:51:24: mismatched input 'in' expecting ')'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:51:73: mismatched input ')' expecting ';'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/examples/OriginalAndDiffTest.java:59:0: extraneous input '<EOF>' expecting {';', '}', 'public', 'private', 'protected', 'static', 'final', 'void', 'boolean', 'int', 'double', 'String', Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java:10:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java:11:2: no viable alternative at input 'Testvoid'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java:11:21: missing '{' at 'throws'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java:11:49: missing ';' at '{'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java:12:38: mismatched input '<' expecting '('`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java:12:39: extraneous input '>' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java:12:42: mismatched input ',' expecting ')'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java:12:62: extraneous input ')' expecting ';'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java:34:2: extraneous input 'private' expecting {';', '{', '}', '(', 'boolean', 'int', 'double', 'String', 'if', 'while', 'for', 'return', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java:35:35: extraneous input '>' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/ChunkTest.java:35:44: extraneous input 'x' expecting ')'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:50:3: JTG2001: duplicate symbol: method com.github.difflib.patch.PatchWithAllDiffAlgorithmsTest.MethodSource`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:66:3: JTG2001: duplicate symbol: method com.github.difflib.patch.PatchWithAllDiffAlgorithmsTest.MethodSource`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:66:3: JTG2001: duplicate symbol: method com.github.difflib.patch.PatchWithAllDiffAlgorithmsTest.MethodSource`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:82:3: JTG2001: duplicate symbol: method com.github.difflib.patch.PatchWithAllDiffAlgorithmsTest.MethodSource`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:82:3: JTG2001: duplicate symbol: method com.github.difflib.patch.PatchWithAllDiffAlgorithmsTest.MethodSource`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:82:3: JTG2001: duplicate symbol: method com.github.difflib.patch.PatchWithAllDiffAlgorithmsTest.MethodSource`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:0:0: JTG2005: import unresolved: com.github.difflib.algorithm.DiffAlgorithmFactory`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:25:35: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:39:48: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:40:46: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:42:42: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:55:48: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:56:46: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:58:42: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:71:48: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:72:46: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:74:42: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:87:48: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:88:46: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithAllDiffAlgorithmsTest.java:90:42: JTG2001: duplicate Go symbol after migration: NewPatchWithAllDiffAlgorithmsTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithMyerDiffTest.java:37:46: JTG2001: duplicate Go symbol after migration: NewPatchWithMyerDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithMyerDiffTest.java:39:42: JTG2001: duplicate Go symbol after migration: NewPatchWithMyerDiffTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithMyerDiffWithLinearSpaceTest.java:49:46: JTG2001: duplicate Go symbol after migration: NewPatchWithMyerDiffWithLinearSpaceTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/patch/PatchWithMyerDiffWithLinearSpaceTest.java:51:42: JTG2001: duplicate Go symbol after migration: NewPatchWithMyerDiffWithLinearSpaceTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java:12:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java:13:2: no viable alternative at input 'Testpublic'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java:25:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java:26:2: no viable alternative at input 'Testpublic'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java:29:33: extraneous input '>' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java:44:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java:45:2: no viable alternative at input 'Testpublic'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java:49:30: extraneous input '>' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java:65:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java:66:2: no viable alternative at input 'Testpublic'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorEqualitiesTest.java:70:30: extraneous input '>' expecting {'(', '!', '-', 'new', BooleanLiteral, 'null', IntegerLiteral, FloatingPointLiteral, StringLiteral, Identifier}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorTest.java:103:23: JTG2001: duplicate Go symbol after migration: NewDiffRowGeneratorTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/DiffRowGeneratorTest.java:867:4: JTG2001: duplicate Go symbol after migration: NewDiffRowGeneratorTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/StringUtilsTest.java:59:53: JTG2001: duplicate symbol: class com.github.difflib.text.StringUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/text/StringUtilsTest.java:59:53: JTG2001: duplicate Go symbol after migration: StringUtils`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:260:26: JTG2001: duplicate symbol: field com.github.difflib.unifieddiff.UnifiedDiffReaderTest.file1`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:290:26: JTG2001: duplicate symbol: field com.github.difflib.unifieddiff.UnifiedDiffReaderTest.file`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:356:26: JTG2001: duplicate symbol: field com.github.difflib.unifieddiff.UnifiedDiffReaderTest.file`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:448:20: JTG2001: duplicate symbol: field com.github.difflib.unifieddiff.UnifiedDiffReaderTest.file1`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:459:20: JTG2001: duplicate symbol: field com.github.difflib.unifieddiff.UnifiedDiffReaderTest.file1`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:469:20: JTG2001: duplicate symbol: field com.github.difflib.unifieddiff.UnifiedDiffReaderTest.file1`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:479:20: JTG2001: duplicate symbol: field com.github.difflib.unifieddiff.UnifiedDiffReaderTest.file1`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:489:20: JTG2001: duplicate symbol: field com.github.difflib.unifieddiff.UnifiedDiffReaderTest.file1`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:500:20: JTG2001: duplicate symbol: field com.github.difflib.unifieddiff.UnifiedDiffReaderTest.file1`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:38:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:40:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:40:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:40:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:95:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:97:15: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:116:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:118:15: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:145:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:147:15: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:161:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:163:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:163:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:163:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:176:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:178:15: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:196:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:198:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:198:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:198:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:211:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:213:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:213:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:213:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:231:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:233:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:233:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:233:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:238:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:254:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:256:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:256:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:256:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:261:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:269:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:271:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:271:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:271:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:274:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:274:35: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:286:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:288:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:288:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:288:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:291:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:291:35: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:305:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:307:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:307:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:307:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:317:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:319:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:319:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:319:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:329:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:331:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:331:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:331:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:339:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:341:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:341:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:341:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:351:86: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:354:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:354:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:354:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:357:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:357:42: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:368:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:370:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:370:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:370:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:425:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:427:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:427:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:427:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:435:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:437:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:437:31: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:437:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:447:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:450:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:450:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:457:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:461:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:461:39: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:467:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:471:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:471:41: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:477:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:481:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:481:40: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:487:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:491:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:491:35: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:498:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:502:4: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffReaderTest.java:502:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffReaderTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripNewLineTest.java:28:0: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripNewLineTest.java:28:1: mismatched input 'Disabled' expecting {<EOF>, 'package', ';', 'import', 'class', 'interface', 'public', 'private', 'protected', 'static', 'final'}`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripNewLineTest.java:30:2: token recognition error at: '@'`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripTest.java:25:29: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffRoundTripTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripTest.java:67:36: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffRoundTripTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripTest.java:72:40: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffRoundTripTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripTest.java:72:51: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffRoundTripTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripTest.java:72:58: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffRoundTripTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripTest.java:87:32: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffRoundTripTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffRoundTripTest.java:94:32: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffRoundTripTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffWriterTest.java:41:2: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffWriterTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffWriterTest.java:47:13: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffWriterTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffWriterTest.java:48:13: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffWriterTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffWriterTest.java:49:16: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffWriterTest`
- `.corpus/java-diff-utils/java-diff-utils/src/test/java/com/github/difflib/unifieddiff/UnifiedDiffWriterTest.java:53:22: JTG2001: duplicate Go symbol after migration: NewUnifiedDiffWriterTest`
