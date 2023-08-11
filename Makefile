.DEFAULT_GOAL := build

build:
	yass queen

install:
	scripts/install.sh

dev:
	scripts/dev.sh

deploy:
	scripts/deploy.sh
