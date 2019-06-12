#!/bin/bash

# GOLANG
GO="export PATH=$PATH:/usr/local/go/bin"
GOPATH="export GOPATH=$HOME/Development/go"

WORKDIR="$HOME/Development/go/src/pluralsight/go-dist-app/distributed"
COORDINATOR="$WORKDIR/coordinator/exec/main.go"
CONSUMER="$WORKDIR/sensors/sensor.go"

FLAGS1="-name=boiler_pressure_out -min=15 -max=15.5 -step=0.05 -freq=1"
FLAGS2="-name=turbine_pressure_out -min=0.9 -max=1.3 -step=0.05 -freq=1"
FLAGS3="-name=condensor_pressure_out -min=0.001 -max=0.002 -step=0.0001 -freq=1"
FLAGS4="-name=boiler_temp_out -min=590 -max=615 -step=1 -freq=1"
FLAGS5="-name=turbine_temp_out -min=100 -max=105 -step=1 -freq=1"
FLAGS6="-name=condensor_temp_out -min=80 -max=98 -step=1 -freq=1"

terminator --new-tab -e "$GO; $GOPATH; go run $COORDINATOR; exec bash"
terminator --new-tab -e "$GO; $GOPATH; go run $CONSUMER $FLAGS1; exec bash"
terminator --new-tab -e "$GO; $GOPATH; go run $CONSUMER $FLAGS2; exec bash"
terminator --new-tab -e "$GO; $GOPATH; go run $CONSUMER $FLAGS3; exec bash"
terminator --new-tab -e "$GO; $GOPATH; go run $CONSUMER $FLAGS4; exec bash"
terminator --new-tab -e "$GO; $GOPATH; go run $CONSUMER $FLAGS5; exec bash"
terminator --new-tab -e "$GO; $GOPATH; go run $CONSUMER $FLAGS6; exec bash"