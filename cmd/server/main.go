package main

import (
	"net/http"

	"twirp-example/internal"
	pb "twirp-example/rpc/haber"
)

func main() {
	server := &internal.Server{} // implements Haberdasher interface
	twirpHandler := pb.NewHaberdasherServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
