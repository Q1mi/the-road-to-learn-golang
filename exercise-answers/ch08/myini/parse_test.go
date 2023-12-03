package myini

import (
	"fmt"
	"os"
	"testing"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MySQLConf  MySQLConfig  `ini:"mysql"`
}

type ServerConfig struct {
	IP   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MySQLConfig struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Debug    bool   `ini:"debug"`
}

func TestParse(t *testing.T) {
	data, err := os.ReadFile("./config.ini")
	if err != nil {
		panic(fmt.Errorf("open file err, err:%v", err))
	}
	var confData Config
	err = Parse(data, &confData)
	if err != nil {
		t.Errorf("Parse failed, err:%v", err)
		return
	}
	t.Log("Parse success!")
	t.Logf("confData:%#v\n", confData)
}
