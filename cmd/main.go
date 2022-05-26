package main

import (
	"context"
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/thathurleyguy/mongo_bench/bencher"
	"github.com/thathurleyguy/mongo_bench/cmd/config"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config := config.Init(ctx)
	defer config.Close()

	bencher := bencher.NewBencher(ctx, config)
	bencher.Start()
}
