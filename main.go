package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

type MySQLConfig struct {
	UserName string
	Password string
	Host     string
	Port     int
	Name     string
	TimeOut  string
}

func (m *MySQLConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		m.UserName, m.Password, m.Host, m.Port, m.Name, m.TimeOut)
}

func main() {
	config := MySQLConfig{
		UserName: "",
		Password: "",
		Host:     "",
		Port:     3306,
		Name:     "",
		TimeOut:  "",
	}
	var err error
	DB, err = gorm.Open(mysql.Open(config.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	people := &People{
		Name: "rory",
		Age:  24,
		AddressList: AddressList{
			{
				Country:  "China",
				Province: "Guangdong",
				City:     "Guangzhou",
				Detail:   "xxx",
			},
		},
	}
	if err = DB.Model(&People{}).Save(&people).Error; err != nil {
		log.Fatalln(err)
	}
}
