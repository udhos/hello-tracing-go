# hello-tracing-go

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
