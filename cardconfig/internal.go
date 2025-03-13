package cardconfig

import (
//        "fmt"
        "strconv"
//        "strings"
        "time"

        "git.vsu.ru/ai/ups-agent/card"
)

func getAlarmTimer(st map[string]any) (t string) {
	if v, ok := st[card.UpsSmartConfigAlarmTimer]; ok {
		d, err := time.ParseDuration(strconv.Itoa(v.(int)) + "s")
		if err == nil {
			t = d.String()
		}
	}
	return
}
