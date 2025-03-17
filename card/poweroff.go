package card

var poweroffOids = map[string]int{
	UpsSmartControlUpsOff: 2,
}

func Poweroff(addr, community string) error {
	return SetOids(addr, community, poweroffOids)
}
