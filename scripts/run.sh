#!/bin/bash
# Path: scripts/run.sh
{
	go run cmd/main.go
} || {
	printf "FAILED\n"
}