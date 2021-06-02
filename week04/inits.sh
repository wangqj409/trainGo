#!/bin/bash
# init app projects with dirs
DIRS="""\
app/cmd \
app/internal \
app/internal/model \
app/internal/dao \
app/internal/service \
app/internal/server \
app/pkg \
app/pkg/cache \
app/pkg/conf \
app/api \
app/config \
app/test \
app/CHANGELOG\
"""

for DIR in $DIRS
do
    if [ ! -e $DIR ]; then
        mkdir -p $DIR
    fi
    if [ ! -e $DIR/README ]; then
        touch  $DIR/README
    fi
done
