# go-gopher-cli

## Cobra

Install `cobra` library

```bash
$ go get -u github.com/spf13/cobra@latest
```

Install `cobra-cli`

```bash
$ go install github.com/spf13/cobra-cli@latest
```

In `$GOPATH/bin`, you will see `cobra-cli`.

Add the following to `.zshsrc` and run `source ~/.zshsrc`.

```pre
export PATH="$GOPATH/bin:$PATH"
```

Generate CLI application

```bash
$ cobra-cli init
```

A `main.go` file and a `cmd` folder has been created like this:

```pre
.
├── LICENSE
├── README.md
├── bin
├── cmd
│   └── root.go
├── go.mod
├── go.sum
└── main.go
```

More details on using the `cobra-cli` generator, please read [The Cobra Generator README](https://github.com/spf13/cobra-cli/blob/main/README.md)
More details on using the `cobra` library, please read [The Cobra User Guide](https://github.com/spf13/cobra/blob/master/user_guide.md)

## Add first command

Use `cobra-cli` to add new command.

```bash
$ cobra-cli add get
```

### Test it

Test the help of `get` command,

```bash
$ go run ./main.go get -h
This get command will call GitHub respository in order to return the desired Gopher.

Usage:
  go-gopher-cli get [flags]

Flags:
  -h, --help   help for get
```

Test get command,

```bash
$ go run ./main.go get friends
Try to get 'friends' Gopher...
Perfect! Just saved in friends.png!
```

Test with an unknown Gopher,

```bash
$ go run ./main.go get awesome
Try to get 'awesome' Gopher...
Error: awesome not exists! :-(
```

## Build app

```bash
$ task build
task: [build] GOFLAGS=-mod=mod go build -o bin/gopher-cli main.go
```

Test the executable binary,

```bash
$ ./bin/gopher-cli get 5th-element
Try to get '5th-element' Gopher...
Perfect! Just saved in 5th-element.png!

$ file 5th-element.png
5th-element.png: PNG image data, 1338 x 590, 8-bit/color RGBA, non-interlaced
```

Clean all *.png files

```bash
$ task clean
task: [clean] rm *.png
```


## Ref

- [Learning Go by examples: part 3 - Create a CLI app in Go](https://dev.to/aurelievache/learning-go-by-examples-part-3-create-a-cli-app-in-go-1h43)
