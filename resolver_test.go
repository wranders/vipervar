package vipervar

import (
	"bytes"
	"strings"
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
diff_config: "%[config_dir]"
diff_domain: $(server::domain)
sc/key: "value"
scref: $(sc/key)
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
		val, err := ResolveKey(test.Key)
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
		val, err := ResolveKeyIn(test.Key, v)
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
		val, err := ResolveValueWith(test.Value, v)
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
		err := ResolveReplaceKey(test.Key)
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
		err := ResolveReplaceKeyIn(test.Key, v)
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

func TestResolverSettingErr(t *testing.T) {
	viper.Reset()
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))

	SetDelimStart("")
	err := ResolveReplaceAll()
	if err == nil {
		t.Error(err)
	}
	if _, ok := err.(*ErrInvalidDelimiter); !ok {
		t.Errorf("Unexpected error: %T", err)
	}
	errmsgStart := "Delimiter `start` invalid: empty"
	if strings.Compare(err.Error(), errmsgStart) != 0 {
		t.Errorf("Unexpected error message: \"%s\". should be \"%s\"", err, errmsgStart)
	}
	Reset()

	SetDelimEnd("")
	err = ResolveReplaceAll()
	if err == nil {
		t.Error(err)
	}
	if _, ok := err.(*ErrInvalidDelimiter); !ok {
		t.Errorf("Unexpected error: %T", err)
	}
	errmsgEnd := "Delimiter `end` invalid: empty"
	if strings.Compare(err.Error(), errmsgEnd) != 0 {
		t.Errorf("Unexpected error message: \"%s\". should be \"%s\"", err, errmsgEnd)
	}
	Reset()

	SetDelimKey("")
	err = ResolveReplaceAll()
	if err == nil {
		t.Error(err)
	}
	if _, ok := err.(*ErrInvalidDelimiter); !ok {
		t.Errorf("Unexpected error: %T", err)
	}
	errmsgKey := "Delimiter `key` invalid: empty"
	if strings.Compare(err.Error(), errmsgKey) != 0 {
		t.Errorf("Unexpected error message: \"%s\". should be \"%s\"", err, errmsgKey)
	}
	Reset()
}

func TestChangeDelim(t *testing.T) {
	viper.Reset()
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))

	v, err := ResolveKey("application.config")
	if err != nil {
		t.Error(err)
	}
	if v != applicationConfig {
		t.Errorf("`application.config` does not match `%s`: `%s`", applicationConfig, v)
	}

	SetDelimStart("%[")
	SetDelimEnd("]")
	v, err = ResolveKey("diff_config")
	if err != nil {
		t.Error(err)
	}
	if v != "/home/user/app/config" {
		t.Errorf("`application.config` does not match `/home/user/app/config`: `%s`", v)
	}

	vip := viper.NewWithOptions(viper.KeyDelimiter("::"))
	vip.SetConfigType("yaml")
	vip.ReadConfig(bytes.NewBuffer(config))
	SetDelimStart("$(")
	SetDelimEnd(")")
	SetDelimKey("::")
	v, err = ResolveKeyIn("diff_domain", vip)
	if err != nil {
		t.Error(err)
	}
	if v != "vipervar.foo" {
		t.Errorf("`diff_domain` does not match `vipervar.foo`: `%s`", v)
	}

	SetKeySpecialCharacters([]byte{'_', '-', '/'})
	v, err = ResolveKey("scref")
	if err != nil {
		t.Error(err)
	}
	if v != "value" {
		t.Errorf("`scref` does not match `value`: `%s`", v)
	}
}
