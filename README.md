# go-daily-log

A thin wrapper around the standard logger. In addition to `Writer` of the
standard logger, go-daily-log also writes to log files timestamped in the
format of `2006-01-02`, and rotates them daily.

## Example

```go
package main

import (
	"net/http"

	"github.com/im-kl/go-daily-log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8080", nil)
}
```

## TODO

- [ ] Unit tests

