.PHONY: build debug

build:
	@rm -Rf ./build/Production/Skeleton.app/
	@mkdir -p ./build/Production/Skeleton.app/Contents/MacOS
	@mkdir -p ./build/Production/Skeleton.app/Contents/Resources
	# Icon files
	# @cp ./assets/appicon.icns ./build/Production/Skeleton.app/Contents/Resources/
	# Meta Data
	@cp ./meta/Info.plist ./build/Production/Skeleton.app/Contents/
	@echo "Building GUI"
	@cd ./ui && yarn build
	@echo "Building go-application"
	@mewn build -o build/Production/Skeleton.app/Contents/MacOS/Skeleton
	@echo "Completed."

debug:
	@rm -Rf ./build/Debug
	@echo "Building GUI"
	@cd ./ui && yarn build
	@echo "Building go-application"
	@mewn build -o build/Debug/Skeleton
	./build/Debug/Skeleton