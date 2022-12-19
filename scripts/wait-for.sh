#!/bin/bash

echo "Waiting service to launch on 3001..."

while ! nc -z localhost 3001; do
  sleep 0.1 # wait for 1/10 of the second before check again
done

echo "Service launched"