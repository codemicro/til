# Type hinting in Python

It might be dynamically typed, but we can still add type hints to make things at least a little more sane.

Some of the below may require extra imports, for example from the `typing` package.

```python
# ----- Variables -----
foo: bool  # uninitialised variable of type bool
bar: str = "hello"  # initialised variable of type string

# ----- Functions -----
def foo(bar: str, baz: int = 5) -> myCustomClass: ...

# ----- Special things -----
foo: Union[str, int]  # either a str or an int
foo: Optional[float]  # alias for Union[float, None]
foo: Any  # literally could be anything
foo: Callable[[Arg1, Arg2], ReturnType]
foo: List[str]
foo: Tuple[str]
foo: Dict[str, myClass]  # dictionary with str keys
```

Extended reference: https://docs.python.org/3/library/typing.html