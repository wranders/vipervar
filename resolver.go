package vipervar

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Resolver struct {
	DelimStart           string
	DelimEnd             string
	DelimKey             string
	KeySpecialCharacters []string
	ExcludeKeys          []string

	regex                   *regexp.Regexp
	recompile               bool
	oldDelimStart           string
	oldDelimEnd             string
	oldDelimKey             string
	oldKeySpecialCharacters []string
}

func (r *Resolver) Resolve(key string) (string, error) {
	defaultViper := viper.GetViper()
	return r.ResolveIn(defaultViper, key)
}

func (r *Resolver) ResolveIn(useViper *viper.Viper, key string) (string, error) {
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
	out, err := r.ResolveValueIn(useViper, kv)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (r *Resolver) ResolveValue(value string) (string, error) {
	defaultViper := viper.GetViper()
	return r.ResolveValueIn(defaultViper, value)
}

func (r *Resolver) ResolveValueIn(useViper *viper.Viper, value string) (string, error) {
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
	out := value
	for _, v := range vars {
		viperVal := useViper.Get(v)
		if viperVal == nil {
			return value, &ErrNonExistentVariableReference{v}
		}
		switch viperVal.(type) {
		case []int, []string, map[string]interface{}, map[string]string:
			return out, &ErrUnsupportedVariableType{v, viperVal}
		}
		delimVar := fmt.Sprintf("%s%s%s", r.DelimStart, v, r.DelimEnd)
		out = strings.Replace(out, delimVar, cast.ToString(viperVal), -1)
	}
	return out, nil
}

func (r *Resolver) ResolveReplace(key string) error {
	defaultViper := viper.GetViper()
	return r.ResolveReplaceIn(defaultViper, key)
}

func (r *Resolver) ResolveReplaceIn(useViper *viper.Viper, key string) error {
	if err := r.validateSettings(); err != nil {
		return err
	}
	val, err := r.ResolveIn(useViper, key)
	if err != nil {
		return err
	}
	useViper.Set(key, val)
	return nil
}

func (r *Resolver) ResolveReplaceAll() error {
	defaultViper := viper.GetViper()
	return r.ResolveReplaceAllIn(defaultViper)
}

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
		value, err := r.ResolveIn(useViper, key)
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
	if len(r.KeySpecialCharacters) != len(r.oldKeySpecialCharacters) {
		r.oldKeySpecialCharacters = r.KeySpecialCharacters
		r.recompile = true
	}
	for i, ksc := range r.KeySpecialCharacters {
		if ksc != r.oldKeySpecialCharacters[i] {
			r.oldKeySpecialCharacters = r.KeySpecialCharacters
			r.recompile = true
			break
		}
	}
	if r.recompile {
		start := regexp.QuoteMeta(r.DelimStart)
		end := regexp.QuoteMeta(r.DelimEnd)
		key := regexp.QuoteMeta(r.DelimKey)
		keychars := make([]string, len(r.KeySpecialCharacters))
		if len(r.KeySpecialCharacters) > 0 {
			for i, c := range r.KeySpecialCharacters {
				keychars[i] = regexp.QuoteMeta(c)
			}
		}
		expr := fmt.Sprintf(
			"%s([A-Za-z0-9-%s%s]+)%s",
			start,
			key,
			strings.Join(keychars, ""),
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
