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
	console.log('UserInfoNavCtrl', $scope);
	//$scope.navItems = [];
	$scope.init = function(name) {
		//$scope.username = name;
		alert('scope' + $scope.username, name);
		$scope.navItems = [
			{
				label:'Groups',
				href:['#/findUser/', $scope.username,'/groups/'].join(''),
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
	}

	
});

AdminNavModule.controller('UserNavCtrl', function($scope, $location) {
	//$scope.activeClass = 'createUser';
});

AdminNavModule.controller('CreateUserCtrl', function($scope) {
	console.log('test');
	$scope.acitveClass = 'createUser';
});

AdminNavModule.controller('FindUserCtrl', function($scope, $routeParams, $http){
	
	//$scope.acitveClass = 'findUser';

	$scope.userInfoDisplay = "";
	
	var testData = {
		'lucas':{
			'name':'Lucas',
			'groups':[
				{
					'name':'Digital Lotion',
					'id':100
				},
				{
					'name':'LoL',
					'id':101
				}
			],
			'games':[],
			'boards':[]
		},
		'ryan':{
			'name': 'Crunch time!',
			'groups':[
				{
					'name':'Digital Lotion',
					'id':100
				},
				{
					'name':'Raiders',
					'id':102
				}
			],
			'games':[],
			'boards':[]
		}
	};

	if($routeParams && $routeParams.username) {
		
		// $.get('/api/getUser')
		// 	.done(function(r){
		// 		console.log('done', r);
		// 		console.log('scope', $scope);
		// 		$scope.userInfoDisplay = 'show-info';
		// 		$scope.searchInput = $routeParams.username;
		// 	})
		// 	.fail(function(err){
		// 		console.log('fail', err);
		// 	});
		$http.get('/api/getUser').success(function(data) {
			//console.log('getUser', data);
			$scope.userInfoDisplay = 'show-info';
			$scope.searchInput = $routeParams.username;
			var myData = testData[$routeParams.username];

			if(myData) {
				$scope.userInfoDisplay = 'show-info';
				$scope.selectedData = myData;
				// $scope.userName = myData.name;
				// $scope.groups = myData.groups;
				// $scope.games = myData.games;
				// $scope.groups = myData.boards;

			} else {
				$scope.userInfoDisplay = 'show-no-info';
			}
		});
	}

});

// ======================= Directives ======================= //

if (!String.prototype.supplant) {
    String.prototype.supplant = function (o) {
        return this.replace(/{([^{}]*)}/g,
            function (a, b) {
                var r = o[b];
                return typeof r === 'string' || typeof r === 'number' ? r : a;
            }
        );
    };
}

AdminNavModule.directive('tabNavDirective', function() {
	return {
		restrict: "A",
		scope: {
			selected: "@",
			navitems: "=",
			supplant: "="
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

