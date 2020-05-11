package services

import (
	"github.com/yezihack/studyGo/org/service/services/repository/mongdbService"
	"github.com/yezihack/studyGo/org/service/services/repository/mysqlService"
	"github.com/yezihack/studyGo/org/service/services/repository/redisService"
)

//定义一个引擎类型
type Engine int

//定义引擎常量
const (
	EngineMysql Engine = iota
	EngineRedis
	EngineMongDB
)

//定义引擎接口
type AppService interface {
	Find()
	Delete()
}

//实例不同引擎
func NewService(engine Engine) AppService {
	switch engine {
	case EngineMysql:
		service := mysqlService.MySQLService{}
		return &service
	case EngineRedis:
		return new(redisService.RedisService)
	case EngineMongDB:
		return new(mongdbService.MongDBService)
	default:
		return nil
	}
}
