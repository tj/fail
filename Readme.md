
# Fail

fail(1) is a dumb little ad-hoc HTTP server which can be configured to fail and increase latency at random.

## Parameters

### max_delay=N

 Max delay in milliseconds. Example of 200ms: `?max_delay=200`.

### error_rate=N

Error rate. Example of 1/3 failures: `?error_rate=3`.

## Badges

[![GoDoc](https://godoc.org/github.com/tj/fail?status.svg)](https://godoc.org/github.com/tj/fail)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)

---

> [tjholowaychuk.com](http://tjholowaychuk.com) &nbsp;&middot;&nbsp;
> GitHub [@tj](https://github.com/tj) &nbsp;&middot;&nbsp;
> Twitter [@tjholowaychuk](https://twitter.com/tjholowaychuk)
