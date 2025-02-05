package logs

import (
	"context"
	"fmt"
	"os"

	"github.com/erickxeno/clib/logs/writer"
)

var (
	V1            *CLogger
	defaultLogger *CompatLogger
)

func init() {
	// TODO: NewAgentWriter relies on init function, refactor agent SDK and supports be defined as variable
	disable := os.Getenv("_DISABLE_LOG_AUTO_INIT")
	if disable == "True" {
		return
	}

	Init()
}

func userHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return home
}

func Init() {
	// TODO: NewAgentWriter relies on init function, refactor agent SDK and supports be defined as variable
	writers := make([]writer.LogWriter, 0)
	level := DebugLevel
	ops := make([]Option, 0)
	isInTCE := true //env.InTCE()
	psm := "psm"    //env.PSM()
	if isInTCE {
		level = InfoLevel
		fileName := fmt.Sprintf("%s/Documents/tiger/log/app/%s.log", userHomeDir(), psm)
		writers = append(writers, writer.NewAsyncWriter(writer.NewFileWriter(fileName, writer.Hourly), true))
		writers = append(writers, writer.NewAgentWriter())
		// for test
		writers = append(writers, writer.NewConsoleWriter(writer.SetColorful(true)))
	} else {
		writers = append(writers, writer.NewConsoleWriter(writer.SetColorful(true)))
	}
	ops = append(ops, SetWriter(level, writers...))
	V1 = NewCLogger(ops...)
	defaultLogger = NewCompatLoggerFrom(V1, WithCallDepthOffset(2))
}

// SetDefaultLogger resets the default logger with specified options,
// it is not thread-safe and please only call it in program initialization.
// Libraries should not config the default logger and leaves it to the application user,
// otherwise loggers would cover each other.
func SetDefaultLogger(ops ...Option) {
	curOps := V1.GetOptions()
	curOps = append(curOps, ops...)
	*V1 = *NewCLogger(curOps...)
	*defaultLogger = *NewCompatLoggerFrom(V1, WithCallDepthOffset(2))
}

// SetLevel sets the minimal level for defaultLogger. It is safe to increase the level.
// Please not decrease the level directly. Use SetLevelForWriters instead.
func SetLevel(newLevel Level) {
	if defaultLogger == nil {
		return
	}
	defaultLogger.v1.SetLevel(newLevel)
}

// SetLevelForWriters updates the minimal level for the writer of the defaultLogger.
// It may also update the loggers' level.
func SetLevelForWriters(newLevel Level, logWriters ...writer.LogWriter) {
	if defaultLogger == nil {
		return
	}
	defaultLogger.v1.SetLevelForWriters(newLevel, logWriters...)
}

// GetWriters returns the level writers
func GetWriters() []leveledWriter {
	if defaultLogger == nil {
		return nil
	}
	return defaultLogger.v1.GetWriter()
}

// EnableDynamicLogLevel enableds dynamic context log level for compatibleLogger.
func EnableDynamicLogLevel() {
	WithDynamicLevel(true)(defaultLogger)
}

// Fatal works like logs.Fatal.
func Fatal(format string, v ...interface{}) {
	defaultLogger.Fatal(format, v...)
}

// Error works like logs.Erorr.
func Error(format string, v ...interface{}) {
	defaultLogger.Error(format, v...)
}

// Warn works like logs.Warn.
func Warn(format string, v ...interface{}) {
	defaultLogger.Warn(format, v...)
}

// Notice works like logs.Notice.
func Notice(format string, v ...interface{}) {
	defaultLogger.Notice(format, v...)
}

// Info works like logs.Info.
func Info(format string, v ...interface{}) {
	defaultLogger.Info(format, v...)
}

// Debug works like logs.Debug.
func Debug(format string, v ...interface{}) {
	defaultLogger.Debug(format, v...)
}

// Trace works like logs.Trace.
func Trace(format string, v ...interface{}) {
	defaultLogger.Trace(format, v...)
}

// Fatalf works like logs.Fatalf.
func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatal(format, v...)
}

// Errorf works like logs.Errorf.
func Errorf(format string, v ...interface{}) {
	defaultLogger.Error(format, v...)
}

// Warnf works like logs.Warnf.
func Warnf(format string, v ...interface{}) {
	defaultLogger.Warn(format, v...)
}

// Noticef works like logs.Noticef.
func Noticef(format string, v ...interface{}) {
	defaultLogger.Notice(format, v...)
}

