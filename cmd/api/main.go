package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"xuanfeng_gin/internal/api/server"
	"xuanfeng_gin/internal/dao"
	"xuanfeng_gin/internal/global"
	config "xuanfeng_gin/pkg/conf"
	"xuanfeng_gin/pkg/id"
	"xuanfeng_gin/pkg/log"
	"xuanfeng_gin/types"
)

var configFile = flag.String("f", "./conf.yaml", "the config file")

func main() {

	flag.Parse()
	// 初始化配置项
	c := config.Init(configFile)
	// 初始化日志
	log.Init(&c.Log)
	defer log.L.Sync()
	// 初始化id生成器
	id.Init(c.NodeId)
	// 获取全局上下文 后续上线问一律使用此上下文派生
	ctx := context.Background()
	// 初始化数据库 redis mongodb
	global.InitDB(c.Mysql, c.Redis)
	// 实例化dao
	dao.Init(c)
	// 启动前台服务
	s := server.NewApiServer(ctx)
	s.Start()

	go func() {
		if err := s.Srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.L.Error(fmt.Sprintf("listen: %s\n", err))
			panic(err)
		}
	}()

	// 监听停止信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-ch

	// 检测到停止信号 开始停止服务
	cancelCtx, cancelFun := context.WithTimeout(ctx, types.ExitShutdown*time.Second)
	defer cancelFun()

	// 停止服务超时,强行结束
	if err := s.Srv.Shutdown(cancelCtx); err != nil {
		log.L.Error("Server forced to shutdown:" + err.Error())
	}

}
