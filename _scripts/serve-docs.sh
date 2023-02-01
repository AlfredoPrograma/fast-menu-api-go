#!/bin/sh

swagger generate spec -o swagger.json
swagger serve -F="$1" ./swagger.json