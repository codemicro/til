# Labels in Golang

Golang supports the use of labels for managing control flow around a program.

For example:

```go
func myThing() {
    fmt.Println("one...")
    goto three
    fmt.Println("two...")
three:
    fmt.Println("three!")
}

myThing() // "one...\nthree!"
```

## Use in loops

Labels can also be used incombination with the `continue` and `break` keywords, for example:

```go
outer:
    for i := 0; i < 5; i += 1 {
        for k := 0; k < 10; k += 1 {
            if k == 5 {
                continue outer // or `break outer`
            }
            fmt.Println(k)
        }
    }

// -> "0 1 2 3 4 0 1 2 3 4 ..."
```

`goto` is used in the standard library in this manner.

## Limitations

1. If defined, a label must be used
2. Labels can only be used within the scope of the function they're defined in
