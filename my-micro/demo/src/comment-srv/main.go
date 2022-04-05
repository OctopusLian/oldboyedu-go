package main

import (
	"oldboyedu-go/my-micro/demo/src/comment-srv/db"
	"oldboyedu-go/my-micro/demo/src/comment-srv/handler"
	"oldboyedu-go/my-micro/demo/src/share/config"
	"oldboyedu-go/my-micro/demo/src/share/pb"
	"oldboyedu-go/my-micro/demo/src/share/utils/log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"go.uber.org/zap"
)

func main() {

	log.Init("comment")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameComment),
		micro.Version("latest"),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("Info", zap.Any("comment-srv", "comment-srv is start ..."))
			// 注册redis
			//redisPool := share.NewRedisPool(3, 3, 1,300*time.Second,":6379","redis")
			// 先注册db
			db.Init(config.MysqlDSN)
			pb.RegisterCommentServiceExtHandler(service.Server(), handler.NewCommentServiceExtHandler(), server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("comment-srv", "comment-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("comment-srv服务启动失败 ...")
	}
}
