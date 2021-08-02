package systemd

import (
	"fmt"
	"log"

	"github.com/godbus/dbus/v5"
)

// systemd dbus client [https://www.freedesktop.org/wiki/Software/systemd/dbus/]
type Systemd struct {
	*dbus.Conn
	dbus.BusObject
}

// create a new systemd client
func NewSystemd() (*Systemd, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	log.Print("connected to systemd")
	return &Systemd{
		Conn:      conn,
		BusObject: conn.Object("org.freedesktop.systemd1", "/org/freedesktop/systemd1"),
	}, nil
}

func (s *Systemd) Call(method string, flags dbus.Flags, args ...interface{}) *dbus.Call {
	call := fmt.Sprint(s.Destination(), ".", method)
	log.Printf("executing call : %s (%+v)\n", call, args)
	return s.BusObject.Call(call, flags, args...)
}

// reboot system
func (s *Systemd) Reboot() error {
	return s.Call("Manager.Reboot", 0).Err
}

// shutdown system
func (s *Systemd) Shutdown() error {
	return s.Call("Manager.PowerOff", 0).Err
}

// restart service
func (s *Systemd) RestartService(name string) error {
	return s.Call("Manager.RestartUnit", 0, name, "replace").Err
}
