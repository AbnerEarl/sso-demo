#!/bin/sh

function check_net() {
  export all_proxy=127.0.0.1:$2
  ping -n -c 3 $1 | grep "rtt"
}

check_net $1 $2
