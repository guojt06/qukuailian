package global

import (
	"modulename/config"

	"github.com/gin-gonic/gin"
	goRedis "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	MysqlLog logger.Interface
	// Engine 全局引擎
	Engine *gin.Engine
	// BackendRouter 后台路由
	BackendRouter *gin.RouterGroup
	RedisCon      *goRedis.Client
)
