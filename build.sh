#!/bin/sh

VERSION="v0.2"  # Set your version here
BUILD_DATE="$(date +%Y%m%d)"  # Set your build date here

GREEN='\033[0;32m'
NC='\033[0m'

mkdir -p ./build

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X 'main.VERSION=${VERSION}' -s -w" -o "build/quran-cli-${VERSION}-r${BUILD_DATE}-linux-amd64"
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-X 'main.VERSION=${VERSION}' -s -w" -o "build/quran-cli-${VERSION}-r${BUILD_DATE}-linux-arm64"
echo "${GREEN}✔${NC} Build for Linux done"

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X 'main.VERSION=${VERSION}' -s -w" -o "build/quran-cli-${VERSION}-r${BUILD_DATE}-darwin-amd64"
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-X 'main.VERSION=${VERSION}' -s -w" -o "build/quran-cli-${VERSION}-r${BUILD_DATE}-darwin-arm64"
echo "${GREEN}✔${NC} Build for Darwin done"

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X 'main.VERSION=${VERSION}' -s -w" -o "build/quran-cli-${VERSION}-r${BUILD_DATE}-windows-amd64.exe"
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-X 'main.VERSION=${VERSION}' -s -w" -o "build/quran-cli-${VERSION}-r${BUILD_DATE}-windows-386.exe"
echo "${GREEN}✔${NC} Build for Windows done"

CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags "-X 'main.VERSION=${VERSION}' -s -w" -o "build/quran-cli-${VERSION}-r${BUILD_DATE}-freebsd-amd64"
CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 go build -ldflags "-X 'main.VERSION=${VERSION}' -s -w" -o "build/quran-cli-${VERSION}-r${BUILD_DATE}-freebsd-arm64"
echo "${GREEN}✔${NC} Build for FreeBSD done"

CGO_ENABLED=0 GOOS=openbsd GOARCH=amd64 go build -ldflags "-X 'main.VERSION=${VERSION}' -s -w" -o "build/quran-cli-${VERSION}-r${BUILD_DATE}-openbsd-amd64"
CGO_ENABLED=0 GOOS=openbsd GOARCH=arm64 go build -ldflags "-X 'main.VERSION=${VERSION}' -s -w" -o "build/quran-cli-${VERSION}-r${BUILD_DATE}-openbsd-arm64"
echo "${GREEN}✔${NC} Build for OpenBSD done"

CGO_ENABLED=0 GOOS=netbsd GOARCH=amd64 go build -ldflags "-X 'main.VERSION=${VERSION}' -s -w" -o "build/quran-cli-${VERSION}-r${BUILD_DATE}-netbsd-amd64"
CGO_ENABLED=0 GOOS=netbsd GOARCH=arm64 go build -ldflags "-X 'main.VERSION=${VERSION}' -s -w" -o "build/quran-cli-${VERSION}-r${BUILD_DATE}-netbsd-arm64"
echo "${GREEN}✔${NC} Build for NetBSD done"
