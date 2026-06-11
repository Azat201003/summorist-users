#!/bin/bash
export USERS_PRIVATE_KEY=$(openssl genrsa 4096 2>/dev/null)
export USERS_PUBLIC_KEY=$(openssl rsa -pubout 2>/dev/null <<<"$USERS_PRIVATE_KEY")
