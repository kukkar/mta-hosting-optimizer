package unittests

import (
	"fmt"
	"log"
	"os"

	"github.com/kukkar/common-golang/pkg/config"
	appConf "github.com/kukkar/mta-hosting-optimizer/conf"
)

//initConfig initialises the Global Test Config
func InitTestConfig() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	cm := new(config.ConfigManager)
	fmt.Printf("### %v", path+"/"+confFile)
	cm.InitializeGlobalConfig(path + "/" + confFile)
	cm.UpdateConfigFromEnv(config.GlobalAppConfig, "global")
	cm.UpdateConfigFromEnv(config.GlobalAppConfig.ApplicationConfig, "")
}

func RegisterTestConfig() {
	config.RegisterConfig(new(appConf.AppConfig))
	config.RegisterConfigEnvUpdateMap(appConf.EnvUpdateMap())
	config.RegisterGlobalEnvUpdateMap(config.GlobalEnvUpdateMap())
}
