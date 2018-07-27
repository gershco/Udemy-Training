$(function () {

// attr(), prop(), val()

var specialLink = $("#specialLink");

specialLink.attr("href", "https://www.google.co.uk");

$("#weirdLink").attr("href", "http://www.snapbubbles.com/");

$("a").attr("target","_blank");

/*
var checkbox = $("input:checkbox");

checkbox.prop("input:checkbox", "");
*/

var textInput = $("input:text");
console.log("textInput is " + textInput.val());
textInput.val("Fred Bloggs")

var rangeInput = $("input[type='range']");
console.log("rangeInput is " + rangeInput.val());

var checkbox = $("input[type=checkbox]");
console.log("checkbox 'checked' is " + checkbox.prop("checked"));

});
