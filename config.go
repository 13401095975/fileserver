package main

import (
	"bytes"
	"encoding/json"
	"os"
)

type Server struct {
	AppName    string
	Port       int
	ServerRoot string
}

var ServerCfg Server

// Loads config information from a JSON file
func LoadConfig(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	b := new(bytes.Buffer)
	_, err = b.ReadFrom(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b.Bytes(), &ServerCfg)
	if err != nil {
		return err
	}

	return nil
}
