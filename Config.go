package commons

import (
	"encoding/json"
	"os"
)

func Load(config interface{}) {
	//filename is the path to the json config file
	file, err := os.Open("maketplay.config")
	if err != nil {
		return
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		panic("invalid format of file config")
	}
}
