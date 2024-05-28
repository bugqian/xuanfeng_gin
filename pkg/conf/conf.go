package config

import (
	"sync"
	"xuanfeng_gin/conf"
	"xuanfeng_gin/pkg/util"
)

var c *conf.Conf
var cMutex sync.Mutex

func Init(configFile *string) *conf.Conf {
	if c == nil {
		cMutex.Lock()
		defer cMutex.Unlock()
		if c == nil {
			err := util.LoadConf(configFile, &c)
			if err != nil {
				panic(err)
			}
		}
	}
	return c
}

func Get() *conf.Conf {
	return c
}
