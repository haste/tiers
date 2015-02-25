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
						'static/vendor/jquery/dist/jquery.js',
						'static/vendor/bootstrap/js/transition.js',
						'static/vendor/bootstrap/js/collapse.js',
						'static/vendor/bootstrap/js/dropdown.js',
						'static/js/app.js'
					]
				}]
			}
		},

		less: {
			app: {
				options: {
					strictMath: true,
					outputSourceFiles: true,
				},

				src: 'static/css/app.less',
				dest: 'static/css/app.css',
			}
		},

		cssmin: {
			combine: {
				files: {
					'static/css/app.min.css': [
						'static/css/app.css'
					]
				}
			}
		},

		watch: {
			js: {
				files: [
					'static/js/app.js'
				],
				tasks: ['jshint', 'uglify']
			},

			lint: {
				files: [
					'static/css/app.less'
				],
				tasks: ['less']
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
	grunt.loadNpmTasks('grunt-contrib-less');

	// Default task(s).
	grunt.registerTask('default', ['jshint', 'csslint', 'uglify', 'cssmin', 'watch']);
};
