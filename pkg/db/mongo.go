package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConf struct {
	Hosts       []string
	Database    string
	AuthSource  string
	Username    string
	Password    string
	MaxPoolSize uint64
	MinPoolSize uint64
}

func NewMongo(conf *MongoConf) (*mongo.Client, *mongo.Database, error) {
	option := &options.ClientOptions{
		Auth: &options.Credential{
			AuthSource:  conf.AuthSource,
			Username:    conf.Username,
			Password:    conf.Password,
			PasswordSet: true,
		},
		Hosts:       conf.Hosts,
		MaxPoolSize: &conf.MaxPoolSize,
		MinPoolSize: &conf.MinPoolSize,
	}
	ctx := context.Background()
	cli, err := mongo.Connect(ctx, option)
	if err != nil {
		return nil, nil, err
	}
	err = cli.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, err
	}
	db := cli.Database(conf.Database)
	return cli, db, nil
}
