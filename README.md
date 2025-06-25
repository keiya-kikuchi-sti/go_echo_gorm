# README

## 前置き

VS Code の Dev Container で開発・実行しやすいようにしています。  

## 起動

Dev Container で開いていれば勝手に起動していると思いますが、  
以下のコマンドで手動起動できます。

```bash
docker compose up -d
```

## API 実行

DevContainer 内のターミナルで以下のコマンドを実行する。  
（必要なパッケージが無ければ勝手にダウンロードします）

```bash
# in /go/src
go run main.go
```

## API呼び出し

`req/req.rest` ファイルを開くと
```
GET http://localhost:8080/todos/
```

などの上に `Send Request` と出ていると思いますので、それをクリック

## DB 初期化

コンテナを Volume ごと消してから再作成すると初期化できます。

```bash
docker compose down -v db
docker compose up -d db
```