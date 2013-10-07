# Set the GOPATH so that we can reference code in the current
# project without having to update the GOPATH for specifically
# to where we've put the project directory

# Save the current GOPATH to later restore it to this value.
# If we continually add to the GOPATH we'd wind up adding PWD
# every time we ran this script -- which is messy.
TEMP_GOPATH=GOPATH
GOPATH=$PWD/server/libs/

# Actual 'build' command that we want to run
# go build -a -o web-server.exe server/server.go

go test -i -parallel=1 \
	./server/libs/src/rt_config \
	./server/libs/src/snap_sql \
	./server/libs/src/models \
	./server/libs/src/sql_utils \
	./server/libs/src/sql_utils/caching

go test -parallel=1 \
	./server/libs/src/rt_config \
	./server/libs/src/snap_sql \
	./server/libs/src/models \
	./server/libs/src/sql_utils \
	./server/libs/src/sql_utils/caching

# Restoring the old GOPATH
GOPATH=TEMP_GOPATH

