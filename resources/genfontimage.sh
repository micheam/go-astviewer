#! /usr/bin/env bash
set -eu

for ttf in `ls *.ttf`;do 
  fontimage $ttf --o fontimages/$ttf.png --fontname
done
