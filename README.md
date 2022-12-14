# 

## 概要
+ goとgrpcとdockerを使って超簡単なものを作ってみた
## 参考サイト
+ [GoでgRPCサーバーを立ててみる](https://zenn.dev/k88t76/books/f3892660871ab2)
	+ このあと、GraphQLも学べるとのことで、これを作る。
	+ ただし、gRPCが古いのか、このサイト通りに作れなかったため、改造。
+ [作ってわかる！ はじめてのgRPC](https://zenn.dev/hsaki/books/golang-grpc-starting)
	+ Stepごとに作れるのでよくわかりやすかった。
## 宣伝
下記の配信で作っていたものです。
+ [【#勉強】プログラミング gRPCを学ぶんだ！#10【新人 Vtuber】ローカルで動いたし、Dockerで動かそー](https://youtu.be/-Biu6rp5MFA)
+ [【#勉強】プログラミング gRPCを学ぶんだ！#10【新人 Vtuber】Dockerで動いたし、デバッグ方法さがそう](https://youtu.be/P82WpfsDddo)

## 利用する場合
### 必須
+ /sample_docker_go_grpc_graphql/article/configに.envを作成し、下記のようにコーディング。
	~~~bash:.env
	SQLITE3_PATH=あなたのローカルのclone先/sample_docker_go_grpc_graphql/article/mydb.sqlite3
	~~~
### 必要ならば
+ もし、これを利用して、自身のGitHubにアップロードする場合、go moduleのパスは変更しましょう。
	```bash
	# go moduleのパス変更
	go mod edit -module github.com/あなたのGitHubアカウント名/sample_docker_go_grpc_graphql/article
	# go moduleを更新
	go mod tidy
	```
	参考: [Goメモ-260 (go.modのモジュール名を変更)(go mod edit -module)](https://devlights.hatenablog.com/entry/2022/10/24/073000)
