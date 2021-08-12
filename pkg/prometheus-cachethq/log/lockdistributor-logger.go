package log

type lockDistributorLogger struct {
	logger Logger
}

func (ldl *lockDistributorLogger) Println(v ...interface{}) {
	// Log as debug to avoid to much log
	ldl.logger.Debugln(v...)
}
