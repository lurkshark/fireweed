---
name: {{name}}
id_prefix: {{prefix}}
depends_on: []
---

# {{name}}

## Specs

When naming Go eval tests, a letter immediately before a spec prefix prevents ID matching. Test_GRT_001_Name matches GRT-001; TestGRT001 does not match GRT-001.

## Decisions
