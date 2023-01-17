package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/shvz0/mmscrap-grpc/scrappers"
	scr "github.com/shvz0/mmscrap/mmscrappers"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedScrappersServer
}

func (s *server) GetTodays24NewsArticles(rect *pb.Empty, stream pb.Scrappers_GetTodays24NewsArticlesServer) error {
	n24 := scr.NewNews24()
	articles, err := n24.ArticleListToday()
	if err != nil {
		stream.Context().Err()
	}
	for _, a := range articles {
		sendA := pb.Article{Author: a.Author, Text: a.Text}
		if err := stream.Send(&sendA); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterScrappersServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
