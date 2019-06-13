# GoでCloud Functionsを利用する

## デプロイ

`functions/helloworld`ディレクトリにある`hello_world.go`の`HelloGet()`をデプロイしたい場合。

```go
package helloworld

import (
	"fmt"
	"net/http"
)

func HelloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
```

`functions/helloworld`まで移動し、以下のコマンドを実行。

```shell
$ gcloud functions deploy HelloGet --runtime go111 --trigger-http --region=asia-northeast1
```

## デプロイした関数を実行

HTTPリクエストで確認できるので、`curl`コマンドで確認できる。

```shell
$ curl "https://[REGION]-[PROJECT_ID].cloudfunctions.net/HelloGet"
```

今回の場合、リクエスト先は以下の感じになる。

```shell
$ curl "https://asia-northeast1-cloud-functions-with-go.cloudfunctions.net/HelloGet" 
Hello World!
```

## 関数の削除

```$
$ gcloud functions delete HelloGet --region=asia-northeast1
```