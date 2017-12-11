#Use this script this way: > . env.shcd ..

echo " "
echo "------------------- "
echo " Script de andres!"
echo "------------------- "
echo " "
echo " We will set GOPATH pointing to '/src' inside this project"
echo " Take into account that the default value was '' "
echo " "

if [ `type -p go` = "" ]; then
    export PATH=$PATH:/usr/local/go/bin
fi
export GOPATH=$PWD/pkg
export PATH=$PATH:$PWD/bin

go env

echo " "
echo "The end... GOPATH was set!"
echo " "