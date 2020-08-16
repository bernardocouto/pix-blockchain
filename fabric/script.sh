#!/bin/bash

set -e

peer channel create -c pix -f pix.tx -o orderer:7050

peer channel join -b pix.block

sleep 600000

exit 0
