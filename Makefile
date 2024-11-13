.PHONY: build clean deploy

build: clean
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bootstrap -tags lambda.norpc main.go && zip app.zip bootstrap

clean:
	rm -rfv ./bin

deploy-dev: build
	sls deploy --verbose --stage dev --force

deploy-staging: build
	sls deploy --verbose --stage staging --force

deploy-prod: build
	sls deploy --verbose --stage prod
