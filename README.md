# Mars
> **WARNING: This is not for regular users**

A super lightweight clone of [api.papermc.io](https://api.papermc.io)

Lightweight does not mean incomplete functionality. On the contrary, Mars has fully implemented all the APIs of bibliothek and further optimized them.

Mars is 1:1 compatible with the bibliothek API (except for database construction and private APIs). You can seamlessly switch to Mars。

[👉Download](https://github.com/LevelTranic/Mars/releases)

[👉Wiki](https://github.com/LevelTranic/Mars/wiki)

## Build
> Build scripts are only useful if cross-platform builds are required.
> 
> To use the build script, you must install the Zig toolchain and PowerShell 7.

Prerequisites for building Mars:
- GCC or clang
- ZigCC (Replacement for gcc and clang, included in the zig toolchain)
- Golang 1.24.3
- PowerShell 7 (No installation required if you don't use the build script)

### Manual build

@Bash
```shell
export CGO_ENABLED=1
go mod tidy
go build -ldflags "-s -w"
```

@CMD
```shell
set CGO_ENABLED=1
go mod tidy
go build -ldflags "-s -w"
```

@PowerShell
```shell
$env:CGO_ENABLED=1
go mod tidy
go build -ldflags "-s -w"
```

### Using build scripts
> Note that the Mars build script requires PowerShell 7 and the Zig toolchain to be installed.
>
> The build script is suitable for cross-compilation.
> 
> The first build will be slow, please be patient.

```shell
./build.ps1
```

## Scan Report
Due to the large number of build products, only reports for the latest versions of amd64 v1 and arm64 are provided.

Version: **v1.4.2**

### AMD64 v1
[👉Windows](https://www.virustotal.com/gui/file/f9ad403aa47a34980865071e60b3dedc46d2933e6dc3bbcc7542fd3d538d48fc?nocache=1)

[👉Linux](https://www.virustotal.com/gui/file/f357959079062db021e2aca8fe158ede0a7e044762df4eabc96132247b94b690?nocache=1)

### ARM64
[👉Windows](https://www.virustotal.com/gui/file/58de988bb1ee771eb39f9e5cc6a64fdd5e40e0659c54eb7e5762a74e07aefa61?nocache=1)

Linux: The toolchain is broken and no binaries can be provided for it.

## Contact Us

Discord: https://discord.gg/dBbSbv2Vuz
