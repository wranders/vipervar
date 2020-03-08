# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

- `Resolver` fields
  - `DelimStart` - Delimiter marking the beginning of a variable field (default: `$(`)
  - `DelimEnd` - Delimiter marking the end of a variable field (default: `)`)
  - `DelimKey` - Delimiter used by Viper to denote subkeys (default: `.`)
  - `KeySpecialCharacters` - Non-alphanumeric characters used in variables (default: `_` & `-`)
- New functions to manipulate the default package-level `Resolver`:
  - Function `New() *Resolver`
  - Function `Reset()`
  - Function `SetDelimStart(string)`
  - Function `SetDelimEnd(string)`
  - Function `SetDelimKey(string)`
  - Function `SetKeySpecialCharacters([]string)`
  - Function `SetExcludedKeys([]string)`
  - Function `Resolve(string) (string, error)`
  - Function `ResolveIn(*viper.Viper, string) (string, error)`
  - Function `ResolveValue(string) (string, error)`
  - Function `ResolveValueIn(*viper.Viper, string) (string, error)`
  - Function `ResolveReplace(string) error`
  - Function `ResolveReplaceIn(*viper.Viper, string) error`
  - Function `ResolveReplaceAll() error`
  - Function `ResolveReplaceAllIn(*viper.Viper) error`
- New `Resolver` methods for more granular control
  - Method `Resolver.ResolveIn(*viper.Viper, string) (string, error)`
  - Method `Resolver.ResolveValue(string) (string, error)`
  - Method `Resolver.ResolveValueIn(*viper.Viper, string) (string, error)`
  - Method `Resolver.ResolveReplace(string) error`
  - Method `Resolver.ResolveReplaceIn(*viper.Viper, string) error`
  - Method `Resolver.ResolveReplaceAll() error`
  - Method `Resolver.ResolveReplaceAllIn(*viper.Viper) error`
- Package-specific error types
  - `ErrInvalidDelimiter` - Delimiter is empty or invalid character
  - `ErrNonExistentVariableReference` - Variable references does not exist
  - `ErrUnsupportedVariableType` - Applies to variables referencing maps and slices
  - `ErrNonExistentKey` - Viper key does not exist

### Changed

- Method `Resolver.Resolve() error` -> `Resolver.Resolve(string) (string, error)`
- Viper configurations are decoupled from the `Resolver`
  - `Resolve*` functions/methods ending in `*In` use the specified Viper configuration
  - `Resolve*` functions/methods not ending in `*In` use the default package-level Viper

### Removed

- Function `NewResolver`
- Function `NewResolverFrom`
- Method `Resolver.SetDelimiters`
