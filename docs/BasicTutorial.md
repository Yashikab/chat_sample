# Basic Tutorial

- [参考](https://grpc.io/docs/languages/go/basics/)

## チュートリアルを通して学ぶこと

- `.proto`ファイルの定義
- protocol buffer compilerをつかったサーバーとクライアントコードの生成
- かんたんなクライアントとサーバーを書くためのgoのgRPC APIの使用

## Why use gRPC ?

exampleはかんたんなルートマッピングアプリケーションで、クライアントはルートの特徴に関する情報を得る、
ルートのサマリーを作る、交通状況などによりルートの情報を変更する。

gRPCを使って `.proto`ファイルを一旦作ればgRPCが対応するどの言語でもクライアントとサーバーを生成できる。

## Get the example code

- grpc-goを使う

```shell
git clone -b v1.57.0 --depth 1 https://github.com/grpc/grpc-go
cd grpc-go/examples/route_guide
```

## Defining the service

最初のステップはgRPCサービスを定義することである。
すでに完成された`.proto`ファイルは `routeguide/route_guide.proto`にある
