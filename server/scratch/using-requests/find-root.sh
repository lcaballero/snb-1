function findAppRoot()
{

	directory="."
	app_root="$directory/server/libs/"
	cwd=$(pwd)
	abs=""

	while [ ! -d "$app_root" ]; do 
		directory="$directory/.."
		app_root="$directory/server/libs/"
		cd $directory
		abs=$(pwd)
		cd $cwd
	done

	echo "${directory}"
}

a="$(findAppRoot)"

cd "$a"
