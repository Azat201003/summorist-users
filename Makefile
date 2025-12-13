
all: run

run: run-develop

run-develop:
	bash -c 'set -a && source secrets.env && source config.develop.env && set +a && go run ./cmd/'

test: test-develop

test-develop:
	bash -c 'set -a && source secrets.env && source config.develop.env && set +a && go test -v ./tests/...'

test-deploy:
	bash -c 'set -a && source secrets.env && source config.deploy.env && set +a && docker compose -f docker-compose.yml up -d --build && docker compose exec app go test ./tests/...'

