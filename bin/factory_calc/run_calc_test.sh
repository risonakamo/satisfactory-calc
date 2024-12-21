set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

go run factory_calc.go computer "Caterium Computer"