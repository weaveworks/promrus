# promrus
Logrus hook to expose Prometheus metrics.

## Usage
```
hook, err := promrus.NewPrometheusHook(appName)
if err != nil {
	log.AddHook(hook)
}
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
