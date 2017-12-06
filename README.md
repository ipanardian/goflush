# GoFlush
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/ipanardian/goflush/issues)
[![GitHub release](https://img.shields.io/github/release/ipanardian/goflush.svg)]()
[![GitHub license](https://img.shields.io/badge/license-Apache-red.svg)](https://raw.githubusercontent.com/ipanardian/goflush/master/LICENSE)

A simple command line app to help you flush or delete Redis keys easily.

## How to install
```
$ go get github.com/ipanardian/goflush
```

## How to use
```
$ goflush //no argument will flush all keys
$ OK

$ goflush follower //will delete key 'follower'
$ (integer) 1
```

## License
Apache License 2.0
