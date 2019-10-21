package vipervar

import "testing"

var tomlDefault = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "%CONFIG_DIR%/app.conf"
[server]
root_url = "%SERVER.SCHEME%://%SERVER.DOMAIN%:%SERVER.PORT%/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlPercent = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "%CONFIG_DIR%/app.conf"
[server]
root_url = "%SERVER.SCHEME%://%SERVER.DOMAIN%:%SERVER.PORT%/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlCaret = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "^CONFIG_DIR^/app.conf"
[server]
root_url = "^SERVER.SCHEME^://^SERVER.DOMAIN^:^SERVER.PORT^/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlDollar = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "$CONFIG_DIR$/app.conf"
[server]
root_url = "$SERVER.SCHEME$://$SERVER.DOMAIN$:$SERVER.PORT$/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlQuestion = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "?CONFIG_DIR?/app.conf"
[server]
root_url = "?SERVER.SCHEME?://?SERVER.DOMAIN?:?SERVER.PORT?/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlAstrisk = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "*CONFIG_DIR*/app.conf"
[server]
root_url = "*SERVER.SCHEME*://*SERVER.DOMAIN*:*SERVER.PORT*/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlPlus = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "+CONFIG_DIR+/app.conf"
[server]
root_url = "+SERVER.SCHEME+://+SERVER.DOMAIN+:+SERVER.PORT+/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlCurlyBracket = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "{CONFIG_DIR}/app.conf"
[server]
root_url = "{SERVER.SCHEME}://{SERVER.DOMAIN}:{SERVER.PORT}/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlDoubleCurlyBracket = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "{{CONFIG_DIR}}/app.conf"
[server]
root_url = "{{SERVER.SCHEME}}://{{SERVER.DOMAIN}}:{{SERVER.PORT}}/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlSquareBracket = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "[CONFIG_DIR]/app.conf"
[server]
root_url = "[SERVER.SCHEME]://[SERVER.DOMAIN]:[SERVER.PORT]/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlDoubleSquareBracket = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "[[CONFIG_DIR]]/app.conf"
[server]
root_url = "[[SERVER.SCHEME]]://[[SERVER.DOMAIN]]:[[SERVER.PORT]]/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlPound = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "#CONFIG_DIR#/app.conf"
[server]
root_url = "#SERVER.SCHEME#://#SERVER.DOMAIN#:#SERVER.PORT#/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlComplexOne = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "[%CONFIG_DIR%]/app.conf"
[server]
root_url = "[%SERVER.SCHEME%]://[%SERVER.DOMAIN%]:[%SERVER.PORT%]/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlComplexTwo = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "[%]CONFIG_DIR[%]/app.conf"
[server]
root_url = "[%]SERVER.SCHEME[%]://[%]SERVER.DOMAIN[%]:[%]SERVER.PORT[%]/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlParentheses = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "(CONFIG_DIR)/app.conf"
[server]
root_url = "(SERVER.SCHEME)://(SERVER.DOMAIN):(SERVER.PORT)/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlDoubleParentheses = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "((CONFIG_DIR))/app.conf"
[server]
root_url = "((SERVER.SCHEME))://((SERVER.DOMAIN)):((SERVER.PORT))/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlAmpersand = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "&CONFIG_DIR&/app.conf"
[server]
root_url = "&SERVER.SCHEME&://&SERVER.DOMAIN&:&SERVER.PORT&/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlCommercialAt = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "@CONFIG_DIR@/app.conf"
[server]
root_url = "@SERVER.SCHEME@://@SERVER.DOMAIN@:@SERVER.PORT@/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

var tomlExclamation = []byte(`
config_dir = "/home/user/app/config"
[application]
config = "!CONFIG_DIR!/app.conf"
[server]
root_url = "!SERVER.SCHEME!://!SERVER.DOMAIN!:!SERVER.PORT!/"
domain = "vipervar.foo"
port = 8080
scheme = "http"
`)

func TestTOMLDefault(t *testing.T) {
	initViper("toml", tomlDefault)
	r := NewResolver()
	err := r.Resolve()
	if err != nil {
		t.Error(err)
	}
	checkViper(t, "application.config")
	checkViper(t, "server.root_url")
}

func TestTOMLPercent(t *testing.T) {
	initViper("toml", tomlPercent)
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

func TestTOMLCaret(t *testing.T) {
	initViper("toml", tomlCaret)
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

func TestTOMLDollar(t *testing.T) {
	initViper("toml", tomlDollar)
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

func TestTOMLQuestion(t *testing.T) {
	initViper("toml", tomlQuestion)
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

func TestTOMLAstrisk(t *testing.T) {
	initViper("toml", tomlAstrisk)
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

func TestTOMLPlus(t *testing.T) {
	initViper("toml", tomlPlus)
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

func TestTOMLCurlyBracket(t *testing.T) {
	initViper("toml", tomlCurlyBracket)
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

func TestTOMLDoubleCurlyBracket(t *testing.T) {
	initViper("toml", tomlDoubleCurlyBracket)
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

func TestTOMLSquareBracket(t *testing.T) {
	initViper("toml", tomlSquareBracket)
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

func TestTOMLDoubleSquareBracket(t *testing.T) {
	initViper("toml", tomlDoubleSquareBracket)
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

func TestTOMLPound(t *testing.T) {
	initViper("toml", tomlPound)
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

func TestTOMLComplexOne(t *testing.T) {
	initViper("toml", tomlComplexOne)
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

func TestTOMLComplexTwo(t *testing.T) {
	initViper("toml", tomlComplexTwo)
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

func TestTOMLParentheses(t *testing.T) {
	initViper("toml", tomlParentheses)
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

func TestTOMLDoubleParentheses(t *testing.T) {
	initViper("toml", tomlDoubleParentheses)
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

func TestTOMLAmpersand(t *testing.T) {
	initViper("toml", tomlAmpersand)
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

func TestTOMLCommercialAt(t *testing.T) {
	initViper("toml", tomlCommercialAt)
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

func TestTOMLExclamation(t *testing.T) {
	initViper("toml", tomlExclamation)
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
