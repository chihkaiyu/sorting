#!/bin/bash

root=`pwd`
folders=("bubble" "insertion" "selection" "heap")

for f in ${folders[@]};
do
    echo $f
    python $root/$f/main.py
    go run $root/$f/main.go
done