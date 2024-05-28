package dao

import (
	"xuanfeng_gin/conf"
	"xuanfeng_gin/internal/global"
	config "xuanfeng_gin/pkg/conf"
	"xuanfeng_gin/pkg/locker"
	"xuanfeng_gin/pkg/mq/kafka"
)

var UserDao *User // 用户dao层 demo

var KafkaCommonSend *kafka.ProducerHandle // kafka

func Init(c *conf.Conf) {
	rdbLock := locker.NewLocker(global.Rdb)              // 分布式锁
	UserDao = NewUser(global.Db, c, global.Rdb, rdbLock) // 用户dao层初始化

	// kafka 初始化
	var kafkaErr error
	KafkaCommonSend = new(kafka.ProducerHandle)
	KafkaCommonSend.Producer, kafkaErr = kafka.NewSyncProducer(
		config.Get().Kafka.Brokers,
		kafka.DefaultProducerConf,
		"common_send",
		c.Kafka.Prefix,
	)

	if kafkaErr != nil {
		panic(kafkaErr)
	}
}
