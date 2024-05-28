package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xuanfeng_gin/internal/api/intranet"
	config "xuanfeng_gin/pkg/conf"
)

var Shutdown func(context.Context)

type Service interface {
	Start()
}

type ApiServer struct {
	ctx context.Context
	Srv *http.Server
}

func NewApiServer(ctx context.Context) *ApiServer {
	return &ApiServer{
		ctx: ctx,
	}
}

func (f *ApiServer) Start() {

	if config.Get().Debug == gin.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	intranet.LoadRoute(router)

	f.Srv = &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Get().Http.Port),
		Handler: router,
	}

	fmt.Println(fmt.Sprintf("analysis api boot listen port:%v", config.Get().Http.Port))

	return
}
