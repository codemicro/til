# List all staged files in a Git repository in plaintext

```
git diff --name-only --cached
```

Alternatively, `cached` can be replaced with `staged` on later versions of Git.

```
➜ headers git:(master)  git diff --name-only --staged
cmd/headers/main.go
go.mod
go.sum
headers
magefile.go
pkg/headers/config.go
pkg/headers/headers.go
pkg/headers/headers_test.go
pkg/headers/template.go
test/headers.toml
test/nested/hi.go
test/nested/otherthing.py
test/thing.py
```
