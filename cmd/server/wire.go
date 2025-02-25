//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zhuguangfeng123/go-chat/cmd/server/app"
	iJwt "github.com/zhuguangfeng123/go-chat/internal/handler/jwt"
	userHdl "github.com/zhuguangfeng123/go-chat/internal/handler/user"
	"github.com/zhuguangfeng123/go-chat/internal/repository"
	"github.com/zhuguangfeng123/go-chat/internal/repository/cache"
	"github.com/zhuguangfeng123/go-chat/internal/repository/dao"
	codeSvc "github.com/zhuguangfeng123/go-chat/internal/service/code"
	smsSvc "github.com/zhuguangfeng123/go-chat/internal/service/sms/ali"
	userSvc "github.com/zhuguangfeng123/go-chat/internal/service/user"
	"github.com/zhuguangfeng123/go-chat/ioc"
)

func InitWebServer() *app.App {
	wire.Build(
		ioc.InitMysql,
		ioc.InitRedisCmd,
		ioc.InitAliSmsClient,

		dao.NewUserDao,

		cache.NewUserCache,
		cache.NewCodeCache,

		repository.NewUserRepository,
		repository.NewCodeRepository,

		userSvc.NewUserService,
		codeSvc.NewCodeService,
		smsSvc.NewAliService,

		iJwt.NewJwtHandler,
		userHdl.NewUserHandler,

		ioc.InitGinMiddleware,

		ioc.InitWebServer,

		wire.Struct(new(app.App), "*"),
	)

	return new(app.App)
}
