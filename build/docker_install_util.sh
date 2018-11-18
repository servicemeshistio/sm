#!/bin/sh
PLAT=`uname -s`
rm -rf tools
rm -rf bin
mkdir tools

ash ./build/build_swagger.sh
ash ./build/build_binary.sh
