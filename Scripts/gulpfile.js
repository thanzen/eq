var browserify = require('gulp-browserify');
var gulp = require('gulp');
var libs = [
  'react',
  'jquery',
  'flux',
  'bluebird'

 // 'react/lib/ReactCSSTransitionGroup',
  //'react/lib/cx',
  //'underscore',
  //'loglevel'
];
gulp.task('build-libs', function () {
    // place code for your default task here
    console.log("test");

    //use fake lib.js to start the file stream
    var stream = gulp.src('./dist/lib.js', { read: false }).pipe(browserify({
        debug: false,  // Don't provide source maps for vendor libs
    })).on('prebundle', function (bundle) {
        // Require vendor libraries and make them available outside the bundle.
        libs.forEach(function (lib) {
            bundle.require(lib);
        });
    }).pipe(gulp.dest('./dist/libs'));
    return stream;
});

gulp.task('app', function () {
      var stream = gulp.src('./app.js', { read: false }).pipe(browserify({
          debug: true,  // If not production, add source maps
          transform: [["reactify",{"es6": false}]],
          extensions: ['.js']
      })).on('prebundle', function (bundle) {
          // The following requirements are loaded from the vendor bundle
          //libs.forEach(function (lib) {
          //    bundle.external(lib);
          //});
      }).pipe(gulp.dest('./dist/apps'));

    return stream;
});

gulp.task('default', [ 'app']);
    //gulp.watch('./**/*.js', function (event) {
    //    console.log('File ' + event.path + ' was ' + event.type + ', running tasks...');
    //});