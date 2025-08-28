package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func triggerAlarm(retryMap map[string]int) {
	servmonDir := os.Getenv("SERVMONDIR")
	fileName := fmt.Sprintf("%s/hush", servmonDir)
	_, err := os.Stat(fileName)
	if err == nil {
		// Hush exists, dont alarm.
	} else if errors.Is(err, os.ErrNotExist) {
		for name, value := range retryMap {
			if value >= 3 {
				msg := fmt.Sprintf("%s: is down", name)
				sendSlack(msg)
				file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
				if err != nil {
					log.Println(err)
				}

				time.Sleep(time.Hour * 1)
				err = os.Remove(fileName)
				if err != nil {
					log.Println(err)
				}
				err = file.Close()
				if err != nil {
					log.Println(err)
				}
			}
		}
	} else {
		log.Println(err)
	}
}
