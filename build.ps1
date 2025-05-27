#!/usr/bin/env pwsh

$GOAMD64Versions = @('v1', 'v2', 'v3', 'v4')

$originalEnv = @{
    "GOOS" = go env GOOS
    "GOARCH" = go env GOARCH
    "GOAMD64" = go env GOAMD64
    "CC" = $env:CC
    "CXX" = $env:CXX
    "CGO_CFLAGS" = $env:CGO_CFLAGS
    "CGO_CXXFLAGS" = $env:CGO_CXXFLAGS
    "CGO_FFLAGS" = $env:CGO_FFLAGS
    "CGO_LDFLAGS" = $env:CGO_LDFLAGS
    "CGO_ENABLED" = $env:CGO_ENABLED
}

function Build-Project {
    param (
        [string]$arch,
        [string]$version,
        [string]$outputFile,
        [string]$os
    )
    $env:GOAMD64 = $version
    $env:GOOS = $os

    $env:GOAMD64=$version

    if ($os -eq "windows") {
            if ($arch -eq "arm64") {
                $env:CC = "zig cc -target aarch64-windows"
                $env:CXX = "zig c++ -target aarch64-windows"
            } else {
                $env:CC = "gcc"
                $env:CXX = "g++"
            }
            $env:CGO_CFLAGS = '-g -flto -O3'
            $env:CGO_CXXFLAGS = '-g -flto -O3'
            $env:CGO_FFLAGS = '-g -flto -O3'
            $env:CGO_LDFLAGS = '-g -flto -O3'
        } else {
            if ($arch -eq "arm64") {
                $env:CC = "zig cc -target aarch64-linux"
                $env:CXX = "zig c++ -target aarch64-linux"
            } else {
                $env:CC = "zig cc -target x86_64-linux-gnu"
                $env:CXX = "zig c++ -target x86_64-linux-gnu"
            }
        }

    Write-Output "Building for OS=$os ARCH=$arch GOAMD64=$version"
    go build -ldflags "-s -w" -o $outputFile

    if ($os -eq "windows") {
            if ($arch -eq "amd64") {
                $zipFileName = "mars_windows_${arch}_$version.zip"
            } else {
                $zipFileName = "mars_windows_${arch}.zip"
            }
            Compress-Archive -Path $outputFile -DestinationPath $zipFileName -Force
    } else {
        if ($arch -eq "amd64") {
            $tarFileName = "mars_linux_${arch}_$version.tar.xz"
        } else {
            $tarFileName = "mars_linux_${arch}.tar.xz"
        }
        tar -cJf $tarFileName $outputFile
    }
    Write-Output "Build completed for OS=$os GOAMD64=$version on $arch"
}

$env:CGO_ENABLED = 1
go mod tidy
foreach ($version in $GOAMD64Versions) {
    Build-Project -arch "amd64" -version $version -outputFile "Mars.exe" -os "windows"
    Build-Project -arch "amd64" -version $version -outputFile "Mars" -os "linux"
}

$version = "v1"
$env:GOARCH = "arm64"
Build-Project -arch "arm64" -version $version -outputFile "Mars.exe" -os "windows"
Build-Project -arch "arm64" -version $version -outputFile "Mars" -os "linux"

Write-Output "All builds completed."

Write-Output "Restoring original environment..."
$env:GOOS = $originalEnv.GOOS
$env:GOARCH = $originalEnv.GOARCH
$env:GOAMD64 = $originalEnv.GOAMD64
$env:CC = $originalEnv.CC
$env:CXX = $originalEnv.CXX
$env:CGO_CFLAGS = $originalEnv.CGO_CFLAGS
$env:CGO_CXXFLAGS = $originalEnv.CGO_CXXFLAGS
$env:CGO_FFLAGS = $originalEnv.CGO_FFLAGS
$env:CGO_LDFLAGS = $originalEnv.CGO_LDFLAGS
$env:CGO_ENABLED = $originalEnv.CGO_ENABLED

Write-Output "Environment restored to original state."
