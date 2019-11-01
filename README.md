# ViperVar

Variable resolution for [Viper](https://github.com/spf13/viper) configurations

## Install

```shell
go get -u github.com/wranders/vipervar
```

---

## The What and Why

ViperVar allows you to resolve variables referencing other Viper settings within your configuration files. The cuts down on the number of lines that need to be editied if different settings contain similar information.

ViperVar has been tested against JSON, TOML, and YAML files. Any loaded configuration should work, but other formats are untested.

Resolving variables replaces the values within Viper, so ***this should be considered destructive***. Unless a key is excluded, variables will be overwritten if Viper was used to save a configuration to file. If Viper is not used to write configuration files, then variables will remain intact within the file.

Your configuration file:

```yaml
app_dir: "/home/user/app"
application:
    name: "MyApp"
server:
    root_url: "%SERVER.PROTOCOL%://%SERVER.DOMAIN%:%SERVER.PORT%/"
    adsr: "%SERVER.LISTEN%:%SERVER.PORT%"
    domain: "myapp.foo"
    protocol: "http"
    listen: "0.0.0.0"
    port: 80
database:
    file: "%APP_DIR%/db/app.db"
    . . .
```

Once resolved, becomes:

```yaml
app_dir: "/home/user/app"
application:
    name: "MyApp"
server:
    root_url: "http://myapp.foo:80/"
    adsr: "0.0.0.0:80"
    domain: "myapp.foo"
    protocol: "http"
    listen: "0.0.0.0"
    port: 80
database:
    file: "/home/user/app/db/app.db"
    . . .
```

By default variable delimiters are the percent (`%`) sign, but are completely customizable with the exception of periods (`.`), since Viper uses those to delimit child keys, and slashes (`/` and `\`), since distiguishing variables from paths would then be too difficult.

---

## Getting Started

The default Viper can be used by simply using:

```go
r := vipervar.NewResolver()
err := r.Resolve()
if err != nil {
    panic(fmt.Errorf("Error resolving configuration:\n%s\n", err))
}
```

Or, a custom Viper can be resolved using `NewResolverFrom(*viper.Viper)`:

```go
v := viper.New()
err := v.ReadInConfig()
if err != nil {
    panic(fmt.Errorf("Error reading config file: %s\n", err))
}
r := vipervar.NewResolverFrom(v)
err = r.Resolve()
if err != nil {
    panic(fmt.Errorf("Error resolving configuration:\n%s\n", err))
}
```

Delimiters can be set using the `SetDelimiters(start, end)` function:

```go
r := vipervar.NewResolver()
r.SetDelimiters(`{{`, `}}`)
```

**Delimiters should be set using string literals**. Under the hood, ViperVar uses regular expressions and all special characters are automatically escaped. Since slashes are considered illegal characters (reason above), escaping them would cause this action to fail and return an error.

All variables must reference an existing Viper setting. If an undefined setting key is encountered, `Resolve` will return an error containing all invalid keys and no changes will be made to the loaded Viper configuration.

## Key Exclusion

If you have a setting that you do not want to be resolved, add it to the `Resolver`'s `ExcludedKeys` slice:

```go
r := vipervar.NewResolver()
r.ExcludedKeys = []string{"some_setting"}
r.ExcludedKeys = append(r.ExcludedKeys, "another_setting")
```

## Watching

To re-resolve variables in a watched configuration file, add the `Resolve` function to your Viper's `OnConfigChange` function:

```go
r := vipervar.NewResolver()
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
    r.Resolve()
})
```
