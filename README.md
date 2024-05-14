# Ptr

Package ptr provides utility functions for converting non-addressable primitive types to pointers.
Its useful in contexts where a variable gives nil primitive type pointers semantics
(often meaning "not set") which can make it annoying to set the value.

[Look at ptr.go for the suite of functions](ptr.go). We recommend using the `To` function for everything unless you can find a really good reason not to.

## Example

```go
type Foo struct {
    A *int
}

func main() {
    foo := Foo{
        A: ptr.To(1)
    }
}
```

## Copy it, don't import it

Given the simplicity of this package, I recommend copying it to your source tree to avoid an unnecessary depedency.
