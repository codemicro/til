# Copy to the global clipboard from the command line

It's possible to write to the global clipboard from the command line using `xclip`.

```bash
sudo apt install xclip

cat file.txt | xclip -selection clipboard
```

You can then paste this content using CTRL+V as usual.
