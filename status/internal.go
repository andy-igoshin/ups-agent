package status

import (
	"fmt"
	"strconv"
	"strings"
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

func getBatteryLevel(bl int) string {
	return fmt.Sprintf("%d%%", bl)
}

func getBatteryTemperature(st map[string]any) (l string) {
	if v, ok := st[card.UpsSmartBatteryTemperature]; ok {
		l = fmt.Sprintf("%.1f C", float32(v.(int))/10)
	}
	return
}

func getBatteryVoltage(st map[string]any) (l string) {
	if v, ok := st[card.UpsSmartBatteryVoltage]; ok {
		l = fmt.Sprintf("%.1f Вольт", float32(v.(int))/10)
	}
	return
}

func getCapacity(st map[string]any) (l string) {
	if v, ok := st[card.UpsSmartBatteryCapacity]; ok {
		l = fmt.Sprintf("%d%%", v)
	}
	return
}

func getFw(st map[string]any) (f string) {
	if v, ok := st[card.UpsSmartIdentFirmwareRevision]; ok {
		f = strings.TrimSpace(v.(string))
	}
	return
}

func getFwCard(st map[string]any) (f string) {
	if v, ok := st[card.UpsSmartIdentAgentFirmwareRevision]; ok {
		f = strings.TrimSpace(v.(string))
	}
	return
}

func getLineFailCause(st map[string]any) (s string) {
	if v, ok := st[card.UpsSmartInputLineFailCause]; ok {
		switch v {
		case 1:
			s = "Причина отсутствует (1)"
		case 2:
			s = "Высокое напряжение (2)"
		case 3:
			s = "Низкое напряжение (3)"
		case 4:
			s = "Отключение питания (4)"
		case 5:
			s = "Кратковременное снижение напряжения (5)"
		case 6:
			s = "Длительное снижение напряжения (6)"
		case 7:
			s = "Кратковременный малый импульс (7)"
		case 8:
			s = "Кратковременный большой импульс (8)"
		}
	}
	return
}

func getLineFreq(st map[string]any) string {
	var f float32
	if v, ok := st[card.UpsSmartInputFrequency]; ok {
		f = float32(v.(int)) / 10
	}
	return fmt.Sprintf("%.1f Гц", f)
}

func getLineVoltage(st map[string]any) string {
	var vcur float32
	var vmin float32
	var vmax float32
	if v, ok := st[card.UpsSmartInputLineVoltage]; ok {
		vcur = float32(v.(int)) / 10
	}
	if v, ok := st[card.UpsSmartInputMaxLineVoltage]; ok {
		vmax = float32(v.(int)) / 10
	}
	if v, ok := st[card.UpsSmartInputMinLineVoltage]; ok {
		vmin = float32(v.(int)) / 10
	}
	return fmt.Sprintf("%.1f Вольт (%.1f/%.1f)", vcur, vmin, vmax)
}

func getLowBatteryRunTime(st map[string]any) (t string) {
	if v, ok := st[card.UpsSmartConfigLowBatteryRunTime]; ok {
		d, err := time.ParseDuration(strconv.Itoa(v.(int)) + "s")
		if err == nil {
			t = d.String()
		}
	}
	return
}

func getHighTransferVolt(st map[string]any) (l string) {
	if v, ok := st[card.UpsSmartConfigHighTransferVolt]; ok {
		l = fmt.Sprintf("%.1f Вольт", float32(v.(int))/10)
	}
	return
}

func getLowTransferVolt(st map[string]any) (l string) {
	if v, ok := st[card.UpsSmartConfigLowTransferVolt]; ok {
		l = fmt.Sprintf("%.1f Вольт", float32(v.(int))/10)
	}
	return
}

func getMinReturnCapacity(st map[string]any) (l string) {
	if v, ok := st[card.UpsSmartConfigMinReturnCapacity]; ok {
		l = fmt.Sprintf("%d%%", v)
	}
	return
}

func getOutputLoad(st map[string]any) (l string) {
	if v, ok := st[card.UpsSmartOutputLoad]; ok {
		l = fmt.Sprintf("%d%%", v)
	}
	return
}

func getOutputVoltage(st map[string]any) (l string) {
	if v, ok := st[card.UpsSmartOutputVoltage]; ok {
		l = fmt.Sprintf("%.1f Вольт", float32(v.(int))/10)
	}
	return
}

func getReturnDelay(st map[string]any) (t string) {
	if v, ok := st[card.UpsSmartConfigReturnDelay]; ok {
		d, err := time.ParseDuration(strconv.Itoa(v.(int)) + "s")
		if err == nil {
			t = d.String()
		}
	}
	return
}

func getRunTimeRemaining(st map[string]any) (t string) {
	if v, ok := st[card.UpsSmartBatteryRunTimeRemaining]; ok {
		d, err := time.ParseDuration(strconv.Itoa(v.(int)) + "s")
		if err == nil {
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
		d, err := time.ParseDuration(strconv.Itoa(v.(int)) + "s")
		if err == nil {
			t = d.String()
		}
	}
	return
}

func getStatus(st map[string]any) (s string) {
	if v, ok := st[card.UpsBaseOutputStatus]; ok {
		switch v {
		case 1:
			s = "Неизвестно (1)"
		case 2:
			s = "Онлайн (2)"
		case 3:
			s = "Работа от АКБ (3)"
		case 4:
			s = "Повышение напряжения (4)"
		case 5:
			s = "Спящий режим (5)"
		case 6:
			s = "Байпас (6)"
		case 7:
			s = "Перезагрузка (7)"
		case 8:
			s = "Ожидание (8)"
		case 9:
			s = "Понижение напряжения (9)"
		}
	}
	return
}

func getUpsModel(st map[string]any, model string) (m string) {
	if v, ok := st[card.UpsBaseIdentModel]; ok {
		m = strings.TrimSpace(v.(string))
		if len(m) == 0 {
			m = model
		}
	}
	return
}

func getUpsName(st map[string]any, upsname string) (n string) {
	if v, ok := st[card.UpsBaseIdentUpsName]; ok {
		n = strings.TrimSpace(v.(string))
		if len(n) == 0 {
			if v, ok := st[card.SysDescr]; ok {
				n = strings.TrimSpace(v.(string))
				if len(n) == 0 {
					n = upsname
				}
			}
		}
	}
	return
}
