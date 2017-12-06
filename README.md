# GoFlush
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/ipanardian/goflush/issues)
[![GitHub release](https://img.shields.io/github/release/ipanardian/goflush.svg)]()
[![GitHub license](https://img.shields.io/badge/license-Apache-red.svg)](https://raw.githubusercontent.com/ipanardian/goflush/master/LICENSE)

A simple command line app to help you flush or delete Redis keys easily.

## How to install
```
$ go get github.com/ipanardian/goflush
```
Make sure your PATH includes the $GOPATH/bin directory

## How to use
```
//flush all keys
$ goflush
$ OK

//delete specific key
$ goflush {key_name}
$ (integer) 1

$ goflush -version / -v
$ goflush help
```

## License
Apache License 2.0
