function Ed2kCtrl($scope, $log) {
  $scope.convert = function() {
    $scope.result = $scope.orgInput.split("%7C").join("|");
  }
}
Ed2kCtrl.$inject = ["$scope", "$log"];

