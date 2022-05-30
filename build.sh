
build="mac"

if [ "$#" -gt 1 ]
then
    build="unknown"
elif [ "$#" == 1 ]
then
    build=$1
fi

if [ $build == "mac" ]
then
    echo "$build"
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o docker-cases
elif [ $build == "linux" ]
then
    echo "$build"
    GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o docker-cases
elif [ $build == "windows" ]
then
    echo "$build"
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o docker-cases.exe
else
    echo "仅一个参数{平台}:linux,mac,windows,默认:mac"
    exit 0
fi
