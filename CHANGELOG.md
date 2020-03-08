# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

- `Resolver` fields
  - `DelimStart` - Delimiter marking the beginning of a variable field (default: `$(`)
  - `DelimEnd` - Delimiter marking the end of a variable field (default: `)`)
  - `DelimKey` - Delimiter used by Viper to denote subkeys (default: `.`)
  - `KeySpecialCharacters` - Non-alphanumeric characters used in variables (default: `_` & `-`)
  - `ExcludeKeys` - Keys to ignore with `Resolver.ResolveReplaceAll` and `Resolver.ResolveReplaceAllIn`
- New functions for the default package-level `Resolver`:
  - Function `New() *Resolver`
  - Function `Reset()`
  - Function `SetDelimStart(string)`
  - Function `SetDelimEnd(string)`
  - Function `SetDelimStartEnd(string, string)`
  - Function `SetDelimKey(string)`
  - Function `SetKeySpecialCharacters([]byte)`
  - Function `SetExcludedKeys([]string)`
  - Function `ResolveKey(string) (string, error)`
  - Function `ResolveKeyIn(string, *viper.Viper) (string, error)`
  - Function `ResolveValue(string) (string, error)`
  - Function `ResolveValueWith(string, *viper.Viper) (string, error)`
  - Function `ResolveReplaceKey(string) error`
  - Function `ResolveReplaceKeyIn(string, *viper.Viper) error`
  - Function `ResolveReplaceAll() error`
  - Function `ResolveReplaceAllIn(*viper.Viper) error`
- New `Resolver` methods for more granular control
  - Method `Resolver.ResolveKey(string) (string, error)`
  - Method `Resolver.ResolveKeyIn(string, *viper.Viper) (string, error)`
  - Method `Resolver.ResolveValue(string) (string, error)`
  - Method `Resolver.ResolveValueWith(string, *viper.Viper) (string, error)`
  - Method `Resolver.ResolveReplaceKey(string) error`
  - Method `Resolver.ResolveReplaceKeyIn(string, *viper.Viper) error`
  - Method `Resolver.ResolveReplaceAll() error`
  - Method `Resolver.ResolveReplaceAllIn(*viper.Viper) error`
- Package-specific error types
  - `ErrInvalidDelimiter` - Delimiter is empty or invalid character
  - `ErrNonExistentVariableReference` - Variable references does not exist
  - `ErrUnsupportedVariableType` - Applies to variables referencing maps and slices
  - `ErrNonExistentKey` - Viper key does not exist

### Removed

- Function `NewResolver`
- Function `NewResolverFrom`
- Method `Resolver.SetDelimiters`
- Method `Resolver.Resolve`
