#!/bin/sh

trap 'kill -9 $(jobs -p)' EXIT
echo -e $DATASTORE_KEY > datastore.key
$@