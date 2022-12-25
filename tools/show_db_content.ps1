# Requires the sqlite CLI
# scoop install sqlite
# 

$base = Join-Path $PSScriptRoot ../

sqlite3 "$base/data.db" .schema
sqlite3 "$base/data.db" .dump
