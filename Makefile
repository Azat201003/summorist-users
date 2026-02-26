USERS_PORT=5252
export USERS_PORT

# Extract the optional argument after "test" from the command line
EXTRA_ARGS := $(filter-out test,$(MAKECMDGOALS))
ifeq ($(EXTRA_ARGS),)
TEST_SUBDIR :=
else
TEST_SUBDIR := $(firstword $(EXTRA_ARGS))
endif

# Treat "all" as no specific subdirectory (i.e., run all tests)
ifeq ($(TEST_SUBDIR),all)
TEST_SUBDIR :=
endif

# Create dummy targets for any extra arguments to avoid "No rule to make target" errors
$(EXTRA_ARGS):
	@# do nothing

test:
	@echo "Generating RSA keys for tests..."
	@USERS_PRIVATE_KEY=$$(openssl genrsa 4096); \
	USERS_PUBLIC_KEY=$$(echo "$$USERS_PRIVATE_KEY" | openssl rsa -pubout); \
	export USERS_PRIVATE_KEY USERS_PUBLIC_KEY; \
	echo "$$USERS_PRIVATE_KEY"; \
	echo "$$USERS_PUBLIC_KEY"; \
	if [ -z "$(TEST_SUBDIR)" ]; then \
		go test -v ./tests/...; \
	else \
		go test -v ./tests/$(TEST_SUBDIR)/...; \
	fi
