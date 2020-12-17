package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"hf_api/models"
	_ "hf_api/routers"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

func main() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		panic(err)
	}
	err = orm.RegisterDataBase("default", "postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
			beego.AppConfig.String("DBUsername"),
			beego.AppConfig.String("DBPassword"),
			beego.AppConfig.String("appname"),
			beego.AppConfig.String("DBHost"),
			beego.AppConfig.String("DBPort")))
	if err != nil {
		panic(err)
	}

	orm.RegisterModel(new(models.User),
		new(models.Profile),
		new(models.Recipe),
		new(models.NutritionalInfo),
		new(models.Ingredient),
		new(models.Menu))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
	orm.RunCommand()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
