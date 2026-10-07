#!/bin/sh

set -e

# Uncomment the line below to see detailed execution logs
#export LOG_LEVEL=debug

rm -rf .out && mkdir .out

go run main.go
echo "Example structure scraped successfully and rendered to: .out/output.mmd"

if [ "$1" = "--test" ]; then
  exit 0
fi

echo "Open .out/output.mmd in a Mermaid renderer (e.g. https://mermaid.live) to preview the diagram."
