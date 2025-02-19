package card

const (
	TrapCommunicationLost             = ".1.3.6.1.4.1.935.0.1"
	TrapUpsOverLoad                   = ".1.3.6.1.4.1.935.0.2"
	TrapUpsDiagnosticsFailed          = ".1.3.6.1.4.1.935.0.3"
	TrapUpsDischarged                 = ".1.3.6.1.4.1.935.0.4"
	TrapUpsOnBattery                  = ".1.3.6.1.4.1.935.0.5"
	TrapBoostOn                       = ".1.3.6.1.4.1.935.0.6"
	TrapLowBattery                    = ".1.3.6.1.4.1.935.0.7"
	TrapCommunicationEstablished      = ".1.3.6.1.4.1.935.0.8"
	TrapPowerRestored                 = ".1.3.6.1.4.1.935.0.9"
	TrapUpsDiagnosticsPassed          = ".1.3.6.1.4.1.935.0.10"
	TrapReturnFromLowBattery          = ".1.3.6.1.4.1.935.0.11"
	TrapUpsTurnedOff                  = ".1.3.6.1.4.1.935.0.12"
	TrapUpsSleeping                   = ".1.3.6.1.4.1.935.0.13"
	TrapUpsWokeUp                     = ".1.3.6.1.4.1.935.0.14"
	TrapUpsRebootStarted              = ".1.3.6.1.4.1.935.0.15"
	TrapEnvOverTemperature            = ".1.3.6.1.4.1.935.0.16"
	TrapEnvTemperatureNormal          = ".1.3.6.1.4.1.935.0.17"
	TrapEnvOverHumidity               = ".1.3.6.1.4.1.935.0.18"
	TrapEnvHumidityNormal             = ".1.3.6.1.4.1.935.0.19"
	TrapEnvSmokeAbnormal              = ".1.3.6.1.4.1.935.0.20"
	TrapEnvWaterAbnormal              = ".1.3.6.1.4.1.935.0.21"
	TrapEnvSecurityAbnormal           = ".1.3.6.1.4.1.935.0.22"
	TrapEnvWaterNormal                = ".1.3.6.1.4.1.935.0.24"
	TrapEnvGasAbnormal                = ".1.3.6.1.4.1.935.0.26"
	TrapUpsTemp                       = ".1.3.6.1.4.1.935.0.27"
	TrapUpsLoadNormal                 = ".1.3.6.1.4.1.935.0.28"
	TrapUpsTempNormal                 = ".1.3.6.1.4.1.935.0.29"
	TrapEnvUnderTemperature           = ".1.3.6.1.4.1.935.0.30"
	TrapEnvUnderHumidity              = ".1.3.6.1.4.1.935.0.31"
	TrapUpsBypass                     = ".1.3.6.1.4.1.935.0.32"
	TrapEnvSecurity1                  = ".1.3.6.1.4.1.935.0.33"
	TrapEnvSecurity2                  = ".1.3.6.1.4.1.935.0.34"
	TrapEnvSecurity3                  = ".1.3.6.1.4.1.935.0.35"
	TrapEnvSecurity4                  = ".1.3.6.1.4.1.935.0.36"
	TrapEnvSecurity5                  = ".1.3.6.1.4.1.935.0.37"
	TrapEnvSecurity6                  = ".1.3.6.1.4.1.935.0.38"
	TrapEnvSecurity7                  = ".1.3.6.1.4.1.935.0.39"
	TrapUpsRecroterror                = ".1.3.6.1.4.1.935.0.47"
	TrapUpsBypassFreFail              = ".1.3.6.1.4.1.935.0.48"
	TrapUpsBypassacnormal             = ".1.3.6.1.4.1.935.0.49"
	TrapUpsBypassacabnormal           = ".1.3.6.1.4.1.935.0.50"
	TrapUpsTest                       = ".1.3.6.1.4.1.935.0.51"
	TrapUpsScheduleShutdown           = ".1.3.6.1.4.1.935.0.52"
	TrapUpsBypassReturn               = ".1.3.6.1.4.1.935.0.53"
	TrapUpsShortCircuitShutdown       = ".1.3.6.1.4.1.935.0.54"
	TrapUpsInverterOutputFailShutdown = ".1.3.6.1.4.1.935.0.55"
	TrapUpsBypassBreadkerOnShutdown   = ".1.3.6.1.4.1.935.0.56"
	TrapUpsHighDCShutdown             = ".1.3.6.1.4.1.935.0.57"
	TrapUpsEmergencyStop              = ".1.3.6.1.4.1.935.0.58"
	TrapUpsInverterMode               = ".1.3.6.1.4.1.935.0.59"
	TrapUpsBypassMode                 = ".1.3.6.1.4.1.935.0.60"
	TrapUpsOverTemperatureShutdown    = ".1.3.6.1.4.1.935.0.61"
	TrapUpsOverLoadShutdown           = ".1.3.6.1.4.1.935.0.62"
	TrapUpsCapacityUnderrun           = ".1.3.6.1.4.1.935.0.63"
	TrapUpsCapacityNormal             = ".1.3.6.1.4.1.935.0.64"
	TrapUpsLowBatteryShutdown         = ".1.3.6.1.4.1.935.0.67"
	TrapBuckOn                        = ".1.3.6.1.4.1.935.0.68"
	TrapReturnFromBuck                = ".1.3.6.1.4.1.935.0.69"
	TrapReturnFromBoost               = ".1.3.6.1.4.1.935.0.70"
)

