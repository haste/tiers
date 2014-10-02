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
