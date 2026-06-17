# Java-To-Go Migration Report

Project: `.corpus/java-diff-utils/java-diff-utils/src/main/java`

## Summary

| Metric | Count |
| --- | ---: |
| Java files | 32 |
| Packages | 7 |
| Classes | 33 |
| Methods | 85 |
| Constructors | 58 |
| Fields | 83 |
| Unsupported features | 0 |
| Diagnostics | 1630 |
| Internal imports | 98 |

## Build Files

No Maven or Gradle build file detected.

## Packages

- `com.github.difflib`
- `com.github.difflib.algorithm`
- `com.github.difflib.algorithm.myers`
- `com.github.difflib.patch`
- `com.github.difflib.text`
- `com.github.difflib.text.deltamerge`
- `com.github.difflib.unifieddiff`

## Feature Counts

| Feature | Count |
| --- | ---: |
| class declarations | 33 |
| class fields | 83 |
| constructors | 58 |
| import declarations | 137 |
| instance methods | 75 |
| package declarations | 32 |
| static methods | 10 |

## Unsupported Feature Counts

None.

## Files

### `com/github/difflib/DiffUtils.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib`
- Imports: `com.github.difflib.algorithm.DiffAlgorithmFactory`, `com.github.difflib.algorithm.DiffAlgorithmI`, `com.github.difflib.algorithm.DiffAlgorithmListener`, `com.github.difflib.algorithm.myers.MyersDiff`, `com.github.difflib.patch.AbstractDelta`, `com.github.difflib.patch.Patch`, `com.github.difflib.patch.PatchFailedException`, `java.util.ArrayList`, `java.util.Arrays`, `java.util.Collections`, `java.util.List`, `java.util.Objects`, `java.util.function.BiPredicate`
- Classes: `com.github.difflib.DiffUtils`
- Fields in `com.github.difflib.DiffUtils`: `DiffAlgorithmFactory DEFAULT_DIFF`
- Constructors in `com.github.difflib.DiffUtils`: `diff(1 params)`, `diff(1 params)`, `diff(1 params)`, `diff(3 params)`, `diff(1 params)`, `diff(1 params)`, `diff(1 params)`, `diff(1 params)`, `diffInline(2 params)`, `patch(1 params)`, `unpatch(1 params)`, `compressLines(0 params)`, `DiffUtils(0 params)`
- Methods in `com.github.difflib.DiffUtils`: `static void withDefaultDiffAlgorithmFactory(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/UnifiedDiffUtils.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib`
- Imports: `com.github.difflib.patch.AbstractDelta`, `com.github.difflib.patch.ChangeDelta`, `com.github.difflib.patch.Chunk`, `com.github.difflib.patch.Patch`, `java.util.ArrayList`, `java.util.HashMap`, `java.util.List`, `java.util.Map`, `java.util.Optional`, `java.util.regex.Matcher`, `java.util.regex.Pattern`, `java.util.stream.Collectors`
- Classes: `com.github.difflib.UnifiedDiffUtils`
- Fields in `com.github.difflib.UnifiedDiffUtils`: `Pattern UNIFIED_DIFF_CHUNK_REGEXP`, `String NULL_FILE_INDICATOR`
- Constructors in `com.github.difflib.UnifiedDiffUtils`: `parseUnifiedDiff(0 params)`, `generateUnifiedDiff(2 params)`, `processDeltas(0 params)`, `getDeltaText(0 params)`, `UnifiedDiffUtils(0 params)`, `generateOriginalAndDiff(0 params)`, `generateOriginalAndDiff(0 params)`, `insertOrig(0 params)`, `getRowMap(1 params)`, `getOrigList(0 params)`
- Methods in `com.github.difflib.UnifiedDiffUtils`: `static void processLinesInPrevChunk(0 params)`, `static void insertOrig(1 params)`, `static void insert(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/algorithm/Change.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm`
- Imports: `com.github.difflib.patch.DeltaType`
- Classes: `com.github.difflib.algorithm.Change`
- Fields in `com.github.difflib.algorithm.Change`: `DeltaType deltaType`, `int startOriginal`, `int endOriginal`, `int startRevised`, `int endRevised`
- Constructors in `com.github.difflib.algorithm.Change`: `Change(5 params)`
- Methods in `com.github.difflib.algorithm.Change`: `instance Change withEndOriginal(1 params)`, `instance Change withEndRevised(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/algorithm/DiffAlgorithmFactory.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm`
- Imports: `java.util.function.BiPredicate`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/algorithm/DiffAlgorithmI.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm`
- Imports: `java.util.Arrays`, `java.util.List`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/algorithm/DiffAlgorithmListener.java`

