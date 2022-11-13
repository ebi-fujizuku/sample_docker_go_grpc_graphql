package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ebi-fujizuku/sample_docker_go_grpc_graphql/article/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client  pb.ArticleServiceClient
)

func main() {
	fmt.Println("start gRPC Client.")

	// gRPCサーバーとのコネクションを確立
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// gRPCクライアントを生成
	client = pb.NewArticleServiceClient(conn)

	Hello()
}

func Hello() {
	name := "Hello World!"

	req := &pb.CreateArticleRequest{
		ArticleInput: name,
	}
	res, err := client.CreateArticle(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetArticle())
	}
}
