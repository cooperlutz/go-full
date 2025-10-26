# Releasing

[ADR-00002](../decisions/00007_releases.md) outlines our release strategy.

Upon merge to `main`, releases are automatically managed and orchestrated via Github Actions, utilizing [GoReleaser](https://goreleaser.com/) for building and packaging releases, [Commitizen](https://commitizen-tools.github.io/commitizen/) for versioning and changelog generation.
