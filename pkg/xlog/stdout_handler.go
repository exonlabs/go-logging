package xlog

import "os"

type StdoutHandler struct {
	*BaseHandler
}

func NewStdoutHandler() *StdoutHandler {
	return &StdoutHandler{
		BaseHandler: NewBaseHandler(),
	}
}

func (h *StdoutHandler) HandleRecord(r Record) error {
	_, err := os.Stdout.WriteString(r.Format(h.Formatter) + "\n")
	return err
}
