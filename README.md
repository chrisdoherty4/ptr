# Ptr

Package ptr provides utility functions for converting non-addressable primitive types to pointers.
Its useful in contexts where a variable gives nil primitive type pointers semantics
(often meaning "not set") which can make it annoying to set the value.

## Example

```go
type Foo struct {
    A *int
}

func main() {
    foo := Foo{
        A: ptr.Int(1)
    }
}
```

## Copy it, don't use-it

Given the simplicity of this package, I recommend copying it to your source tree to avoid an unnecessary depdency.