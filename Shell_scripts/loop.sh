#!/bin/bash

echo "Enter number:"
read n
sleep 0.5
echo "printing $n numbers..."
sleep 0.5

for ((i=1;i<=n;i++))
do
echo "$i"
done
