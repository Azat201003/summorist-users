
all: run

run: run-develop

run-develop:
	bash -c 'set -a && source secrets.env && source config.env.develop && set +a && go run ./cmd/'

run-deploy:
	bash -c 'set -a && source secrets.env && source config.env.deploy && set +a && go run ./cmd/'

test: test-develop

test-develop:
	bash -c 'set -a && source secrets.env && source config.env.develop && set +a && go test -v ./tests/...'

test-deploy:
	bash -c 'set -a && source secrets.env && source config.env.deploy && set +a && go test -v ./tests/...'

