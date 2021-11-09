package config

/*
iniでConfigの設定ファイルを読み込む。
設定ファイルを作成する。

インストール　済
コンフィグのファイルを読み込む為のパッケージ
https://godoc.org/gopkg.in/go-ini/ini.v1
go get "gopkg.in/go-ini/ini.v1"

congig.iniを作成する
下記、よくある設定
[web]
port = 8080

[db]
name = test.sql
driver = sqlite3

これをGoで読み込む。osopenなんかも使えるが、iniの方が便利
*/

import (
	"github.com/RyoAsai-git/todo_app/utils"
	"gopkg.in/go-ini/ini.v1"
	"log"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
