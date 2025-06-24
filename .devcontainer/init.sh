#!/bin/bash

echo "run .devcontainer/init/sh"
echo "whoami and id"
whoami
id
echo "which go"
which go
echo "install packages"
go clean -cache
go clean -modcache
go install golang.org/x/tools/gopls@latest
go install github.com/cweill/gotests/gotests@v1.6.0
go install github.com/fatih/gomodifytags@v1.17.0
go install github.com/josharian/impl@v1.4.0
go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0
go install github.com/go-delve/delve/cmd/dlv@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
go install golang.org/x/tools/cmd/goimports@latest
echo "finish installing packages"