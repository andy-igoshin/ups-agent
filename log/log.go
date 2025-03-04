package log

import (
	"log"
	"log/syslog"
	"os"
	"path"

	"git.vsu.ru/ai/ups-agent/config"
)

type Logger struct {
	file     *os.File
	logger   *log.Logger
	level    syslog.Priority
	logdir   string
	filename string
}

func New(conf *config.Config) *Logger {

	var err error

	l := Logger{}
	l.logdir = conf.Logger.LogDir
	l.filename = conf.Logger.FileName

	l.file, err = os.OpenFile(path.Join(l.logdir, l.filename), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Open log error: ", err)
	}

	l.logger = log.New(l.file, "", log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix)
	l.level = l.getLevel(conf.Logger.Level)

	return &l
}

func (l *Logger) Close() {
	l.file.Close()
}

func (l *Logger) Reopen() {

	var err error

	l.file.Close()
	l.file, err = os.OpenFile(path.Join(l.logdir, l.filename), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Reopen log error: %s", err)
	}

	l.logger.SetOutput(l.file)
}

func (l *Logger) Alert(format string, v ...any) {
	if l.level >= syslog.LOG_ALERT {
		l.logger.Printf("[alert] "+format, v...)
	}
}

func (l *Logger) Crit(format string, v ...any) {
	if l.level >= syslog.LOG_CRIT {
		l.logger.Printf("[crit] "+format, v...)
	}
}

func (l *Logger) Debug(format string, v ...any) {
	if l.level >= syslog.LOG_DEBUG {
		l.logger.Printf("[debug] "+format, v...)
	}
}

func (l *Logger) Emerg(v ...any) {
	if l.level >= syslog.LOG_EMERG {
		l.logger.Print("[emerg]", v)
	}
}

func (l *Logger) Emergf(format string, v ...any) {
	if l.level >= syslog.LOG_EMERG {
		l.logger.Printf("[emerg] "+format, v...)
	}
}

func (l *Logger) Error(format string, v ...any) {
	if l.level >= syslog.LOG_ERR {
		l.logger.Printf("[error] "+format, v...)
	}
}

func (l *Logger) Fatalf(format string, v ...any) {
	l.logger.Fatalf("[crit] "+format, v...)
}

func (l *Logger) Fatalln(v ...any) {
	l.logger.Fatalln("[crit]", v)
}

func (l *Logger) Info(format string, v ...any) {
	if l.level >= syslog.LOG_INFO {
		l.logger.Printf("[info] "+format, v...)
	}
}

func (l *Logger) Notice(format string, v ...any) {
	if l.level >= syslog.LOG_NOTICE {
		l.logger.Printf("[notice] "+format, v...)
	}
}

func (l *Logger) Warn(format string, v ...any) {
	if l.level >= syslog.LOG_WARNING {
		l.logger.Printf("[warn] "+format, v...)
	}
}
