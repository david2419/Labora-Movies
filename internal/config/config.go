package config

import "github.com/spf13/viper"

type Configuration struct {
	Server_Adress string
	DatabaseUrl   string
	Access_Token  string
	Secret_Key    string
}

func Config() (*Configuration, error) {
	viper.SetConfigName("configuracion")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg = &Configuration{
		Server_Adress: viper.GetString("server_Adress"),
		DatabaseUrl:   viper.GetString("dataBase_url"),
		Access_Token:  viper.GetString("access_token"),
		Secret_Key:    viper.GetString("secret_key"),
	}
	return cfg, nil
}
