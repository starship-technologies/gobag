Gobag: a standalone ROS bag reader

### What is it?

Gobag is exploration around the idea of fast [ROS](http://www.ros.org/ "Robot Operating System") bag parser that can easily be deployed without ROS stack.

### Installation

It uses Go modules system for dependency management.

In older versions of Go, you need to do the usual `go get <all the dependencies>`.

In Go 1.11 (and working in GOPATH) you can turn on the modules support by this environment variable `export GO111MODULE=on` and then just build. Much more on this [here](https://github.com/golang/go/wiki/Modules#how-to-define-a-module).

In future Go versions (starting 1.12) it should just work automagically as the module support will be switched on by default.

Then build the CLI tool:

```
cd gobag-cli/
go build                                      # or go install
./gobag-cli help
```

**Note**: it works 99.99% of the time with the newest version of [lz4 package](https://github.com/pierrec/lz4), but in case of very large _single_ message (more than 2^10 bytes) the version 1.1 needs to be used (currently newest is 2.0.7).

### Usage

CLI usage is shown below.

Code usage examples are shown in `gobag-cli/actions.go`.

### TODO

* Provide out-of-the-box code samples with a sample bag file
* Benchmarking tests
* Work with lz4 package author to track down the large message regression

---
# `gobag-cli`

1.0.0 - Starship Technologies OÜ <technology@starship.xyz>

## Commands (3)

### `gobag-cli dump`

Dump indicated content of the bag 

#### Subcommands (6)

### `gobag-cli dump chunks`

dump uncompressed chunks

### `gobag-cli dump chunksinfo`

dump chunk information

### `gobag-cli dump messagedefinitions`

dump message definitions

### `gobag-cli dump tabledefinitions`

dump HIVE DDL table definitions for all topics into separate .sql files

### `gobag-cli dump json`

dump full bag to JSON

### `gobag-cli dump topics`

dump bag messages to JSON by topic

#### Flags

- `--time value`: Comma separated Unix epoch timestamps for start and end to filter by time
- `--filter value`: Comma separated list of topics to limit output (including '/' prefix if needed)


---

### `gobag-cli docs`

Usage: `gobag-cli docs > documentation.md`

Generate documentation in markdown format and print to standard out.

---

### `gobag-cli help`

Usage: `Shows a list of commands or help for one command`

---

## Global Flags

- `--help, -h`: show help
- `--version, -v`: print the version