// Infof works like logs.Infof.
func Infof(format string, v ...interface{}) {
	defaultLogger.Info(format, v...)
}

// Debugf works like logs.Debugf.
func Debugf(format string, v ...interface{}) {
	defaultLogger.Debug(format, v...)
}

// Tracef works like logs.Tracef.
func Tracef(format string, v ...interface{}) {
	defaultLogger.Trace(format, v...)
}

// CtxFatal works like logs.CtxFatal.
func CtxFatal(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxFatal(ctx, format, v...)
}

// CtxError works like logs.CtxError.
func CtxError(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxError(ctx, format, v...)
}

// CtxWarn works like logs.CtxWarn.
func CtxWarn(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxWarn(ctx, format, v...)
}

// CtxNotice works like logs.CtxNotice.
func CtxNotice(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxNotice(ctx, format, v...)
}

// CtxInfo works like logs.CtxInfo.
func CtxInfo(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxInfo(ctx, format, v...)
}

// CtxDebug works like logs.CtxDebug.
func CtxDebug(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxDebug(ctx, format, v...)
}

// CtxTrace works like logs.CtxTrace.
func CtxTrace(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxTrace(ctx, format, v...)
}

// CtxFatalsf works like logs.CtxFatalsf.
func CtxFatalsf(ctx context.Context, format string, v ...string) {
	defaultLogger.CtxFatalsf(ctx, format, v...)
}

// CtxErrorsf works like logs.CtxErrorsf.
func CtxErrorsf(ctx context.Context, format string, v ...string) {
	defaultLogger.CtxErrorsf(ctx, format, v...)
}

// CtxErrorsf works like logs.CtxErrorsf.
func CtxWarnsf(ctx context.Context, format string, v ...string) {
	defaultLogger.CtxWarnsf(ctx, format, v...)
}

// CtxNoticesf works like logs.CtxNoticesf.
func CtxNoticesf(ctx context.Context, format string, v ...string) {
	defaultLogger.CtxNoticesf(ctx, format, v...)
}

// CtxErrorsf works like logs.CtxErrorsf.
func CtxInfosf(ctx context.Context, format string, v ...string) {
	defaultLogger.CtxInfosf(ctx, format, v...)
}

// CtxDebugsf works like logs.CtxDebugsf.
func CtxDebugsf(ctx context.Context, format string, v ...string) {
	defaultLogger.CtxDebugsf(ctx, format, v...)
}

// CtxTracesf works like logs.CtxDebugsf.
func CtxTracesf(ctx context.Context, format string, v ...string) {
	defaultLogger.CtxTracesf(ctx, format, v...)
}

// CtxFatalKVs provides function like logs.CtxFatalKVs.
func CtxFatalKVs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxFatalKVs(ctx, kvs...)
}

// CtxErrorKVs provides function like logs.CtxErrorKVs.
func CtxErrorKVs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxErrorKVs(ctx, kvs...)
}

// CtxWarnKVs provides function like logs.CtxWarnKVs.
func CtxWarnKVs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxWarnKVs(ctx, kvs...)
}

// CtxNoticeKVs provides function like logs.CtxNoticeKVs.
func CtxNoticeKVs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxNoticeKVs(ctx, kvs...)
}

// CtxInfoKVs provides function like logs.CtxInfoKVs.
func CtxInfoKVs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxInfoKVs(ctx, kvs...)
}

// CtxDebugKVs provides function like logs.CtxDebugKVs.
func CtxDebugKVs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxDebugKVs(ctx, kvs...)
}

// CtxTraceKVs provides function like logs.CtxTraceKVs.
func CtxTraceKVs(ctx context.Context, kvs ...interface{}) {
	defaultLogger.CtxTraceKVs(ctx, kvs...)
}

// CtxFlushNotice provides function like logs.CtxFlushNotice
func CtxFlushNotice(ctx context.Context) {
	ntc := GetNotice(ctx)
	if ntc == nil {
		return
	}
	kvs := ntc.KVs()
	if len(kvs) == 0 {
		return
	}
	defaultLogger.CtxNoticeKVs(ctx, kvs...)
}

// PrintStack prints the stacks in Info level.
// If printAllGoroutines is true, it prints the stacks of all goroutines.
// Otherwise, it just prints the current goroutine's stack.
func PrintStack(printAllGoroutines bool) {
	defaultLogger.PrintStack(printAllGoroutines)
}

func Flush() {
	defaultLogger.Flush()
}

func Stop() {
	defaultLogger.Close()
}
