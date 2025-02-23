package card

func NeedPoweroff(oid string) (need bool) {
	if ts, ok := traps[oid]; ok {
		need = ts.shutdown
	}
	return
}
