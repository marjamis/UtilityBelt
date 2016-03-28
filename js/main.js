var app = angular.module('servicetesting', []);

app.controller('specifics', function($scope, $http) {

  $scope.redisDisplay = function () { $http.get("/redis?action=display").success(function(data, status, headers, config) {
    $scope.redisData = data.RedisItems; 
  });}

  $scope.removeRedisItem = function(key) {
    $http.get("/redis?action=del&key="+key).success(function(data, status, headers, config) { });
    $scope.redisDisplay();
  };

  $scope.addRedisItem = function(key, value) {
    $http.get("/redis?action=add&key="+key+"&value="+value).success(function(data, status, headers, config) { });
    $scope.newdata.key = ""
    $scope.newdata.value = ""
    $scope.redisDisplay();
  };

  $scope.redisDisplay();
});
