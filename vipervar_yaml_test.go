package vipervar

import "testing"

var yamlDefault = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "%CONFIG_DIR%/app.conf"
server:
  root_url: "%SERVER.SCHEME%://%SERVER.DOMAIN%:%SERVER.PORT%/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlPercent = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "%CONFIG_DIR%/app.conf"
server:
  root_url: "%SERVER.SCHEME%://%SERVER.DOMAIN%:%SERVER.PORT%/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlCaret = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "^CONFIG_DIR^/app.conf"
server:
  root_url: "^SERVER.SCHEME^://^SERVER.DOMAIN^:^SERVER.PORT^/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlDollar = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "$CONFIG_DIR$/app.conf"
server:
  root_url: "$SERVER.SCHEME$://$SERVER.DOMAIN$:$SERVER.PORT$/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlQuestion = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "?CONFIG_DIR?/app.conf"
server:
  root_url: "?SERVER.SCHEME?://?SERVER.DOMAIN?:?SERVER.PORT?/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlAstrisk = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "*CONFIG_DIR*/app.conf"
server:
  root_url: "*SERVER.SCHEME*://*SERVER.DOMAIN*:*SERVER.PORT*/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlPlus = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "+CONFIG_DIR+/app.conf"
server:
  root_url: "+SERVER.SCHEME+://+SERVER.DOMAIN+:+SERVER.PORT+/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlCurlyBracket = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "{CONFIG_DIR}/app.conf"
server:
  root_url: "{SERVER.SCHEME}://{SERVER.DOMAIN}:{SERVER.PORT}/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlDoubleCurlyBracket = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "{{CONFIG_DIR}}/app.conf"
server:
  root_url: "{{SERVER.SCHEME}}://{{SERVER.DOMAIN}}:{{SERVER.PORT}}/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlSquareBracket = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "[CONFIG_DIR]/app.conf"
server:
  root_url: "[SERVER.SCHEME]://[SERVER.DOMAIN]:[SERVER.PORT]/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlDoubleSquareBracket = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "[[CONFIG_DIR]]/app.conf"
server:
  root_url: "[[SERVER.SCHEME]]://[[SERVER.DOMAIN]]:[[SERVER.PORT]]/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlPound = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "#CONFIG_DIR#/app.conf"
server:
  root_url: "#SERVER.SCHEME#://#SERVER.DOMAIN#:#SERVER.PORT#/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlComplexOne = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "[%CONFIG_DIR%]/app.conf"
server:
  root_url: "[%SERVER.SCHEME%]://[%SERVER.DOMAIN%]:[%SERVER.PORT%]/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlComplexTwo = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "[%]CONFIG_DIR[%]/app.conf"
server:
  root_url: "[%]SERVER.SCHEME[%]://[%]SERVER.DOMAIN[%]:[%]SERVER.PORT[%]/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlParentheses = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "(CONFIG_DIR)/app.conf"
server:
  root_url: "(SERVER.SCHEME)://(SERVER.DOMAIN):(SERVER.PORT)/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlDoubleParentheses = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "((CONFIG_DIR))/app.conf"
server:
  root_url: "((SERVER.SCHEME))://((SERVER.DOMAIN)):((SERVER.PORT))/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlAmpersand = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "&CONFIG_DIR&/app.conf"
server:
  root_url: "&SERVER.SCHEME&://&SERVER.DOMAIN&:&SERVER.PORT&/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlCommercialAt = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "@CONFIG_DIR@/app.conf"
server:
  root_url: "@SERVER.SCHEME@://@SERVER.DOMAIN@:@SERVER.PORT@/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

var yamlExclamation = []byte(`
config_dir: "/home/user/app/config"
application:
  config: "!CONFIG_DIR!/app.conf"
server:
  root_url: "!SERVER.SCHEME!://!SERVER.DOMAIN!:!SERVER.PORT!/"
  domain: "vipervar.foo"
  port: 8080
  scheme: "http"
`)

func TestYAMLDefault(t *testing.T) {
	initViper("yaml", yamlDefault)
	r := NewResolver()
	err := r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLPercent(t *testing.T) {
	initViper("yaml", yamlPercent)
	r := NewResolver()
	err := r.SetDelimiters(`%`, `%`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLCaret(t *testing.T) {
	initViper("yaml", yamlCaret)
	r := NewResolver()
	err := r.SetDelimiters(`^`, `^`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLDollar(t *testing.T) {
	initViper("yaml", yamlDollar)
	r := NewResolver()
	err := r.SetDelimiters(`$`, `$`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLQuestion(t *testing.T) {
	initViper("yaml", yamlQuestion)
	r := NewResolver()
	err := r.SetDelimiters(`?`, `?`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLAstrisk(t *testing.T) {
	initViper("yaml", yamlAstrisk)
	r := NewResolver()
	err := r.SetDelimiters(`*`, `*`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLPlus(t *testing.T) {
	initViper("yaml", yamlPlus)
	r := NewResolver()
	err := r.SetDelimiters(`+`, `+`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLCurlyBracket(t *testing.T) {
	initViper("yaml", yamlCurlyBracket)
	r := NewResolver()
	err := r.SetDelimiters(`{`, `}`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLDoubleCurlyBracket(t *testing.T) {
	initViper("yaml", yamlDoubleCurlyBracket)
	r := NewResolver()
	err := r.SetDelimiters(`{{`, `}}`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLSquareBracket(t *testing.T) {
	initViper("yaml", yamlSquareBracket)
	r := NewResolver()
	err := r.SetDelimiters(`[`, `]`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLDoubleSquareBracket(t *testing.T) {
	initViper("yaml", yamlDoubleSquareBracket)
	r := NewResolver()
	err := r.SetDelimiters(`[[`, `]]`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLPound(t *testing.T) {
	initViper("yaml", yamlPound)
	r := NewResolver()
	err := r.SetDelimiters(`#`, `#`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLComplexOne(t *testing.T) {
	initViper("yaml", yamlComplexOne)
	r := NewResolver()
	err := r.SetDelimiters(`[%`, `%]`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLComplexTwo(t *testing.T) {
	initViper("yaml", yamlComplexTwo)
	r := NewResolver()
	err := r.SetDelimiters(`[%]`, `[%]`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLParentheses(t *testing.T) {
	initViper("yaml", yamlParentheses)
	r := NewResolver()
	err := r.SetDelimiters(`(`, `)`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLDoubleParentheses(t *testing.T) {
	initViper("yaml", yamlDoubleParentheses)
	r := NewResolver()
	err := r.SetDelimiters(`((`, `))`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLAmpersand(t *testing.T) {
	initViper("yaml", yamlAmpersand)
	r := NewResolver()
	err := r.SetDelimiters(`&`, `&`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLCommercialAt(t *testing.T) {
	initViper("yaml", yamlCommercialAt)
	r := NewResolver()
	err := r.SetDelimiters(`@`, `@`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestYAMLExclamation(t *testing.T) {
	initViper("yaml", yamlExclamation)
	r := NewResolver()
	err := r.SetDelimiters(`!`, `!`)
	if err != nil {
		t.Error(err)
	}
	err = r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}
