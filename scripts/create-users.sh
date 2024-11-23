#!/bin/bash

curl -X POST http://localhost:8080/api/v1/users/sync \
  -H "Content-Type: application/json" \
  -d '{"id": "user1", "name": "John Doe", "email": "john.doe@example.com"}'

curl -X POST http://localhost:8080/api/v1/users/sync \
  -H "Content-Type: application/json" \
  -d '{"id": "user2", "name": "Jane Smith", "email": "jane.smith@example.com"}'

curl -X POST http://localhost:8080/api/v1/users/sync \
  -H "Content-Type: application/json" \
  -d '{"id": "user3", "name": "Bob Johnson", "email": "bob.johnson@example.com"}'

curl -X POST http://localhost:8080/api/v1/users/sync \
  -H "Content-Type: application/json" \
  -d '{"id": "user4", "name": "Alice Brown", "email": "alice.brown@example.com"}' 