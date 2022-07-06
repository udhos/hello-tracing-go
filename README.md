# hello-tracing-go

# Open Telemetry tracing with Gin

See https://github.com/udhos/kubecloudconfigserver/blob/main/cmd/kubeconfigserver/tracing.go

1) Initialize the tracing (see main.go)

2) Enable trace propagation (see tracePropagation below)

3) Use handler middleware (see main.go)
   import "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
   router.Use(otelgin.Middleware("virtual-service"))

4) For http client, create a Request from Context (see backend.go)
   newCtx, span := b.tracer.Start(ctx, "backendHTTP.fetch")
   req, errReq := http.NewRequestWithContext(newCtx, "GET", u, nil)
   client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
   resp, errGet := client.Do(req)

## fib

Trace output: traces-fib.txt

https://opentelemetry.io/docs/instrumentation/go/getting-started/

```
Run
├── Poll
└── Write
    └── Fibonacci
```

## web1

Trace output: traces-web1.txt

https://opentelemetry.io/docs/instrumentation/go/libraries/

curl localhost:3001/hello-instrumented

## web2

Trace output: traces-web2.txt

curl localhost:3002/hello-instrumented

## web3

Trace output: traces-web3.txt

curl localhost:3003/hello-instrumented

# References

[Traces](https://opentelemetry.io/docs/concepts/signals/traces/)

[Export to Jaeger](https://github.com/open-telemetry/opentelemetry-go/blob/main/example/jaeger/main.go)
