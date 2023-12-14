#!/bin/bash


#function

add_user()
{
USER=$1
USER=$2

useradd -m -p $PASS $USER && echo "Successfully added user"

}

#Main Body to call function

add_user DipenR test@123
