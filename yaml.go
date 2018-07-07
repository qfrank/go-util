package goutil

import (
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func LoadYaml(file string, targetType interface{}) {
	f, err := os.Open(file)
	PanicIfErr(err)
	c, err := ioutil.ReadAll(f)
	PanicIfErr(err)
	err = yaml.Unmarshal(c, targetType)
	PanicIfErr(err)
}
