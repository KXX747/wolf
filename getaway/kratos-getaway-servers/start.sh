#!/usr/bin/env bash
cd cmd

go build

 ./cmd -conf ./../configs/  -http.perf=tcp://0.0.0.0:38791