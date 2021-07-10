# Configure credential storage globally

1. ```bash
   git config --global credential.helper store
   ```
2. The next time credentials are required for a remote, they will be stored in `~/.git-credentials`, and you will not be prompted for them.
