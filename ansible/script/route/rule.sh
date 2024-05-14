#!/bin/bash

netPrefix=$1
ips=($(ip rule show | grep "${netPrefix}." | awk '{print $(NF-2)}'))
tables=($(ip rule show | grep "${netPrefix}." | awk '{print $NF}'))

for ((i = 0; i < ${#tables[@]}; i++)); do
  nic=$(ip route show table ${tables[i]} | awk '{print $NF}')
  if [ ! -z $nic ]; then
    echo ${tables[i]},${nic},${ips[i]}
  fi
done
