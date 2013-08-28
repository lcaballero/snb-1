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
							//activeClass:'groups',
							controller:['$scope', '$stateParams', function($scope, $stateParams) {
								$scope.activeMenu = 'groups';
							}],
							templateUrl: '/views/partials/admin/userGroupPartial.html'
						})
						.state('findUser.username.games', {
							url:'/games',
							//activeClass:'games',
							templateUrl:'/views/partials/admin/userGamesPartial.html',
							//controller: function($stateParams){}
							controller:['$scope', '$state', function($scope, $state) {
								$scope.activeMenu = 'games';
							}]
						});
		}
	]);


// ======================= Controllers ======================= //

AdminNavModule.controller('UserNavCtrl', function($scope, $state, $location) {
	//$scope.activeClass = 'createUser';
	//console.log('UserNavCtrl', $scope, $state);
});

AdminNavModule.controller('CreateUserCtrl', function($scope, $http) {

	$scope.submitUser = function() {
		console.log('scope', $scope);
		var email = $scope.email;
		var pw = $scope.pw;
		var confirm = $scope.confirm;
		var info = {email:email, pw:pw};
		console.log('info', {data:info});
		if(pw == confirm) {
			$http({
				url: '/api/addUser/',
				method: "POST",
				params: info
			})
			//$http.post('/api/addUser', info)
				.success(function(res) {
					console.log('addUser data', res);
				})
				.error(function(err){
					alert('error');
				});
		}
		// TODO: else show error message
	}
});

//AdminNavModule.controller('FindUserCtrl', function($scope, $state){
var FindUserCtrl = function($scope, $state, $stateParams, $http){	
	//console.log("FindUserCtrl stateParams", $stateParams, $state);
	//console.log('FindUserCtrl scope', $scope);

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
			'games':[
				{
					name:'Breweries of Boulder',
					id:'g100'
				},
				{
					name:'Hipster Hippies',
					id:'g101'
				}
			],
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
			'games':[
				{
					name:'Breweries of Boulder',
					id:'g100'
				},
				{
					name:'Hipster Hippies',
					id:'g101'
				}
			],
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

AdminNavModule.directive('mainNavDirective', function() {
	return {
		restrict: "A",
		scope: {
			selected: "@"
		},
		templateUrl: "/views/partials/admin/tabNav.html",
		link: function (scope, elem, attrs) {
			//console.log('scope', scope);
		},
		controller: function($scope){

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

			$scope.navitems = $scope.navItems;
			console.log('mainNavDirective', $scope);
		}
	};
});

AdminNavModule.directive('findNavDirective', function() {
	
	return {
		restrict: "A",
		scope: {
			selected: "@"
		},
		templateUrl: "/views/partials/admin/tabNav.html",
		link: function (scope, elem, attrs) {
			//console.log('scope', scope);
		},
		controller: function($scope, $stateParams){
			
			$scope.navItems = [
				{
					label:'Groups',
					href:['#/findUser/', $stateParams.username, '/groups'].join(''),
					activeId:'groups'
				},
				{
					label:'Games',
					href:['#/findUser/', $stateParams.username, '/games'].join(''),
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
			
			$scope.navitems = $scope.navItems;
		}
	};
});

AdminNavModule.directive('usersNavDirective', function($location) {

	return {
		restrict: "A",
		scope: true,
		templateUrl: "/views/partials/admin/usersMenuPartial.html",
		link: function (scope, elem, attrs) {
			console.log('usersNavDirective', scope);

			scope.$on( '$stateChangeSuccess', function () {
				var menuEl = $(elem).find('a[href="#'+$location.path()+'"]');
				$(elem).find('li.active').removeClass('active');

				if(menuEl) menuEl.parent().addClass('active');
			});
		}
	};
});


AdminNavModule.directive('createUserDirective', function() {
	alert('t')
	return {
		restrict: "A",
		scope: { },
		templateUrl: "/views/partials/admin/createUserPartial.html",
		link: function (scope, elem, attrs) {
			console.log('scope', scope);

			elem.bind('click', function(e){
				alert('click');
			});
		}
	};
});

