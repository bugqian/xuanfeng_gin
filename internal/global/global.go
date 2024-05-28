package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"xuanfeng_gin/pkg/db"
)

var Db *gorm.DB
var Rdb *redis.Client

//var Mdb *mongo.Database // mongodb初始化

func InitDB(dbConf db.MysqlConf, rdbConf db.RedisConf) {
	var err error
	Db, _, err = db.NewMysql(&dbConf)
	if err != nil {
		panic(err)
	}

	Rdb, err = db.NewRedis(&rdbConf)
	if err != nil {
		panic(err)
	}

	// mongodb初始化
	//mOptions := &options.ClientOptions{
	//	Auth: &options.Credential{
	//		AuthSource:  mdbConf.AuthSource,
	//		Username:    mdbConf.Username,
	//		Password:    mdbConf.Password,
	//		PasswordSet: true,
	//	},
	//	Hosts:       mdbConf.Hosts,
	//	MaxPoolSize: &mdbConf.MaxPoolSize,
	//	MinPoolSize: &mdbConf.MinPoolSize,
	//}
	////mOptions := options.Client().ApplyURI("mongodb://root:NeN6wNZJk3vuP0Pg@192.168.2.84:18637")
	//mCtx := context.Background()
	//mClient, err := mongo.Connect(mCtx, mOptions)
	//if err != nil {
	//	panic(err)
	//}
	////rp := readpref.ReadPref{}
	//err = mClient.Ping(mCtx, readpref.Primary())
	//if err != nil {
	//	panic(err)
	//}
	//Mdb = mClient.Database(mdbConf.Database)
}
