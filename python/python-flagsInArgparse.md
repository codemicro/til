# Defining a flag using `argparse`

```python
import argparse

parser = argparse.ArgumentParser()

# actions="store_const" is used to make this argument behave like a flag (you don't have to specify `-d True` or
# `-d False`, just `-d`)
# See https://docs.python.org/3/library/argparse.html#action
parser.add_argument(
    "-d", "--debug", help="enables debug mode", action="store_const", const=True
)
```
