package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
)

type HandlerFunc func(msg []byte)

type ConsumerHandle struct {
	ConsumerGroup *KConsumerGroup
	Handler       HandlerFunc
}

func (a *ConsumerHandle) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (a *ConsumerHandle) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (a *ConsumerHandle) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		a.Handler(msg.Value)
		sess.MarkMessage(msg, "")
	}
	return nil
}

type ProducerHandle struct {
	Producer *KSyncProducer
}

func (a *ProducerHandle) SendMsg(msg []byte, partition int32, key ...string) (int32, int64, error) {
	kMsg := &sarama.ProducerMessage{}
	kMsg.Topic = a.Producer.prefix + a.Producer.topic
	kMsg.Partition = partition
	if len(key) == 1 {
		kMsg.Key = sarama.StringEncoder(key[0])
	}
	kMsg.Value = sarama.ByteEncoder(msg)
	return a.Producer.SendMessage(kMsg)
}

func (a *ProducerHandle) SendMsgByDynamicTopic(topic string, msg []byte, partition int32, key ...string) (int32, int64, error) {
	kMsg := &sarama.ProducerMessage{}
	if topic == "" {
		topic = a.Producer.topic
	}
	kMsg.Topic = a.Producer.prefix + topic
	kMsg.Partition = partition
	if len(key) == 1 {
		kMsg.Key = sarama.StringEncoder(key[0])
	}
	kMsg.Value = sarama.ByteEncoder(msg)
	return a.Producer.SendMessage(kMsg)
}

func (a *ProducerHandle) SendMsgJson(msg interface{}, partition int32, key ...string) (int32, int64, error) {
	var msgJSON []byte
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		return 0, 0, err
	}

	return a.SendMsg(msgJSON, partition, key...)
}

func (a *ProducerHandle) SendMsgByDynamicTopicJson(topic string, msg interface{}, partition int32, key ...string) (int32, int64, error) {
	var msgJSON []byte
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		return 0, 0, err
	}

	return a.SendMsgByDynamicTopic(topic, msgJSON, partition, key...)
}

func (a *ProducerHandle) SendMsgByDynamicTopicPartitionJson(topic string, msg interface{}, partitions []int32, key ...string) (int32, int64, error) {
	var msgJSON []byte
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		return 0, 0, err
	}

	var partition int32
	tmpMap := make(map[int]int32)
	for k, v := range partitions {
		tmpMap[k] = v
	}
	for _, v := range tmpMap {
		partition = v
	}

	return a.SendMsgByDynamicTopic(topic, msgJSON, partition, key...)
}
