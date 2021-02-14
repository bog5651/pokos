package lib

import (
	"pokos/lib/funcs"
	"pokos/lib/hooks"

	"github.com/guark/guark/app"
	"github.com/guark/plugins/clipboard"
	"github.com/guark/plugins/dialog"
	"github.com/guark/plugins/notify"
)

// Exposed functions to guark Javascript api.
var Funcs = app.Funcs{
	"getKKM":       funcs.GetKKM,
	"getClients":   funcs.GetClients,
	"GetModelsKKM": funcs.GetModelsKKM,
}

// App hooks.
var Hooks = app.Hooks{
	"created": hooks.Created,
	"mounted": hooks.Mounted,
}

// App plugins.
var Plugins = app.Plugins{
	"dialog":    &dialog.Plugin{},
	"notify":    &notify.Plugin{},
	"clipboard": &clipboard.Plugin{},
}
