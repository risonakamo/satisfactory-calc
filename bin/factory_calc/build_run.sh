set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

go build -o ../../factory_calc.exe factory_calc.go
cd ../..
./factory_calc.exe $@