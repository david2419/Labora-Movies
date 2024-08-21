package config

import "github.com/spf13/viper"

type Configuration struct {
	Server_Adress          string
	DatabaseUrl            string
	ApiMovies_Access_Token string
	JWT_Secret_Key         string
}

func Config() (*Configuration, error) {
	viper.SetConfigName("configuracion")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg = &Configuration{
		Server_Adress:          viper.GetString("server_Adress"),
		DatabaseUrl:            viper.GetString("dataBase_url"),
		ApiMovies_Access_Token: viper.GetString("apiMovies_access_token"),
		JWT_Secret_Key:         viper.GetString("jwt_secret_key"),
	}
	return cfg, nil
}
