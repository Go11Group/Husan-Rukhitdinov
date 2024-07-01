package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "lesson44/genproto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTranslateServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Translate(ctx, &pb.TranslateRequest{Words: []string{"salom", "dunyo", "kitob"}})
	if err != nil {
		log.Fatalf("could not translate: %v", err)
	}

	log.Printf("Translated Words: %v", r.Translations)
}
