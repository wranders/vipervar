package vipervar

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func initViper(conftype string, conf []byte) {
	viper.Reset()
	viper.SetConfigType(conftype)
	viper.ReadConfig(bytes.NewBuffer(conf))
}

func initThisViper(v *viper.Viper, conftype string, conf []byte) {
	v.SetConfigType(conftype)
	v.ReadConfig(bytes.NewBuffer(conf))
}

var expectedValues = map[string]string{
	"application.config":     `/home/user/app/config/app.conf`,
	"application.config_var": `%CONFIG_DIR%/app.conf`,
	"server.root_url":        `http://vipervar.foo:8080/`,
}

func checkViper(t *testing.T, key string) {
	value := viper.GetString(key)
	if strings.Compare(value, expectedValues[key]) != 0 {
		t.Errorf("Got `%s`, expected `%s`", value, expectedValues[key])
	}
}

func checkThisViper(t *testing.T, v *viper.Viper, key string) {
	value := v.GetString(key)
	if strings.Compare(value, expectedValues[key]) != 0 {
		t.Errorf("Got `%s`, expected `%s`", value, expectedValues[key])
	}
}

func TestInvalidCharacter(t *testing.T) {
	r := NewResolver()
	err := r.SetDelimiters(`.`, `%`)
	if err == nil {
		t.Error("Invalid character used, should return error")
	}
	err = r.SetDelimiters(`%`, `.`)
	if err == nil {
		t.Error("Invalid character used, should return error")
	}
}

func TestNonDefaultViper(t *testing.T) {
	v := viper.New()
	initThisViper(v, "yaml", yamlDefault)
	r := NewResolverFrom(v)
	err := r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkThisViper(t, v, "application.config")
	checkThisViper(t, v, "server.root_url")
}

func TestKeyExclude(t *testing.T) {
	yamlExclude := []byte(`
config_dir: "/home/user/app/config"
application:
  config: "%CONFIG_DIR%/app.conf"
  config_var: "%CONFIG_DIR%/app.conf"
server:
  root_url: "%SERVER.SCHEME%://%SERVER.DOMAIN%:%SERVER.PORT%/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)
	initViper("yaml", yamlExclude)
	r := NewResolver()
	r.ExcludeKeys = []string{
		"application.config_var",
	}
	err := r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "application.config_var")
	checkViper(t, "server.root_url")
}

func TestInvalidKey(t *testing.T) {
	yamlInvalidKey := []byte(`
config_dir: "/home/user/app/config"
application:
  config: "%CONFIG_DIR%/app.conf"
server:
  root_url: "%SERVER.SCHEME%://%SERVER.DOMAIN%:%SERVER.PORT%/"
  listen: "%SERVER.IP%:%SERVER.PORT%"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)
	initViper("yaml", yamlInvalidKey)
	r := NewResolver()
	err := r.Resolve()
	if err == nil {
		t.Error("An invalid key was used, this should return an error")
	}
	errMsg := err.Error()
	expectedErrMsg := ("Errors were found in the specified configuration, no changes were made\n" +
		"Configuration variable `SERVER.IP` references a non-existent value")
	if strings.Compare(errMsg, expectedErrMsg) != 0 {
		t.Errorf("Unexpected error message.\nGot:\n%s\n\nExpected:\n%s", errMsg, expectedErrMsg)
	}
}
