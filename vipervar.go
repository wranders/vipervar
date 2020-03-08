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

// New creates a Resolver with default delimiter and special character values
func New() *Resolver {
	r := new(Resolver)
	r.DelimStart = defaultDelimStart
	r.DelimEnd = defaultDelimEnd
	r.DelimKey = defaultDelimKey
	r.KeySpecialCharacters = []byte{'_', '-'}
	return r
}

// Reset the package-level Resolver with default values
func Reset() {
	defaultResolver = New()
}

// SetDelimStart sets the characters that mark the beginning of a variable
func SetDelimStart(delim string) {
	defaultResolver.DelimStart = delim
}

// SetDelimEnd sets the characters that mark the end of a variable
func SetDelimEnd(delim string) {
	defaultResolver.DelimEnd = delim
}

// SetDelimStartEnd sets the start and end delimiters marking variables
func SetDelimStartEnd(start, end string) {
	defaultResolver.DelimStart = start
	defaultResolver.DelimEnd = end
}

// SetDelimKey is the delimiter used by Viper to denote subkeys.
// Default is `.` to match Viper and changes to Viper's key delimiter
// must be reflected here.
func SetDelimKey(delim string) {
	defaultResolver.DelimKey = delim
}

// SetKeySpecialCharacters sets the non-alphanumeric characters that are
// used in variable names. Default are `_` and `-`.
func SetKeySpecialCharacters(chars []byte) {
	defaultResolver.KeySpecialCharacters = chars
}

// SetExcludedKeys is a list of keys ignored by the `ResolveReplaceAll`
// and `ResolveReplaceAllIn` methods
func SetExcludedKeys(keys []string) {
	defaultResolver.ExcludeKeys = keys
}

// ResolveKey in the default Viper
func ResolveKey(key string) (string, error) {
	return defaultResolver.ResolveKey(key)
}

// ResolveKeyIn the specified Viper
func ResolveKeyIn(key string, useViper *viper.Viper) (string, error) {
	return defaultResolver.ResolveKeyIn(key, useViper)
}

// ResolveValue with the default Viper
func ResolveValue(value string) (string, error) {
	return defaultResolver.ResolveValue(value)
}

// ResolveValueWith the specified Viper
func ResolveValueWith(value string, useViper *viper.Viper) (string, error) {
	return defaultResolver.ResolveValueWith(value, useViper)
}

// ResolveReplaceKey in the default Viper
func ResolveReplaceKey(key string) error {
	return defaultResolver.ResolveReplaceKey(key)
}

// ResolveReplaceKeyIn the specified Viper
func ResolveReplaceKeyIn(key string, useViper *viper.Viper) error {
	return defaultResolver.ResolveReplaceKeyIn(key, useViper)
}

// ResolveReplaceAll values containting variables in the default Viper
func ResolveReplaceAll() error {
	return defaultResolver.ResolveReplaceAll()
}

// ResolveReplaceAllIn the specified Viper
func ResolveReplaceAllIn(useViper *viper.Viper) error {
	return defaultResolver.ResolveReplaceAllIn(useViper)
}
