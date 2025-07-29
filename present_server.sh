#!/bin/bash

PORT=${1:-3999}

echo "🚀 Starting Go Course Presentation Server..."

go run present_server.go $PORT
