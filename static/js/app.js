/*global $:false */
'use strict';

// Landing
$(".btn-signin").click(function(event) {
	event.preventDefault();

	$(".signin-box").show();
	$(".signup-box").hide();
	$(".forgot-box").hide();
});

$(".btn-register").click(function(event) {
	event.preventDefault();

	$(".signin-box").hide();
	$(".signup-box").show();
	$(".forgot-box").hide();
});

$(".btn-forgot").click(function(event) {
	event.preventDefault();

	$(".signin-box").hide();
	$(".signup-box").hide();
	$(".forgot-box").show();
});

// Upload
$(".form-upload").submit(function(event) {
	event.preventDefault();

	$.ajax({
		url: "/upload",
		type: "POST",
		data: new FormData(this),
		dataType: "json",
		processData: false,
		contentType: false,

		statusCode: {
			413: function() {
				$(".progress").hide();
				$(".alert")
				.hide()
				.removeClass("alert-success")
				.removeClass("hide")
				.addClass("alert-warning")
				.text("File size(s) exceed total upload limit.");
			}
		},

		beforeSend: function() {
			var files = $(".form-upload input")[0].files;
			var maxSize = 1024 * 1024 * 10;
			var totalSize = 0;
			for (var i = 0, numFiles = files.length; i < numFiles; i++) {
				var file = files[i];
				totalSize += file.size;
			}

			if(totalSize > maxSize) {
				$(".progress").hide();
				$(".alert")
				.hide()
				.removeClass("alert-success")
				.removeClass("hide")
				.addClass("alert-warning")
				.text("File size(s) exceed total upload limit.")
				.fadeIn();

				return false;
			}

			$(".form-upload button").prop("disabled", true);
			$(".progress").hide().removeClass("hide").fadeIn();
		},

		success: function(data, textStatus, xhr) {
			console.log(data, textStatus, xhr);
			$(".alert")
			.hide()
			.addClass("alert-success")
			.removeClass("alert-warning")
			.removeClass("hide")
			.text(data.message)
			.fadeIn();
		},

		xhr: function() {
			var xhr = new window.XMLHttpRequest();

			xhr.upload.addEventListener("progress", function(event) {
				if (event.lengthComputable) {
					var percentage = Math.floor(event.loaded / event.total * 100);
					$(".progress-bar").width(percentage + "%");

					if(percentage === 100) {
						$(".progress").delay(250).fadeOut(function() {
							$(".form-upload").find("button").prop("disabled", false);
							$(".progress-bar").width(0);
						});
					}
				}
			}, false);

			return xhr;
		}
	});
});
