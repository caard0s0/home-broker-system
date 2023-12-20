package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/caard0s0/home-broker-system/internal/dto"
	"github.com/caard0s0/home-broker-system/internal/infra/kafka"
	"github.com/caard0s0/home-broker-system/internal/model"
	"github.com/caard0s0/home-broker-system/internal/transformer"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	ordersIn := make(chan *model.Order)
	ordersOut := make(chan *model.Order)
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	kafkaMsgChan := make(chan *ckafka.Message)

	pConfigMap := kafka.ProducerConfigMap()
	cConfigMap := kafka.ConsumerConfigMap()

	producer := kafka.NewKafkaProducer(pConfigMap)
	consumer := kafka.NewKafkaConsumer(cConfigMap, []string{"ordersIn"})

	go consumer.Consume(kafkaMsgChan)

	offerBook := model.NewOfferBook(ordersIn, ordersOut, wg)
	go offerBook.Trade()

	go func() {
		for msg := range kafkaMsgChan {
			wg.Add(1)
			fmt.Println(string(msg.Value))
			tradeInput := dto.TradeInput{}
			err := json.Unmarshal(msg.Value, &tradeInput)
			if err != nil {
				panic(err)
			}
			order := transformer.TransformInput(tradeInput)
			ordersIn <- order
		}
	}()

	for res := range ordersOut {
		output := transformer.TransformOutput(res)
		outputJson, err := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(outputJson))
		if err != nil {
			fmt.Println(err)
		}
		producer.Publish(outputJson, []byte("orders"), "ordersOut")
	}
}
