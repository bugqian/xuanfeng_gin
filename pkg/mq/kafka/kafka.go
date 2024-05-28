package kafka

import (
	"context"

	"github.com/IBM/sarama"
)

type Conf struct {
	Brokers []string `yaml:"brokers" json:"brokers"`
	Prefix  string   `yaml:"prefix" json:"prefix"`
}

type ConsumerConf struct {
	ConsumerReturnErrors           bool
	Version                        sarama.KafkaVersion
	ConsumerGroupRebalanceStrategy sarama.BalanceStrategy
	ConsumerOffsetsInitial         int64
}

type ProducerConf struct {
	ProducerRequiredAcks    sarama.RequiredAcks
	ProducerPartitioner     sarama.PartitionerConstructor
	ProducerReturnSuccesses bool
	ProducerReturnErrors    bool
	Version                 sarama.KafkaVersion
}

type KConsumerGroup struct {
	sarama.ConsumerGroup
	groupID string
	topics  []string
	prefix  string
}

type KSyncProducer struct {
	sarama.SyncProducer
	topic  string
	prefix string
}

var DefaultProducerConf = ProducerConf{
	ProducerRequiredAcks:    sarama.WaitForAll,
	ProducerPartitioner:     sarama.NewHashPartitioner,
	ProducerReturnSuccesses: true,
	ProducerReturnErrors:    true,
	Version:                 sarama.MaxVersion,
}

var DefaultConsumerConf = ConsumerConf{
	ConsumerReturnErrors:           true,
	Version:                        sarama.MaxVersion,
	ConsumerGroupRebalanceStrategy: sarama.BalanceStrategySticky,
	ConsumerOffsetsInitial:         sarama.OffsetNewest,
}

func NewSyncProducer(brokers []string, producerConf ProducerConf, topic, prefix string) (*KSyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = producerConf.ProducerRequiredAcks
	config.Producer.Partitioner = producerConf.ProducerPartitioner
	config.Producer.Return.Successes = producerConf.ProducerReturnSuccesses
	config.Producer.Return.Errors = producerConf.ProducerReturnErrors
	config.Version = producerConf.Version
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &KSyncProducer{
		producer,
		topic,
		prefix,
	}, nil
}

func NewConsumerGroup(brokers []string, consumerConf ConsumerConf, topics []string, groupId, prefix string) (*KConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = consumerConf.ConsumerReturnErrors
	config.Version = consumerConf.Version
	//config.Consumer.Group.Rebalance.Strategy = consumerConf.ConsumerGroupRebalanceStrategy
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{consumerConf.ConsumerGroupRebalanceStrategy}
	config.Consumer.Offsets.Initial = consumerConf.ConsumerOffsetsInitial
	newGroupId := prefix + groupId
	consumerGroup, err := sarama.NewConsumerGroup(brokers, newGroupId, config)
	if err != nil {
		return nil, err
	}
	return &KConsumerGroup{
		consumerGroup,
		newGroupId,
		topics,
		prefix,
	}, nil
}

func (a *KConsumerGroup) RunConsumer(handler sarama.ConsumerGroupHandler) {
	newTopics := []string{}
	for _, v := range a.topics {
		newTopics = append(newTopics, a.prefix+v)
	}

	ctx := context.Background()
	for {
		err := a.ConsumerGroup.Consume(ctx, newTopics, handler)
		if err != nil {
			panic(err.Error())
		}
	}
}

func (a *KConsumerGroup) RunConsumerCtx(ctx context.Context, handler sarama.ConsumerGroupHandler) {
	newTopics := []string{}
	for _, v := range a.topics {
		newTopics = append(newTopics, a.prefix+v)
	}
	for {
		err := a.ConsumerGroup.Consume(ctx, newTopics, handler)
		if err != nil {
			panic(err.Error())
		}
	}
}
