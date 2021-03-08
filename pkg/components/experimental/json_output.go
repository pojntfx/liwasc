package experimental

import (
	"encoding/json"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type JSONOutputComponent struct {
	app.Compo

	Object interface{}
}

func (c *JSONOutputComponent) Render() app.UI {
	output, err := json.Marshal(c.Object)
	if err != nil {
		panic(err)
	}

	return app.Code().Text(
		string(output),
	)
}
