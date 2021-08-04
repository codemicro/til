# Copy files from another branch into the current branch

```
git checkout otherbranch myfile.txt
```

General versions:

```
git checkout <commit_hash> <relative_path_to_file_or_dir>
git checkout <remote_name>/<branch_name> <file_or_dir>
```

* Using the commit hash, you can pull files from any commit
* This works for files and directories
* Overwrites the file myfile.txt and mydir
* Wildcards don't work, but relative paths do
* Multiple paths can be specified

---

[Source](https://stackoverflow.com/a/307872)
