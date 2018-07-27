$(function () {

$(".gallery").css("display", "none");

var redBox = $(".red-box");
console.log("Width of red-box is " + redBox.css("width"));
console.log("Without units, using .width is " + redBox.width());

redBox.css("backgroundColor", "#E9590D");
redBox.css("width", "+=20px");

$("p").css("fontSize", "+=20px");

$(".blue-box").css("color", "#B3F1EF");
});
