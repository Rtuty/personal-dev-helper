package main

import (
	"context"
	"dev-helper/internal/bot"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	bot.Start(ctx)
}
