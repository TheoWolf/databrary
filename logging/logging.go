package log

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/pressly/chi/middleware"
	"github.com/rifflock/lfshook"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"net/http"
)

var Logger *logrus.Logger
var errFmtString string

func InitLgr(conf *viper.Viper) *logrus.Logger {
	if Logger != nil {
		return Logger
	}

	log_path := conf.GetString("log.path")
	log_level := conf.GetString("log.level")

	lvl, err := logrus.ParseLevel(log_level)
	if err != nil {
		panic("can't parse log level")
	}

	if lvl == logrus.DebugLevel {
		errFmtString = "%+v"
	} else {
		errFmtString = "%v"
	}

	Logger = logrus.New()
	Logger.Level = lvl

	// roratelogs config
	writer, err := rotatelogs.New(
		log_path+".%Y%m%d%H%M", // rotation pattern
		rotatelogs.WithLinkName(log_path),
		rotatelogs.WithRotationTime(time.Duration(86400)*time.Second), // rotate once a day
		rotatelogs.WithMaxAge(time.Duration(604800)*time.Second),      // keep one week of log files
	)
	if err != nil {
		panic("couldn't create rotate logs writer")
	}

	hook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.InfoLevel:  writer,
		logrus.ErrorLevel: writer,
	})
	hook.SetFormatter(&logrus.JSONFormatter{})
	Logger.Hooks.Add(hook)
	return Logger
}

func buildError(err error, msgf string, args []interface{}) error {
	msg := fmt.Sprintf(msgf, args...)
	if err != nil {
		err = errors.Wrap(err, msg)
	} else {
		err = errors.New(msg)
	}
	return err
}

// if you pass a nil error to any of these it will create a new error (so don't need to use errors.New at call site
func EntryWrapErr(entry *logrus.Entry, err error, msgf string, args ...interface{}) (error, string) {
	err = buildError(err, msgf, args)
	errorUuid := uuid.NewV4().String()
	entry.WithField("error_uuid", errorUuid).Errorf(errFmtString, err)
	return err, errorUuid
}

func LogWrapErr(err error, msgf string, args ...interface{}) (error, string) {
	err = buildError(err, msgf, args)
	errorUuid := uuid.NewV4().String()
	Logger.WithField("error_uuid", errorUuid).Errorf(errFmtString, err)
	return err, errorUuid
}

func EntryWrapErrLogWarn(entry *logrus.Entry, err error, msgf string, args ...interface{}) {
	err = buildError(err, msgf, args)
	entry.Warnf(errFmtString, err)
}

func WrapErrLogWarn(err error, msgf string, args ...interface{}) {
	err = buildError(err, msgf, args)
	Logger.Warnf(errFmtString, err)
}

func EntryWrapErrLogInfo(entry *logrus.Entry, err error, msgf string, args ...interface{}) {
	err = buildError(err, msgf, args)
	entry.Infof(errFmtString, err)
}

func WrapErrLogInfo(err error, msgf string, args ...interface{}) {
	err = buildError(err, msgf, args)
	Logger.Infof(errFmtString, err)
}

func EntryWrapErrLogFatal(entry *logrus.Entry, err error, msgf string, args ...interface{}) {
	err = buildError(err, msgf, args)
	entry.Fatalf(errFmtString, err)
}

func WrapErrLogFatal(err error, msgf string, args ...interface{}) {
	err = buildError(err, msgf, args)
	Logger.Fatalf(errFmtString, err)
}

// stolen from somewhere chi example

func NewStructuredLogger(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

type StructuredLogger struct {
	Logger *logrus.Logger
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{Logger: logrus.NewEntry(l.Logger)}
	logFields := logrus.Fields{}

	logFields["ts"] = time.Now().UTC().Format(time.RFC1123)

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	logFields["http_scheme"] = scheme
	logFields["http_proto"] = r.Proto
	logFields["http_method"] = r.Method

	logFields["remote_addr"] = r.RemoteAddr
	logFields["user_agent"] = r.UserAgent()

	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	entry.Logger = entry.Logger.WithFields(logFields)

	entry.Logger.Infoln("request started")

	return entry
}

type StructuredLoggerEntry struct {
	Logger logrus.FieldLogger
}

func (l *StructuredLoggerEntry) Write(status, bytes int, elapsed time.Duration) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"resp_status": status, "resp_bytes_length": bytes,
		"resp_elasped_ms": float64(elapsed.Nanoseconds()) / 1000000.0,
	})

	l.Logger.Infoln("request complete")
}

func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}

// Helper methods used by the application to get the request-scoped
// logger entry and set additional fields between handlers.
//
// This is a useful pattern to use to set state on the entry as it
// passes through the handler chain, which at any point can be logged
// with a call to .Print(), .Info(), etc.

func GetLogEntry(r *http.Request) logrus.FieldLogger {
	entry := middleware.GetLogEntry(r).(*StructuredLoggerEntry)
	return entry.Logger
}

func LogEntrySetField(r *http.Request, key string, value interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*StructuredLoggerEntry); ok {
		entry.Logger = entry.Logger.WithField(key, value)
	}
}

func LogEntrySetFields(r *http.Request, fields map[string]interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*StructuredLoggerEntry); ok {
		entry.Logger = entry.Logger.WithFields(fields)
	}
}
