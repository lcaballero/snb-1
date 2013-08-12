# First Draft of App Repo

## Building

Run `shell\build.bat` or `shell\build.sh`

Then run web-server.exe found at the root of the directory.

```
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
```