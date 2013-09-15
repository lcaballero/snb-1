;(function($) {

	var tabs = [
		{
			href:"#",
			text:"Search Users",
			name:"search-users"
		},
		{
			href:"#",
			text:"Search Groups",
			name:"search-groups"
		}
	];


	function main() {
		console.log("here");

		$("#promise-ajax-base").promiseJade({
			url:"/views/base-page.Jade",
			data: { tabs: tabs }			
		})
		.done(function() {

			return $("nav.top").delegate(
				"li",
				"click",
				function(ev) {
					ev.preventDefault();
					console.log("show search inputs");
				});
		});

		$("#search-users").promiseJade({
			url:"/views/users/search-users.jade",
			data:{}
		})
		.done(
			function() {
				$("#search-users").delegate(
					"button",
					"click",
					function(ev) {
						ev.preventDefault();
						console.log("running a search (username and email)");
						
					})
		})
	};

	window.requestUsers = function() {
		console.log("requesting users");
		
		$.ajax({
			url:"/app/get-users",
			type:"GET",
			data:{},
			success:function(results) {
				console.log("requestUsers", results)
				$("#users-table").promiseJade({
					url:"/views/users/users-table.jade",
					data:{users:results}
				})
			}
		})
	};

	main();
	window.requestUsers();

})(jQuery);
