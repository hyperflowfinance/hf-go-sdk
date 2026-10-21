package main

import (
	"context"
	"fmt"
	"log"

	hyperflow "github.com/hyperflowfinance/hf-go-sdk"
	"github.com/hyperflowfinance/hf-go-sdk/examples/metadata"
)

func main() {
	client, err := hyperflow.Connect()
	if err != nil {
		log.Fatal(err)
	}
	result, err := metadata.Get(context.Background(), client)
	if err != nil {
		log.Fatal(err)
	}
	if result.ChainID == nil {
		log.Fatal("metadata.chainId is null")
	}
	fmt.Println(*result.ChainID)
}
