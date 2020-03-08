package vipervar

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
)

var config = []byte(`---
config_dir: /home/user/app/config
application:
  config: $(config_dir)/app.conf
server:
  domain: vipervar.foo
  port: 8080
  scheme: http
  root_url: "$(server.scheme)://$(server.domain):$(server.port)/"
`)

const (
	applicationConfig string = "/home/user/app/config/app.conf"
	serverRootURL     string = "http://vipervar.foo:8080/"
)

func TestResolveDefaultOneKey(t *testing.T) {
	viper.Reset()
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))

	tests := []struct {
		Key      string
		Expected string
	}{
		{"application.config", applicationConfig},
		{"server.root_url", serverRootURL},
	}

	for _, test := range tests {
		val, err := Resolve(test.Key)
		if err != nil {
			t.Error(err)
		}
		if val != test.Expected {
			t.Errorf("`%s` does not match `%s`: `%s`", test.Key, test.Expected, val)
		}
	}
}

func TestResolveOneKey(t *testing.T) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.ReadConfig(bytes.NewBuffer(config))

	tests := []struct {
		Key      string
		Expected string
	}{
		{"application.config", applicationConfig},
		{"server.root_url", serverRootURL},
	}

	for _, test := range tests {
		val, err := ResolveIn(v, test.Key)
		if err != nil {
			t.Error(err)
		}
		if val != test.Expected {
			t.Errorf("`%s` does not match `%s`: `%s`", test.Key, test.Expected, val)
		}
	}
}

func TestResolveDefaultOneValue(t *testing.T) {
	viper.Reset()
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))

	tests := []struct {
		Value    string
		Expected string
	}{
		{"$(config_dir)/app.conf", applicationConfig},
		{
			"$(server.scheme)://$(server.domain):$(server.port)/",
			serverRootURL,
		},
	}

	for _, test := range tests {
		val, err := ResolveValue(test.Value)
		if err != nil {
			t.Error(err)
		}
		if val != test.Expected {
			t.Errorf("`%s` does not match `%s`: `%s`", test.Value, test.Expected, val)
		}
	}
}

func TestResolveOneValue(t *testing.T) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.ReadConfig(bytes.NewBuffer(config))

	tests := []struct {
		Value    string
		Expected string
	}{
		{"$(config_dir)/app.conf", applicationConfig},
		{
			"$(server.scheme)://$(server.domain):$(server.port)/",
			serverRootURL,
		},
	}

	for _, test := range tests {
		val, err := ResolveValueIn(v, test.Value)
		if err != nil {
			t.Error(err)
		}
		if val != test.Expected {
			t.Errorf("`%s` does not match `%s`: `%s`", test.Value, test.Expected, val)
		}
	}
}

func TestResolveReplaceDefaultOneKey(t *testing.T) {
	viper.Reset()
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))

	tests := []struct {
		Key      string
		Expected string
	}{
		{"application.config", applicationConfig},
		{"server.root_url", serverRootURL},
	}

	for _, test := range tests {
		err := ResolveReplace(test.Key)
		if err != nil {
			t.Error(err)
		}
		val := viper.GetString(test.Key)
		if val != test.Expected {
			t.Errorf("`%s` does not match `%s`: `%s`", test.Key, test.Expected, val)
		}
	}
}

func TestResolveReplaceOneKey(t *testing.T) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.ReadConfig(bytes.NewBuffer(config))

	tests := []struct {
		Key      string
		Expected string
	}{
		{"application.config", applicationConfig},
		{"server.root_url", serverRootURL},
	}

	for _, test := range tests {
		err := ResolveReplaceIn(v, test.Key)
		if err != nil {
			t.Error(err)
		}
		val := v.GetString(test.Key)
		if val != test.Expected {
			t.Errorf("`%s` does not match `%s`: `%s`", test.Key, test.Expected, val)
		}
	}
}

func TestResolveReplaceAllDefault(t *testing.T) {
	viper.Reset()
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))

	tests := []struct {
		Key      string
		Expected string
	}{
		{"application.config", applicationConfig},
		{"server.root_url", serverRootURL},
	}

	err := ResolveReplaceAll()
	if err != nil {
		t.Error(err)
	}

	for _, test := range tests {
		val := viper.GetString(test.Key)
		if val != test.Expected {
			t.Errorf("`%s` does not match `%s`: `%s`", test.Key, test.Expected, val)
		}
	}
}
