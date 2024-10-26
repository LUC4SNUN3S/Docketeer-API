#!/bin/bash

docker run -it --rm --name tarelai-webhook-receive \
--network tarelai_webhook_network \
--env-file $(pwd)/.env \
-v $(pwd):/go/src/tarelai-webhook-receive \
-p 8080:8080 \
golang:1.22.1