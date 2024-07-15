package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	pb "twirp-example/rpc/haber"
)

func main() {
	// client := pb.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})
	client := pb.NewHaberdasherJSONClient("http://localhost:8080", &http.Client{})

	hat, err := client.MakeHat(context.Background(), &pb.Size{Inches: 12})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}

	Println(GreenColor, "I have a nice new hat:", Pretty(hat))
}
