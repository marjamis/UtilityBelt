angular.module('utilitybelt_redis', [])

.config(function($stateProvider, $urlRouterProvider) {
  $urlRouterProvider.otherwise("/redis");

  $stateProvider.state('redis', {
    url: '/redis',
    templateUrl: './static/templates/partials/redis.html',
    controller: 'redisContr'
  });
})

.controller('redisContr', ['$scope', '$http', function ($scope, $http){
  $scope.redisDisplay = function () { $http.get("/redis?action=display").success(function(data, status, headers, config) {
    $scope.redisData = data.RedisItems;
  });}

  $scope.removeRedisItem = function(key) {
    $http.get("/redis?action=del&key="+key).success(function(data, status, headers, config) {
      $scope.redisDisplay();
   });
  };

  $scope.addRedisItem = function(key, value) {
//    if ( angular.isUndefined($scope.newdata) != false && $scope.newdata.key != "" && $scope.newdata.value != "" ) {
      $http.get("/redis?action=add&key="+key+"&value="+value).success(function(data, status, headers, config) {
        $scope.newdata.key = ""
        $scope.newdata.value = ""
        $scope.redisDisplay();
      });
 //   }
  };

  $scope.redisDisplay();
}]);
