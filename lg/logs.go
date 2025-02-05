package lg

var (
	serverLogRatio  int64 = 1000
	clientLogRatio  int64 = 1000
	enableMatchMask int64 = 1
	//textFinder      *sensitive_finder_engine.StreamTextFinder
	//V2              *logs.ByteDLogger
)

//func recovery() {
//	err := recover()
//	if err == nil {
//		return
//	}
//	LarkError(context.Background(), "panic accur", Data("err", err), Data("Stack", string(debug.Stack())))
//}
//
//func init() {
//	writers := make([]writer.LogWriter, 0)
//}
//
//func LarkError(ctx context.Context, title string, extras ...MessageOpt) {
//	f := LogMessage(ctx, LevelLarkError, title, extras)
//	V2.Error().With(ctx).KVs(f.Format().LogContent()...).Emit()
//
//	if canSend(f.ErrInfo()) {
//		noticeDefault(title, f.FormatNotice().NoticeContent())
//	}
//	hlog.CtxError()
//}
