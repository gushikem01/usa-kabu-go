package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gushikem01/usa-kabu-go/server/ent"
	_ "github.com/lib/pq"
)

// https://future-architect.github.io/articles/20210728a/
func main() {
	client, err := ent.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			"localhost", "5432", "postgres", "postgres", "postgres"))

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool
	ctx := context.Background()
	// if err := client.Schema.WriteTo(ctx, os.Stdin); err != nil {
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("ent sample done.")
}
