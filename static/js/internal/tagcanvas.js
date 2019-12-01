angular.module('utilitybelt_tagcanvas', [])

.config(function($stateProvider, $urlRouterProvider) {
  $urlRouterProvider.otherwise("/tagcanvas");

  $stateProvider.state('tagcanvas', {
    url: '/tagcanvas',
    templateUrl: './static/templates/partials/tagcanvas.html',
    controller: 'tagcanvasContr'
  });
})

.controller('tagcanvasContr', ['$scope', '$http', function ($scope, $http){
}]);
