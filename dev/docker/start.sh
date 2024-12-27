#!/bin/bash

exec reflex -r '(.go$|.go\.mod|.env$)' --decoration=none -s go run "$@"
