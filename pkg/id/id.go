package id

import (
	"github.com/bwmarrin/snowflake"
	"sync"
)

var node *snowflake.Node

var idMutex sync.Mutex

// Init 初始化
// n 不能和集群里其他节点重复，否则有可能会产生相同的ID
func Init(n int64) {
	if node == nil {
		idMutex.Lock()
		defer idMutex.Unlock()
		if node == nil {
			var err error
			node, err = snowflake.NewNode(n)
			if err != nil {
				// ID生成器初始化失败，一定要panic
				panic(err)
			}

		}
	}
}

// GenString 生成string类型的ID
func GenString() string {
	return node.Generate().String()
}

// GenInt64 生成int64的ID
func GenInt64() int64 {
	return node.Generate().Int64()
}
