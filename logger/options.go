package logger

// Option es un tipo de función para configurar el logger
type Option func(*loggerImpl)

func WithJSON(enabled bool) Option {
	return func(l *loggerImpl) {
		l.jsonOutput = enabled
	}
}

func WithPrettyJSON(enabled bool) Option {
	return func(l *loggerImpl) {
		l.prettyJSON = enabled
	}
}

func WithLevel(level Level) Option {
	return func(l *loggerImpl) {
		l.level = level
	}
}

func WithLayer(layer string) Option {
	return func(l *loggerImpl) {
		l.layer = layer
	}
}

func WithField(field Field) Option {
	return func(l *loggerImpl) {
		l.fields = append(l.fields, field)
	}
}

func WithCaller(enable bool) Option {
	return func(l *loggerImpl) {
		l.showCaller = enable
	}
}
