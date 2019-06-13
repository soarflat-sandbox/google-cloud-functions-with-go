package helloworld

import (
	"fmt"
	"net/http"
)

// ResponseWriter（https://golang.org/pkg/net/http/#ResponseWriter）インターフェース型の変数wと、
// Request（https://golang.org/pkg/net/http/#Request）ポインタ型のポインタ変数rを引数にとる
func HelloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
