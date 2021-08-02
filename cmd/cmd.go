package cmd

import (
	"github.com/med8bra/gosystem/cmd/networkd"
	"github.com/med8bra/gosystem/cmd/systemd"
)

type SystemdCmd systemd.Cmd
type NetworkdCmd networkd.Cmd

type Cmd struct {
	SystemdCmd
	NetworkdCmd
}
