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

var AdminNavModule = angular.module('adminNavModule', ['ui.state']);
AdminNavModule.config(['$stateProvider', '$routeProvider', '$urlRouterProvider',
    	function ($stateProvider, $routeProvider, $urlRouterProvider){
			//$urlRouterProvider.otherwise('/addUser');

			$stateProvider
				.state('addUser', {
					url:'/addUser',
					controller:'CreateUserCtrl',
					templateUrl:'/views/partials/admin/createUserPartial.html'
				})
				.state('findUser', {
					url:'/findUser',
					//controller:['$scope', '$state', '$stateParams', '$http', FindUserCtrl],
					templateUrl: '/views/partials/admin/findUserPartial.html'
				})
					.state('findUser.username', {
						url:'/:username',
						controller:['$scope', '$state', '$stateParams', '$http', FindUserCtrl],
						templateUrl: '/views/partials/admin/userInfoPartial.html'
					})
						.state('findUser.username.groups', {
							url:'/groups',
							//controller:['$scope', '$state', '$http', FindUserCtrl],
							templateUrl: '/views/partials/admin/userGroupPartial.html'
						});
		}
	]);


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

AdminNavModule.controller('UserInfoNavCtrl', function($scope, $location) {
	$scope.navItems = [
		{
			label:'Groups',
			href:['#', $location.$$path, '/groups'].join(''),
			activeId:'groups'
		},
		{
			label:'Games',
			href:['#', $location.$$path, '/games'].join(''),
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

//AdminNavModule.controller('FindUserCtrl', function($scope, $state){
var FindUserCtrl = function($scope, $state, $stateParams, $http){	
	console.log("stateParams", $stateParams);
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
	
	if($state && $state.params) {
		
		var username = $state.params.username;

		$http.get('/api/getUser').success(function(data) {
			//console.log('getUser', data);
			$scope.userInfoDisplay = 'show-info';
			$scope.searchInput = username;
			var myData = testData[username];

			if(myData) {
				$scope.userInfoDisplay = 'show-info';
				$scope.selectedData = myData;
			} else {
				$scope.userInfoDisplay = 'show-no-info';
			}
		});
	}

};

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

