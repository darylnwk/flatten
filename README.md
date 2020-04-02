# flatten
--
    import "github.com/darylnwk/flatten"


## Usage

```go
var TimeFormat = time.RFC3339Nano
```
TimeFormat defines default time format when unmarshalling time.Time Override
this to unmarshal time.Time to a different format

#### func  Struct

```go
func Struct(s interface{}, m map[string]interface{})
```
Struct parses a struct `s` with JSON tags and flattens nested parameters to only
one level and passes the result to `m`.
