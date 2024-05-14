#!/bin/bash

netPrefix=$1
ip rule show | grep "${netPrefix}." | awk '{print $NF}'
