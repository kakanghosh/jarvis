platforms := linux darwin_amd64 darwin_arm64

all: ${platforms}
	echo '####### All build completed! ${platforms} ########'

linux:
	echo '>>>>>> Building for linux <<<<<<<<<'
	go build -o jarvis-linux

darwin_amd64:
	echo '>>>>>> Building for darwin amd64 <<<<<<<<<'
	env GOOS=darwin GOARCH=amd64 go build -o jarvis-darwin-amd64

darwin_arm64:
	echo '>>>>>> Building for darwin arm64 <<<<<<<<<'
	env GOOS=darwin GOARCH=arm64 go build -o jarvis-darwin-arm64

clean:
	rm jarvis-*