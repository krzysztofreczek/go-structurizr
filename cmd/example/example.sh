#!/bin/sh

set -e

# Uncomment the line below to see detailed execution logs
#export LOG_LEVEL=debug

rm -rf .out && mkdir .out

go run main.go
echo "Example structure scraped successfully and rendered to: .out/output.plantuml"

if [ "$1" = "--test" ]; then
  exit 0
fi

if ! which plantuml > /dev/null; then
  echo "Plantuml is not available. In order to generate example PNG, please install it and run the script again."
  exit 1
fi

plantuml .out/output.plantuml
echo "Example structure rendered to PDF successfully: .out/output.png"

open .out/output.png
