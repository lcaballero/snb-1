

directory="."
d="$directory/server/libs/"
cwd=$(pwd)
abs=""

echo "cwd: $cwd"
echo "d: $d"

while [ ! -d "$d" ]; do 
	directory="$directory/.."
	d="$directory/server/libs/"
	echo $directory
	echo $d
	cd $directory
	abs=$(pwd)
	echo "abs: $abs"
	cd $cwd
done
