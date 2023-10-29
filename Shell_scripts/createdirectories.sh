#!/bin/bash

# take input

echo "Enter base name for Directory"
read name
echo "Enter the starting number"
read start_num
echo "Enter the ending number"
read end_num

# for loop to create directories

for (( i = start_num; i <= end_num; i++ ))
do

	mkdir $name$i

done
echo "Directories created ..."


