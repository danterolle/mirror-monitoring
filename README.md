# ParrotOS Mirror Monitoring

This utility is used to monitor the status of each mirror of the ParrotOS project, it makes a simple HEAD request using Go's http.Head() library for each url present in the mirrors.json file.

To start it, `go run *.go` or build it with `go build`, once started it will expose an API on the `/mirrors/status` address (test it locally and check `http://localhost:8080/mirrors/status`) where each mirror will be associated with the status (online, offline or unknown).
