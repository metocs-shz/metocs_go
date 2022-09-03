package manager

import (
	"fmt"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"metocs/common/application"
	"metocs/common/database"
	"metocs/oauth/models"
	"src.techknowlogick.com/oauth2-gorm"
	"strconv"
	"time"
)

var (
	OauthServer *server.Server
)

func OauthInit() {

	mysqlConfig := application.Application.MysqlConfig
	redisConfig := application.Application.Redis
	//	auth := application.Application.Auth

	//h获取权限管理器
	manager := manage.NewDefaultManager()
	//设置授权码过期时间 10分钟
	manager.SetAuthorizeCodeExp(time.Minute * 10)
	//配置token相关参数
	cfg := &manage.Config{
		// 访问令牌过期时间（默认为2小时）
		AccessTokenExp: time.Hour * 2,
		// 更新令牌过期时间（默认为72小时）
		RefreshTokenExp: time.Hour * 24 * 3,
		// 是否生成更新令牌（默认为true）
		IsGenerateRefresh: true,
	}
	manager.SetAuthorizeCodeTokenCfg(cfg)
	//配置token1生成器
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	// 设置客户端存储器为mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", mysqlConfig.UserName,
		mysqlConfig.PassWord,
		mysqlConfig.Ip,
		mysqlConfig.Port,
		mysqlConfig.DataBase,
		"10s")
	config := oauth2gorm.NewConfig(dsn, oauth2gorm.MySQL, "")
	oauthStore := oauth2gorm.NewClientStoreWithDB(config, database.DB)
	manager.MapClientStorage(oauthStore)

	//设置token管理器 可升级为集群模式
	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr:     redisConfig.Ip + ":" + strconv.Itoa(redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DataBase,
	}))

	//创建 gin server oatuh 服务管理器
	OauthServer = ginserver.InitServer(manager)
	//设置允许的授权请求类型
	OauthServer.SetAllowedResponseType(oauth2.Code)
	//设置允许的授权模式类型
	OauthServer.SetAllowedGrantType(oauth2.AuthorizationCode)
	OauthServer.SetAllowGetAccessRequest(true)
	OauthServer.SetClientInfoHandler(server.ClientFormHandler)
	OauthServer.SetUserAuthorizationHandler(models.SerAuthorizationHandler)

}
