// @author AlphaSnow

package g

import (
	"go.uber.org/zap"
	"server/pkg/xlog"
)

func Log(names ...string) (l *zap.Logger) {
	var err error
	var logName string
	if len(names) == 0 {
		logName = "app"
	} else {
		logName = names[0]
	}
	insKey := "logger." + logName
	if globalContainer.Has(insKey) {
		l = globalContainer.Get(insKey).(*zap.Logger)
	} else {
		l, err = xlog.ProvideZapLogger(
			Config(),
			logName,
			Config().GetBool("app.debug"),
		)
		if err != nil {
			panic(err)
		}
		globalContainer.Set(insKey, l)
	}
	return l
}
