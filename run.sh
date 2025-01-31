#!/usr/bin/env bash
make init                      # Initializes 
make                           # Builds the binary
docker compose up --build -d   # Runs the docker to start the API and database.