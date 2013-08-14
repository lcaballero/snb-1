angular.module('phonecat', ['phonecatFilters'])
	.config(['$routeProvider', function( $routeProvider ) {
		$routeProvider
			.when("/phones", {
				templateUrl:'partials/phone-list.html',
				controller:PhoneListCtrl
			})
			.when("/phones/:phoneId", {
				templateUrl:'partials/phone-detail.html',
				controller:PhoneDetailCtrl
			})
			.otherwise({
				redirectTo:"/phones"
			})
	}])
/*
angular.module('myApp', [])
	.config(['$routeProvider', function($routeProvider){
		$routeProvider
			.when('/fullView', {
				templateUrl: 'partials/fullView',
				controller:FullView
			})
			.when('/picView', {
				templateUrl: 'partials/viewPic',
				controller:PicView
			})
			.otherwise({redirectTo:'/fullView'});
	}]);

angular.module('myAdmin', [])
	.config(['$routeProvider', function($routeProvider){
		$routeProvider
			.when('/users', {
				templateUrl: 'partials/usersPartial',
				controller:UsersCtrl
			})
			.when('/gourps', {
				templateUrl: 'partials/groupsPartial',
				controller:GroupsCtrl
			})
			.otherwise({redirectTo:'/'});
	}]);

angular.module('login', [])
	.config(['$routeProvider', function($routeProvider){
		$routeProvider
			.when('/login', {
				templateUrl: 'partials/login.html',
				controller:Login
			})
		}]);
*/

// ======================= Routes ======================= //

var AdminNavModule = angular.module('adminNavModule', [])
	.config(['$routeProvider', function($routeProvider){
		$routeProvider
			.when('/addUser', {
				templateUrl: '/views/partials/admin/createUserPartial.html',
				controller:'CreateUserCtrl'
			})
			.when('/findUser', {
				templateUrl: '/views/partials/admin/findUserPartial.html',
				controller:'FindUserCtrl'
			})
			.when('/findUser/:username', {
				templateUrl: '/views/partials/admin/findUserPartial.html',
				controller:'FindUserCtrl'
			});
	}]);

// ======================= Controllers ======================= //

AdminNavModule.controller('AdminMainNavCtrl', function($scope) {
	$scope.navItems = [
		{
			label:'Users',
			href:'/views/pages/admin/userAdmin.html',
			activeId:'users'
		},
		{
			label:'Groups',
			href:'/views/pages/admin/groupsAdmin.html',
			activeId:'groups'
		},
		{
			label: 'Games',
			href:'javascript:void(0);',
			activeId:'games'
		},
		{
			label: 'Boards',
			href:'javascript:void(0);',
			activeId:'boards'
		},
		{
			label: 'Criteria',
			href:'javascript:void(0);',
			activeId:'criteria'
		}
	];
});

AdminNavModule.controller('UserInfoNavCtrl', function($scope) {
	$scope.navItems = [
		{
			label:'Groups',
			href:'/views/pages/admin/groupsAdmin.html',
			activeId:'groups'
		},
		{
			label:'Games',
			href:'javascript:void(0);',
			activeId:'games'
		},
		{
			label:'Boards',
			href:'javascript:void(0);',
			activeId:'boards'
		},
		{
			label:'Tags',
			href:'javascript:void(0);',
			activeId:'tags'
		}
	];
});

AdminNavModule.controller('UserNavCtrl', function($scope, $location) {
	//$scope.activeClass = 'createUser';
});

AdminNavModule.controller('CreateUserCtrl', function($scope) {
	console.log('test');
	$scope.acitveClass = 'createUser';
});

AdminNavModule.controller('FindUserCtrl', function($scope, $routeParams, $http){
	console.log('FindUserCtrl');
	$scope.acitveClass = 'findUser';

	console.log('route params', $routeParams);
	if($routeParams && $routeParams.username) {
		$scope.userInfoDisplay = 'show-info';
	}
});

// ======================= Directives ======================= //

AdminNavModule.directive('tabNavDirective', function() {
	return {
		restrict: "A",
		scope: {
			selected: "@",
			navitems: "="
		},
		templateUrl: "/views/partials/admin/tabNav.html",
		link: function (scope, elem, attrs) {
			//console.log('scope', scope);
		}
	};
});

AdminNavModule.directive('usersNavDirective', function() {
	return {
		restrict: "A",
		scope: {
			selected: "@",
			navitems: "="
		},
		templateUrl: "/views/partials/admin/usersMenuPartial.html",
		link: function (scope, elem, attrs) {
			//console.log('scope', scope);
		}
	};
});


AdminNavModule.directive('createUserDirective', function() {
	return {
		restrict: "A",
		scope: { },
		templateUrl: "/views/partials/admin/createUserPartial.html",
		link: function (scope, elem, attrs) {
			//console.log('scope', scope);
		}
	};
});

