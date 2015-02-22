/// <reference path="../packages/React.TypeScript.DefinitelyTyped.0.1.12/react.d.ts" />

//import React = require('react');


interface Scope {
    greetingText: string;
}

class Controller {
    constructor($scope: Scope) {
        $scope.greetingText = "Helloddd from TypeScrssddsipt + AngularJS";
    }

}



