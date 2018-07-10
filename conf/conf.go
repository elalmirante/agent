package conf

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	DefaultPort = "9095"
)

type Configuration struct {
	Port   string
	Key    string
	Script []string
}

func (c Configuration) ScriptLine(ref string) string {
	ref = `"` + ref + `"`
	cmd := strings.Join(c.Script, " && ")
	cmd = strings.Replace(cmd, "$REF", ref, -1)
	cmd = strings.Replace(cmd, "${REF}", ref, -1)
	return cmd
}

func Parse(path string) (*Configuration, error) {
	confBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// unmarshall config
	var conf *Configuration
	err = yaml.Unmarshal(confBytes, &conf)
	if err != nil {
		return nil, err
	}

	// add default port if not present.
	if conf.Port == "" {
		conf.Port = DefaultPort
	}

	return conf, nil
}
