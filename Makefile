build: build-wasm build-server

build-wasm:
	@ echo "🔨 Building Wasm modules..."
	@ cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./web
	@ GOARCH=wasm GOOS=js go build -o web/app.wasm ./cmds/wasm
	@ du -sh ./web/app.wasm
	@ echo "✅ Done"

build-server:
	@ echo "🔧 Building server..."
	@ go build -o server ./cmds/server
	@ echo "✅ Done"

run: build
	@ echo "" && ./server
