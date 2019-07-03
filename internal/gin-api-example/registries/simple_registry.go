package registries

import (
	"github.com/coby-spotim/go-application-example/internal/gin-api-example/db"
	"github.com/spf13/viper"
)

type SimpleRegistry struct {
	Env    string
	Config *viper.Viper
	DB     db.DB
}

func NewSimpleRegistry(env string) *SimpleRegistry {
	return &SimpleRegistry{
		Env:    env,
		Config: NewConfig(env),
		DB:     db.NewMock(),
	}
}

func NewConfig(env string) *viper.Viper {
	config := viper.New()
	config.AddConfigPath("./configs")
	config.SetConfigName(env)
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	config.AutomaticEnv()
	return config
}
