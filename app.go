package main

import (
	"pokos/lib"
	"pokos/lib/database"

	"github.com/guark/guark/app"
	"github.com/guark/guark/engine"
	"github.com/guark/guark/log"
)

func main() {
	a := &app.App{
		Log:     log.New("app"),
		Hooks:   lib.Hooks,
		Funcs:   lib.Funcs,
		Embed:   lib.Embeds,
		Plugins: lib.Plugins,
	}

	if err := a.Use(engine.New(a)); err != nil {
		a.Log.Fatal(err)
	}
	defer a.Quit()
	defer database.CloseDb()

	if err := a.Run(); err != nil {
		a.Log.Fatal(err)
	}
}
