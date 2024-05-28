package intranet

import (
	"xuanfeng_gin/internal/dao"
	"xuanfeng_gin/pkg/ecode"
	ginext "xuanfeng_gin/pkg/gin-ext"
	"xuanfeng_gin/pkg/log"
	"xuanfeng_gin/types/request"
)

func userCreate(ctx *ginext.Context) {
	var req request.UserCreateReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		log.L.Error(err.Error())
		ctx.JSON(nil, ecode.RequestErr)
		return
	}
	ctx.JSON(nil, dao.UserDao.Create(&req))
}
