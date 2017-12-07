
function loadData() {

    var $body = $('body');
    var $wikiElem = $('#wikipedia-links');
    var $nytHeaderElem = $('#nytimes-header');
    var $nytElem = $('#nytimes-articles');
    var $greeting = $('#greeting');

    // clear out old data before new request
    $wikiElem.text("");
    $nytElem.text("");

    // Streetview
    var street = $( "#street" ).val();
    var city = $( "#city" ).val();
    var loc = street + ', ' + city;
 
    $greeting.text('So, you want to live at ' + loc + '?');


    var streetviewUrl = 'http://maps.googleapis.com/maps/api/streetview?size=600x400&location=' + loc + '';

    $body.append('<img class="bgimg" src="' + streetviewUrl + '">');

    // NY Times
    var nytimesUrl = 'http://api.nytimes.com/svc/search/v2/articlesearch.json?q=' + city + '&sort=newest&api-key=3a98ec7202bf49c08b6de1cc32b4331f'

    $.getJSON(nytimesUrl, function (data) {
        
        $nytHeaderElem.text('New York Times Articles About ' + city);

        articles = data.response.docs;
        for (var i = 0; i < articles.length; i++) {
            var article = articles[i];
            $nytElem.append('<li class="article">'+ '<a href = "'+ article.web_url + '">' + article.headline.main + '</a>' + '<p>' + article.snippet + '</p>' + '</li>');
        };
    }).error(function(e){
        $nytHeaderElem.text('New York Times articles could not be loaded.');
    });

    //Wikipedia
    var wikiUrl = 'http://en.wikipedia.org/w/api.php?action=opensearch&search=' + city + '&format=json&callback=wikiCallback';

    var wikiRequestTimeout = setTimeout(function(){
        $wikiElem.text("Failed to load Wikipedia resources");
    }, 8000);

    $.ajax({
        url: wikiUrl,
        dataType: "jsonp",
        success: function( response) {
            var articleList = response[1];
            
            for (var i = 0; i < articleList.length; i++) {
                articleStr = articleList[i];
                var url = 'http://en.wikipedia.org/wiki/' + articleStr;
                $wikiElem.append('<li><a href="' + url + '">' + articleStr + '</a></li>');
            };
            clearTimeout(wikiRequestTimeout);
        }
    })

    return false;

  
};

$('#form-container').submit(loadData);
