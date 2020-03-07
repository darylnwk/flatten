# flatten
--
    import "github.com/darylnwk/flatten"


## Usage

#### func  Struct

```go
func Struct(s interface{}, m map[string]interface{})
```
Struct parses a struct `s` with JSON tags and flattens nested parameters to only
one level and passes the result to `m`.
