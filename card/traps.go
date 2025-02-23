package card

const (
	trapCommunicationLost             = ".1.3.6.1.4.1.935.0.1"
	trapUpsOverLoad                   = ".1.3.6.1.4.1.935.0.2"
	trapUpsDiagnosticsFailed          = ".1.3.6.1.4.1.935.0.3"
	trapUpsDischarged                 = ".1.3.6.1.4.1.935.0.4"
	trapUpsOnBattery                  = ".1.3.6.1.4.1.935.0.5"
	trapBoostOn                       = ".1.3.6.1.4.1.935.0.6"
	trapLowBattery                    = ".1.3.6.1.4.1.935.0.7"
	trapCommunicationEstablished      = ".1.3.6.1.4.1.935.0.8"
	trapPowerRestored                 = ".1.3.6.1.4.1.935.0.9"
	trapUpsDiagnosticsPassed          = ".1.3.6.1.4.1.935.0.10"
	trapReturnFromLowBattery          = ".1.3.6.1.4.1.935.0.11"
	trapUpsTurnedOff                  = ".1.3.6.1.4.1.935.0.12"
	trapUpsSleeping                   = ".1.3.6.1.4.1.935.0.13"
	trapUpsWokeUp                     = ".1.3.6.1.4.1.935.0.14"
	trapUpsRebootStarted              = ".1.3.6.1.4.1.935.0.15"
	trapEnvOverTemperature            = ".1.3.6.1.4.1.935.0.16"
	trapEnvTemperatureNormal          = ".1.3.6.1.4.1.935.0.17"
	trapEnvOverHumidity               = ".1.3.6.1.4.1.935.0.18"
	trapEnvHumidityNormal             = ".1.3.6.1.4.1.935.0.19"
	trapEnvSmokeAbnormal              = ".1.3.6.1.4.1.935.0.20"
	trapEnvWaterAbnormal              = ".1.3.6.1.4.1.935.0.21"
	trapEnvSecurityAbnormal           = ".1.3.6.1.4.1.935.0.22"
	trapEnvWaterNormal                = ".1.3.6.1.4.1.935.0.24"
	trapEnvGasAbnormal                = ".1.3.6.1.4.1.935.0.26"
	trapUpsTemp                       = ".1.3.6.1.4.1.935.0.27"
	trapUpsLoadNormal                 = ".1.3.6.1.4.1.935.0.28"
	trapUpsTempNormal                 = ".1.3.6.1.4.1.935.0.29"
	trapEnvUnderTemperature           = ".1.3.6.1.4.1.935.0.30"
	trapEnvUnderHumidity              = ".1.3.6.1.4.1.935.0.31"
	trapUpsBypass                     = ".1.3.6.1.4.1.935.0.32"
	trapEnvSecurity1                  = ".1.3.6.1.4.1.935.0.33"
	trapEnvSecurity2                  = ".1.3.6.1.4.1.935.0.34"
	trapEnvSecurity3                  = ".1.3.6.1.4.1.935.0.35"
	trapEnvSecurity4                  = ".1.3.6.1.4.1.935.0.36"
	trapEnvSecurity5                  = ".1.3.6.1.4.1.935.0.37"
	trapEnvSecurity6                  = ".1.3.6.1.4.1.935.0.38"
	trapEnvSecurity7                  = ".1.3.6.1.4.1.935.0.39"
	trapUpsRecroterror                = ".1.3.6.1.4.1.935.0.47"
	trapUpsBypassFreFail              = ".1.3.6.1.4.1.935.0.48"
	trapUpsBypassacnormal             = ".1.3.6.1.4.1.935.0.49"
	trapUpsBypassacabnormal           = ".1.3.6.1.4.1.935.0.50"
	trapUpsTest                       = ".1.3.6.1.4.1.935.0.51"
	trapUpsScheduleShutdown           = ".1.3.6.1.4.1.935.0.52"
	trapUpsBypassReturn               = ".1.3.6.1.4.1.935.0.53"
	trapUpsShortCircuitShutdown       = ".1.3.6.1.4.1.935.0.54"
	trapUpsInverterOutputFailShutdown = ".1.3.6.1.4.1.935.0.55"
	trapUpsBypassBreadkerOnShutdown   = ".1.3.6.1.4.1.935.0.56"
	trapUpsHighDCShutdown             = ".1.3.6.1.4.1.935.0.57"
	trapUpsEmergencyStop              = ".1.3.6.1.4.1.935.0.58"
	trapUpsInverterMode               = ".1.3.6.1.4.1.935.0.59"
	trapUpsBypassMode                 = ".1.3.6.1.4.1.935.0.60"
	trapUpsOverTemperatureShutdown    = ".1.3.6.1.4.1.935.0.61"
	trapUpsOverLoadShutdown           = ".1.3.6.1.4.1.935.0.62"
	trapUpsCapacityUnderrun           = ".1.3.6.1.4.1.935.0.63"
	trapUpsCapacityNormal             = ".1.3.6.1.4.1.935.0.64"
	trapUpsLowBatteryShutdown         = ".1.3.6.1.4.1.935.0.67"
	trapBuckOn                        = ".1.3.6.1.4.1.935.0.68"
	trapReturnFromBuck                = ".1.3.6.1.4.1.935.0.69"
	trapReturnFromBoost               = ".1.3.6.1.4.1.935.0.70"
)

type trapSpec struct {
	name     string // имя трапа
	send     bool   // отправлять ли оповещение о трапе в канал
	shutdown bool   // должен ли приводить трап к немедленному завершению работы ОС
}

