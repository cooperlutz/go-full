# Branching Strategy

We follow [Trunk Based Development](https://trunkbaseddevelopment.com/) practices, enforcing and encouraging short-lived feature branches that are continuously merged into `main` (trunk).

![branching](../../_img/branching.drawio.png)

## Justification & Context

Long-lived branches can lead to significant merge conflicts, integration challenges, and delays in delivering value to end users. By adopting Trunk Based Development, we aim to minimize these issues and promote a more collaborative and efficient development process.

[Example of long-lived feature branches](https://github.com/cooperlutz/go-full/pull/110)

In the above example, a feature branch was created for a new feature, the exam library. Luckily, the code itself was modularized enough to avoid significant merge conflicts, along with only having an individual maintainer at the time. but the branch itself adds 9415 lines of code and 134 file changes. Imagine the amount of time and effort required to review and merge this code if multiple maintainers were working on the codebase simultaneously, or if the branch had been open for several weeks or months, and the `main` branch had diverged significantly, and the feature branch had to be rebased multiple times....

By keeping feature branches short-lived and merging them into `main` frequently, we can reduce the risk of integration issues, ensure that code changes are reviewed and tested promptly, and maintain a high level of code quality and stability in the `main` branch.

We treat `main` as the source of truth for production-ready code, and all changes to `main` should be treated as production-ready, thoroughly reviewed, tested, and validated before being merged.
