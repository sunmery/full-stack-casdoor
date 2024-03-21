package main

func initAuthConfig() {
	casdoorsdk.InitConfig(
		GlobalConfig.Server.Endpoint,
		GlobalConfig.Server.ClientID,
		GlobalConfig.Server.ClientSecret,
		GlobalConfig.Certificate,
		GlobalConfig.Server.Organization,
		GlobalConfig.Server.Application,
	)
}

func main() {
	viper.SetConfigFile("./config.yaml")

}
