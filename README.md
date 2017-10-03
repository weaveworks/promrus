<p align="center">
	<a href="https://github.com/marccarre/promrus/releases/latest">
		<img src="https://img.shields.io/github/release/marccarre/promrus.svg"/>
	</a>
	<a href="https://travis-ci.org/marccarre/promrus">
		<img src="https://img.shields.io/travis/marccarre/promrus.svg"/>
	</a>
	<a href="https://coveralls.io/github/marccarre/promrus?branch=master">
		<img src="https://img.shields.io/coveralls/marccarre/promrus.svg"/>
	</a>
	<a href="https://goreportcard.com/report/github.com/marccarre/promrus">
		<img src="https://goreportcard.com/badge/github.com/marccarre/promrus&x=1"/>
	</a>
	<a href="LICENSE">
		<img src="https://img.shields.io/badge/license-Apache%202.0-blue.svg"/>
	</a>
</p>

# promrus
Logrus hook to expose the number of log messages as Prometheus metrics:
```
log_messages{level="debug"}
log_messages{level="info"}
log_messages{level="warning"}
log_messages{level="error"}
```

## Usage
```
hook, err := promrus.NewPrometheusHook(appName)
if err != nil {
	log.AddHook(hook)
}
```

## Setup development environment
```
$ go get github.com/golang/dep/cmd/dep
$ dep ensure
```

## Compile
```
$ go build
```

## Test
```
$ go test
DEBU[0000] this is at debug level!                      
INFO[0000] this is at info level!                       
WARN[0000] this is at warning level!                    
ERRO[0000] this is at error level!                      
PASS
ok  	github.com/marccarre/promrus	0.011s
```
