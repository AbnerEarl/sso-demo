#!/bin/sh

destIP="8.8.8.8"
function check_net() {
  proxy_ip=$(ifconfig | grep $1 -A 2 | grep netmask | awk '{print $2}')
  if [ -z $proxy_ip ]; then
    exit 3
  fi
  ip route add $destIP via $proxy_ip dev $1
  receive=$(ping -n -c 5 $destIP | grep received | awk -F"," '{print $2}' | awk '{print $1}')
  ip route del $destIP via $proxy_ip dev $1
  if [ $receive -gt 2 ]; then
    exit 0
  else
    exit 2
  fi
}

check_net $1
