package dhcp

import (
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/client4"
	"github.com/insomniacslk/dhcp/dhcpv4/nclient4"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/dhcp", new(DHCP))
}

type DHCP struct {
	client *client4.Client
	iface  string
}

func New(iface string, opts ...nclient4.ClientOpt) *DHCP {
	return &DHCP{
		iface:  iface,
		client: client4.NewClient(),
	}
}

func (d *DHCP) Exchange(modifiers ...dhcpv4.Modifier) (*dhcpv4.DHCPv4, error) {
	results, err := d.client.Exchange(d.iface, modifiers...)
	if err != nil {
		return results[0], nil
	}
	return nil, err
}
