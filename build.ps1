Write-Output "Generating build/app.exe ..."

# -s Omit the symbol table and debug information.
# -w Omit the DWARF symbol table.
go build -ldflags "-w -s" -o build/app.exe main.go
Write-Output "Done"

Write-Output "Generating build/app-debug.exe ..."
go build -o build/app-debug.exe main.go
Write-Output "Done"

Write-Output "Coping the application resources..."
Copy-Item app.config.toml build/
Write-Output "Done"
