# Get the last modified date of a file

To get the last modified date of a committed file in a Git repository, the following command can be run:

```bash
git log -1 --format=%cd $FILENAME
# Fri May 21 00:13:26 2021 +0100
```