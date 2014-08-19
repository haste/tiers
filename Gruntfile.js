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
						'static/vendor/bootstrap/dist/js/bootstrap.js',
						'static/js/app.js',
						'static/js/bullet.js'
					]
				}]
			}
		},

		cssmin: {
			combine: {
				files: {
					'static/css/app.min.css': [
						'static/vendor/bootstrap/dist/css/bootstrap.css',
						'static/css/app.css'
					]
				}
			}
		},

		watch: {
			js: {
				files: [
					'static/js/app.js',
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
