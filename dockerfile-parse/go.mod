module github.com/gucchisk/go-samples/dockerfile-parse

go 1.14

replace github.com/containerd/containerd => github.com/containerd/containerd v1.3.1-0.20200227195959-4d242818bf55

replace github.com/docker/docker => github.com/docker/docker v1.4.2-0.20200227233006-38f52c9fec82

require (
	github.com/asottile/dockerfile v3.1.0+incompatible // indirect
	github.com/moby/buildkit v0.7.1 // indirect
)
