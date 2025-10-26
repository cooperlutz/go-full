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

## Drawio Diagrams

We use [Drawio VSCode Extension](https://marketplace.visualstudio.com/items?itemName=hediet.vscode-drawio) for creating and editing diagrams in our documentation.

Utilizing the extension, diagrams are saved as `.drawio.svg` files which contain both the SVG image data and the embedded Drawio XML data.
