package vipervar

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// Resolver contains settings and methods for resolving varaibles in
// Viper configurations
type Resolver struct {
	DelimStart           string
	DelimEnd             string
	DelimKey             string
	KeySpecialCharacters []byte
	ExcludeKeys          []string

	regex                   *regexp.Regexp
	recompile               bool
	oldDelimStart           string
	oldDelimEnd             string
	oldDelimKey             string
	oldKeySpecialCharacters []byte
}

// ResolveKey in the default Viper
func (r *Resolver) ResolveKey(key string) (string, error) {
	defaultViper := viper.GetViper()
	return r.ResolveKeyIn(key, defaultViper)
}

// ResolveKeyIn the specified Viper
func (r *Resolver) ResolveKeyIn(key string, useViper *viper.Viper) (string, error) {
	if err := r.validateSettings(); err != nil {
		return "", err
	}
	k := useViper.Get(key)
	if k == nil {
		return "", &ErrNonExistentKey{key}
	}
	kv, ok := k.(string)
	if !ok {
		return kv, &ErrUnsupportedVariableType{key, k}
	}
	out, err := r.ResolveValueWith(kv, useViper)
	if err != nil {
		return out, err
	}
	return out, nil
}

// ResolveValue with the default Viper
func (r *Resolver) ResolveValue(value string) (string, error) {
	defaultViper := viper.GetViper()
	return r.ResolveValueWith(value, defaultViper)
}

// ResolveValueWith the specified Viper
func (r *Resolver) ResolveValueWith(value string, useViper *viper.Viper) (string, error) {
	if err := r.validateSettings(); err != nil {
		return "", err
	}
	ok, vars, err := r.containsVariables(value)
	if err != nil {
		return "", err
	}
	if !ok {
		return value, nil
	}
	// Verify all variables are resolvable before making changes
	err = checkVariableReference(useViper, vars)
	if err != nil {
		return "", err
	}
	out := value
	for _, v := range vars {
		viperVal := useViper.Get(v)
		delimVar := fmt.Sprintf("%s%s%s", r.DelimStart, v, r.DelimEnd)
		out = strings.Replace(out, delimVar, cast.ToString(viperVal), -1)
	}
	return out, nil
}

// ResolveReplaceKey in the default Viper
func (r *Resolver) ResolveReplaceKey(key string) error {
	defaultViper := viper.GetViper()
	return r.ResolveReplaceKeyIn(key, defaultViper)
}

// ResolveReplaceKeyIn the specified Viper
func (r *Resolver) ResolveReplaceKeyIn(key string, useViper *viper.Viper) error {
	if err := r.validateSettings(); err != nil {
		return err
	}
	val, err := r.ResolveKeyIn(key, useViper)
	if err != nil {
		return err
	}
	useViper.Set(key, val)
	return nil
}

// ResolveReplaceAll values containting variables in the default Viper
func (r *Resolver) ResolveReplaceAll() error {
	defaultViper := viper.GetViper()
	return r.ResolveReplaceAllIn(defaultViper)
}

// ResolveReplaceAllIn the specified Viper
func (r *Resolver) ResolveReplaceAllIn(useViper *viper.Viper) error {
	if err := r.validateSettings(); err != nil {
		return err
	}
	keys := useViper.AllKeys()
	for _, key := range keys {
		excluded := false
		for _, exclude := range r.ExcludeKeys {
			if key == exclude {
				excluded = true
				break
			}
		}
		if excluded {
			continue
		}
		value, err := r.ResolveKeyIn(key, useViper)
		if err != nil {
			if _, ok := err.(*ErrUnsupportedVariableType); ok {
				continue
			}
			return err
		}
		useViper.Set(key, value)
	}
	return nil
}

func (r *Resolver) containsVariables(v string) (bool, []string, error) {
	if err := r.verifyRegex(); err != nil {
		return false, nil, err
	}
	matches := r.regex.FindAllStringSubmatch(v, -1)
	if matches == nil {
		return false, nil, nil
	}
	vars := make([]string, len(matches))
	for i, g := range matches {
		vars[i] = g[1]
	}
	return true, vars, nil
}

func (r *Resolver) validateSettings() error {
	if r.DelimStart == "" {
		return &ErrInvalidDelimiter{"start", "empty"}
	}
	if r.DelimEnd == "" {
		return &ErrInvalidDelimiter{"end", "empty"}
	}
	if r.DelimKey == "" {
		return &ErrInvalidDelimiter{"key", "empty"}
	}
	return nil
}

func (r *Resolver) verifyRegex() error {
	if r.regex == nil {
		r.oldDelimStart = r.DelimStart
		r.oldDelimEnd = r.DelimEnd
		r.oldDelimKey = r.DelimKey
		r.oldKeySpecialCharacters = r.KeySpecialCharacters
		r.recompile = true
	}
	if r.DelimStart != r.oldDelimStart {
		r.oldDelimStart = r.DelimStart
		r.recompile = true
	}
	if r.DelimEnd != r.oldDelimEnd {
		r.oldDelimEnd = r.DelimEnd
		r.recompile = true
	}
	if r.DelimKey != r.oldDelimKey {
		r.oldDelimKey = r.DelimKey
		r.recompile = true
	}
	if !bytes.Equal(r.KeySpecialCharacters, r.oldKeySpecialCharacters) {
		dedupedKSC := removeDuplicateBytes(r.KeySpecialCharacters)
		r.KeySpecialCharacters = dedupedKSC
		r.oldKeySpecialCharacters = dedupedKSC
		r.recompile = true
	}
	if r.recompile {
		start := regexp.QuoteMeta(r.DelimStart)
		end := regexp.QuoteMeta(r.DelimEnd)
		key := regexp.QuoteMeta(r.DelimKey)
		keychars := escapeBytesToString(r.KeySpecialCharacters)
		expr := fmt.Sprintf(
			"%s([A-Za-z0-9%s%s]+)%s",
			start,
			key,
			keychars,
			end,
		)
		// All user definable inputs are sanitized, so unless
		// the world is on fire, no error should be returned here.
		// Ignoring errors outright is bad, so we'll check and
		// return anyways.
		regexpr, err := regexp.Compile(expr)
		if err != nil {
			return err
		}
		r.regex = regexpr
		r.recompile = false
	}
	return nil
}

func removeDuplicateBytes(byteSlice []byte) []byte {
	keys := make(map[byte]bool)
	out := []byte{}
	for _, entry := range byteSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			out = append(out, entry)
		}
	}
	return out
}

func escapeBytesToString(byteSlice []byte) string {
	var out strings.Builder
	for _, b := range byteSlice {
		out.WriteString(`\`)
		out.WriteByte(b)
	}
	return out.String()
}

func checkVariableReference(useViper *viper.Viper, vars []string) error {
	for _, v := range vars {
		viperVal := useViper.Get(v)
		if viperVal == nil {
			return &ErrNonExistentVariableReference{v}
		}
		switch viperVal.(type) {
		case []int, []string, map[string]interface{}, map[string]string:
			return &ErrUnsupportedVariableType{v, viperVal}
		}
	}
	return nil
}
