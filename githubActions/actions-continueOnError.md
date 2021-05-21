# Continue a GitHub Actions workflow if a step errors

If a step in a GitHub Actions workflow step fails (ie. returns an exit code that's not zero), the entire workflow run will abort and report as failed.

In some situations, it may be impossible to avoid a step returning a non-zero exit code. This can be dealt with using the [`continue-on-error` parameter](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepscontinue-on-error) for a step.

```yaml
- name: Git commit and push
  run: |
    git config user.email 'actions@github.com'
    git config user.name 'github-actions'
    git add README.md
    git commit -m 'Update README.md'
    git push origin HEAD:${{ github.ref }}
  continue-on-error: true # if there's nothing changed, Git would cause an error here
```

In this example, the workflow will complete as normal, even if Git complains that there is `nothing to commit` and that the `working tree clean`, like it would if nothing had been changed.