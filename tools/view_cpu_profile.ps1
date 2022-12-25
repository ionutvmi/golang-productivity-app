# Run the `app-debug.exe cpu-profile` and perform the needed actions
# Then run this script
# It expects the CPU profile to be located in the build directory

$base = Join-Path $PSScriptRoot ../

pushd $base/build

go tool pprof -http=:8080 ./app-debug.exe ./cpu.prof

popd
