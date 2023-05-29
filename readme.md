If you start this app with `go run cmd/web/main.go` then you get the following error:

durdekm@mbp2014 38-using-pat-for-routing % go run cmd/web/main.go
# command-line-arguments
cmd/web/main.go:46:12: undefined: routes
durdekm@mbp2014 38-using-pat-for-routing %

Start the app with `go run cmd/web/*.go` and it should work

durdekm@mbp2014 38-using-pat-for-routing % go run cmd/web/*.go
Starting application on port 127.0.0.1:8080