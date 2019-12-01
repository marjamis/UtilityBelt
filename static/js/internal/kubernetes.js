angular.module('utilitybelt_kubernetes', [])

.config(function($stateProvider, $urlRouterProvider) {
  $urlRouterProvider.otherwise("/kubernetes");

  $stateProvider.state('kubernetes', {
    url: '/kubernetes',
    templateUrl: './static/templates/partials/kubernetes.html',
    controller: 'kubernetesContr'
  });
})

.controller('kubernetesContr', ['$scope', '$http', function ($scope, $http){
  $scope.kubernetesDisplay = function () { $http.get("/kubernetes").success(function(data, status, headers, config) {
    $scope.data = data;
    console.log(data);
  }).error(function(data, status, headers, config){
    $scope.data = data;
    console.log(data);
  });}

  $scope.kubernetesDisplay();
}]);
