# Prerequisites

- [Install Go](https://golang.org/doc/install)
- [Install lib2git](https://libgit2.github.com)

```
$ # If on Mac OS
$ brew install libgit2
```

# Getting started

```
$ go get github.com/ijsnow/gitup
$ cd $GOPATH/src/github.com/ijsnow/gitup
```

## API

```
$ cd api
```

Add a `.env` file such as this

```
PORT=8080
REPO_DIR=~/ws/tmp/git
DB_PATH=~/ws/tmp/gitup.db
```

```
$ go run main.go
```

## CLI

```
$ cd cli
$ go build gitup.go
$ ./gitup help
```
