# unhttpx

[![Go Reference](https://pkg.go.dev/badge/github.com/melvinsh/unhttpx.svg)](https://pkg.go.dev/github.com/melvinsh/unhttpx)

Turns a list of URLs into hostnames.

## Installation

``` bash
go install -v github.com/melvinsh/unhttpx@latest
```

## Usage

``` bash
$ echo "https://google.com/yo" | unhttpx
google.com

$ cat urls.txt | unhttpx
hackerone.com
zerocopter.com
```
