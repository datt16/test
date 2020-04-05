# test
## 概要
go module の練習

## 使い方

```
$ cd test/cmd/server
$ go build
$ go run main.go
------------------------------------------------
$ curl localhost:8081/hello
> HelloWorld
```

## 構成

```
test
├ cmd 
│   └ server
│       └ main.go  ・  ・  ・  ・「サーバー起動用」
├ internal
│   └ server
│       └ server.go    ・  ・  ・「cmd/server/main.go で使用する処理群、ハンドラーなど」
```
