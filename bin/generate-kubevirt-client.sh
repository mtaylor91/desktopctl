#!/bin/sh
client-gen \
  --input-base="kubevirt.io/api/" \
  --input="core/v1" \
  --output-package="github.com/mtaylor91/desktopctl/pkg/kubevirt/client" \
  --clientset-name="versioned" \
  --go-header-file boilerplate.go.txt
