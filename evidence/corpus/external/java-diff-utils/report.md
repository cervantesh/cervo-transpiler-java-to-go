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
