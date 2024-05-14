#!/bin/sh

function add_next_hop_route() {
  for (( i=0; i<20; i++ ))
    do
      vpn_ip=$(ip a | grep $1 | awk '{print $2}' | cut -d'/' -f1)
      if [ -z $vpn_ip ];then
        sleep 2
      else
        break
      fi
    done

    if [ -z $vpn_ip ];then
      exit 2
    fi

    route add $2 gw $vpn_ip

}

add_next_hop_route $1 $2
