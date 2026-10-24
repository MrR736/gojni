if (-not $env:JAVA_HOME) {
    $java = (Get-Command java -ErrorAction Stop).Source
    $env:JAVA_HOME = Split-Path (Split-Path $java)
}

$jni = Join-Path $env:JAVA_HOME "include"
$platform = Join-Path $jni "win32"

if (-not (Test-Path (Join-Path $jni "jni.h"))) {
    throw "jni.h not found. Install a JDK and set JAVA_HOME."
}

$env:CGO_ENABLED = "1"
$env:CGO_CFLAGS = "-I$jni -I$platform"

Write-Host "JAVA_HOME=$env:JAVA_HOME"
Write-Host "CGO_CFLAGS=$env:CGO_CFLAGS"

go test ./...