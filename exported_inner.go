package logs

import (
	"context"
	"os"
	"time"
)

var (
	inner = &Logger{
		Out: stderr(),
		Formatter: &TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		},
		Hooks:        make(LevelHooks),
		Level:        InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
)

func InnerLog() *Logger {
	return inner
}

func InnerSetReportCaller(include bool) {
	inner.SetReportCaller(include)
}

func InnerSetLevel(level Level) {
	inner.SetLevel(level)
}

func InnerGetLevel() Level {
	return inner.GetLevel()
}

func InnerAddHook(hook Hook) {
	inner.AddHook(hook)
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func InnerWithError(err error) *Entry {
	return inner.WithField(ErrorKey, err)
}

// WithContext creates an entry from the standard logger and adds a context to it.
func InnerWithContext(ctx context.Context) *Entry {
	return inner.WithContext(ctx)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func InnerWithField(key string, value interface{}) *Entry {
	return inner.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func InnerWithFields(fields Fields) *Entry {
	return inner.WithFields(fields)
}

// WithTime creats an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func InnerWithTime(t time.Time) *Entry {
	return inner.WithTime(t)
}

// Trace logs a message at level Trace on the standard logger.
func InnerTrace(args ...interface{}) {
	inner.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func InnerDebug(args ...interface{}) {
	inner.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func InnerPrint(args ...interface{}) {
	inner.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func InnerInfo(args ...interface{}) {
	inner.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func InnerWarn(args ...interface{}) {
	inner.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func InnerWarning(args ...interface{}) {
	inner.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func InnerError(args ...interface{}) {
	inner.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func InnerPanic(args ...interface{}) {
	inner.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func InnerFatal(args ...interface{}) {
	inner.Fatal(args...)
}

// Tracef logs a message at level Trace on the standard logger.
func InnerTracef(format string, args ...interface{}) {
	inner.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func InnerDebugf(format string, args ...interface{}) {
	inner.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func InnerPrintf(format string, args ...interface{}) {
	inner.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func InnerInfof(format string, args ...interface{}) {
	inner.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func InnerWarnf(format string, args ...interface{}) {
	inner.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func InnerWarningf(format string, args ...interface{}) {
	inner.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func InnerErrorf(format string, args ...interface{}) {
	inner.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func InnerPanicf(format string, args ...interface{}) {
	inner.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func InnerFatalf(format string, args ...interface{}) {
	inner.Fatalf(format, args...)
}

// Traceln logs a message at level Trace on the standard logger.
func InnerTraceln(args ...interface{}) {
	inner.Traceln(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func InnerDebugln(args ...interface{}) {
	inner.Debugln(args...)
}

// Println logs a message at level Info on the standard logger.
func InnerPrintln(args ...interface{}) {
	inner.Println(args...)
}

// Infoln logs a message at level Info on the standard logger.
func InnerInfoln(args ...interface{}) {
	inner.Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func InnerWarnln(args ...interface{}) {
	inner.Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func InnerWarningln(args ...interface{}) {
	inner.Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func InnerErrorln(args ...interface{}) {
	inner.Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func InnerPanicln(args ...interface{}) {
	inner.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func InnerFatalln(args ...interface{}) {
	inner.Fatalln(args...)
}
