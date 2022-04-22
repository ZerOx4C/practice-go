pushd `dirname $0`
go build -ldflags "-H windowsgui"
popd
