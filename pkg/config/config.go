package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	items map[string]map[interface{}]interface{}
}

func Factory() interface{} {
	godotenv.Load()

	fmt.Print("config Factory\n")

	items := make(map[string]map[interface{}]interface{})
	c := Config{items}
	return &c
}

/*
get from yaml file
if you want to get env key value, use os.GetEnv("key")
*/
func (c *Config) Get(keyChain string) interface{} {
	keys := strings.Split(keyChain, ".")
	if _, ok := c.items[keys[0]]; !ok {
		//do something here
		c.items[keys[0]] = load(keys[0])
	}

	cnfItem := (map[interface{}]interface{})(c.items[keys[0]])
	for i, key := range keys {
		if i > 0 {
			if v, ok := cnfItem[key].(map[interface{}]interface{}); ok {
				cnfItem = v
			} else {
				if i == len(keys)-1 {
					return cnfItem[key]
				} else {
					return nil
				}
			}
		}
	}
	return cnfItem
}

func load(path string) map[interface{}]interface{} {
	filename, _ := filepath.Abs("conf//" + path + ".yml")
	yamlContent, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	yamlContent = []byte(os.ExpandEnv(string(yamlContent)))

	var item = make(map[interface{}]interface{})
	err = yaml.Unmarshal(yamlContent, &item)
	if err != nil {
		panic(err)
	}

	return item
}
