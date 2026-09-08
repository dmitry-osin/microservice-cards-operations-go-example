package main

import (
	"cards-operations/kafka"
	"context"
	"fmt"
	"log"
	"net/http"

	"cards-operations/core"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "{ \"status\" : \"ok\" }")
}

func main() {
	cfg, err := core.Load(".")
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	client, err := kafka.NewKafkaClient(*cfg)
	if err != nil {
		log.Fatalf("create kafka client: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	go kafka.Listen(ctx, client)

	http.HandleFunc("/health", handler)

	addr := fmt.Sprintf(":%d", cfg.HTTP.Port)
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
