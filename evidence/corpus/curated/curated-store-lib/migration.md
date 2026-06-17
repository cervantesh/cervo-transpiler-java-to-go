# Java-To-Go Migration Report

Project: `corpus/projects/curated-store-lib`

## Summary

| Metric | Count |
| --- | ---: |
| Java files | 5 |
| Packages | 3 |
| Classes | 5 |
| Methods | 8 |
| Constructors | 1 |
| Fields | 1 |
| Unsupported features | 0 |
| Diagnostics | 0 |
| Internal imports | 1 |

## Build Files

- `pom.xml` (maven)

## Packages

- `com.cervo.store.math`
- `com.cervo.store.model`
- `com.cervo.store.service`

## Feature Counts

| Feature | Count |
| --- | ---: |
| class declarations | 5 |
| class fields | 1 |
| constructors | 1 |
| import declarations | 1 |
| instance methods | 2 |
| package declarations | 5 |
| static methods | 6 |

## Unsupported Feature Counts

None.

## Files

### `src/main/java/com/cervo/store/math/PriceMath.java`

- Source root: `src/main/java`
- Risk: `medium`
- Package: `com.cervo.store.math`
- Classes: `com.cervo.store.math.PriceMath`
- Methods in `com.cervo.store.math.PriceMath`: `static int addTax(2 params)`, `static int discount(2 params)`

### `src/main/java/com/cervo/store/model/Counter.java`

- Source root: `src/main/java`
- Risk: `medium`
- Package: `com.cervo.store.model`
- Classes: `com.cervo.store.model.Counter`
- Fields in `com.cervo.store.model.Counter`: `int value`
- Constructors in `com.cervo.store.model.Counter`: `Counter(1 params)`
- Methods in `com.cervo.store.model.Counter`: `instance void add(1 params)`, `instance int current(0 params)`

### `src/main/java/com/cervo/store/model/Label.java`

- Source root: `src/main/java`
- Risk: `medium`
- Package: `com.cervo.store.model`
- Classes: `com.cervo.store.model.Label`
- Methods in `com.cervo.store.model.Label`: `static String currency(0 params)`

### `src/main/java/com/cervo/store/service/OrderService.java`

- Source root: `src/main/java`
- Risk: `medium`
- Package: `com.cervo.store.service`
- Imports: `com.cervo.store.math.PriceMath`
- Classes: `com.cervo.store.service.OrderService`
- Methods in `com.cervo.store.service.OrderService`: `static int total(2 params)`, `static int discountedTotal(2 params)`

### `src/test/java/com/cervo/store/service/OrderServiceTest.java`

- Source root: `src/test/java`
- Risk: `medium`
- Package: `com.cervo.store.service`
- Classes: `com.cervo.store.service.OrderServiceTest`
- Methods in `com.cervo.store.service.OrderServiceTest`: `static void smoke(0 params)`

## Recommended Migration Order

- `src/main/java/com/cervo/store/math/PriceMath.java` (medium risk)
- `src/main/java/com/cervo/store/model/Counter.java` (medium risk)
- `src/main/java/com/cervo/store/model/Label.java` (medium risk)
- `src/main/java/com/cervo/store/service/OrderService.java` (medium risk)
- `src/test/java/com/cervo/store/service/OrderServiceTest.java` (medium risk)

## Migration Execution Summary

| Metric | Count |
| --- | ---: |
| Java files | 5 |
| Generated files | 5 |
| Skipped files | 0 |
| Diagnostics | 0 |
| Dry run | false |

### Migration Diagnostics

None.
