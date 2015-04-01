{
    "compilerOptions": {
        "target": "es5",
        "module": "commonjs",
        "declaration": false,
        "noImplicitAny": false,
        "removeComments": true,
        "noLib": false
    },
    "filesGlob": [
        "./**/*.ts",
        "!./node_modules/**/*.ts"
    ],
    "files": [
        "./modules/admin/*.ts",
        "./libs/**/*.d.ts"
        // "./linter.ts",
        // "./main/atom/atomUtils.ts",
        // "./main/atom/autoCompleteProvider.ts",
        // "./worker/messages.ts",
        // "./worker/parent.ts"
    ]
