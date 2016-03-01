cli arguments library
=====================

Yet another go cli arguments parsing library.

Installation
------------

```
go get github.com/shagabutdinov/arguments
```

Usage
-----

```
var (
  options = arguments.Arguments{
    "config": arguments.Argument{
      "config",
      "configuration file",
      arguments.String,
      "c",
      false,
    },
  }
)

func main() {
  arguments, err := options.Parse(optionsArray)
  if err != nil {
    log.Println(err.Error())
    os.Exit(1)
  }

  configFile, isDefault, err = arguments.String("local", false)
  if err != nil {
    log.Println(err.Error())
    os.Exit(1)
  }

  log.Println(configFile)
}
```
