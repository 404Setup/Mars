package utils

import (
	"bytes"
	"os"
	"runtime"

	"github.com/3JoB/unsafeConvert"
	"github.com/klauspost/cpuid/v2"
)

var (
	HasIntelSha = runtime.GOARCH == "amd64" && cpuid.CPU.Supports(cpuid.SHA, cpuid.SSSE3, cpuid.SSE4)
	HasAvx256   = runtime.GOARCH == "amd64" && cpuid.CPU.Supports(cpuid.AVX, cpuid.AVX2)
	HasAvx512   = runtime.GOARCH == "amd64" && cpuid.CPU.Supports(cpuid.AVX512F, cpuid.AVX512DQ, cpuid.AVX512BW, cpuid.AVX512VL, cpuid.AVX512CD)
	HasArmSha2  = false
)

func init() {
	HasArmSha2 = hasArmSha2()
}

func hasArmSha2() bool {
	if cpuid.CPU.Has(cpuid.SHA2) {
		return true
	}
	if runtime.GOARCH != "arm64" || runtime.GOOS != "linux" {
		return false
	}

	// Fall back to hacky cpuinfo parsing...
	const procCPUInfo = "/proc/cpuinfo"

	// Feature to check for.
	const sha256Feature = "sha2"

	cpuInfo, err := os.ReadFile(procCPUInfo)
	if err != nil {
		return false
	}
	return bytes.Contains(cpuInfo, unsafeConvert.BytePointer(sha256Feature))
}
