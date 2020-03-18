# xgo
This is cli tool to build all binary for multi platform.
See the following page about combination of $GOOS and $GOARCH.

** I noticed that you should use gox after I made this tool! hahaha
https://github.com/mitchellh/gox

# Usage
If you want to specified binary name, specify the name in first argument (ex. $ xgo hogehogebin)
```
$ xgo
binary name: xgo
made binary: /.../dtmu/xgo/bin/aix/ppc64/xgo
made binary fail: /.../dtmu/xgo/bin/android/386/xgo
made binary fail: /.../dtmu/xgo/bin/android/amd64/xgo
made binary fail: /.../dtmu/xgo/bin/android/arm/xgo
made binary fail: /.../dtmu/xgo/bin/android/arm64/xgo
made binary: /.../dtmu/xgo/bin/darwin/386/xgo
made binary: /.../dtmu/xgo/bin/darwin/amd64/xgo
made binary fail: /.../dtmu/xgo/bin/darwin/arm/xgo
made binary fail: /.../dtmu/xgo/bin/darwin/arm64/xgo
made binary: /.../dtmu/xgo/bin/dragonfly/amd64/xgo
made binary: /.../dtmu/xgo/bin/freebsd/386/xgo
made binary: /.../dtmu/xgo/bin/freebsd/amd64/xgo
made binary: /.../dtmu/xgo/bin/freebsd/arm/xgo
made binary: /.../dtmu/xgo/bin/illumos/amd64/xgo
made binary: /.../dtmu/xgo/bin/js/wasm/xgo
made binary: /.../dtmu/xgo/bin/linux/386/xgo
made binary: /.../dtmu/xgo/bin/linux/amd64/xgo
made binary: /.../dtmu/xgo/bin/linux/arm/xgo
made binary: /.../dtmu/xgo/bin/linux/arm64/xgo
made binary: /.../dtmu/xgo/bin/linux/ppc64/xgo
made binary: /.../dtmu/xgo/bin/linux/ppc64le/xgo
made binary: /.../dtmu/xgo/bin/linux/mips/xgo
made binary: /.../dtmu/xgo/bin/linux/mipsle/xgo
made binary: /.../dtmu/xgo/bin/linux/mips64/xgo
made binary: /.../dtmu/xgo/bin/linux/mips64le/xgo
made binary: /.../dtmu/xgo/bin/linux/s390x/xgo
made binary: /.../dtmu/xgo/bin/netbsd/386/xgo
made binary: /.../dtmu/xgo/bin/netbsd/amd64/xgo
made binary: /.../dtmu/xgo/bin/netbsd/arm/xgo
made binary: /.../dtmu/xgo/bin/openbsd/386/xgo
made binary: /.../dtmu/xgo/bin/openbsd/amd64/xgo
made binary: /.../dtmu/xgo/bin/openbsd/arm/xgo
made binary: /.../dtmu/xgo/bin/openbsd/arm64/xgo
made binary: /.../dtmu/xgo/bin/plan9/386/xgo
made binary: /.../dtmu/xgo/bin/plan9/amd64/xgo
made binary: /.../dtmu/xgo/bin/plan9/arm/xgo
made binary: /.../dtmu/xgo/bin/solaris/amd64/xgo
made binary: /.../dtmu/xgo/bin/windows/386/xgo.exe
made binary: /.../dtmu/xgo/bin/windows/amd64/xgo.exe
```
