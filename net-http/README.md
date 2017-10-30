# About This

[net/http](https://go-talks.appspot.com/github.com/voyagegroup/talks/2017/treasure-go/intro.slide#33) を実装。

ただし指定したURLの以下の項目を取得するものとする。

* HTTP Response Code
* Content Type
* Content-Length
* Title

# Usage

* --urlで取得対象のURLを指定して起動
```
$ go run main.go --url=https://golang.org
```

* curlでlocalhostにアクセス
```
$ curl http://localhost:8080 | jq .
```
