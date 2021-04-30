build:
	@go build -ldflags "-s -w"

install: build
	sudo mv ttime /usr/local/bin

buildall:
	@echo "creating directories"
	@mkdir -p dist/win/amd64
	@mkdir -p dist/darwin/amd64
	@mkdir -p dist/darwin/amd64
	@mkdir -p dist/linux/amd64
	@mkdir -p dist/linux/arm64

	@echo "start building"
	@GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/win/amd64/ttime

	@GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/darwin/amd64/ttime
	@GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o dist/darwin/arm64/ttime

	@GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/linux/amd64/ttime
	@GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o dist/linux/arm64/ttime

	@echo "zipping binaries"
	@zip dist/win/amd64/ttime-win-amd64 dist/win/amd64/ttime

	@zip dist/linux/arm64/ttime-linux-arm64 dist/linux/arm64/ttime
	@zip dist/linux/amd64/ttime-linux-amd64 dist/linux/amd64/ttime

	@zip dist/darwin/arm64/ttime-darwin-arm64 dist/darwin/arm64/ttime
	@zip dist/darwin/amd64/ttime-darwin-amd64 dist/darwin/amd64/ttime

	@echo "done"

clean:
	@if [ -d "$dist" ]; then rm -r dist; fi
	@if [ -f "$ttime" ]; then rm ttime; fi
	@echo "cleaning"
