# Python/uv recipe

Prerequisites: Python 3.11 or newer and uv. Create a root pyproject.toml with `[project]` (name, version and requires-python) and `[tool.uv.workspace]` with `members = ['packages/*']`. Each module has its own pyproject.toml and a retained `__init__.py` at `packages/<module>/__init__.py`, making it importable as `<module>` when `packages` is on PYTHONPATH. The module's `package = false` metadata tells uv to manage workspace dependencies without installing the module, so no build backend is needed. Run `uv sync --offline` when dependencies are cached, or use the normal package index first. The eval hook runs Python unittest through uv.

Derive the dependency graph from `FIREWEED_BIN modules --json`: use each module `dir`, `depends_on`, and resolved `zones`. Do not maintain a second dependency graph. Use the supplied import linter boundary hook to check Python imports against those live Fireweed modules and retained interfaces. Put implementation in `packages/<module>/internal/` and expose retained functions from `packages/<module>/__init__.py`. The eval hook adds the workspace's `packages` directory to Python's import path.

Run the setup hook after creating the root workspace files. Name unittest methods with spec IDs, for example `test_ALC001Pass` for ALC-001.
