package promrus

import "github.com/sirupsen/logrus"

type options struct {
	levels []logrus.Level
}

var defaultLevels = []logrus.Level{logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel}

func newDefaultOptions() options {
	return options{
		levels: defaultLevels,
	}
}

type Option func(*options)

// ForLevels allows overriding the default levels which are counted. All other levels are not counted.
// The default levels are DebugLevel, InfoLevel, WarnLevel, & ErrorLevel.
func ForLevels(levels ...logrus.Level) Option {
	return func(opts *options) {
		opts.levels = levels
	}
}
