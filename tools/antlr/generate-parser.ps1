param(
    [Parameter(Mandatory = $true)]
    [string] $Jar,
    [Parameter(Mandatory = $true)]
    [string] $Grammar,
    [Parameter(Mandatory = $true)]
    [string] $OutputDir
)

$ErrorActionPreference = "Stop"

$Root = Resolve-Path (Join-Path $PSScriptRoot "..\..")
$JarPath = Join-Path $Root $Jar
$GrammarPath = Join-Path $Root $Grammar
$OutputPath = Join-Path $Root $OutputDir

$JavaCandidates = @(
    "C:\dev\jdk-24.0.1\bin\java.exe",
    "java"
)

$Java = $null
foreach ($Candidate in $JavaCandidates) {
    $Command = Get-Command $Candidate -ErrorAction SilentlyContinue
    if ($Command) {
        $Java = $Command.Source
        break
    }
    if (Test-Path -LiteralPath $Candidate) {
        $Java = $Candidate
        break
    }
}

if (-not $Java) {
    throw "JDK 17+ is required to run ANTLR."
}

New-Item -ItemType Directory -Force $OutputPath | Out-Null

& $Java -jar $JarPath -Dlanguage=Go -visitor -package gen -Xexact-output-dir -o $OutputPath $GrammarPath
if ($LASTEXITCODE -ne 0) {
    throw "ANTLR parser generation failed with exit code $LASTEXITCODE"
}
