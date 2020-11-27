package main

import (
	"log"

	"gopkg.in/yaml.v2"
)

type goyaml struct {
	data map[string]string
}

// YamlStruct for yaml data
type YamlStruct []struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

var dataYaml = `
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
- path: /urlshort
  url: https://github.com/gophercises/urlshort
`

func newYaml() (*goyaml, error) {
	y := YamlStruct{}
	err := yaml.Unmarshal([]byte(dataYaml), &y)
	log.Println("yaml file:", y)
	if err != nil {
		log.Println("error while unmarshalng yaml: ", err)
		return nil, err
	}
	d := make(map[string]string)
	yamlDict := goyaml{
		data: d,
	}
	for _, yy := range y {
		yamlDict.data[yy.Path] = yy.URL
	}
	return &yamlDict, nil
}

func (g *goyaml) getData(path string) (string, bool) {
	value, ok := g.data[path]
	return value, ok
}
