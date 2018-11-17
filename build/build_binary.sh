#!/bin/bash

dep ensure
scriptdir=`dirname $0`
cd ${scriptdir}
scriptdir=`pwd`

mkdir -p ${scriptdir}/../bin
cd ${scriptdir}/../source/cmd/sm-server
LDFLAGS="-ldflags='-w -s'"

cmd="go build $LDFLAGS"
eval $cmd
mv ${scriptdir}/../source/cmd/sm-server/sm-server ${scriptdir}/../bin
upx ${scriptdir}/../bin/sm-server
