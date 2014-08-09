module.exports = function(grunt) {

	// Project configuration.
	grunt.initConfig({
		pkg: grunt.file.readJSON('package.json'),

		jshint: {
			sources: 'static/js/app.js',
			options: {
				jshintrc: '.jshintrc'
			}
		},

		csslint: {
			options: {
				"adjoining-classes": false,
			},
			src: ['css/app.css']
		},

		uglify: {
			options: {
				compress: true,
				sourceMap: true,
				sourceMapName: "static/js/app.map",
			},

			build: {
				files: [{
					'static/js/app.min.js': [
						'vendor/bootstrap/dist/js/bootstrap.js',
						'vendor/angular/angular.js',
						'vendor/angular-route/angular-route.js',
						'static/js/app.js',
						'static/js/services.js',
						'static/js/controllers.js',
						'static/js/filters.js',
						'static/js/directives.js',

						'static/js/bullet.js'
					]
				}]
			}
		},

		cssmin: {
			combine: {
				files: {
					'static/css/app.min.css': [
						'vendor/bootstrap/dist/css/bootstrap.css',
						'vendor/nvd3/nv.d3.css',
						'static/css/app.css'
					]
				}
			}
		},

		watch: {
			js: {
				files: [
					'static/js/app.js',
					'static/js/services.js"',
					'static/js/controllers.js',
					'static/js/filters.js',
					'static/js/directives.js',

					'static/js/bullet.js'
				],
				tasks: ['jshint', 'uglify']
			},

			css: {
				files: [
					'static/css/app.css'
				],
				tasks: ['csslint', 'cssmin']
			},

			config: {
				options: {
					reload: true,
				},

				files: [ 'Gruntfile.js']
			},
		}
	});

	grunt.loadNpmTasks('grunt-contrib-jshint');
	grunt.loadNpmTasks('grunt-contrib-csslint');
	grunt.loadNpmTasks('grunt-contrib-uglify');
	grunt.loadNpmTasks('grunt-contrib-cssmin');
	grunt.loadNpmTasks('grunt-contrib-watch');

	// Default task(s).
	grunt.registerTask('default', ['jshint', 'csslint', 'uglify', 'cssmin']);
};
