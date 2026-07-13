#!/usr/bin/env bash
set -e

# Navigate to project root
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

echo "Building Goodbye Gmail..."

# Rename previous stable to last if it exists
if [ -f "build/stable" ]; then
    echo "Backing up previous stable build to 'last'..."
    mv build/stable build/last
fi

# Compile the new stable build
echo "Compiling..."
go build -o build/stable ./src/cmd/goodbye-gmail

echo "Build successful! Executable is at build/stable"
