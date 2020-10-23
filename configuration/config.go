package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Subscription represents a subscription to a single account.
type Subscription struct {
	UserID        string `json:"userID"`
	ChannelID     string `json:"channelID"`
	LastTweetTime int64  `json:"lastTweetTime"`
}

// Configuration represents the program configuration schema.
type Configuration struct {
	Subscriptions  []Subscription `json:"subscriptions"`
	ModRoles       []string       `json:"modRoles"`
	Prefix         string         `json:"prefix"`
	NativeLanguage string         `json:"nativeLanguage"`
}

// Save saves the configuration file.
func (config *Configuration) Save() {
	bytes, err := json.Marshal(*config)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("config.json", bytes, 0755)
	if err != nil {
		log.Println(err)
		return
	}
}

// LoadConfig loads the configuration file.
func LoadConfig() *Configuration {
	data, err := ioutil.ReadFile("config.json")
	if err != nil || len(data) == 0 {
		log.Println("Created new configuration data.")
		return makeConfigDefaults()
	}

	out := Configuration{}

	err = json.Unmarshal(data, &out)
	if err != nil {
		log.Fatalln(err)
	}

	return &out
}

func makeConfigDefaults() *Configuration {
	config := &Configuration{
		Subscriptions:  make([]Subscription, 0),
		ModRoles:       make([]string, 0),
		Prefix:         "^",
		NativeLanguage: "en",
	}
	config.Save()
	return config
}
