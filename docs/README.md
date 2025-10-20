# Docs

Mkdocs is a Static Site Generator (SSG) that reads and compiles markdown files as a static web site for the purposes of hosting documentation.

## Running from project root

```shell
make docs
```

## Setup (within /docs dir)

as Mkdocs is Python based, python is required

uv it utilized for python package mgmt and is also required.

```shell
uv sync

uv run mkdocs serve
```
