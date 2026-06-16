# Spring Exploration Fixtures

These fixtures document common Spring patterns for a future adapter. They are intentionally excluded from the deterministic transpiler golden suite.

The current decision is to classify Spring behavior before translating it:

- controllers can become route inventory and optional scaffolding
- services can use the pure Java migration path only for supported method bodies
- repositories remain advisory/manual by default
- dependency injection becomes explicit constructor wiring notes
- config properties can become config struct scaffolding
- exception handlers require manual HTTP/error mapping

See `docs/spring-adapter-decision.md`.
