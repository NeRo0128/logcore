package writer

import "os"

type Writer interface {
	Write([]byte) (int, error)
}

// todo aclopar al logger
type StdoutWriter struct{}

func (StdoutWriter) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}
