#!/bin/bash

cd goserv/
gomake.sh
cd ../
stopContainers.sh
removeImages.sh
buildImages.sh
startContainers.sh