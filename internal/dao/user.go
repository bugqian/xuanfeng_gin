package dao

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"xuanfeng_gin/conf"
	"xuanfeng_gin/model"
	"xuanfeng_gin/pkg/ecode"
	"xuanfeng_gin/pkg/id"
	"xuanfeng_gin/pkg/locker"
	"xuanfeng_gin/pkg/log"
	"xuanfeng_gin/types/request"
)

type User struct {
	db      *gorm.DB
	c       *conf.Conf
	rdb     *redis.Client
	rdbLock *locker.RedisLocker
}

func NewUser(db *gorm.DB, c *conf.Conf, rdb *redis.Client, rdbLock *locker.RedisLocker) *User {
	return &User{
		db:      db,
		c:       c,
		rdb:     rdb,
		rdbLock: rdbLock,
	}
}

// Create .demo
func (a *User) Create(req *request.UserCreateReq) (err error) {

	user := &model.User{
		Id:   id.GenString(),
		Name: req.Name,
		Age:  req.Age,
	}

	err = a.db.Model(&model.User{}).Create(user).Error

	if err != nil {
		log.L.Error(err.Error())
		err = ecode.ServerErr
	}
	return
}
