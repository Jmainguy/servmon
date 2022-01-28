package main

import (
	"bufio"
	"log"
	"os"
)

func toFile(file, data string) {
	// Truncate file
	f, err := os.Create(file)
	if err != nil {
		log.Println(err)
		return
	}
	// Write lines to file
	w := bufio.NewWriter(f)
	_, err = w.WriteString(data)
	if err != nil {
		log.Println(err)
		return
	}
	w.Flush()
	err = f.Sync()
	if err != nil {
		log.Println(err)
		return
	}
	f.Close()

}
