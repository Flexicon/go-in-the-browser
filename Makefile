build: build-wasm build-server

build-wasm: web/app.wasm

build-server:
	@ echo "🔧 Building server..."
	@ go build -o server ./cmds/server
	@ echo "✅ Done"

run: build
	@ echo "" && ./server

web/app.wasm: cmds/wasm/*.go
	@ echo "🔨 Building Wasm modules..."
	@ cp $(tinygo env GOROOT)/targets/wasm_exec.js ./web
	@ tinygo build -o web/tiny.wasm -target wasm ./cmds/wasm/main.go
	@ du -sh ./web/app.wasm
	@ echo "✅ Done"
