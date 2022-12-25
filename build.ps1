$base = $PSScriptRoot
Write-Output "Generating build/app.exe ..."

# -s Omit the symbol table and debug information.
# -w Omit the DWARF symbol table.
go build -ldflags "-w -s" -o "$base/build/app.exe" "$base/main.go"
Write-Output "Done"

Write-Output "Generating build/app-debug.exe ..."
go build -o "$base/build/app-debug.exe" "$base/main.go"
Write-Output "Done"

Write-Output "Coping the application resources..."
Copy-Item "$base/app.config.toml" "$base/build/"
Write-Output "Done"
