package vipervar

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

var specialCharacters = []string{
	`%`, `^`, `$`, `?`, `*`, `+`, `[`, `]`, `(`, `)`,
}

// Periods are considered illegal since Viper uses them internally
// to delimit key children
// Forward and backward slashes are considered illegal since they
// are used in paths
var illegalCharacters = []string{
	`.`, `/`, `\`,
}

// Resolver contains the settings used to resolve variables in a
// Viper configuration
type Resolver struct {
	ExcludeKeys []string

	viper          *viper.Viper
	regex          *regexp.Regexp
	startDelimiter string
	endDelimiter   string
}

// NewResolver uses the default Viper configuration
func NewResolver() Resolver {
	return Resolver{
		viper: viper.GetViper(),
	}
}

// NewResolverFrom uses the specified Viper configuration
func NewResolverFrom(v *viper.Viper) Resolver {
	return Resolver{
		viper: v,
	}
}

// SetDelimiters specifies the strings that mark the beginning and end of a variable
func (r *Resolver) SetDelimiters(start string, end string) error {
	sd, err := validateDelimiter(start)
	if err != nil {
		return err
	}
	ed, err := validateDelimiter(end)
	if err != nil {
		return err
	}
	r.startDelimiter = start
	r.endDelimiter = end

	exp := fmt.Sprintf("%s([a-zA-Z_.]*?)%s", sd, ed)
	r.regex = regexp.MustCompile(exp)

	return nil
}

// Resolve validates all non-excluded configuration keys and replaces them with
// their specified values.
// If no delimiters are specified, the default is `%` for both start and end
// (ie. `%SOME_VAR%`)
func (r *Resolver) Resolve() error {
	if r.regex == nil {
		r.SetDelimiters(`%`, `%`)
	}

	var errs []string
	queue := make(map[string][]string)

	for _, key := range r.viper.AllKeys() {
		// If key is excluded, skip it
		if isIn(key, r.ExcludeKeys) {
			continue
		}

		// Only work with string valued keys
		var strValue string
		value := r.viper.Get(key)
		if val, ok := value.(string); !ok {
			continue
		} else {
			strValue = val
		}

		// Find all variables in string key
		match := r.regex.FindAllStringSubmatch(strValue, -1)
		vars := make([]string, len(match))
		for i, group := range match {
			vars[i] = group[1]
		}
		if len(vars) <= 0 {
			// No variables found, skip key
			continue
		}
		for _, v := range vars {
			if ok := r.viper.IsSet(v); !ok {
				msg := fmt.Sprintf(
					"Configuration variable `%s` references a non-existent value",
					v,
				)
				errs = append(errs, msg)
			}
		}
		queue[key] = vars
	}

	// If any errors were found, return them so no changes are made
	if len(errs) > 0 {
		errsStr := strings.Join(errs[:], "\n")
		return fmt.Errorf(
			"Errors were found in the specified configuration, no changes were made\n%s",
			errsStr,
		)
	}

	// Configuration variables are valid, commit the changes
	for key, vars := range queue {
		for _, v := range vars {
			k := r.viper.GetString(key)
			ref := r.viper.GetString(v)
			delimvar := fmt.Sprintf(
				"%s%s%s",
				r.startDelimiter,
				v,
				r.endDelimiter,
			)
			replaced := strings.Replace(k, delimvar, ref, -1)
			r.viper.Set(key, replaced)
		}
	}

	return nil
}

func validateDelimiter(delim string) (vdelim string, err error) {
	for _, c := range strings.Split(delim, "") {
		if isIn(c, illegalCharacters) {
			msg := fmt.Sprintf("Delimeter `%s` contains illegal character: `%s`", delim, c)
			err = errors.New(msg)
			return
		}
		if isIn(c, specialCharacters) {
			vdelim = vdelim + `\` + c
		} else {
			vdelim = vdelim + c
		}
	}
	return
}

func isIn(key string, stringslice []string) bool {
	for _, value := range stringslice {
		if key == value {
			return true
		}
	}
	return false
}
