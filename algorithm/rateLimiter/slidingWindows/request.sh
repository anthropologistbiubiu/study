#!/bin/bash

for i in {1..3}; do
    curl -X POST http://localhost:8080/rate
done