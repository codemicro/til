# Get the PID of the last run command

If the previous command is still running, it's possible to get the PID of the last run commad using the `$!` variable.

```bash
➜ til git:(master)  sleep 15 &
[1] 7652
➜ til git:(master)  kill $!
[1]+  Terminated              sleep 15
```

This also works if the command has terminated.

```bash
➜ til git:(master)  sleep 2
➜ til git:(master)  echo $!
7652
```