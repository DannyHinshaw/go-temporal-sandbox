# go-temporal-sandbox

Sandbox for testing [Temporal](https://github.com/temporalio/temporal).

## Architecture

This is a simple mono-repo design geared to house multiple services and aims to:

1. Give a sandbox for working with multiple workers (pretend they aren't carbon copies of each other 😉)
2. Give a REST API layer for further testing inter-service communication as well as provide an easy way to test locally.

## Run

Run the stack:

```bash
docker-compose up -d
```
