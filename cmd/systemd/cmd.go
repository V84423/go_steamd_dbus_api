package systemd

var client *Systemd

func getSystemd() (*Systemd, error) {
	if client != nil {
		return client, nil
	}
	client, err := NewSystemd()
	return client, err
}

type Cmd struct {
	Reboot   Reboot   `kong:"cmd,help='reboot the system'"`
	Shutdown Shutdown `kong:"cmd,help='shutdown the system'"`
	Restart  Restart  `kong:"cmd,help='restart a service'"`
}

type Reboot struct{}

func (*Reboot) Run() error {
	client, err := getSystemd()
	if err != nil {
		return err
	}
	return client.Reboot()
}

type Shutdown struct{}

func (*Shutdown) Run() error {
	client, err := getSystemd()
	if err != nil {
		return err
	}
	return client.Shutdown()
}

type Restart struct {
	Service string `kong:"arg,required,help='the service to restart'"`
}

func (r *Restart) Run() error {
	client, err := getSystemd()
	if err != nil {
		return err
	}
	return client.RestartService(r.Service)
}
