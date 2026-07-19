package wireguard

import "fmt"

type FirewallManager struct {
	Interface string
}

func NewFirewallManager() FirewallManager {

	return FirewallManager{
		Interface: "wg0",
	}
}

func (f FirewallManager) EnableForwardingCommand() string {

	return "sysctl -w net.ipv4.ip_forward=1"
}

func (f FirewallManager) EnableNATCommand(
	interfaceName string,
) string {

	return fmt.Sprintf(
		"iptables -t nat -A POSTROUTING -o %s -j MASQUERADE",
		interfaceName,
	)
}

func (f FirewallManager) AllowForwardCommand() string {

	return fmt.Sprintf(
		"iptables -A FORWARD -i %s -j ACCEPT",
		f.Interface,
	)
}
