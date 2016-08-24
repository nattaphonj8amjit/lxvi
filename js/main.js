MainJS = (function ($) {

    function init() {
        $('.ui.selection.dropdown').dropdown();
    }


    function digits(digit){
         return digit.toString().replace(/(\d)(?=(\d\d\d)+(?!\d))/g, "$1,");
    }







    function getAllGoods() {
        $.ajax({
                method: "GET",
                cache: false,
                url: $(location).attr('origin') + "/findAll",
                data: {name: "John", location: "Boston"}
            })
            .done(function (objs) {
                var tmp = '';
                for (obj of objs) {
                    console.log(JSON.stringify(obj))
                    tmp += ' <div class="column">' +
                        '<div class="ui card">' +
                        '<div class="ui slide masked reveal image">' +
                        '   <img src="'+obj.Imgsrc+'" class="visible content">' +
                       // '   <img src="http://semantic-ui.com/images/avatar/large/elliot.jpg" class="hidden content">' +
                        '   </div>' +
                        '   <div class="content">' +
                        '   <div class="header"> <a href="/findDetail"> ' + obj.Name + '</a></div>' +
                        '<div class="meta">' +
                        '   <a>' + obj.Collection + '</a>' +
                        '   </div>' +
                        '   <div class="description">' +obj.Detail+ '</div>' +
                        '   <div class="description">' +
                        '   <div class="ui mini horizontal green statistic">' +
                        '<div class="value price">' +
                        ' '+digits(obj.Price)+' </div> <div class="label"> Baht </div>  </div>' +
                        '</div>' +
                        '</div>' +
                        '<div class="extra content">' +
                        '   <span class="right floated">' +
                        '   Joined in 2013' +
                        '   </span>' +
                        '   <span>' +
                        '<div class="fb-share-button" data-href="https://lxvi.co" data-layout="button_count" data-mobile-iframe="true"></div>' +
                        '</span>' +
                        '</div>' +
                        '</div>' +
                        '</div>';

                }
                $('#goodsGrid').html(tmp);
               // $('.value.price').digits();
            })
            .fail(function () {
                console.log("error");
            })
            .always(function () {
                console.log("complete");
            });
    }


    return {
        init: init,
        getAllGoods: getAllGoods
    }

})(jQuery);




$(function () {
    MainJS.init();
    MainJS.getAllGoods();
})
