package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	KeyID            string `json:"keyId"`
	ServiceAccountID string `json:"serviceAccountId"`
	KeyFile          string `json:"keyFile"`
	DNSZoneID        string `json:"dnsZoneId"`
	RecordName       string `json:"name"`
	RecordType       string `json:"type"`
	RecordTTL        string `json:"ttl"`
}

var config *Config = nil

func loadConfig() {
	if config != nil {
		return
	}
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	config = new(Config)
	err = json.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
}

func GetDNSRecord() (string, string, string) {
	loadConfig()
	return config.RecordName, config.RecordType, config.RecordTTL
}

func GetZoneID() string {
	loadConfig()
	return config.DNSZoneID
}

func GetServiceAccountID() string {
	loadConfig()
	return config.ServiceAccountID
}

func GetKeyID() string {
	loadConfig()
	return config.KeyID
}

func GetKeyFileName() string {
	loadConfig()
	return config.KeyFile
}
