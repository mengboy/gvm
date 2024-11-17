#!/bin/bash

#if [ ! -d "~/.gvm/" ]; then
#  mkdir ~/.gvm/
#fi

version=$1
os=$2
arch=$3
defaultUrl="https://golang.org/dl/go"
localURL="https://studygolang.com/dl/golang/go"
suffix=".tar.gz"
connect="-"
point="."
downUrl=$defaultUrl$version$point$os$connect$arch$suffix

if [[ $4 == "local" ]]; then
  downUrl=$localURL$version$point$os$connect$arch$suffix
fi

cmd_exists() {
  if [ -x "$(command -v "$1")" ]; then
    return 0
  else
    return 1
  fi
}

print_error() {
  # Print output in red
  printf "\e[0;31m  [✖] $1 $2\e[0m\n"
}

print_success() {
  # Print output in green
  printf "\e[0;32m  [✔] $1\e[0m\n"
}

print_result() {
  [ $1 -eq 0 ] &&
    print_success "$2" ||
    print_error "$2"

  [ "$3" == "true" ] && [ $1 -ne 0 ] &&
    exit
}

execute() {
  $1 &>/dev/null
  print_result $? "${2:-$1}"
}

if ! cmd_exists "wget"; then
  if ! cmd_exists "curl"; then
    print_error "not found wget or curl "
    exit 1
  else
    execute "curl $downUrl -o $HOME/.gvm/go$version$point$os$connect$arch$suffix"
  fi
else
  execute "wget -P $HOME/.gvm/ $downUrl"
fi

execute "sudo tar -zxf $HOME/.gvm/go$version$point$os$connect$arch$suffix -C /usr/local"

echo "please add 'export PATH=\$PATH:/usr/local/go/bin' in your shell setting file. e.g.: .bashrc\.zshrc."
