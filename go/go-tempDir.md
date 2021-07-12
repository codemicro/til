# Working with the `/tmp` directory

All files within the `/tmp` directory require the sticky bit to be set.

When creating a directory, this can be achieved as follows:

```go
tempDir := getTemporaryDir()
_ = os.Mkdir(tempDir, os.ModeDir|os.ModeSticky|os.FileMode(0700))
```
