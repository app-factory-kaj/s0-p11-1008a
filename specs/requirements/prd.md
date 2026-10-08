# Greeter

## Problem Statement

Teams building or testing other services often need a tiny, dependable HTTP
endpoint to greet a caller by name — useful as a smoke-test target, an example
service, or a building block to wire other things against. Standing one up
from scratch each time is needless repeated work.

## Solution

Greeter is a small Go HTTP service with a single endpoint: given a name, it
returns a JSON greeting. It follows this organization's standard Go service
conventions (project layout, JSON error shape, configuration via environment
variables) so it behaves predictably alongside other services.

## Actors

- API Caller — any HTTP client that calls the service; there is no sign-in and
no distinct roles. \[org default: no sign-in for a service with no user-facing
stories\]

## Features

- F1 [Greeting](features/F1-greeting.md)

## Product-wide

See [Product-wide](product-wide.md).

## Out of Scope

- No authentication or authorization — the endpoint is open.
- No user interface — this is an API-only service.
- No multi-language or localized greetings.
- No persistence — the service holds no data between requests.