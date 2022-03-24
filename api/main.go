package api

import (
	"apiGetaway/config"
	"apiGetaway/package/logger"
	"apiGetaway/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// @Summary
	// @Description
	// @Produce json
	// @Param body body controllers.LoginParams true "body"
	// @Success 200 {string} string "ok" ""
	// @Failure 400 {string} string "err_code：10002 ； err_code：10003 ""
	// @Failure 401 {string} string "err_code：10001 "
	// @Failure 500 {string} string "err_code：20001 ；err_code：20002 ；err_code：20003 ；err_code：20004 ；err_code：20005 "
	// @Router /user/person/login [post]
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServiceManager
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	router.GET("/config", handlerV1.GetConfig)

	apiV1 := router.Group("/v1")
	apiV1.GET("/ping", handlerV1.Ping)

	// profession
	apiV1.POST("/profession", handlerV1.CreateProfession)
	apiV1.GET("/profession/:profession_id", handlerV1.GetProfession)
	apiV1.GET("/profession", handlerV1.GetAllProfessions)

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
