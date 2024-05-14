#!/bin/sh

function compute_flow() {
  cat /proc/net/dev | grep $1 | awk '{for(i=1;i<=NF;i++){print $i}}'
}

compute_flow $1
