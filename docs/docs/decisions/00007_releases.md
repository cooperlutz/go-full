# 00007: Releases

## Status

Accepted

## Context

In the context of releasing code, a mechanism to track and document releases is essential. This decision outlines the approach for managing releases, including versioning, changelogs, and distribution methods.

## Decision

The decision is to adopt semantic versioning (SemVer) for versioning, maintain a detailed changelog in the repository, and use Git tags to mark release points. Release distributions and packages will be managed via Github Releases.

### Implementation

- Commitizen will handle conventional commit messages to automate versioning and changelog generation.
- GoReleaser will be used to automate the build and release process, ensuring consistent and reproducible releases.
- Releases will be documented in a `CHANGELOG.md` file, following the "Keep a Changelog" format. This will be implemented via Commitizen.
- Git tags will be created for each release, following the SemVer format (e.g., v1.0.0). Commitizen will handle tag creation.
- Github Actions will handle orchestration of the release process, integrating with GoReleaser for building and publishing releases.
- **Releases will be triggered upon merging to the main branch, without utilizing release branches.**

## Alternatives Considered

Release branches were considered and initially implemented but were ultimately removed in favor of a more streamlined approach that is more aligned to Trunk Based Development.

## Consequences

- Development Flow is streamlined by removing the need to manage release branches, changelogs and automated versioning.
- Releases are more frequent and easier to manage, reducing overhead and complexity in the release process.
- All merges to `main` should be considered fully functional code, as they will directly trigger releases that could be consumed by end users.
- Exceptions may be required in situations when merging to main without triggering a release is required, and will need to be solutioned for.
