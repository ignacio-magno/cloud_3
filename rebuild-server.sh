#!/bin/bash
set -e

cd "$(dirname "$0")"

echo "Rebuilding main server..."
go build -o bin/server cmd/server/main.go

echo "Restarting main service..."
systemctl restart servidor-main

echo "Done. Check with: systemctl status servidor-main"
