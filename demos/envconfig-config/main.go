package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kelseyhightower/envconfig"
)

const (
	EnvVarRedisHost = "REDIS_HOST"
	EnvVarRedisPort = "REDIS_PORT"
)

type Config struct {
	Redis *RedisConfig
}

func (c *Config) String() string {
	return fmt.Sprintf("[redis=%v]", c.Redis)
}

type RedisConfig struct {
	Host     string `envconfig:"REDIS_HOST" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD" required:"true"`
	Port     int    `envconfig:"REDIS_PORT" required:"true"`
}

func (rc *RedisConfig) String() string {
	return fmt.Sprintf("[host=%s,port=%d,password=%s]", rc.Host, rc.Port, rc.Password)
}

func main() {
	config, err := initConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", config)
	for {
	}
}

func initConfig() (c *Config, err error) {
	c = &Config{}
	err = envconfig.Process("app", c)
	return
}

func lookupEnvInt(envvar string) (i int, ok bool) {
	var value string
	var err error

	if value, ok = os.LookupEnv(envvar); !ok {
		return
	}

	if i, err = strconv.Atoi(value); err != nil {
		panic(err)
	}
	return
}

func formatConfigError(envvar string) (err error) {
	return fmt.Errorf("%s is not defined in the environment.", envvar)
}
