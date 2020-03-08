package vipervar

import "github.com/spf13/viper"

const (
	defaultDelimStart string = `$(`
	defaultDelimEnd   string = `)`
	defaultDelimKey   string = `.`
)

var defaultResolver *Resolver

func init() {
	defaultResolver = New()
}

func New() *Resolver {
	r := new(Resolver)
	r.DelimStart = defaultDelimStart
	r.DelimEnd = defaultDelimEnd
	r.DelimKey = defaultDelimKey
	r.KeySpecialCharacters = []string{`_`, `-`}
	return r
}

func Reset() {
	defaultResolver = New()
}

func SetDelimStart(delim string) {
	defaultResolver.DelimStart = delim
}

func SetDelimEnd(delim string) {
	defaultResolver.DelimEnd = delim
}

func SetDelimKey(delim string) {
	defaultResolver.DelimKey = delim
}

func SetKeySpecialCharacters(chars []string) {
	defaultResolver.KeySpecialCharacters = chars
}

func SetExcludedKeys(keys []string) {
	defaultResolver.ExcludeKeys = keys
}

func Resolve(key string) (string, error) {
	return defaultResolver.Resolve(key)
}

func ResolveIn(useViper *viper.Viper, key string) (string, error) {
	return defaultResolver.ResolveIn(useViper, key)
}

func ResolveValue(value string) (string, error) {
	return defaultResolver.ResolveValue(value)
}

func ResolveValueIn(useViper *viper.Viper, value string) (string, error) {
	return defaultResolver.ResolveValueIn(useViper, value)
}

func ResolveReplace(key string) error {
	return defaultResolver.ResolveReplace(key)
}

func ResolveReplaceIn(useViper *viper.Viper, key string) error {
	return defaultResolver.ResolveReplaceIn(useViper, key)
}

func ResolveReplaceAll() error {
	return defaultResolver.ResolveReplaceAll()
}

func ResolveReplaceAllIn(useViper *viper.Viper) error {
	return defaultResolver.ResolveReplaceAllIn(useViper)
}
