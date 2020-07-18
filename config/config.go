package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

type db struct {
	Database   string `json:"database"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AuthSource string `json:"authsource"`
}

type server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type configuration struct {
	// DB
	DB     db     `json:"database"`
	Server server `json:"server"`
}

// DB configuration
var (
	DB     *db
	Server *server
)

// Init - Initialize application
func Init(fname *string) {
	file, _ := os.Open(*fname)
	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	fmt.Println(*fname)
	fmt.Println(string(data))
	config := configuration{}
	err := json.Unmarshal(data, &config)
	if nil != err {
		logrus.Fatalln("error:", err)
	}

	fmt.Println(config)
	DB = &config.DB
	Server = &config.Server
}
