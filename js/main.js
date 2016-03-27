var app = angular.module('servicetesting', []);

app.controller('specifics', function($scope, $http) {

  $scope.redisDisplay = function () { $http.get("/redis?action=display").success(function(data, status, headers, config) {
    $scope.redisData = data.RedisItems; 
  });}

  $scope.removeRedisItem = function(key) {
    console.log(key);
    $scope.redisDisplay();
  };

  $scope.addRedisItem = function() {

  };

  //Initlisation - check if this is standard and OK way to do this.
  $scope.redisDisplay();
});
