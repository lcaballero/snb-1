# Set the GOPATH so that we can reference code in the current
# project without having to update the GOPATH for specifically
# to where we've put the project directory

# Save the current GOPATH to later restore it to this value.
# If we continually add to the GOPATH we'd wind up adding PWD
# every time we ran this script -- which is messy.
TEMP_GOPATH=GOPATH
GOPATH=$PWD/server/libs/:$GOPATH

# Actual 'build' command that we want to run
go build -a -o web-server.exe server/server.go

# Restoring the old GOPATH
GOPATH=TEMP_GOPATH
