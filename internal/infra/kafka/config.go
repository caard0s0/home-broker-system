package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func ConsumerConfigMap() *ckafka.ConfigMap {
	return &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "ConsumerGroup_1",
		"auto.offset.reset": "earliest",
	}
}

func ProducerConfigMap() *ckafka.ConfigMap {
	return &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
	}
}
