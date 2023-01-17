package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/shvz0/mmscrap-grpc/scrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func printArticles(client pb.ScrappersClient, rect *pb.Empty) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	stream, err := client.GetTodays24NewsArticles(ctx, rect)
	if err != nil {
		log.Fatalf("client.ScrappersClient failed: %v", err)
	}
	for {
		article, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ScrappersClient failed: %v", err)
		}
		fmt.Println(article)
	}
}

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewScrappersClient(conn)

	printArticles(client, &pb.Empty{})
}
