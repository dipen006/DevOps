#!/bin/bash

echo "Enter the 1st number"
read number1

echo "Enter the 2nd number"
read number2

if [ $number1 -gt $number2 ]
then
 echo "$number1 is greater than $number2"
elif [ $number1 -lt $number2 ]
 then echo "$number1 is less than $number2"
else
 echo "$number1 is equal to $number2"
fi
