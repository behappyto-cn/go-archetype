package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1"
	port     = 3306
	database = ""
)

// DbDefault 默认数据库
var DbDefault *sqlx.DB

// DB 初始化数据库
func DB() {
	// 初始化数据库
	db, err := connMysql(username, password, hostname, database, port)
	if err != nil {
		log.Error(err)
		return
	}
	DbDefault = db
}

// 连接MySQL数据库
func connMysql(username, password, hostname, database string, port int) (*sqlx.DB, error) {
	connConf := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local&allowOldPasswords=1",
		username, password, hostname, port, database)
	db, err := sqlx.Open("mysql", connConf)
	if err != nil {
		log.Errorf("连接数据库[%s:%d/%s]失败, err = %s", hostname, port, database, err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Errorf("连接数据库[%s:%d/%s]失败, err = %s", hostname, port, database, err)
		return nil, err
	}
	log.Infof("连接数据库[%s:%d/%s]成功", hostname, port, database)
	return db, nil
}
