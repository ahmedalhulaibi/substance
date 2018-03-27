#!/bin/sh
echo "pre-commit hook starting"
if go test ./...; then
    exit 0
else
    echo "Aborting commit go test failed"
    exit 1
fi
echo "pre-commit hook finished"