type trapSpec struct {
	name string // имя трапа
	desc string // описание трапа
	send bool   // отправлять ли оповещение о трапе в канал
}

var traps = map[string]trapSpec{
	TrapCommunicationLost:             trapSpec{name: "communicationLost", desc: "Communication to the UPS has been lost", send: false},
	TrapUpsOverLoad:                   trapSpec{name: "upsOverLoad", desc: "The UPS has sensed a load greater than 100 percent of its rated capacity", send: false},
	TrapUpsDiagnosticsFailed:          trapSpec{name: "upsDiagnosticsFailed", desc: "The UPS has failed its internal diagnostic self-test", send: false},
	TrapUpsDischarged:                 trapSpec{name: "upsDischarged", desc: "The UPS has just started a runtime calibration discharge", send: false},
	TrapUpsOnBattery:                  trapSpec{name: "upsOnBattery", desc: "The UPS has switched to battery backup power", send: false},
	TrapBoostOn:                       trapSpec{name: "boostOn", desc: "The UPS has enabled Boost", send: false},
	TrapLowBattery:                    trapSpec{name: "lowBattery", desc: "The UPS batteries are low and will soon be exhausted", send: false},
	TrapCommunicationEstablished:      trapSpec{name: "communicationEstablished", desc: "Communication with the UPS has been established", send: false},
	TrapPowerRestored:                 trapSpec{name: "powerRestored", desc: "Normal power has been restored to the UPS", send: false},
	TrapUpsDiagnosticsPassed:          trapSpec{name: "upsDiagnosticsPassed", desc: "The UPS has passed its internal self-test", send: false},
	TrapReturnFromLowBattery:          trapSpec{name: "returnFromLowBattery", desc: "The UPS has returned from a low battery condition", send: false},
	TrapUpsTurnedOff:                  trapSpec{name: "upsTurnedOff", desc: "The UPS is being switched off by a management station", send: false},
	TrapUpsSleeping:                   trapSpec{name: "upsSleeping", desc: "The UPS is going into sleep mode", send: false},
	TrapUpsWokeUp:                     trapSpec{name: "upsWokeUp", desc: "The UPS has woke up from sleep mode", send: false},
	TrapUpsRebootStarted:              trapSpec{name: "upsRebootStarted", desc: "The UPS has started its reboot sequence", send: false},
	TrapEnvOverTemperature:            trapSpec{name: "envOverTemperature", desc: "The environment temperature exceed the normal value", send: false},
	TrapEnvTemperatureNormal:          trapSpec{name: "envTemperatureNormal", desc: "The environment temperature is normal", send: false},
	TrapEnvOverHumidity:               trapSpec{name: "envOverHumidity", desc: "The environment humidity exceed the normal value", send: false},
	TrapEnvHumidityNormal:             trapSpec{name: "envHumidityNormal", desc: "The environment humidity is normal", send: false},
	TrapEnvSmokeAbnormal:              trapSpec{name: "envSmokeAbnormal", desc: "Smoke is abnormal", send: false},
	TrapEnvWaterAbnormal:              trapSpec{name: "envWaterAbnormal", desc: "Water is abnormal", send: false},
	TrapEnvSecurityAbnormal:           trapSpec{name: "envSecurityAbnormal", desc: "Security is abnormal", send: false},
	TrapEnvWaterNormal:                trapSpec{name: "envWaterNormal", desc: "Water is normal", send: false},
	TrapEnvGasAbnormal:                trapSpec{name: "envGasAbnormal", desc: "Gas Alarm", send: false},
	TrapUpsTemp:                       trapSpec{name: "upsTemp", desc: "UPS Temperature Overrun", send: false},
	TrapUpsLoadNormal:                 trapSpec{name: "upsLoadNormal", desc: "UPS Load Normal", send: false},
	TrapUpsTempNormal:                 trapSpec{name: "upsTempNormal", desc: "UPS temperature Normal", send: false},
	TrapEnvUnderTemperature:           trapSpec{name: "envUnderTemperature", desc: "The environment temperature below the normal value", send: false},
	TrapEnvUnderHumidity:              trapSpec{name: "envUnderHumidity", desc: "The environment humidity below the normal value", send: false},
	TrapUpsBypass:                     trapSpec{name: "upsBypass", desc: "The UPS is entering 'bypass' mode", send: false},
	TrapEnvSecurity1:                  trapSpec{name: "envSecurity1", desc: "Security1 Alarm", send: false},
	TrapEnvSecurity2:                  trapSpec{name: "envSecurity2", desc: "Security2 Alarm", send: false},
	TrapEnvSecurity3:                  trapSpec{name: "envSecurity3", desc: "Security3 Alarm", send: false},
	TrapEnvSecurity4:                  trapSpec{name: "envSecurity4", desc: "Security4 Alarm", send: false},
	TrapEnvSecurity5:                  trapSpec{name: "envSecurity5", desc: "Security5 Alarm", send: false},
	TrapEnvSecurity6:                  trapSpec{name: "envSecurity6", desc: "Security6 Alarm", send: false},
	TrapEnvSecurity7:                  trapSpec{name: "envSecurity7", desc: "Security7 Alarm", send: false},
	TrapUpsRecroterror:                trapSpec{name: "upsRecroterror", desc: "Rectifier Rotation Error", send: false},
	TrapUpsBypassFreFail:              trapSpec{name: "upsBypassFreFail", desc: "Bypass Frequency Fail", send: false},
	TrapUpsBypassacnormal:             trapSpec{name: "upsBypassacnormal", desc: "Bypass AC Normal", send: false},
	TrapUpsBypassacabnormal:           trapSpec{name: "upsBypassacabnormal", desc: "Bypass AC Abnormal", send: false},
	TrapUpsTest:                       trapSpec{name: "upsTest", desc: "UPS Test", send: false},
	TrapUpsScheduleShutdown:           trapSpec{name: "upsScheduleShutdown", desc: "UPS Schedule Shutdown", send: false},
	TrapUpsBypassReturn:               trapSpec{name: "upsBypassReturn", desc: "The UPS return from Bypass Mode", send: false},
	TrapUpsShortCircuitShutdown:       trapSpec{name: "upsShortCircuitShutdown", desc: "The UPS Short Circuit Shutdown", send: false},
	TrapUpsInverterOutputFailShutdown: trapSpec{name: "upsInverterOutputFailShutdown", desc: "The UPS Inverter Output Fail Shutdown", send: false},
	TrapUpsBypassBreadkerOnShutdown:   trapSpec{name: "upsBypassBreadkerOnShutdown", desc: "The UPS Manual Bypass Breaker on Shutdown", send: false},
	TrapUpsHighDCShutdown:             trapSpec{name: "upsHighDCShutdown", desc: "The UPS High DC Shutdown", send: false},
	TrapUpsEmergencyStop:              trapSpec{name: "upsEmergencyStop", desc: "The UPS Emergency Stop", send: false},
	TrapUpsInverterMode:               trapSpec{name: "upsInverterMode", desc: "The UPS Static Switch in Inverter Mode", send: false},
	TrapUpsBypassMode:                 trapSpec{name: "upsBypassMode", desc: "The UPS Static Switch in Bypass Mode", send: false},
	TrapUpsOverTemperatureShutdown:    trapSpec{name: "upsOverTemperatureShutdown", desc: "The UPS Over Temperature Shutdown", send: false},
	TrapUpsOverLoadShutdown:           trapSpec{name: "upsOverLoadShutdown", desc: "The UPS Overload Shutdown", send: false},
	TrapUpsCapacityUnderrun:           trapSpec{name: "upsCapacityUnderrun", desc: "The UPS Capacity Underrun", send: false},
	TrapUpsCapacityNormal:             trapSpec{name: "upsCapacityNormal", desc: "The UPS Capacity Normal", send: false},
	TrapUpsLowBatteryShutdown:         trapSpec{name: "upsLowBatteryShutdown", desc: "The UPS Low Battery Shutdown", send: false},
	TrapBuckOn:                        trapSpec{name: "buckOn", desc: "The UPS has enabled Buck", send: false},
	TrapReturnFromBuck:                trapSpec{name: "returnFromBuck", desc: "The UPS has return from Buck", send: false},
	TrapReturnFromBoost:               trapSpec{name: "returnFromBoost", desc: "The UPS has return from Boost", send: false},
}
