# Handle SIGINT

```bash
trap my_function INT
my_function() { echo "hello!"; }

trap "echo 'hello';" INT
```

Both function identially.