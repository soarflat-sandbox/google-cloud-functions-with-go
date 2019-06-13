package triggerhttp

import (
	"fmt"
	"net/http"
)

func TriggerHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // HTTPメソッドにより処理を分岐する
	case http.MethodGet: // GETの場合
		fmt.Fprint(w, "Cloud Functions より Hello World!（GET メソッド）")
	case http.MethodPost: //4.POST の場合。 14 fmt.Fprint(w, "Cloud Functions より Hello World!（POST メソッド）")
		fmt.Fprint(w, "Cloud Functions より Hello World!（POST メソッド）")
	default:
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
