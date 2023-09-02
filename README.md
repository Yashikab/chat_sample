# chat_sample

## goでgRPCをQuick Start

- [参考](https://grpc.io/docs/languages/go/quickstart/)

### 事前準備

- goのインストール(済)
- Protocol buffer compiler (`protoc`)のインストール

    ```shell
    sudo apt install -y protobuf-compiler
    protc --version
    ```

- protocol buffer用のgoのplugin

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### Example Code

- [grpc-go](https://github.com/grpc/grpc-go)をクローン
  - `git clone -b v1.57.0 --depth 1 https://github.com/grpc/grpc-go`
- `grpc-go/examples/helloworld`に移動
- `go run greeter_server/main.go`でサーバーを実行
- 別の端末から `go run greeter_client/main.go`を実行してクライアントを起動
  - `Greeting: Hello World` が出力される
