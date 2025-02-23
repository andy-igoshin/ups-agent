package log

import (
	"log/syslog"
)

func (l *Logger) getLevel(ls string) syslog.Priority {

	var lev syslog.Priority

	switch ls {
	case "emerg":
		lev = syslog.LOG_EMERG

	case "alert":
		lev = syslog.LOG_ALERT

	case "crit":
		lev = syslog.LOG_CRIT

	case "error":
		lev = syslog.LOG_ERR

	case "warn":
		lev = syslog.LOG_WARNING

	case "notice":
		lev = syslog.LOG_NOTICE

	case "info":
		lev = syslog.LOG_INFO

	case "debug":
		lev = syslog.LOG_DEBUG

	default:
		lev = syslog.LOG_INFO
		l.Warn("Logger has incorrect log level '%s', default level 'info' assumed", ls)
	}

	return lev
}

func getLevelString(lev syslog.Priority) string {

	var ls string

	switch lev {
	case syslog.LOG_EMERG:
		ls = "emerg"

	case syslog.LOG_ALERT:
		ls = "alert"

	case syslog.LOG_CRIT:
		ls = "crit"

	case syslog.LOG_ERR:
		ls = "error"

	case syslog.LOG_WARNING:
		ls = "warn"

	case syslog.LOG_NOTICE:
		ls = "notice"

	case syslog.LOG_INFO:
		ls = "info"

	case syslog.LOG_DEBUG:
		ls = "debug"
	}

	return ls
}
