build:
	@go build -ldflags "-s -w"

buildall:
	@echo "start building"
	@mkdir -p dist/win/amd64
	@mkdir -p dist/darwin/amd64
	@mkdir -p dist/darwin/amd64
	@mkdir -p dist/linux/amd64
	@mkdir -p dist/linux/arm64

	@GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/win/amd64/ttime

	@GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/darwin/amd64/ttime
	@GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o dist/darwin/arm64/ttime

	@GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/darwin/amd64/ttime
	@GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o dist/darwin/arm64/ttime
	@echo "done"

clean:
	@rm -r dist
	@rm ttime
	@echo "cleaning"
