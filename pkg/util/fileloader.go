package util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func LoadConf[T any](configFile *string, c T) (err error) {
	bf, err := os.ReadFile(*configFile)
	if err != nil {
		log.Fatalf(fmt.Sprintf("error: config file %v, %s", *configFile, err.Error()))
		return
	}

	if err = yaml.Unmarshal(bf, &c); err != nil {
		log.Fatalf(fmt.Sprintf("error: config format error %s, %s", bf, err.Error()))
		return
	}
	return
}
