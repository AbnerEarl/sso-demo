#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app_linux
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o app_windows
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -o app_macos

#1. GOARCH（目标平台架构）可能的值：
#   - `amd64`：64位 x86 架构
#   - `386`：32位 x86 架构
#   - `arm`：ARM 架构（32位）
#   - `arm64`：ARM64 架构（64位）
#   - `ppc64`：64位 PowerPC 架构
#   - `ppc64le`：64位小端 PowerPC 架构
#   - `mips64`：64位 MIPS 架构
#   - `mips64le`：64位小端 MIPS 架构
#   - `s390x`：64位 IBM z/Architecture
#
#2. GOOS（目标平台操作系统）可能的值：
#   - `linux`：Linux 操作系统
#   - `windows`：Windows 操作系统
#   - `darwin`：macOS 操作系统
#   - `freebsd`：FreeBSD 操作系统
#   - `netbsd`：NetBSD 操作系统
#   - `openbsd`：OpenBSD 操作系统
#   - `dragonfly`：DragonFly BSD 操作系统
#   - `solaris`：Solaris 操作系统
#   - `plan9`：Plan 9 操作系统
#   - `aix`：IBM AIX 操作系统
