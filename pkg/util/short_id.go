package util

import (
	"github.com/teris-io/shortid"
	"xuanfeng_gin/pkg/log"
)

func GetInviteCode() {
	sid, err := shortid.New(1, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 2432)
	if err != nil {
		log.L.Error(err.Error())
	}
	sid.Generate()
}
