# TriggerHTTP

HTTPリクエストから実行する関数。

リクエストタイプに応じて処理を分岐する。

## デプロイ

```shell
$ gcloud functions deploy TriggerHTTP --runtime go111 --trigger-http --region=asia-northeast1
```

## デプロイした関数を実行

```shell
$ curl -X GET "https://asia-northeast1-cloud-functions-with-go.cloudfunctions.net/TriggerHTTP" 
Cloud Functions より Hello World!（GET メソッド）
```

```shell
$ curl -X POST "https://asia-northeast1-cloud-functions-with-go.cloudfunctions.net/TriggerHTTP" 
Cloud Functions より Hello World!（POST メソッド）
```

```shell
$ curl -X PUT "https://asia-northeast1-cloud-functions-with-go.cloudfunctions.net/TriggerHTTP" 
405 - Method Not Allowed
```