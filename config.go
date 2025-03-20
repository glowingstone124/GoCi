package main

import (
	"encoding/json"
	"os"
	"sync"
)

type Config struct {
	Port  int  `json:"port"`
	Debug bool `json:"debug"`
}

var (
	ConfigPath string = "config.json"
	config     *Config
	once       sync.Once
)

func loadConfig() <-chan *Config {
	result := make(chan *Config, 1)
	go func() {
		defer close(result)
		file, err := os.Open(ConfigPath)
		if err != nil {
			log("Error opening config file")
			return
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		var cfg Config
		err = decoder.Decode(&cfg)
		if err != nil {
			log("Error parsing config file", err.Error())
		}

		once.Do(func() {
			config = &cfg
		})
		result <- &cfg
	}()
	return result
}
