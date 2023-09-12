package xlog

import "os"

type FileHandler struct {
	*BaseHandler
	FilePath string
}

func NewFileHandler(path string) *FileHandler {
	return &FileHandler{
		BaseHandler: NewBaseHandler(),
		FilePath:    path,
	}
}

func (h *FileHandler) HandleRecord(r Record) error {
	fh, err := os.OpenFile(
		h.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		return err
	}
	defer fh.Close()
	_, err = fh.WriteString(r.Format(h.Formatter) + "\n")
	return err
}
