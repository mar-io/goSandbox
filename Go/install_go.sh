#! /usr/bin/env bash
[ $# -eq 0 ] && set -- -d date > /dev/null

while getopts 'v:d:h' OPT; do
  case $OPT in
    v)  VERSION=$OPTARG;;
    d)  VERSION=$(curl -s https://golang.org/doc/devel/release.html | grep 'Build version' | awk '{ print $3 }' | grep -o "[0-9.][0-9.][0-9.]..");;
    h)  hlp="yes";;
    *)  unknown="yes";;
  esac
done

# usage
HELP="
    usage: $0 [ -v version -d --h ]
        -n --> Particular Go version < ex. 0.4.2
        -d --> Latest Stable Go version
        -h --> print this help screen
"

if [ "$hlp" = "yes" ]; then
  echo "$HELP"
  exit 0
fi

BINARY="go$VERSION.linux-amd64.tar.gz"
LINK="https://storage.googleapis.com/golang/$BINARY"
TMP=/tmp
GOGOGO="$HOME/.go"

curl -Oo $TMP $LINK
mkdir -p $GOGOGO
tar -C $GOGOGO -xzf $TMP/$BINARY
rm -rf $TMP/$BINARY

if [ "$SHELL" = "/bin/bash" ] ; then
  echo 'export GOROOT=$HOME/.go' >> "$HOME/.bashrc"
  echo 'export PATH=$PATH:$GOROOT/bin' >> "$HOME/.bashrc"
  source "$HOME/.bashrc"
elif [ "$SHELL" = "/bin/zsh" ]; then
  echo 'export GOROOT=$HOME/.go' >> "$HOME/.zshrc"
  echo 'export PATH=$PATH:$GOROOT/bin' >> "$HOME/.zshrc"
  source "$HOME/.zshrc"
else 
  echo "Not a valid shell. Use bash or zsh"
fi






