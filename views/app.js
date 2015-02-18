var Products;
(function (Products) {
    var Controller = (function () {
        function Controller($scope) {
            $scope.greetingText = "Hello from TypeScript + AngularJS";
        }
        return Controller;
    })();
    Products.Controller = Controller;
})(Products || (Products = {}));
