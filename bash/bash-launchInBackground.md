# Run a command in a background process

To run a command in a new background process, it's possible to use the `&` in Bash.

```bash
nemo . &
# [2] 69062
```

This, for example, will open the Nemo file manager in the background. It'll continue to send the contents of STDOUT and STDERR to our Bash session, so we can pipe all of that into `/dev/null` instead.

```bash
nemo . &>/dev/null &
# [2] 68983
```

(These created processes can be killed with `kill <pid>`, eg `kill 68983`)
