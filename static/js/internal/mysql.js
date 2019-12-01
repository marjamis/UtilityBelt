angular.module('utilitybelt_mysql', [])

.config(function($stateProvider, $urlRouterProvider) {
  $urlRouterProvider.otherwise("/mysql");

  $stateProvider.state('mysql', {
    url: '/mysql',
    templateUrl: './static/templates/partials/mysql.html',
    controller: 'mysqlContr'
  });
})

.controller('mysqlContr', ['$scope', '$http', function ($scope, $http){
}]);
