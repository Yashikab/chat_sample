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

- serviceを定義するには serviceという名前を `.proto`ファイルの中で特定する

```proto
service RouteGuide {
    ...
}
```

- serviceのなかに、`rpc`メソッドを定義し、リクエストとレスポンスタイプを特定する。gRPCは4種類のサービスメソッドをRouteGuideのなかに定義している

  - stubをつかってクライアントがサーバーにリクエストを送信したときに、普通の関数呼び出しのように返信するもの

  ```proto
  service RouteGuide {
    rpc GetFeature(Point) returns(Feature) {}
  }
  ```

  - サーバーサイドのストリーミングRPCでクライアントがリクエストを送信したときにメッセージのシーケンスをストリームする。
    クライアントは返信されたストリームをメッセージがなくなるまで読み続ける。

  ```proto
  service RouteGuide {
    rpc ListFeatures(Rectangle) returns(stream Feature) {}
  }
  ```

  - クライアント側のストリーミングRPCで、クライアントがメッセージのシーケンスを送信する。一旦クライアント側のメッセージが書き終わったら、サーバーが全部読むのをまち、その後レスポンスをもらう。

  ```proto
  service RouteGuide {
    rpc RecordRoute(stream Point) returns(RouteSummary) {}
  }
  ```

  - 双方向のストリーミングRPC。2つのストリームを独立に操作する。クライアントとサーバーは読み書きが自由にできる。
  例えば、サーバーがクライアントメッセージをまってからレスポンスを書くことや、交互に読み書きを行ったりできる。

  ```proto
  service RouteGuide {
    rpc RouteChat(stream RouteNote) returns(stream RouteNote) {}
  }
  ```

## Generating client and server code

```shell
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    routeguide/route_guide.proto
```

- `route_guide.pb.go`はすべてのprotocol bufferコードのリクエストとレスポンスタイプに関する記述をする
- `route_guide_grpc.pb.go`は以下を含む
  - `RouteGUide`サービスのメソッド定義を呼ぶためのクライアントのためのインタフェースタイプ
  - `RouteGuide`サービスメソッド定義のサーバー側の実装

## Creating the server

`RouteGuide`サーバーがどのように作られているか見る。
2つのパートがある

- サービス定義によって生成されたインタフェースの実装
- クラアントからのリクエストを聞き、正しいサービスの実装を返す

### RouteGuideの実装

生成された`routeGuideServer`インタフェースを実装した`routeGuideServer`structタイプを私達のサーバーは持っている。

#### Simple RPC

`routeGuideServer`実装はすべてのサービスメソッドを実装する。
`GetFeature`はクライアントからの`Point`を取得するだけの単純なもので、`Feature`内のDBにある対応する情報を返却する。

```go
func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
  for _, feature := range s.savedFeatures {
    if proto.Equal(feature.Location, point) {
      return feature, nil
    }
  }
  // No feature was found, return an unnamed feature
  return &pb.Feature{Location: point}, nil
}
```

#### Server-side streaming RPC

`ListFeatures`はサーバーサイドのRPCでmultiple featuresをクライアントに送る
simple RPCとは違いサーバーは request objectをうけとり`RouteGuide_ListFeaturesServer`オブジェクトを返却する

メソッド内では、たくさんの戻り値に必要な`Feature`オブジェクトを格納し、`Send()`をつかって `RouteGuide_ListFeaturesServer`に書き込む。

```go
func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
  for _, feature := range s.savedFeatures {
    if inRange(feature.Location, rect) {
      if err := stream.Send(feature); err != nil {
        return err
      }
    }
  }
  return nil
}
```

#### client-side streaming RPC

クライアントから、`Point`のストリームをうけ、`RouteSummary`を返す。
メソッドはrequest parameterをもたないかわりに `RouteGuide_RecordRouteServer`ストリームを取得する。
このサーバーはメッセージの読み書き両方で使用可能である。`Recv()`をつかうことでメッセージを受信し、single responseを`SendAndClose()`を使って送信する。

```go
func (s *routeGuideServer) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
  var pointCount, featureCount, distance int32
  var lastPoint *pb.Point
  startTime := time.Now()
  for {
    point, err := stream.Recv()
    if err == io.EOF {
      endTime := time.Now()
      return stream.SendAndClose(&pb.RouteSummary{
        PointCount:   pointCount,
        FeatureCount: featureCount,
        Distance:     distance,
        ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
      })
    }
    if err != nil {
      return err
    }
    pointCount++
    for _, feature := range s.savedFeatures {
      if proto.Equal(feature.Location, point) {
        featureCount++
      }
    }
    if lastPoint != nil {
      distance += calcDistance(lastPoint, point)
    }
    lastPoint = point
  }
}
```
