package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func main() {
	log.Println("🚀 Starting ngrok with custom domain...")
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	listener, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(
			config.WithDomain("kundhavi.ngrok.io"), // 💡 Replace with your actual reserved domain
		),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return fmt.Errorf("❌ could not start tunnel: %w", err)
	}

	log.Println("✅ Tunnel started at:", listener.URL())

	return http.Serve(listener, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 Received request:", r.URL.Path)
	fmt.Fprintln(w, "Hello from ngrok-go with custom domain!")
}
