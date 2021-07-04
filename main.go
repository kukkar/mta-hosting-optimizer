package main

import (
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/kukkar/common-golang/pkg/config"
	appConf "github.com/kukkar/mta-hosting-optimizer/conf"

	_ "github.com/kukkar/mta-hosting-optimizer/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"github.com/kukkar/common-golang/pkg/healthcheck"
	"github.com/kukkar/common-golang/pkg/middleware"
	routes "github.com/kukkar/mta-hosting-optimizer/src/routes"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const (
	confFile       = "conf/config.json"
	envFilePathEnv = "ENV_FILE_PATH"
)

// @title mta-hosting-optimizer Swagger API
// @version 1.0
// @description Swagger API for mta-hosting-optimizer Project.

// @contact.name API Support
// @contact.email sahil.kukkar99@gmail.com

// @license.name SAHIL

// @BasePath /mta-hosting-optimizer/
func main() {
	//load env
	loadEnv()
	router := gin.New()
	//registering appconfig to global config
	registerConfig()
	//taking config into memory
	initConfig()
	//initiating logger
	initLogger(router)

	router.Use(gin.Recovery())

	//register default routes
	registerDefaultRoutes(router)
	// registerning middlewares
	registerMiddleware(router)
	// registering apis
	registerApis(router)

	//register health check
	registerHealthCheck(router)
	initServer(router)
}

func loadEnv() {
	filePath := os.Getenv(envFilePathEnv)
	if filePath == "" {
		godotenv.Load()
	} else {
		godotenv.Load(filePath)
	}
}

func registerApis(router *gin.Engine) {
	// register routing
	routes.Routes(router)
}

func initLogger(router *gin.Engine) {
	conf, err := appConf.GetGlobalConfig()
	if err != nil {
		panic(err)
	}
	err = conf.LogConfig.InitiateLogger()
	if err != nil {
		panic(err)
	}
}

//initConfig initialises the Global Application Config
func initConfig() {
	cm := new(config.ConfigManager)

	cm.InitializeGlobalConfig(confFile)
	cm.UpdateConfigFromEnv(config.GlobalAppConfig, "global")
	cm.UpdateConfigFromEnv(config.GlobalAppConfig.ApplicationConfig, "")
}

func registerConfig() {
	config.RegisterConfig(new(appConf.AppConfig))
	config.RegisterConfigEnvUpdateMap(appConf.EnvUpdateMap())
	config.RegisterGlobalEnvUpdateMap(config.GlobalEnvUpdateMap())
}

func initServer(router *gin.Engine) {
	conf, err := appConf.GetGlobalConfig()
	if err != nil {
		panic(err)
	}
	router.Run(conf.ServerHost + ":" + conf.ServerPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func registerMiddleware(router *gin.Engine) {
	defaultMiddleware := middleware.DefaultMiddleware{}
	router.Use(middleware.DebugMiddleware())
	router.Use(defaultMiddleware.CORSMiddleware())
}

func registerHealthCheck(router *gin.Engine) {

	gConf, err := appConf.GetGlobalConfig()
	if err != nil {
		panic(err)
	}
	hConfig := healthcheck.Config{}
	group := router.Group(string(gConf.AppName))
	{
		group.GET("/healthcheck", healthcheck.HealthCheckHandler(healthcheck.GetHealthCheck(hConfig)))
	}
}

func registerDefaultRoutes(router *gin.Engine) {
	conf, err := appConf.GetGlobalConfig()
	if err != nil {
		panic(err)
	}
	if strings.ToLower(conf.Environment) == "local" {
		pprof.Register(router)
		url := ginSwagger.URL("http://localhost:" + conf.ServerPort + "/swagger/doc.json") // The url pointing to API definition
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
