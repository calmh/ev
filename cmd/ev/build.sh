#!/bin/bash
set -euo pipefail

{
    echo -e 'package main\n\nvar indexHtml = []byte(`'
    cat index.html
    echo '`)'
} > index.html.go

{
    echo -e 'package main\n\nvar appJs = []byte(`'
    cat app.js
    echo '`)'
} > app.js.go
