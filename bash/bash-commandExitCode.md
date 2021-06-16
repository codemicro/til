# Exit codes in Bash scripts

The last exit code of a command can be obtained through the `$?` variable.

```
➜ ~: echo $?
0
```

In a Bash script, the `exit` command can be used to, well, exit with a given exit code.

```bash
exit 2
```
