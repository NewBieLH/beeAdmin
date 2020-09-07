package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
)

func init()  {
	runmode:=beego.AppConfig.String("runmode")
	confpath:=""
	if runmode == "dev"{
		confpath="dev.conf"
	}else{
		confpath="pro.conf"
	}
	conf,err:=config.NewConfig("conf",confpath)
	if err !=nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	dbhost := conf.String("dbhost")
	dbport := conf.String("dbport")
	dbuser := conf.String("dbuser")
	dbpassword := conf.String("dbpassword")
	dbname := conf.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dsn)
	//orm.RegisterModel(new(User), new(Category),new(Post), new(Config), new(Comment))
}
//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}