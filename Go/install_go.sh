#! /usr/bin/env bash
set -e
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
    usage: $0 [ -v version -d -h ]
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
TMP="/tmp"
GOGOGO="$HOME/.go/go-$VERSION"

function goswitch {

if [ "$SHELL" = "/bin/bash" ] ; then
  if cat $HOME/.bashrc | grep "GOROOT" > /dev/null ; then
    OLD=$(cat $HOME/.bashrc | grep "GOROOT" | grep -o "[0-9.][0-9.][0-9.]..")
    sed -i "s@GOROOT=$HOME/.go/go-$OLD@GOROOT=$GOGOGO@g" $HOME/.bashrc
    source "$HOME/.bashrc"
    echo "Current shell using go version $VERSION"
    echo "Type '$source .bashrc' to start using go$VERSION"
  fi
elif [ "$SHELL" = "/bin/zsh" ]; then
  if cat $HOME/.zshrc | grep "GOROOT" > /dev/null ; then
    OLD=$(cat $HOME/.zshrc | grep "GOROOT" | grep -o "[0-9.][0-9.][0-9.]..")
    sed -i "s@GOROOT=$HOME/.go/go-$OLD@GOROOT=$GOGOGO@g" $HOME/.zshrc
    source "$HOME/.zshrc"
    echo "Current shell using go version $VERSION"
    echo "Type '$source .zshrc' to start using go$VERSION"
  fi
elif [ "$SHELL" != "/bin/bash" ] || [ "$SHELL" != "/bin/zsh" ] ; then
    echo "Not a valid shell. Use bash or zsh"
    exit 1
fi
}

function goshell {

if [ "$SHELL" = "/bin/bash" ] ; then
  if ! cat $HOME/.bashrc | grep "GOROOT" > /dev/null ; then
    echo "export GOROOT=$GOGOGO" >> "$HOME/.bashrc"
    echo 'export PATH=$PATH:$GOROOT/bin' >> "$HOME/.bashrc"
    source "$HOME/.bashrc"
    echo "Current shell using go version $VERSION"
    echo "Type '$source .bashrc' to start using go$VERSION"
    exit 0
  fi
elif [ "$SHELL" = "/bin/zsh" ]; then
  if ! cat $HOME/.zshrc | grep "GOROOT" > /dev/null ; then
    echo "export GOROOT=$GOGOGO" >> "$HOME/.zshrc"
    echo 'export PATH=$PATH:$GOROOT/bin' >> "$HOME/.zshrc"
    source "$HOME/.zshrc"
    echo "Current shell using go version $VERSION"
    echo "Type '$source .zshrc' to start using go$VERSION"
    exit 0
  fi
elif [ "$SHELL" != "/bin/bash" ] || [ "$SHELL" != "/bin/zsh" ] ; then
    echo "Not a valid shell. Use bash or zsh"
    exit 1
fi
}

if go version 2> /dev/null | grep $VERSION > /dev/null ; then
  echo "Go version $VERSION already installed."
  goshell
  goswitch
  exit 0
elif ls $HOME/.go 2> /dev/null | grep $VERSION > /dev/null ; then
  echo "Switching Go to version $VERSION"
  goshell
  goswitch
  exit 0
fi

cd $TMP
curl -O $LINK
tar -C $HOME -xzf $TMP/$BINARY
rm -rf $TMP/$BINARY
mkdir -p $HOME/.go
mv $HOME/go $HOME/.go/go-$VERSION
echo "Go version $VERSION installed!"
goshell
goswitch