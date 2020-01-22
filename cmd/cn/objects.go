package main

import (
	"github.com/codenotary/objects/pkg/extractor"
	"github.com/codenotary/objects/pkg/extractor/docker"
	"github.com/codenotary/objects/pkg/extractor/file"
	"github.com/codenotary/objects/pkg/extractor/git"
)

var _ = (func() interface{} {
	extractor.Register(file.Scheme, file.Extract)
	extractor.Register(git.Scheme, git.Extract)
	extractor.Register(docker.Scheme, docker.Extract)
	return nil
})()
