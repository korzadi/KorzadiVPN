package wireguard

import "fmt"

type InterfaceController struct {
	Name string
}

func NewInterfaceController() InterfaceController {

	return InterfaceController{
		Name: "wg0",
	}
}

func (i InterfaceController) CreateCommand() string {

	return fmt.Sprintf(
		"ip link add %s type wireguard",
		i.Name,
	)
}

func (i InterfaceController) SetAddressCommand(
	address string,
) string {

	return fmt.Sprintf(
		"ip address add %s dev %s",
		address,
		i.Name,
	)
}

func (i InterfaceController) UpCommand() string {

	return fmt.Sprintf(
		"ip link set %s up",
		i.Name,
	)
}
