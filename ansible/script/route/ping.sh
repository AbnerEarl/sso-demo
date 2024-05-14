#!/bin/sh

function check_net() {
  receive=$(ping -n -c 5 $1 | grep received | awk -F"," '{print $2}' | awk '{print $1}')
  if [ $receive -gt 2 ]; then
    exit 0
  else
    exit 2
  fi
}

check_net $1
