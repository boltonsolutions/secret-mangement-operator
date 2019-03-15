#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
deepcopy \
github.com/boltonsolutions/secret-management-operator/pkg/generated \
github.com/boltonsolutions/secret-management-operator/pkg/apis \
secret:v1alpha1 \
--go-header-file "./tmp/codegen/boilerplate.go.txt"
