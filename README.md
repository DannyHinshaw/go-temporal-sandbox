# go-temporal-sandbox

Sandbox for testing [Temporal](https://github.com/temporalio/temporal).

## Overview

This is a simple mono-repo design geared to house multiple services and aims to:

1. Give a sandbox for working with multiple workers (pretend they aren't carbon copies of each other 😉)
2. Give a REST API layer for further testing inter-service communication as well as provide an easy way to test locally.
3. Provide useful patterns and utilities for working with Temporal in a production environment.

## Run

Run formatter:

```bash
make fmt
```

Run the stack:

```bash
make run
```

Stop the stack:

```bash
make stop
```

## Endpoints

Health:

```bash
curl 'localhost:8080/health'
```

Worker A:

```bash
curl 'localhost:8080/workflow-a'
```

Worker B:

```bash
curl 'localhost:8080/workflow-b'
```