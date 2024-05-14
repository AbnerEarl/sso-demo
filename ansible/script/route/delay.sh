#!/bin/sh

function check_net() {
  ping -n -c 3 $1 | grep "rtt"
}

check_net $1
