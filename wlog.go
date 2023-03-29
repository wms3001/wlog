package wlog

import (
	"log"
)

type Wlog struct {
}

func (wlog *Wlog) Show(message string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix)
	log.SetPrefix("Message:")
	log.Println(message)
}
