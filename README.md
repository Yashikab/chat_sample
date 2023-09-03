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

### update the gRPC service

拡張メソッドで上記サンプルコードをアップデートする。gRPCサービスは `protocol buffers`を使って定義される。
詳しくは Basic tutorialを参考にする。

- サンプルコードでは、サーバー・クライントともに `SayHello()` stubをもっている。
- helloworld/helloworld.protoに `SayHelloAgain()`メソッドを追加する。

```proto
// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
  // Sends another greeting
  rpc SayHelloAgain (HelloRequest) returns (HelloReply) {}
}
```

- 新しいサービスを使う前に、`.proto`ファイルの更新を再コンパイルする必要がある。

```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```

- これで `helloworld/helloworld.pb.go`が再生成される
- サーバーに関数を加えてアップデートする

```go
func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
        return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}
```

- クライアントをアップデートする

```go
r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: *name})
if err != nil {
        log.Fatalf("could not greet: %v", err)
}
log.Printf("Greeting: %s", r.GetMessage())
```
