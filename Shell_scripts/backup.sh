#!/bin/bash

# Backup Directory

Backup_dir="/home/sh/Desktop/Devops"

# where to backup

Backup="/home/sh/Desktop/Backup"

# date of the backup

date=$(date + "%d-%b-%y")
cp -r $Backup_dir $Backup/$date

echo "Backup created in $Backup/$date"


