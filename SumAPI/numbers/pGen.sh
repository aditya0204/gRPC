#!/bin/bash
protoc  numbers.proto --go_out=plugins=grpc:.

