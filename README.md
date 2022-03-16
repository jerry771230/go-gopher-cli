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


## Ref

- [Learning Go by examples: part 3 - Create a CLI app in Go](https://dev.to/aurelievache/learning-go-by-examples-part-3-create-a-cli-app-in-go-1h43)
