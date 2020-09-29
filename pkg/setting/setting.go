package setting

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	configs map[string]map[string]interface{}
)

func init() {
	configs = make(map[string]map[string]interface{})
}

func get(keyChain string) interface{} {
	keys := strings.Split(keyChain, ".")
	if _, ok := configs[keys[0]]; !ok {
		//do something here
		configs[keys[0]] = load(keys[0])
	}
	return configs[keys[0]]
}

func load(path string) map[string]interface{} {
	filename, _ := filepath.Abs("conf//" + path)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config = make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value: %#v\n", config)
	return config
}
