package main

import (
	"github.com/alecthomas/kong"
	"github.com/med8bra/gosystem/cmd"
)

func main() {
	ctx := kong.Parse(&cmd.Cmd{},
		kong.Name("gosystem"),
		kong.Description("execute systemd actions via Go"),
		kong.ConfigureHelp(kong.HelpOptions{Compact: true}),
	)
	if err := ctx.Run(); err != nil {
		ctx.Fatalf("%s: %s", ctx.Selected().Name, err.Error())
	}
}