var traps = map[string]trapSpec{
	trapCommunicationLost:             trapSpec{name: "communicationLost"},
	trapUpsOverLoad:                   trapSpec{name: "upsOverLoad"},
	trapUpsDiagnosticsFailed:          trapSpec{name: "upsDiagnosticsFailed"},
	trapUpsDischarged:                 trapSpec{name: "upsDischarged", send: true},
	trapUpsOnBattery:                  trapSpec{name: "upsOnBattery", send: true},
	trapBoostOn:                       trapSpec{name: "boostOn"},
	trapLowBattery:                    trapSpec{name: "lowBattery", send: true},
	trapCommunicationEstablished:      trapSpec{name: "communicationEstablished"},
	trapPowerRestored:                 trapSpec{name: "powerRestored", send: true},
	trapUpsDiagnosticsPassed:          trapSpec{name: "upsDiagnosticsPassed"},
	trapReturnFromLowBattery:          trapSpec{name: "returnFromLowBattery", send: true},
	trapUpsTurnedOff:                  trapSpec{name: "upsTurnedOff", send: true},
	trapUpsSleeping:                   trapSpec{name: "upsSleeping", send: true},
	trapUpsWokeUp:                     trapSpec{name: "upsWokeUp"},
	trapUpsRebootStarted:              trapSpec{name: "upsRebootStarted", send: true},
	trapEnvOverTemperature:            trapSpec{name: "envOverTemperature"},
	trapEnvTemperatureNormal:          trapSpec{name: "envTemperatureNormal"},
	trapEnvOverHumidity:               trapSpec{name: "envOverHumidity"},
	trapEnvHumidityNormal:             trapSpec{name: "envHumidityNormal"},
	trapEnvSmokeAbnormal:              trapSpec{name: "envSmokeAbnormal"},
	trapEnvWaterAbnormal:              trapSpec{name: "envWaterAbnormal"},
	trapEnvSecurityAbnormal:           trapSpec{name: "envSecurityAbnormal"},
	trapEnvWaterNormal:                trapSpec{name: "envWaterNormal"},
	trapEnvGasAbnormal:                trapSpec{name: "envGasAbnormal"},
	trapUpsTemp:                       trapSpec{name: "upsTemp"},
	trapUpsLoadNormal:                 trapSpec{name: "upsLoadNormal"},
	trapUpsTempNormal:                 trapSpec{name: "upsTempNormal"},
	trapEnvUnderTemperature:           trapSpec{name: "envUnderTemperature"},
	trapEnvUnderHumidity:              trapSpec{name: "envUnderHumidity"},
	trapUpsBypass:                     trapSpec{name: "upsBypass"},
	trapEnvSecurity1:                  trapSpec{name: "envSecurity1"},
	trapEnvSecurity2:                  trapSpec{name: "envSecurity2"},
	trapEnvSecurity3:                  trapSpec{name: "envSecurity3"},
	trapEnvSecurity4:                  trapSpec{name: "envSecurity4"},
	trapEnvSecurity5:                  trapSpec{name: "envSecurity5"},
	trapEnvSecurity6:                  trapSpec{name: "envSecurity6"},
	trapEnvSecurity7:                  trapSpec{name: "envSecurity7"},
	trapUpsRecroterror:                trapSpec{name: "upsRecroterror"},
	trapUpsBypassFreFail:              trapSpec{name: "upsBypassFreFail"},
	trapUpsBypassacnormal:             trapSpec{name: "upsBypassacnormal"},
	trapUpsBypassacabnormal:           trapSpec{name: "upsBypassacabnormal"},
	trapUpsTest:                       trapSpec{name: "upsTest"},
	trapUpsScheduleShutdown:           trapSpec{name: "upsScheduleShutdown", send: true, shutdown: true},
	trapUpsBypassReturn:               trapSpec{name: "upsBypassReturn"},
	trapUpsShortCircuitShutdown:       trapSpec{name: "upsShortCircuitShutdown", send: true, shutdown: true},
	trapUpsInverterOutputFailShutdown: trapSpec{name: "upsInverterOutputFailShutdown", send: true, shutdown: true},
	trapUpsBypassBreadkerOnShutdown:   trapSpec{name: "upsBypassBreadkerOnShutdown", send: true},
	trapUpsHighDCShutdown:             trapSpec{name: "upsHighDCShutdown", send: true, shutdown: true},
	trapUpsEmergencyStop:              trapSpec{name: "upsEmergencyStop", send: true, shutdown: true},
	trapUpsInverterMode:               trapSpec{name: "upsInverterMode"},
	trapUpsBypassMode:                 trapSpec{name: "upsBypassMode"},
	trapUpsOverTemperatureShutdown:    trapSpec{name: "upsOverTemperatureShutdown", send: true, shutdown: true},
	trapUpsOverLoadShutdown:           trapSpec{name: "upsOverLoadShutdown", send: true, shutdown: true},
	trapUpsCapacityUnderrun:           trapSpec{name: "upsCapacityUnderrun"},
	trapUpsCapacityNormal:             trapSpec{name: "upsCapacityNormal"},
	trapUpsLowBatteryShutdown:         trapSpec{name: "upsLowBatteryShutdown", send: true, shutdown: true},
	trapBuckOn:                        trapSpec{name: "buckOn"},
	trapReturnFromBuck:                trapSpec{name: "returnFromBuck"},
	trapReturnFromBoost:               trapSpec{name: "returnFromBoost"},
}
