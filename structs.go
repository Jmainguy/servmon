package main

type service struct {
	HTTP servHTTP `yaml:"http"`
	TCP  servTCP  `yaml:"tcp"`
}

type servTCP struct {
	Hostname string `yaml:"hostname"`
	Ports    []port `yaml:"ports"`
}

type servHTTP struct {
	URL        string `yaml:"url"`
	Statuscode int    `yaml:"statuscode"`
	Status     string `yaml:"status"`
}

type port struct {
	Port    string `yaml:"port"`
	Network string `yaml:"network"`
	Status  string `yaml:"status"`
}

type services []service
