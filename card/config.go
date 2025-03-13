package card

var configOids = []string{
	UpsSmartConfigRatedOutputVoltage,
	UpsSmartConfigHighTransferVolt,
	UpsSmartConfigLowTransferVolt,
	UpsSmartConfigAlarm,
	UpsSmartConfigAlarmTimer,
	UpsSmartConfigMinReturnCapacity,
	UpsSmartConfigSensitivity,
	UpsSmartConfigLowBatteryRunTime,
	UpsSmartConfigReturnDelay,
	UpsSmartConfigShutoffDelay,
	UpsSmartConfigUpsSleepTime,
}

func GetConfig(addr, community string) (map[string]any, error) {
	return getOids(addr, community, configOids)
}
