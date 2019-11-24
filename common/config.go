package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"time"
)

type ServiceConfig struct {
	ServiceName   string
	LogFile       string
	JaegerAddr    string
	ConsulAddr    string
	ConsulPort    int
	HttpAddr      string
	DbUri         string
	Port          int
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

func ConfigFromFileName(config string) (serviceConfig ServiceConfig, err error) {
	viper.SetConfigFile(config)
	if err = viper.ReadInConfig(); err != nil {
		return
	}

	serviceConfig = ServiceConfig{
		ServiceName:   viper.GetString("service_name"),
		LogFile:       viper.GetString("log_file"),
		JaegerAddr:    viper.GetString("jaeger_addr"),
		ConsulAddr:    viper.GetString("consul_addr"),
		ConsulPort:    viper.GetInt("consul_port"),
		HttpAddr:      viper.GetString("address"),
		DbUri:         viper.GetString("db_uri"),
		Port:          viper.GetInt("port"),
		RedisHost:     viper.GetString("redis_host"),
		RedisPort:     viper.GetString("redis_port"),
		RedisPassword: viper.GetString("redis_password"),
		RedisDB:       viper.GetInt("redis_db"),
	}
	return
}

func (serviceConfig ServiceConfig) GetDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", serviceConfig.DbUri)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	return
}

func (conf ServiceConfig) Migrate(db *gorm.DB, models []BaseModel) (err error) {
	for _, value := range models {
		db.AutoMigrate(value)
	}
	return nil
}
