function PhoneListCtrl($scope, $http) {

	$http.get('/app/phones/phones.json').success(
		function(data) {
			$scope.phones = data.slice(0, 5);
		})

	$scope.orderProp = 'age'
}

function PhoneDetailCtrl($scope, $routeParams, $http) {
	$http.get('phones/' + $routeParams.phoneId + '.json')
		.success(function(data) {
			$scope.phone = data;
			$scope.mainImageUrl =
				data.images && data.images.length
					? data.images[0]
					: data.imageUrl;
		})


	$scope.setImage = function(imageUrl) {
		$scope.mainImageUrl = imageUrl;
	}

}