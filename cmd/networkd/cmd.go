package networkd

var client *Networkd

func getNetworkd() (*Networkd, error) {
	if client != nil {
		return client, nil
	}
	client, err := NewNetworkd()
	return client, err
}

type Cmd struct {
	Net   Net   `kong:"cmd,help='list links'"`
	Renew Renew `kong:"cmd,help='Renew dynamic configurations'"`
}

type Net struct{}

func (*Net) Run() error {
	client, err := getNetworkd()
	if err != nil {
		return err
	}
	return client.ListLinks()
}

type Renew struct {
	Devices []string `kong:"help='interfaces to renew'"`
}

func (*Renew) Run() error {
	client, err := getNetworkd()
	if err != nil {
		return err
	}
	//TODO:
	return client.ListLinks()
}
