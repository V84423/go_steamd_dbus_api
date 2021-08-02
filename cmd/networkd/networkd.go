package networkd

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/vishvananda/netlink"
)

// netlink client
// @see [https://github.com/systemd/systemd/blob/e18f21e34924d02dd7c330a644149d89fcc38042/src/network/networkctl.c#L791]
type Networkd struct {
	*netlink.Handle
}

// create a new networkd client
func NewNetworkd() (*Networkd, error) {
	h, err := netlink.NewHandle(netlink.FAMILY_ALL)
	if err != nil {
		return nil, err
	}
	log.Print("connected to Networkd")
	return &Networkd{
		Handle: h,
	}, nil
}

// get links
func (n *Networkd) ListLinks() error {
	links, err := n.LinkList()
	if err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"IDX", "LINK", "TYPE", "OPERATIONAL", "MAC", "IPs"})

	for _, link := range links {
		attrs := link.Attrs()
		addrs, err := n.AddrList(link, netlink.FAMILY_V4)
		if err != nil {
			log.Printf("failed to get address for link[%d:%s]: %s", attrs.Index, attrs.Name, err.Error())
			continue
		}
		var ips bytes.Buffer
		for _, addr := range addrs {
			ips.WriteString(addr.IP.String())
		}

		table.Append([]string{
			fmt.Sprint(attrs.Index),
			attrs.Name,
			link.Type(),
			attrs.OperState.String(),
			attrs.HardwareAddr.String(),
			ips.String(),
		})
	}
	table.Render()
	return nil
}
