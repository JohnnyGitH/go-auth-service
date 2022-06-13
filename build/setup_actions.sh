#!/bin/bash
mkdir ~/.ssh
echo "$COMMON_KEY" > ~/.ssh/id_common
chmod 600 ~/.ssh/id_common
printf "Host githbut.com \n\t Hostname github.com \n\t StrictHostKeyChecking no \n\t User git \n\t IdentityFile ~/.ssh/id_common" >> ~/.ssh/config

git config --global user.email "jctotalportal@gmail.com"
git config --global user.name "Jonathan Cockrem"
git config --global url."ssh://gi@github.com/".insteadOF "https//github.com/"