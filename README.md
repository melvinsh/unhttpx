# unhttpx
unhttpx is not a fast and multi-purpose HTTP toolkit that allows running multiple probes using the retryablehttp library. In fact, it does the exact opposite. 

**TL;DR: turns a list of URLs into hostnames.**

## Installation

``` bash
go install -v github.com/melvinsh/unhttpx@latest
```

## Usage

``` bash
cat urls.txt | unhttpx
```
