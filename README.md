# go-querystring

[![GoDoc](https://godoc.org/github.com/saadullahsaeed/go-querystring?status.svg)](https://godoc.org/github.com/saadullahsaeed/go-querystring) 

Go library to decode a query string to a struct.

## Usage 

```go
import querystring "github.com/saadullahsaeed/go-querystring"

type someStruct struct {
	A string  `url:"a"`
	B uint64  `url:"b"`
	C uint64  `url:"-"`
	D float32 `url:"d"`
	E float64 `url:"e"`
}

func handler(w http.ResponseWriter, req *http.Request) {
    ts := &someStruct{}
    querystring.Decode(req.URL.Query(), ts)
    fmt.Println(ts)
}
```
