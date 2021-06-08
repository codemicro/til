# Show the used expression in an f-string

*Python 3.8+*

```py
myVar = list(range(4))
print(f"{myVar=}") # -> "myVar=[1,2,3,4]"

thing = 4
print(f"{thing == 4=}") # -> "thing == 4=True"

# why you would want to do this, I do not know - however..
print(f"{'hi=}") # -> "'hi'='hi'"
```
