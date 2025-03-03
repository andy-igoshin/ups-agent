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
		l.logger.SetPrefix("[alert] ")
		l.logger.Printf(format, v...)
	}
}

func (l *Logger) Crit(format string, v ...any) {
	if l.level >= syslog.LOG_CRIT {
		l.logger.SetPrefix("[crit] ")
		l.logger.Printf(format, v...)
	}
}

func (l *Logger) Debug(format string, v ...any) {
	if l.level >= syslog.LOG_DEBUG {
		l.logger.SetPrefix("[debug] ")
		l.logger.Printf(format, v...)
	}
}

func (l *Logger) Emerg(v ...any) {
	if l.level >= syslog.LOG_EMERG {
		l.logger.SetPrefix("[emerg] ")
		l.logger.Print(v...)
	}
}

func (l *Logger) Emergf(format string, v ...any) {
	if l.level >= syslog.LOG_EMERG {
		l.logger.SetPrefix("[emerg] ")
		l.logger.Printf(format, v...)
	}
}

func (l *Logger) Error(format string, v ...any) {
	if l.level >= syslog.LOG_ERR {
		l.logger.SetPrefix("[error] ")
		l.logger.Printf(format, v...)
	}
}

func (l *Logger) Fatalf(format string, v ...any) {
	l.logger.SetPrefix("[crit] ")
	l.logger.Fatalf(format, v...)
}

func (l *Logger) Fatalln(v ...any) {
	l.logger.SetPrefix("[crit] ")
	l.logger.Fatalln(v...)
}

func (l *Logger) Info(format string, v ...any) {
	if l.level >= syslog.LOG_INFO {
		l.logger.SetPrefix("[info] ")
		l.logger.Printf(format, v...)
	}
}

func (l *Logger) Notice(format string, v ...any) {
	if l.level >= syslog.LOG_NOTICE {
		l.logger.SetPrefix("[notice] ")
		l.logger.Printf(format, v...)
	}
}

func (l *Logger) Warn(format string, v ...any) {
	if l.level >= syslog.LOG_WARNING {
		l.logger.SetPrefix("[warn] ")
		l.logger.Printf(format, v...)
	}
}
