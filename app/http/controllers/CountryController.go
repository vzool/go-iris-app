package controllers

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type countryController struct {
	Logger LoggerService

	Session *sessions.Session
}

// CountryController controller
func CountryController(app *mvc.Application) {

	app.Handle(new(countryController))
}

// Routes

func (c *countryController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "Get")
	b.Handle("GET", "/create", "Create")
	b.Handle("POST", "/", "Store")
	b.Handle("GET", "/{country_id}", "Show")
	b.Handle("GET", "/{country_id}/edit", "Edit")
	b.Handle("PUT", "/{country_id}", "Put")
	b.Handle("DELETE", "/{country_id}", "Delete")
}

// Funcs

func (c *countryController) Get(ctx iris.Context) mvc.Result {

	return hero.View{
		Name: "country/index.html",
		Data: map[string]interface{}{
			"Title":     "Hello Page",
			"MyMessage": "Welcome to my awesome website",
		},
	}
}

func (c *countryController) Create(ctx iris.Context) string {
	return "Country Create Form"
}

func (c *countryController) Store(ctx iris.Context) string {
	return "Country Store"
}

func (c *countryController) Show(ctx iris.Context) string {
	countryID, _ := ctx.Params().GetInt("country_id")
	return fmt.Sprint("Country Show: ", countryID)
}

func (c *countryController) Edit(ctx iris.Context) string {
	countryID, _ := ctx.Params().GetInt("country_id")
	return fmt.Sprint("Country Edit Form: ", countryID)
}

func (c *countryController) Put(ctx iris.Context) string {
	countryID, _ := ctx.Params().GetInt("country_id")
	return fmt.Sprint("Country Put: ", countryID)
}

func (c *countryController) Delete(ctx iris.Context) string {
	countryID, _ := ctx.Params().GetInt("country_id")
	return fmt.Sprint("Country Delete: ", countryID)
}
