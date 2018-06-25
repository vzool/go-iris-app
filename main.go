package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/vzool/vZoolShipment/app/http/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.RegisterView(iris.HTML("./app/views", ".html"))

	app.StaticWeb("/public", "./public")

	// Web Routes

	mvc.Configure(app.Party("/basic"), controllers.BasicController)
	mvc.Configure(app.Party("/country"), controllers.CountryController)
	mvc.Configure(app.Party("/country/{country_id}/city"), controllers.CityController)

	app.Run(iris.Addr("localhost:8080"))
}
