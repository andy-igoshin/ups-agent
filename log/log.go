package log

import (
	"log"
	"os"
	"path"

	"git.vsu.ru/ai/ups-agent/config"
)

const (
	log_EMERG = iota + 1
	log_ALERT
	log_CRIT
	log_ERR
	log_WARNING
	log_NOTICE
	log_INFO
	log_DEBUG
)

type Logger struct {
	file     *os.File
	logger   *log.Logger
	level    int
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
	if l.level >= log_ALERT {
		l.logger.Printf("[alert] "+format+eol, v...)
	}
}

func (l *Logger) Crit(format string, v ...any) {
	if l.level >= log_CRIT {
		l.logger.Printf("[crit] "+format+eol, v...)
	}
}

func (l *Logger) Debug(format string, v ...any) {
	if l.level >= log_DEBUG {
		l.logger.Printf("[debug] "+format+eol, v...)
	}
}

func (l *Logger) Emerg(v ...any) {
	if l.level >= log_EMERG {
		l.logger.Print("[emerg]", v, eol)
	}
}

func (l *Logger) Emergf(format string, v ...any) {
	if l.level >= log_EMERG {
		l.logger.Printf("[emerg] "+format+eol, v...)
	}
}

func (l *Logger) Error(format string, v ...any) {
	if l.level >= log_ERR {
		l.logger.Printf("[error] "+format+eol, v...)
	}
}

func (l *Logger) Fatalf(format string, v ...any) {
	l.logger.Fatalf("[crit] "+format+eol, v...)
}

func (l *Logger) Fatalln(v ...any) {
	l.logger.Fatal("[crit]", v, eol)
}

func (l *Logger) Info(format string, v ...any) {
	if l.level >= log_INFO {
		l.logger.Printf("[info] "+format+eol, v...)
	}
}

func (l *Logger) Notice(format string, v ...any) {
	if l.level >= log_NOTICE {
		l.logger.Printf("[notice] "+format+eol, v...)
	}
}

func (l *Logger) Warn(format string, v ...any) {
	if l.level >= log_WARNING {
		l.logger.Printf("[warn] "+format+eol, v...)
	}
}
