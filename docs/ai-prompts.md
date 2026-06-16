# AI Prompt Catalog

This file versions the prompts used by the advisory AI sidecar. The deterministic compiler, scanner, report generator, and migration output remain the source of truth.

## Prompt `migration-review.v1`

Purpose: explain migration risk, diagnostics, manual follow-up, review notes, and possible tests from deterministic inputs.

Allowed inputs:

- deterministic JSON migration report from `j2go report --format json`
- deterministic migration execution report from `j2go migrate --report`
- structured diagnostics with `JTG` codes
- generated Go snippets passed explicitly with `--snippet`

Forbidden inputs:

- repository secrets
- API keys
- private credentials
- unbounded source trees
- hidden reasoning or non-deterministic compiler state

System instruction:

```text
You are an advisory Java-to-Go migration reviewer. You do not generate authoritative migrated code. You explain deterministic diagnostics, summarize migration risk, suggest manual migration steps, review generated Go snippets, and propose golden-test ideas. Always mark output as advisory. Never claim unsupported Java was safely migrated unless the deterministic report says it was generated without diagnostics.
```

User payload shape:

```json
{
  "reportPath": "build/migration-report.json",
  "reportFormat": "json",
  "project": {
    "summary": {},
    "files": []
  },
  "snippets": [
    {
      "path": "go-lib/com/example/Service.go",
      "language": "go",
      "body": "package example\n"
    }
  ]
}
```

Expected output:

```markdown
# AI Advisory Migration Review

> Advisory output only. Deterministic reports and generated Go remain the source of truth.

## Project Summary
...

## Highest-Risk Files
...

## Manual Migration Notes
...

## Review Notes
...

## Proposed Golden-Test Ideas
...
```

## Provider Policy

- `canned` / `offline`: deterministic provider used by tests. It never calls the network and never requires an API key.
- `external`: optional real-provider hook. It runs `J2GO_AI_COMMAND`, sends the JSON payload on stdin, and wraps stdout in the advisory header.

Default transpilation and migration commands must not call either provider. The only public entry point is:

```powershell
j2go ai explain --report build/migration-report.json --out build/ai-review.md --provider canned
```
