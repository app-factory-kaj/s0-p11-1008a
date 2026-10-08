# Greeting

## Purpose

Lets an API caller get a JSON greeting for a given name via the service's
HTTP endpoint.

## User Stories

- F1.1 As an API caller, I send a name and get back a JSON greeting for it.
- F1.2 As an API caller, if I omit the name or send an empty one, I get a 400 error telling me the name is required.

## Decisions

- The greeting text is "Hello, {name}!".
- The response is JSON with the greeting in a `message` field: `{"message": "Hello, Ada!"}`.
- A missing or empty `name` is rejected with HTTP 400 and a JSON error body: `{"error": "name is required"}`.

 E2E marker p11-1008a.