#!/usr/bin/env sh
set -eu

APP="mywebfont"
CMD="./cmd/mywebfont"
ARCH="${ARCH:-amd64}"
OUT_DIR="bin"

build_one() {
	platform="$1"
	goos="$2"
	ext="$3"
	work_dir="$OUT_DIR/$APP-$platform"
	binary="$APP$ext"
	zip_file="$OUT_DIR/$APP-$platform.zip"

	rm -rf "$work_dir" "$zip_file"
	mkdir -p "$work_dir"

	echo "==> $platform ($goos/$ARCH)"
	GOOS="$goos" GOARCH="$ARCH" CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o "$work_dir/$binary" "$CMD"

	cp README.md "$work_dir/"
	( cd "$OUT_DIR" && zip -qr "$APP-$platform.zip" "$APP-$platform" )
	rm -rf "$work_dir"
}

mkdir -p "$OUT_DIR"

build_one "linux" "linux" ""
build_one "mac" "darwin" ""
build_one "windows" "windows" ".exe"

echo "Archives generees dans $OUT_DIR/:"
ls -1 "$OUT_DIR"/"$APP"-*.zip
