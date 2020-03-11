#!/bin/bash
set -euo pipefail

cd encoder
go test -v
go test --bench=.
