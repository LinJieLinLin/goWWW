var gulp = require("gulp");
var babel = require("gulp-babel");
var shell = require("gulp-shell");

gulp.task('check', shell.task(['es-checker']));
gulp.task("default", function () {
    return gulp.src("es6Test.js")
        .pipe(babel({
            presets: ['es2015']
        }))
        .pipe(gulp.dest("dist"));
});