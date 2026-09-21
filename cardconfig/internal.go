package cardconfig

import (
	"fmt"
	"time"

	"git.vsu.ru/ai/ups-agent/card"
)

func getAlarm(st map[string]any) (s string) {
	if v, ok := st[card.UpsSmartConfigAlarm]; ok {
		switch v {
		case 1:
			s = "По времени (1)"
		case 2:
			s = "При низком заряде АКБ (2)"
		case 3:
			s = "Никогда (3)"
		}
	}
	return
}

func getAlarmTimer(st map[string]any) (t string) {
	if v, ok := st[card.UpsSmartConfigAlarmTimer]; ok {
		if val, ok := v.(int); ok {
			d := time.Duration(val) * time.Second
			t = d.String()
		}
	}
	return
}

func getHighTransferVolt(st map[string]any) (h string) {
	if v, ok := st[card.UpsSmartConfigHighTransferVolt]; ok {
		if val, ok := v.(int); ok {
			h = fmt.Sprintf("%.1f В", float64(val)/10)
		}
	}
	return
}

func getLowTransferVolt(st map[string]any) (l string) {
	if v, ok := st[card.UpsSmartConfigLowTransferVolt]; ok {
		if val, ok := v.(int); ok {
			l = fmt.Sprintf("%.1f В", float64(val)/10)
		}
	}
	return
}

func getLowBatteryRunTime(st map[string]any) (t string) {
	if v, ok := st[card.UpsSmartConfigLowBatteryRunTime]; ok {
		if val, ok := v.(int); ok {
			d := time.Duration(val) * time.Second
			t = d.String()
		}
	}
	return
}

func getMinReturnCapacity(st map[string]any) (l string) {
	if v, ok := st[card.UpsSmartConfigMinReturnCapacity]; ok {
		if val, ok := v.(int); ok {
			l = fmt.Sprintf("%d%%", val)
		}
	}
	return
}

func getRatedOutputVoltage(st map[string]any) (l string) {
	if v, ok := st[card.UpsSmartConfigRatedOutputVoltage]; ok {
		if val, ok := v.(int); ok {
			l = fmt.Sprintf("%.1f В", float64(val)/10)
		}
	}
	return
}

func getReturnDelay(st map[string]any) (t string) {
	if v, ok := st[card.UpsSmartConfigReturnDelay]; ok {
		if val, ok := v.(int); ok {
			d := time.Duration(val) * time.Second
			t = d.String()
		}
	}
	return
}

func getSensitivity(st map[string]any) (s string) {
	if v, ok := st[card.UpsSmartConfigSensitivity]; ok {
		switch v {
		case 1:
			s = "Авто (1)"
		case 2:
			s = "Низкая (2)"
		case 3:
			s = "Средняя (3)"
		case 4:
			s = "Высокая (4)"
		}
	}
	return
}

func getShutoffDelay(st map[string]any) (t string) {
	if v, ok := st[card.UpsSmartConfigShutoffDelay]; ok {
		if val, ok := v.(int); ok {
			d := time.Duration(val) * time.Second
			t = d.String()
		}
	}
	return
}

func getUpsSleepTime(st map[string]any) (t string) {
	if v, ok := st[card.UpsSmartConfigUpsSleepTime]; ok {
		if val, ok := v.(int); ok {
			d := time.Duration(val) * time.Minute
			t = d.String()
		}
	}
	return
}
