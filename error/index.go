package error

import "log"

func Handle(err error) {
	if err != nil {
		log.Fatalln("ERROR: ", err, "\nnow you need to restart...")
	}
}
