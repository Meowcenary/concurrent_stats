#!/usr/bin/env bash

for ((i = 0 ; i < 100 ; i++)); do
	go test -bench=.
done
