package log

func (l *Logger) getLevel(ls string) (lev int) {

	switch ls {
	case "emerg":
		lev = log_EMERG

	case "alert":
		lev = log_ALERT

	case "crit":
		lev = log_CRIT

	case "error":
		lev = log_ERR

	case "warn":
		lev = log_WARNING

	case "notice":
		lev = log_NOTICE

	case "info":
		lev = log_INFO

	case "debug":
		lev = log_DEBUG

	default:
		lev = log_INFO
		l.Warn("Logger has incorrect log level '%s', default level 'info' assumed", ls)
	}

	return
}

func getLevelString(lev int) string {

	var ls string

	switch lev {
	case log_EMERG:
		ls = "emerg"

	case log_ALERT:
		ls = "alert"

	case log_CRIT:
		ls = "crit"

	case log_ERR:
		ls = "error"

	case log_WARNING:
		ls = "warn"

	case log_NOTICE:
		ls = "notice"

	case log_INFO:
		ls = "info"

	case log_DEBUG:
		ls = "debug"
	}

	return ls
}
