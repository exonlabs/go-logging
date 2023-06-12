package xlog

type Handler interface {
	InitFormatter(string)
	SetFormatter(string)
	HandleRecord(*Record) error
}

type BaseHandler struct {
	Formatter string
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

func (h *BaseHandler) InitFormatter(s string) {
	if h.Formatter == "" {
		h.Formatter = s
	}
}

func (h *BaseHandler) SetFormatter(s string) {
	h.Formatter = s
}
