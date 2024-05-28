package conf

import (
	"xuanfeng_gin/pkg/db"
	"xuanfeng_gin/pkg/log"
	"xuanfeng_gin/pkg/mq/kafka"
)

type Conf struct {
	Http struct {
		Port string `yaml:"port"`
	} `yaml:"http"`
	Debug  string       `yaml:"debug"`
	NodeId int64        `yaml:"nodeId"`
	Log    log.Config   `yaml:"log"`
	Kafka  kafka.Conf   `yaml:"kafka"`
	Redis  db.RedisConf `yaml:"redis"`
	Mysql  db.MysqlConf `yaml:"mysql"`
}
