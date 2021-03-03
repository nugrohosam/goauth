#!/bin/bash

sudo apt install sshpass -y

sshpass -p $1 ssh -p $2 $3@$4

cd $5

git pull