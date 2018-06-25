package controllers

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type cityController struct {
	Logger LoggerService

	Session *sessions.Session
}

// CityController controller
func CityController(app *mvc.Application) {

	app.Handle(new(cityController))
}

// Routes

func (c *cityController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "Get")
	b.Handle("GET", "/create", "Create")
	b.Handle("POST", "/", "Store")
	b.Handle("GET", "/{city_id}", "Show")
	b.Handle("GET", "/{city_id}/edit", "Edit")
	b.Handle("PUT", "/{city_id}", "Put")
	b.Handle("DELETE", "/{city_id}", "Delete")
}

// Funcs

func (c *cityController) Get(ctx iris.Context) string {
	countryID, _ := ctx.Params().GetInt("country_id")
	return fmt.Sprint("City Index - Country Show: ", countryID)
}

func (c *cityController) Create(ctx iris.Context) string {
	countryID, _ := ctx.Params().GetInt("country_id")
	return fmt.Sprint("City Create Form - Country Show: ", countryID)
}

func (c *cityController) Store(ctx iris.Context) string {
	countryID, _ := ctx.Params().GetInt("country_id")
	return fmt.Sprint("City Store - Country Show: ", countryID)
}

func (c *cityController) Show(ctx iris.Context) string {
	countryID, _ := ctx.Params().GetInt("country_id")
	cityID, _ := ctx.Params().GetInt("city_id")
	return fmt.Sprint("City Show - Country Show: ", cityID, countryID)
}

func (c *cityController) Edit(ctx iris.Context) string {
	countryID, _ := ctx.Params().GetInt("country_id")
	cityID, _ := ctx.Params().GetInt("city_id")
	return fmt.Sprint("City Edit Form - Country Show: ", cityID, countryID)
}

func (c *cityController) Put(ctx iris.Context) string {
	countryID, _ := ctx.Params().GetInt("country_id")
	cityID, _ := ctx.Params().GetInt("city_id")
	return fmt.Sprint("City Put - Country Show: ", cityID, countryID)
}

func (c *cityController) Delete(ctx iris.Context) string {
	countryID, _ := ctx.Params().GetInt("country_id")
	cityID, _ := ctx.Params().GetInt("city_id")
	return fmt.Sprint("City Delete - Country Show: ", cityID, countryID)
}
