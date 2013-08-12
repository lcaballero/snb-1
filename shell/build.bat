/*
/app-location
	/public
		/js
			/controllers
				/admin.js
				/filters.js
			/angular.js
			/jquery.js
		/css
			/bootstrap.css
		/images
			/icons
				/pdf.png
				/tear-off.png
			/
	/views
		/pages
			/admin.html
		/partials
			/admin
				/admin-partial-1.html
			/shared
				/nav-tabs.html
	/config
		/redis.conf
		/web-server.conf
		/sql.conf
	/shell
		/start-redis.sh
		/build.sh
	/server
		/server.go

	server.exe
*/

rmdir /S /Q seed
mkdir seed
cd seed

mkdir public
mkdir public\js
mkdir public\css
mkdir public\images
mkdir public\images\icons
mkdir views
mkdir views\pages
mkdir views\partials
mkdir views\partials\admin
mkdir views\partials\shared
mkdir config
mkdir shell
mkdir server

dir