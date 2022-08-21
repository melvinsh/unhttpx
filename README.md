# unhttpx

[![Go Reference](https://pkg.go.dev/badge/github.com/melvinsh/unhttpx.svg)](https://pkg.go.dev/github.com/melvinsh/unhttpx)

unhttpx is **not** a fast and multi-purpose HTTP toolkit that allows running multiple probes using the retryablehttp library. In fact, it does the exact opposite. 

**TL;DR: turns a list of URLs into hostnames.**

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
