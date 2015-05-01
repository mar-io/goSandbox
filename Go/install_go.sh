#! /usr/bin/env bash
set -e
[ $# -eq 0 ] && set -- -d date > /dev/null

while getopts 'v:b:d:r:h' OPT; do
  case $OPT in
    v)  VERSION=$OPTARG;;
    b)  BIT=$OPTARG;;
    d)  VERSION=$(curl -s https://golang.org/doc/devel/release.html | grep 'Build version' | awk '{ print $3 }' | grep -o "[0-9.][0-9.][0-9.]..");;
    p)  PURGE=true;;
    h)  hlp="yes";;
    *)  unknown="yes";;
  esac
done

# usage
HELP="
    usage: $0 [ -v version -d -h ]
        -n --> Particular Go version < ex. 0.4.2
        -b --> Bit Version of Go < 32 or 64(default)
        -d --> Latest Stable Go version
        -p --> Deletes Go in its entirety from your system
        -h --> print this help screen
"

if [ "$hlp" = "yes" ]; then
  echo "$HELP"
  exit 0
fi

if [ "$PURGE" = true ] ; then
  read -r -p "Are you sure you want to delete go? This will delete EVERYTHING [y/N]"
  case $yn in
        [Yy]* ) rm -rf "$HOME/.go/"
                rm -rf "$HOME/go/"
                if [ "$SHELL" = "/bin/bash" ] ; then
                  sed -i '/# GoLang/d' "$HOME/.bashrc"
                  sed -i '/export GOROOT/d' "$HOME/.bashrc"
                  sed -i '/:$GOROOT/d' "$HOME/.bashrc"
                  sed -i '/export GOPATH/d' "$HOME/.bashrc"
                  sed -i '/:$GOPATH/d' "$HOME/.bashrc"
                  echo "Go purged!"
                elif [ "$SHELL" = "/bin/zsh" ] ; then
                  sed -i '/# GoLang/d' "$HOME/.zshrc"
                  sed -i '/export GOROOT/d' "$HOME/.zshrc"
                  sed -i '/:$GOROOT/d' "$HOME/.bashrc"
                  sed -i '/export GOPATH/d' "$HOME/.zshrc"
                  sed -i '/:$GOPATH/d' "$HOME/.zshrc"
                  echo "Go purged!"
                fi
                exit 0;;
        [Nn]* ) exit;;
        * ) echo "Please answer yes or no.";;
    esac
  exit
fi

if [ "$BIT" = '64' ] || [ "$BIT" = '32' ] ; then
    echo $BIT > /dev/null
elif [ -z "$BIT" ] ; then
    BIT='64'
else
  echo 'Not a valid bit version. Please select 32,64 or nothing at all.Exiting.'
  exit 1
fi

BINARY="go$VERSION.linux-amd$BIT.tar.gz"
LINK="https://storage.googleapis.com/golang/$BINARY"
TMP="/tmp"
GOGOGO="$HOME/.go/go-$VERSION-$BIT"

function goswitch {

if [ "$SHELL" = "/bin/bash" ] ; then
  if cat $HOME/.bashrc | grep "GOROOT" > /dev/null ; then
    OLD=$(cat $HOME/.bashrc | grep "GOROOT" | grep -o "[0-9.][0-9.][0-9.]..")
    sed -i "s@GOROOT=$HOME/.go/go-$OLD@GOROOT=$GOGOGO@g" $HOME/.bashrc
    echo "Type 'source ~/.bashrc' to start using go$VERSION"
  fi
elif [ "$SHELL" = "/bin/zsh" ]; then
  if cat $HOME/.zshrc | grep "GOROOT" > /dev/null ; then
    OLD=$(cat $HOME/.zshrc | grep "GOROOT" | grep -o "[0-9.][0-9.][0-9.]..")
    sed -i "s@GOROOT=$HOME/.go/go-$OLD@GOROOT=$GOGOGO@g" $HOME/.zshrc
    echo "Type 'source ~/.zshrc' to start using go$VERSION"
  fi
elif [ "$SHELL" != "/bin/bash" ] || [ "$SHELL" != "/bin/zsh" ] ; then
    echo "Not a valid shell. Use bash or zsh"
    exit 1
fi
}

function goshell {

if [ "$SHELL" = "/bin/bash" ] ; then
  if ! cat $HOME/.bashrc | grep "GOROOT" > /dev/null ; then
    {
    echo '# GoLang'
    echo "export GOROOT=$GOGOGO"
    echo 'export PATH=$PATH:$GOROOT/bin'
    echo 'export GOPATH=$HOME/go'
    echo 'export PATH=$PATH:$GOPATH/bin'
    } >> "$HOME/.bashrc"
    echo "Type 'source ~/.bashrc' to start using go$VERSION"
    exit 0
  fi
elif [ "$SHELL" = "/bin/zsh" ]; then
  if ! cat $HOME/.zshrc | grep "GOROOT" > /dev/null ; then
 {
    echo '# GoLang'
    echo "export GOROOT=$GOGOGO"
    echo 'export PATH=$PATH:$GOROOT/bin'
    echo 'export GOPATH=$HOME/go'
    echo 'export PATH=$PATH:$GOPATH/bin'
    } >> "$HOME/.zshrc"
    echo "Type 'source ~/.zshrc' to start using go$VERSION"
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
if [ $? -ne 0 ]; then
    echo "Download failed! Are you sure you are using a valid version? Exiting..."
    exit 1
fi
tar -C $HOME -xzf $TMP/$BINARY
rm -rf $TMP/$BINARY
mkdir -p $HOME/go
mkdir -p $HOME/.go
mv $HOME/go $HOME/.go/go-$VERSION
echo "Go version $VERSION installed!"
goshell
goswitch