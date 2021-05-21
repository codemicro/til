# Skip running CI on `push` or `pull_request` events

Commits will not run a GitHub Actions workflow if one of the following strings are present in a specific place of your Git commit.

* `[skip ci]`
* `[ci skip]`
* `[no ci]`
* `[skip actions]`
* `[actions skip]`

## `push` events

If one of the above strings is present in any of the commit messages in your push, no workflows will be run on any of those commits.

## `pull_request` events

If the `HEAD` commit in a PR includes any of the following strings in the commit message, no workflows will be run on that pull request.

---

[Source](https://github.blog/changelog/2021-02-08-github-actions-skip-pull-request-and-push-workflows-with-skip-ci/)