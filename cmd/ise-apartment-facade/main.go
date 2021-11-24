package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ozonmp/ise-apartment-facade/internal/app/receiver"
	"github.com/ozonmp/ise-apartment-facade/internal/config"
	apartment "github.com/ozonmp/ise-apartment-facade/internal/model"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ConsumeEvent(ctx context.Context, message *sarama.ConsumerMessage) error {
	var event apartment.Event

	err := json.Unmarshal(message.Value, &event)
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("event consumed: %s", event.String()))

	return nil
}

func main() {

	sigs := make(chan os.Signal, 1)

	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal("Failed init configuration")
	}
	mainGfg := config.GetConfigInstance()

	err := receiver.StartConsuming(ctx, mainGfg.Kafka.Brokers, mainGfg.Kafka.Topic, mainGfg.Kafka.GroupID, ConsumeEvent)
	if err != nil {
		log.Fatal("Failed to create EventReceiver")
	}

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}
