package vipervar

import "testing"

var jsonDefault = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "%CONFIG_DIR%/app.conf"
	},
	"server": {
		"root_url": "%SERVER.SCHEME%://%SERVER.DOMAIN%:%SERVER.PORT%/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonPercent = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "%CONFIG_DIR%/app.conf"
	},
	"server": {
		"root_url": "%SERVER.SCHEME%://%SERVER.DOMAIN%:%SERVER.PORT%/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonCaret = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "^CONFIG_DIR^/app.conf"
	},
	"server": {
		"root_url": "^SERVER.SCHEME^://^SERVER.DOMAIN^:^SERVER.PORT^/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonDollar = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "$CONFIG_DIR$/app.conf"
	},
	"server": {
		"root_url": "$SERVER.SCHEME$://$SERVER.DOMAIN$:$SERVER.PORT$/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonQuestion = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "?CONFIG_DIR?/app.conf"
	},
	"server": {
		"root_url": "?SERVER.SCHEME?://?SERVER.DOMAIN?:?SERVER.PORT?/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonAstrisk = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "*CONFIG_DIR*/app.conf"
	},
	"server": {
		"root_url": "*SERVER.SCHEME*://*SERVER.DOMAIN*:*SERVER.PORT*/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonPlus = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "+CONFIG_DIR+/app.conf"
	},
	"server": {
		"root_url": "+SERVER.SCHEME+://+SERVER.DOMAIN+:+SERVER.PORT+/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonCurlyBracket = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "{CONFIG_DIR}/app.conf"
	},
	"server": {
		"root_url": "{SERVER.SCHEME}://{SERVER.DOMAIN}:{SERVER.PORT}/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonDoubleCurlyBracket = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "{{CONFIG_DIR}}/app.conf"
	},
	"server": {
		"root_url": "{{SERVER.SCHEME}}://{{SERVER.DOMAIN}}:{{SERVER.PORT}}/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonSquareBracket = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "[CONFIG_DIR]/app.conf"
	},
	"server": {
		"root_url": "[SERVER.SCHEME]://[SERVER.DOMAIN]:[SERVER.PORT]/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonDoubleSquareBracket = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "[[CONFIG_DIR]]/app.conf"
	},
	"server": {
		"root_url": "[[SERVER.SCHEME]]://[[SERVER.DOMAIN]]:[[SERVER.PORT]]/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonPound = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "#CONFIG_DIR#/app.conf"
	},
	"server": {
		"root_url": "#SERVER.SCHEME#://#SERVER.DOMAIN#:#SERVER.PORT#/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonComplexOne = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "[%CONFIG_DIR%]/app.conf"
	},
	"server": {
		"root_url": "[%SERVER.SCHEME%]://[%SERVER.DOMAIN%]:[%SERVER.PORT%]/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonComplexTwo = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "[%]CONFIG_DIR[%]/app.conf"
	},
	"server": {
		"root_url": "[%]SERVER.SCHEME[%]://[%]SERVER.DOMAIN[%]:[%]SERVER.PORT[%]/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonParentheses = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "(CONFIG_DIR)/app.conf"
	},
	"server": {
		"root_url": "(SERVER.SCHEME)://(SERVER.DOMAIN):(SERVER.PORT)/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonDoubleParentheses = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "((CONFIG_DIR))/app.conf"
	},
	"server": {
		"root_url": "((SERVER.SCHEME))://((SERVER.DOMAIN)):((SERVER.PORT))/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonAmpersand = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "&CONFIG_DIR&/app.conf"
	},
	"server": {
		"root_url": "&SERVER.SCHEME&://&SERVER.DOMAIN&:&SERVER.PORT&/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonCommercialAt = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "@CONFIG_DIR@/app.conf"
	},
	"server": {
		"root_url": "@SERVER.SCHEME@://@SERVER.DOMAIN@:@SERVER.PORT@/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

var jsonExclamation = []byte(`
{
	"config_dir": "/home/user/app/config",
	"application": {
		"config": "!CONFIG_DIR!/app.conf"
	},
	"server": {
		"root_url": "!SERVER.SCHEME!://!SERVER.DOMAIN!:!SERVER.PORT!/",
		"domain": "vipervar.foo",
		"port": 8080,
		"scheme": "http"
	}
}
`)

func TestJSONDefault(t *testing.T) {
	initViper("json", jsonDefault)
	r := NewResolver()
	err := r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestJSONPercent(t *testing.T) {
	initViper("json", jsonPercent)
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

func TestJSONCaret(t *testing.T) {
	initViper("json", jsonCaret)
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

func TestJSONDollar(t *testing.T) {
	initViper("json", jsonDollar)
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

func TestJSONQuestion(t *testing.T) {
	initViper("json", jsonQuestion)
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

func TestJSONAstrisk(t *testing.T) {
	initViper("json", jsonAstrisk)
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

func TestJSONPlus(t *testing.T) {
	initViper("json", jsonPlus)
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

func TestJSONCurlyBracket(t *testing.T) {
	initViper("json", jsonCurlyBracket)
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

func TestJSONDoubleCurlyBracket(t *testing.T) {
	initViper("json", jsonDoubleCurlyBracket)
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

func TestJSONSquareBracket(t *testing.T) {
	initViper("json", jsonSquareBracket)
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

func TestJSONDoubleSquareBracket(t *testing.T) {
	initViper("json", jsonDoubleSquareBracket)
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

func TestJSONPound(t *testing.T) {
	initViper("json", jsonPound)
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

func TestJSONComplexOne(t *testing.T) {
	initViper("json", jsonComplexOne)
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

func TestJSONComplexTwo(t *testing.T) {
	initViper("json", jsonComplexTwo)
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

func TestJSONParentheses(t *testing.T) {
	initViper("json", jsonParentheses)
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

func TestJSONDoubleParentheses(t *testing.T) {
	initViper("json", jsonDoubleParentheses)
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

func TestJSONAmpersand(t *testing.T) {
	initViper("json", jsonAmpersand)
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

func TestJSONCommercialAt(t *testing.T) {
	initViper("json", jsonCommercialAt)
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

func TestJSONExclamation(t *testing.T) {
	initViper("json", jsonExclamation)
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
