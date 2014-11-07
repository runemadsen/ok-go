var gulp = require("gulp");
var coffee = require("gulp-coffee");
var include = require("gulp-include");
var rev = require('gulp-rev');
var gulpif = require('gulp-if');
var sass = require("gulp-sass");
var uglify = require("gulp-uglify");
var minify = require("gulp-minify-css");

//var merge = require("merge-stream");
//var del = require("del");
//var exec = require("child_process").exec;
//var concat = require("gulp-concat");
//var rename = require("gulp-rename");
//var uglify = require("gulp-uglify");

//var open = require("gulp-open");
//var connect = require("gulp-connect");
//
//var directoryMap = require("gulp-directory-map");
//var liquify = require('gulp-liquify');
//var through = require('through2');
//var _ = require("underscore");

//
//gulp.task("assets", function() {
//  return gulp.src("./scss/oreilly.scss").pipe(sass()).pipe(gulp.dest("./public/css"));
//});
//
//gulp.task("serve", ["docs"], function() {
//  return connect.server({
//    port: 8001,
//    root: 'public'
//  });
//});

var assets = [
  'app/assets/javascripts/application.coffee', 
  'app/assets/stylesheets/application.scss'
]

// Compilation for development
gulp.task("assets:compile", function() {
  gulp.src(assets)
    .pipe(include())
    .pipe(gulpif('*.coffee', coffee()))
    .pipe(gulpif('*.scss', sass()))
    .pipe(gulp.dest("public/assets"))
});

// Compilation for production
gulp.task("assets:precompile", function() {
  gulp.src(assets)
    .pipe(include())
    .pipe(gulpif('*.coffee', coffee()))
    .pipe(gulpif('*.scss', sass()))
    .pipe(gulpif('*.css', minify()))
    .pipe(gulpif('*.js', uglify()))
    .pipe(gulp.dest("public/assets")) // file without digest
    .pipe(rev())
    .pipe(gulp.dest("public/assets")) // file with digest
    .pipe(rev.manifest({path: 'manifest.json'}))
    .pipe(gulp.dest("public/assets")) // manifest.json
});