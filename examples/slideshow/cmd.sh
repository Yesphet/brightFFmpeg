#!/usr/bin/env bash

cd $(dirname $0)

ffmpeg -loop 1 -t 3 -i 1.jpg \
-loop 1 -t 3 -i 2.jpg \
-loop 1 -t 3 -i 3.jpg \
-loop 1 -t 4 -i 4.jpg \
-filter_complex '[0:v][1:v][2:v][3:v] concat=n=4:v=1:a=0,format=yuv420p[v]' \
-map "[v]" output.mp4