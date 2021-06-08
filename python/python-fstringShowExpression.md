# Show the used expression in an f-string

*Available in Python 3.8+*

Useful for debugging!

```py
myVar = list(range(4))
print(f"{myVar=}") # -> "myVar=[1,2,3,4]"

thing = 4
print(f"{thing == 4=}") # -> "thing == 4=True"

# why you would want to do this, I do not know - however..
print(f"{'hi=}") # -> "'hi'='hi'"
```

See also: [Formatted string literals](https://docs.python.org/3/reference/lexical_analysis.html#f-strings)
Source: [Reddit](https://reddit.com/r/learnpython/comments/nur6o9/til_ive_been_making_debugging_statements_harder/)
