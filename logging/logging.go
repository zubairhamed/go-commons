package logging

import "log"

func LogError(e error) {
	if e != nil {
		log.Println("[ERROR] ", e)
	}
}

func LogWarn(msg string) {
	log.Println("[WARN] ", msg)
}

func LogInfo(msg string) {
	log.Println("[INFO] ", msg)
}