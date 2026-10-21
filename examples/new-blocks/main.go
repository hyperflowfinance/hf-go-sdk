package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"

	hyperflow "github.com/hyperflowfinance/hf-go-sdk"
)

func main() {
	client, err := hyperflow.Connect()
	if err != nil {
		log.Fatal(err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := watch(ctx, client, os.Stdout); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal(err)
	}
}

func watch(ctx context.Context, client *hyperflow.Client, output io.Writer) (err error) {
	document := hyperflow.NewSubscription(func(s *hyperflow.SubscriptionFields) {
		s.NewBlocks(func(b *hyperflow.BlockFields) {
			b.Number()
			b.Hash()
			b.Timestamp()
			b.GasUsed()
		})
	})
	stream, err := client.Subscribe(ctx, document)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := stream.Close(); err == nil {
			err = closeErr
		}
	}()
	return printBlocks(stream, output)
}

func printBlocks(stream *hyperflow.Subscription, output io.Writer) error {
	for {
		var event hyperflow.SubscriptionEvent
		if err := stream.Next(&event); errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return err
		}
		if _, err := fmt.Fprintf(output, "block %s hash=%s timestamp=%s gasUsed=%s\n",
			event.NewBlocks.Number, event.NewBlocks.Hash, event.NewBlocks.Timestamp, event.NewBlocks.GasUsed); err != nil {
			return err
		}
	}
}
