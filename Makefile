.PHONY: build debug

build-mewn:
	@go build -o build/mewn cmd/mewn/main.go

build:
	@rm -Rf ./build/Production/Loda.app/
	@mkdir -p ./build/Production/Loda.app/Contents/MacOS
	@mkdir -p ./build/Production/Loda.app/Contents/Resources
	# Icon files
	# @cp ./assets/appicon.icns ./build/Production/Loda.app/Contents/Resources/
	# Meta Data
	@cp ./meta/Info.plist ./build/Production/Loda.app/Contents/
	@echo "Building GUI"
	@cd ./ui && yarn build
	@echo "Building go-application"
	@mewn build -o build/Production/Loda.app/Contents/MacOS/Loda
	@echo "Completed."

debug:
	@rm -Rf ./build/Debug
	@echo "Building GUI"
	@cd ./ui && yarn build
	@echo "Building go-application"
	@mewn build -o build/Debug/Loda
	./build/Debug/Loda