# QueryParams

Utility package that provides the ability to more conveniently work with `URL` parameters.

### Install

````
go get github.com/mozg1984/queryparams
````

### Import
````go
import (
        "github.com/mozg1984/queryparams"
)
````

Usage example
-----------------

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/mozg1984/queryparams"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	// p1=hello&p2=1&p3=2.3&p4=0&p4=1&p4=2
	params := NewQueryParams(r.URL)

	fmt.Println(params.GetString("p1")) // hello
	
	p2, _ := params.GetUint("p2")
	fmt.Println(p2) // 1

	p3, _ := params.GetFloat("p3")
	fmt.Println(p3) // 2.3
	
	fmt.Println(params.GetStrings("p4")) // [0 1 2]

	unk, _ := params.GetInt32("unknown", -1)
	fmt.Println(unk) // -1
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}
```

License

MIT