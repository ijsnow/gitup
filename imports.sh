#!/bin/bash

go list -f '{{join .Deps "\n"}}' > deps.txt
go list -f '{{join .TestImports "\n"}}' >> deps.txt
sort deps.txt -u -o deps.txt
go list std | sort | comm -23 ./deps.txt - | grep -v `go list` |  xargs -I{} $GOPATH/bin/gvt fetch {}
rm ./deps.txt
