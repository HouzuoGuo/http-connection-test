# http-connection-test
A basic HTTP client for diagnosis of HTTP/S traffic interception behaviour.

Build using go compiler:

`env CGO_ENABLED=0 go build`

And transport the result executable into destination host, then invoke using this syntax:

`./http-connection-test -host hz.gl -port 443 -url / -tls`
