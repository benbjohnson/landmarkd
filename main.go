package main

import (
	"flag"
	"log"
	"runtime"

	. "github.com/skylandlabs/landmarkd/config"
	"github.com/skylandlabs/landmarkd/server"
)

var config *Config
var configPath string

func init() {
	config = NewConfig()
	flag.StringVar(&configPath, "c", "", "the path to the config file")
}

func main() {
	flag.Parse()
	if configPath != "" {
		if err := config.DecodeFile(configPath); err != nil {
			log.Fatal("Config error:", err)
		}
	}

	// Initialize.
	runtime.GOMAXPROCS(runtime.NumCPU())
	writePidFile()

	// Run server.
	s := server.New(config.Port)
	log.Fatal(s.ListenAndServe(c))
}

func writePidFile() {
	pid := strconv.Itoa(os.Getpid())
	ioutil.WriteFile(config.PidPath, []byte(pid), 0644)
}
