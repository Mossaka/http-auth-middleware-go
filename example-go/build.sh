tinygo build -o main_2023_10_18.wasm -target=wasi main_2023_10_18.go
wasm-tools component embed wit/2023_10_18 main_2023_10_18.wasm > main_2023_10_18.embed.wasm
wasm-tools component new main_2023_10_18.embed.wasm -o example.wasm --adapt wasi_snapshot_preview1.reactor.2023_10_18.wasm
wasm-tools compose ../github-oauth/target/wasm32-wasi/release/github_oauth.wasm -d example.wasm -o service.wasm