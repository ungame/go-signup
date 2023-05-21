package utils

import (
	"io"
	"log"
)

var DefaultCloser = CloseHandler{verbose: true}

type CloseHandler struct {
	verbose bool
}

func NewCloseHandler(verbose bool) *CloseHandler {
	return &CloseHandler{verbose: verbose}
}

func (h CloseHandler) Close(closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if h.verbose {
			if err != nil {
				log.Printf("error on close %T: %s\n", closer, err.Error())
			} else {
				log.Printf("%T closed", closer)
			}
		}
	}
}
