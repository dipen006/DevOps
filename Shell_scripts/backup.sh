#!/bin/bash

# source Directory

src_dir=/home/sh/Desktop/DevOps
tgt_dir=/home/sh/Desktop/Backups

curr_timestamp=$(date "+%Y-%m-%d-%H-%M-%S")
backup_file=$tgt_dir/$curr_timestamp.tgz

echo "Taking backup on $curr_timestamp"

tar czf $backup_file --absolute-names  $src_dir

echo "Backup Complete"

