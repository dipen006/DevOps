#!/bin/bash

read i

echo "installing... $i "

sudo apt update -y
sudo apt -y install $i
