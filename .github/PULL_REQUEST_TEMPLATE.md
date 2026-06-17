## Summary

- Describe the change.

## Type

- [ ] Bug fix
- [ ] Feature
- [ ] Documentation
- [ ] Diagnostic/reporting change
- [ ] Corpus/evidence update
- [ ] Refactor

## Verification

- [ ] `go test ./...`
- [ ] `go run ./tools/legacytest`
- [ ] `go run ./tools/corpus`
- [ ] Parser regenerated when grammar changed
- [ ] Golden outputs updated when generated Go changed
- [ ] Documentation updated when CLI behavior changed

## Migration Safety

- [ ] Deterministic output remains reproducible
- [ ] Unsupported behavior emits structured diagnostics
- [ ] No proprietary or confidential code included
- [ ] AI output, if present, is advisory only
