# Loading JSON into a class using Marshmallow Dataclasses

## Simple example

Marshal the following JSON in to a class:

```python
from marshmallow_dataclass import dataclass

@dataclass
class Person:
    name: str
    perferred_colour: str
    has_availability: bool

def main():
    jsonString = '{"name": "Abigail", "preferred_colour": "orange", "has_availability": true}'

    thePerson = Person.Schema().loads(jsonString)
    # or
    # thePerson = Person.Schema().load(json.loads(parsedJSON))

    thePerson.name # -> "Abigail"
```

## Loading arrays

```python
def main():
    jsonString = """
    [
        {"name": "Abigail", "preferred_colour": "orange", "has_availability": true},
        {"name": "Keira", "preferred_colour": "red", "has_availability": false}
    ]
    """

    persons = Person.Schema(many=True).loads(jsonString)

    persons[0].name # -> "Abigail"
    persons[1].name # -> "Keira"
```

## Selectively loading fields

```python
import marshmallow

def main():
    jsonString = '{"name": "Abigail", "preferred_colour": "orange", "has_availability": true, "height": 172}'

    thePerson = Person.Schema(unknown=marshmallow.EXCLUDE).loads(jsonString)
    
    thePerson.name # -> "Abigail"
```

## Optional fields

```python
from typing import Optional

@dataclass
class Person:
    name: str
    perferred_colour: str
    has_availability: bool
    eye_colour: Optional[str]

def main():
    jsonString = """
    [
        {"name": "Abigail", "preferred_colour": "orange", "has_availability": true},
        {"name": "Keira", "preferred_colour": "red", "has_availability": false, "eye_colour": "green"}
    ]
    """

    persons = Person.Schema(many=True).loads(jsonString)

    persons[0].eye_colour # -> None
    persons[1].eye_colour # -> "green"
```

## Custom schemas

Any arguments passed to `Classname.Schema` can also be put into their own, dedicated class.

```python
class _DefaultSchema(marshmallow.Schema):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.unknown = marshmallow.EXCLUDE
        self.many = True

@dataclass(base_schema=_DefaultSchema)
class Person:
    name: str
    perferred_colour: str
    has_availability: bool
    eye_colour: Optional[str]

def main():
    persons = Person.Schema().loads(jsonString)

    persons[0].eye_colour # -> None
    persons[1].eye_colour # -> "green"
```

## Reference

https://github.com/lovasoa/marshmallow_dataclass
