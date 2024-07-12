package main

import (
	"log"
	"openfeature-simple-rest/controller"
	pkgFlagsmith "openfeature-simple-rest/pkg/flagsmith"

	"github.com/labstack/echo/v4"
	"github.com/open-feature/go-sdk/openfeature"
)

func main() {
	flagsmithProvider, err := pkgFlagsmith.NewProvider(&pkgFlagsmith.Config{
		EnvKey: `ser.ioYojLppkgpd53LrRCmUz9`,
	})
	if err != nil {
		log.Fatal(`Failed to create Flagsmith provider`, err)
	}

	openfeature.SetProvider(flagsmithProvider)
	oc := openfeature.NewClient(flagsmithProvider.Metadata().Name)

	greetController, err := controller.NewGreet(oc)
	if err != nil {
		log.Fatal(`Failed to create Greet controller`, err)
	}

	e := echo.New()
	group := e.Group("/greet")
	group.GET("", greetController.Greet)
	e.Logger.Fatal(e.Start(":1323"))
}
