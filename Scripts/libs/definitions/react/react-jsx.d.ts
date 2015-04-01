 /// <reference path="react-addons.d.ts"/>
declare module "react/addons" {
		function jsx(jsx?: string): ReactElement<any>;
		function __spread(...args: any[]): any; // for JSX Spread Attributes
		function jsxFile(filename: string): ReactElement<any>;
}
// This file is intended to be used with the DefinitelyTyped React addons definition (react-addons.d.ts)
