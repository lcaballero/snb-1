# First Draft of App Repo

## Building

Run `shell\build.bat` or `shell\build.sh`

Then run web-server.exe found at the root of the directory.

````
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

## Dependencies

So far the project has been written using these tools, and they are listed
here just for completeness, and are not strict dependencies.

Tools:
1. Go (Golang distribution).
1. Redis
1. Sublime Text 3
1. Postgres Sql Server
1. Postgres Admin

Git Packages:
1. Postgres Sql Driver: gitchub.com/lib/pq
1. Redis Sql: github.com/garyburd/redigo/redis

go get github.com/bmizerany/pq
go get github.com/nu7hatch/gouuid
go get github.com/russross/blackfriday
