#!/bin/bash

grpcurl -plaintext -format text -d 'name: "Vincent"' localhost:8080 tutorial.HelloService.Hello