- Source root: `.`
- Risk: `medium`
- Package: `com.github.difflib.algorithm`

### `com/github/difflib/algorithm/myers/MyersDiff.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm.myers`
- Imports: `com.github.difflib.algorithm.Change`, `com.github.difflib.algorithm.DiffAlgorithmFactory`, `com.github.difflib.algorithm.DiffAlgorithmI`, `com.github.difflib.algorithm.DiffAlgorithmListener`, `com.github.difflib.patch.DeltaType`, `com.github.difflib.patch.Patch`, `java.util.ArrayList`, `java.util.List`, `java.util.Objects`, `java.util.function.BiPredicate`
- Classes: `com.github.difflib.algorithm.myers.MyersDiff`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/algorithm/myers/MyersDiffWithLinearSpace.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm.myers`
- Imports: `com.github.difflib.algorithm.Change`, `com.github.difflib.algorithm.DiffAlgorithmFactory`, `com.github.difflib.algorithm.DiffAlgorithmI`, `com.github.difflib.algorithm.DiffAlgorithmListener`, `com.github.difflib.patch.DeltaType`, `java.util.ArrayList`, `java.util.List`, `java.util.Objects`, `java.util.function.BiPredicate`, `java.util.function.Consumer`
- Classes: `com.github.difflib.algorithm.myers.MyersDiffWithLinearSpace`, `com.github.difflib.algorithm.myers.DiffData`, `com.github.difflib.algorithm.myers.Snake`
- Fields in `com.github.difflib.algorithm.myers.DiffData`: `int size`, `int[] vDown`, `int[] vUp`
- Constructors in `com.github.difflib.algorithm.myers.DiffData`: `DiffData(1 params)`
- Fields in `com.github.difflib.algorithm.myers.Snake`: `int start`, `int end`, `int diag`
- Constructors in `com.github.difflib.algorithm.myers.Snake`: `Snake(3 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/algorithm/myers/PathNode.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.algorithm.myers`
- Classes: `com.github.difflib.algorithm.myers.PathNode`
- Fields in `com.github.difflib.algorithm.myers.PathNode`: `int i`, `int j`, `PathNode prev`, `boolean snake`, `boolean bootstrap`
- Constructors in `com.github.difflib.algorithm.myers.PathNode`: `PathNode(5 params)`
- Methods in `com.github.difflib.algorithm.myers.PathNode`: `instance boolean isSnake(0 params)`, `instance boolean isBootstrap(0 params)`, `instance PathNode previousSnake(0 params)`, `instance String toString(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/patch/AbstractDelta.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.io.Serializable`, `java.util.List`, `java.util.Objects`
- Classes: `com.github.difflib.patch.AbstractDelta`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/patch/ChangeDelta.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.util.List`, `java.util.Objects`
- Classes: `com.github.difflib.patch.ChangeDelta`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/patch/Chunk.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.io.Serializable`, `java.util.ArrayList`, `java.util.Arrays`, `java.util.List`, `java.util.Objects`
- Classes: `com.github.difflib.patch.Chunk`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/patch/ConflictOutput.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.io.Serializable`, `java.util.List`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/patch/DeleteDelta.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.util.List`
- Classes: `com.github.difflib.patch.DeleteDelta`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/patch/DeltaType.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Diagnostics: `JTG0001`

### `com/github/difflib/patch/DiffException.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Classes: `com.github.difflib.patch.DiffException`
- Fields in `com.github.difflib.patch.DiffException`: `long serialVersionUID`
- Constructors in `com.github.difflib.patch.DiffException`: `DiffException(0 params)`, `DiffException(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`

### `com/github/difflib/patch/EqualDelta.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.util.List`
- Classes: `com.github.difflib.patch.EqualDelta`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/patch/InsertDelta.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `java.util.List`
- Classes: `com.github.difflib.patch.InsertDelta`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/patch/Patch.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Imports: `staticjava.util.Comparator.comparing`, `com.github.difflib.algorithm.Change`, `java.io.Serializable`, `java.util.ArrayList`, `java.util.Collections`, `java.util.List`, `java.util.ListIterator`
- Classes: `com.github.difflib.patch.Patch`, `com.github.difflib.patch.PatchApplyingContext`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/patch/PatchFailedException.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Classes: `com.github.difflib.patch.PatchFailedException`
- Fields in `com.github.difflib.patch.PatchFailedException`: `long serialVersionUID`
- Constructors in `com.github.difflib.patch.PatchFailedException`: `PatchFailedException(0 params)`, `PatchFailedException(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`

### `com/github/difflib/patch/VerifyChunk.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.patch`
- Diagnostics: `JTG0001`

### `com/github/difflib/text/DiffRow.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.text`
- Imports: `java.io.Serializable`, `java.util.Objects`
- Classes: `com.github.difflib.text.DiffRow`
- Fields in `com.github.difflib.text.DiffRow`: `Tag tag`, `String oldLine`, `String newLine`
- Constructors in `com.github.difflib.text.DiffRow`: `DiffRow(3 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/text/DiffRowGenerator.java`

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
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/text/StringUtils.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.text`
- Imports: `staticjava.util.stream.Collectors.toList`, `java.util.List`
- Classes: `com.github.difflib.text.StringUtils`
- Constructors in `com.github.difflib.text.StringUtils`: `wrapText(0 params)`, `StringUtils(0 params)`
- Methods in `com.github.difflib.text.StringUtils`: `static String htmlEntites(1 params)`, `static String normalize(1 params)`, `static String wrapText(2 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/text/deltamerge/DeltaMergeUtils.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.text.deltamerge`
- Imports: `com.github.difflib.patch.AbstractDelta`, `com.github.difflib.patch.ChangeDelta`, `com.github.difflib.patch.Chunk`, `java.util.ArrayList`, `java.util.List`, `java.util.function.Predicate`
- Classes: `com.github.difflib.text.deltamerge.DeltaMergeUtils`
- Constructors in `com.github.difflib.text.deltamerge.DeltaMergeUtils`: `mergeInlineDeltas(2 params)`, `DeltaMergeUtils(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/text/deltamerge/InlineDeltaMergeInfo.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.text.deltamerge`
- Imports: `com.github.difflib.patch.AbstractDelta`, `java.util.List`
- Classes: `com.github.difflib.text.deltamerge.InlineDeltaMergeInfo`
- Constructors in `com.github.difflib.text.deltamerge.InlineDeltaMergeInfo`: `InlineDeltaMergeInfo(1 params)`, `getDeltas(0 params)`, `getOrigList(0 params)`, `getRevList(0 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/unifieddiff/UnifiedDiff.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Imports: `com.github.difflib.patch.PatchFailedException`, `java.util.ArrayList`, `java.util.Collections`, `java.util.List`, `java.util.function.Predicate`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiff`
- Fields in `com.github.difflib.unifieddiff.UnifiedDiff`: `String header`, `String tail`
- Constructors in `com.github.difflib.unifieddiff.UnifiedDiff`: `getFiles(0 params)`, `applyPatchTo(0 params)`
- Methods in `com.github.difflib.unifieddiff.UnifiedDiff`: `instance String getHeader(0 params)`, `instance void setHeader(1 params)`, `instance void addFile(1 params)`, `instance void setTailTxt(1 params)`, `instance String getTail(0 params)`, `static UnifiedDiff from(2 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/unifieddiff/UnifiedDiffFile.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Imports: `com.github.difflib.patch.Patch`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiffFile`
- Fields in `com.github.difflib.unifieddiff.UnifiedDiffFile`: `String diffCommand`, `String fromFile`, `String fromTimestamp`, `String toFile`, `String renameFrom`, `String renameTo`, `String copyFrom`, `String copyTo`, `String toTimestamp`, `String index`, `String newFileMode`, `String oldMode`, `String newMode`, `String deletedFileMode`, `String binaryAdded`, `String binaryDeleted`, `String binaryEdited`, `boolean noNewLineAtTheEndOfTheFile`, `Integer similarityIndex`
- Constructors in `com.github.difflib.unifieddiff.UnifiedDiffFile`: `getPatch(0 params)`
- Methods in `com.github.difflib.unifieddiff.UnifiedDiffFile`: `instance String getDiffCommand(0 params)`, `instance void setDiffCommand(1 params)`, `instance String getFromFile(0 params)`, `instance void setFromFile(1 params)`, `instance String getToFile(0 params)`, `instance void setToFile(1 params)`, `instance void setIndex(1 params)`, `instance String getIndex(0 params)`, `instance String getFromTimestamp(0 params)`, `instance void setFromTimestamp(1 params)`, `instance String getToTimestamp(0 params)`, `instance void setToTimestamp(1 params)`, `instance Integer getSimilarityIndex(0 params)`, `instance void setSimilarityIndex(1 params)`, `instance String getRenameFrom(0 params)`, `instance void setRenameFrom(1 params)`, `instance String getRenameTo(0 params)`, `instance void setRenameTo(1 params)`, `instance String getCopyFrom(0 params)`, `instance void setCopyFrom(1 params)`, `instance String getCopyTo(0 params)`, `instance void setCopyTo(1 params)`, `static UnifiedDiffFile from(2 params)`, `instance void setNewFileMode(1 params)`, `instance String getNewFileMode(0 params)`, `instance String getDeletedFileMode(0 params)`, `instance void setDeletedFileMode(1 params)`, `instance String getOldMode(0 params)`, `instance void setOldMode(1 params)`, `instance String getNewMode(0 params)`, `instance void setNewMode(1 params)`, `instance String getBinaryAdded(0 params)`, `instance void setBinaryAdded(1 params)`, `instance String getBinaryDeleted(0 params)`, `instance void setBinaryDeleted(1 params)`, `instance String getBinaryEdited(0 params)`, `instance void setBinaryEdited(1 params)`, `instance boolean isNoNewLineAtTheEndOfTheFile(0 params)`, `instance void setNoNewLineAtTheEndOfTheFile(1 params)`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/unifieddiff/UnifiedDiffParserException.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiffParserException`
- Constructors in `com.github.difflib.unifieddiff.UnifiedDiffParserException`: `UnifiedDiffParserException(0 params)`, `UnifiedDiffParserException(1 params)`, `UnifiedDiffParserException(2 params)`, `UnifiedDiffParserException(1 params)`, `UnifiedDiffParserException(4 params)`
- Diagnostics: `JTG0001`

### `com/github/difflib/unifieddiff/UnifiedDiffReader.java`

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

### `com/github/difflib/unifieddiff/UnifiedDiffWriter.java`

- Source root: `.`
- Risk: `high`
- Package: `com.github.difflib.unifieddiff`
- Imports: `com.github.difflib.patch.AbstractDelta`, `java.io.IOException`, `java.io.Writer`, `java.util.ArrayList`, `java.util.List`, `java.util.Objects`, `java.util.function.Consumer`, `java.util.function.Function`, `java.util.logging.Level`, `java.util.logging.Logger`
- Classes: `com.github.difflib.unifieddiff.UnifiedDiffWriter`, `com.github.difflib.unifieddiff.getName`
- Fields in `com.github.difflib.unifieddiff.UnifiedDiffWriter`: `Logger LOG`
- Diagnostics: `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`, `JTG0001`

### `com/github/difflib/unifieddiff/package-info.java`

- Source root: `.`
- Risk: `medium`
- Package: `com.github.difflib.unifieddiff`

## Recommended Migration Order

- `com/github/difflib/algorithm/DiffAlgorithmListener.java` (medium risk)
- `com/github/difflib/unifieddiff/package-info.java` (medium risk)
- `com/github/difflib/DiffUtils.java` (high risk)
- `com/github/difflib/UnifiedDiffUtils.java` (high risk)
- `com/github/difflib/algorithm/Change.java` (high risk)
- `com/github/difflib/algorithm/DiffAlgorithmFactory.java` (high risk)
- `com/github/difflib/algorithm/DiffAlgorithmI.java` (high risk)
- `com/github/difflib/algorithm/myers/MyersDiff.java` (high risk)
- `com/github/difflib/algorithm/myers/MyersDiffWithLinearSpace.java` (high risk)
- `com/github/difflib/algorithm/myers/PathNode.java` (high risk)
- `com/github/difflib/patch/AbstractDelta.java` (high risk)
- `com/github/difflib/patch/ChangeDelta.java` (high risk)
- `com/github/difflib/patch/Chunk.java` (high risk)
- `com/github/difflib/patch/ConflictOutput.java` (high risk)
- `com/github/difflib/patch/DeleteDelta.java` (high risk)
- `com/github/difflib/patch/DeltaType.java` (high risk)
- `com/github/difflib/patch/DiffException.java` (high risk)
- `com/github/difflib/patch/EqualDelta.java` (high risk)
- `com/github/difflib/patch/InsertDelta.java` (high risk)
- `com/github/difflib/patch/Patch.java` (high risk)
- `com/github/difflib/patch/PatchFailedException.java` (high risk)
- `com/github/difflib/patch/VerifyChunk.java` (high risk)
- `com/github/difflib/text/DiffRow.java` (high risk)
- `com/github/difflib/text/DiffRowGenerator.java` (high risk)
- `com/github/difflib/text/StringUtils.java` (high risk)
- `com/github/difflib/text/deltamerge/DeltaMergeUtils.java` (high risk)
- `com/github/difflib/text/deltamerge/InlineDeltaMergeInfo.java` (high risk)
- `com/github/difflib/unifieddiff/UnifiedDiff.java` (high risk)
- `com/github/difflib/unifieddiff/UnifiedDiffFile.java` (high risk)
- `com/github/difflib/unifieddiff/UnifiedDiffParserException.java` (high risk)
- `com/github/difflib/unifieddiff/UnifiedDiffReader.java` (high risk)
- `com/github/difflib/unifieddiff/UnifiedDiffWriter.java` (high risk)
