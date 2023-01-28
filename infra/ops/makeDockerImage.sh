#!/bin/bash

cd ../goserv/
gomake.sh
cd ../ops/
stopContainers.sh
removeImages.sh
buildImages.sh
startContainers.sh