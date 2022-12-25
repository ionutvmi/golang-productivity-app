# tools used
# go install github.com/kisielk/godepgraph@latest
# scoop install graphviz

$base = Join-Path $PSScriptRoot ../

pushd $pase

godepgraph -s -p github.com,lukechampine.com,modernc.org main.go | dot -Tpng -o "$base/build/local-project-graph.png"
godepgraph -s main.go | dot -Tpng -o "$base/build/project-with-extenral-deps-graph.png"

popd
