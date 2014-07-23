#!/bin/sh
TOOLS=`dirname $0`
DIRS=`find $TOOLS/../problems/*/*/ -iname Makefile -exec dirname {} \;`
for DIR in $DIRS; do
    make -C $DIR $1
done
