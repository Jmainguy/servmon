package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"gopkg.in/yaml.v3"
)

func readPlay(filename string) (services services) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &services)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return services
}

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

func sendSlack(msg string) {
	token := os.Getenv("SLACK_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")
	api := slack.New(token)
	channelID, timestamp, err := api.PostMessage(
		channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
	log.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)

}

func triggerAlarm(msg string) {
	servmonDir := os.Getenv("SERVMONDIR")
	fileName := fmt.Sprintf("%s/hush", servmonDir)
	_, err := os.Stat(fileName)
	if err == nil {
		// Hush exists, dont alarm.
	} else if errors.Is(err, os.ErrNotExist) {
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
		file.Close()

	} else {
		log.Println(err)
	}
}

func checkHTTPService(servHTTP servHTTP) servHTTP {
	resp, err := http.Get(servHTTP.URL)
	if err != nil {
		log.Println(err)
		servHTTP.Status = "down"
		return servHTTP
	}
	if servHTTP.Statuscode == resp.StatusCode {
		servHTTP.Status = "up"
	} else {
		servHTTP.Status = "down"
	}
	resp.Body.Close()
	return servHTTP
}

func main() {
	servmonDir := os.Getenv("SERVMONDIR")
	timeout := time.Second
	go func() {
		for {
			services := readPlay(fmt.Sprintf("%s/monitor.yml", servmonDir))
			for sk, service := range services {
				if service.HTTP != (servHTTP{}) {
					services[sk].HTTP = checkHTTPService(service.HTTP)
				}
				if len(service.TCP.Ports) > 0 {
					for pk, port := range service.TCP.Ports {
						conn, err := net.DialTimeout(port.Network, net.JoinHostPort(service.TCP.Hostname, port.Port), timeout)
						if err != nil {
							services[sk].TCP.Ports[pk].Status = "down"
							msg := fmt.Sprintf("%s:%s %s is down", service.TCP.Hostname, port.Port, port.Network)
							go triggerAlarm(msg)

						} else {
							if conn != nil {
								services[sk].TCP.Ports[pk].Status = "up"
								conn.Close()
							}
						}
					}
				}
			}
			var doc bytes.Buffer
			t, err := template.ParseFiles(fmt.Sprintf("%s%s", servmonDir, "/template.html"))
			if err != nil {
				log.Println(err)
				continue
			}
			err = t.Execute(&doc, services)
			if err != nil {
				log.Println(err)
				continue
			}

			s := doc.String()
			toFile(fmt.Sprintf("%s%s", servmonDir, "/index.html"), s)
			time.Sleep(5 * time.Second)

		}
	}()
	// Http stuff
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.StaticFile("/", fmt.Sprintf("%s%s", servmonDir, "/index.html"))

	// Listen and server on 0.0.0.0:8080
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
