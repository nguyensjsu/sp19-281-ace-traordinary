
var userid = "";
var name = "Ramya";
var myList = new Array();
var connection;

//myList.push("Ravi");
Array.remove = function(array, from, to) {
    var rest = array.slice((to || from) + 1 || array.length);
    array.length = from < 0 ? array.length + from : from;
    return array.push.apply(array, rest);
};

//this variable represents the total number of popups can be displayed according to the viewport width
var total_popups = 0;

//arrays of popups ids
var popups = [];

userid = localStorage.getItem("username");
console.log("user id is",userid);
function  open_connection(){



}

  //  window.onload = localStorage.getItem("connection",userid);

function register_popup(id)
{

    for(var iii = 0; iii < popups.length; iii++)
    {
        //already registered. Bring it to front.
        if(id == popups[iii])
        {
            Array.remove(popups, iii);
            popups.unshift(id);
            calculate_popups();
            return;
        }
    }
   var chatPopup =  '<div class="msg_box" id="'+id+'" style="right:270px" rel="'+id+'">'+
        '<div class="msg_head">'+id+
        '<div class="close">x<a href="javascript:close_popup(\''+ id +'\');">&#10005;</a></div> </div>'+
        '<div class="msg_wrap"> <div class="msg_body"> <div class="msg_push"></div> </div>'+
        '<div class="msg_footer"><textarea class="msg_input" rows="4"></textarea>></div>  </div>  </div>' ;

    var element = '<div class="popup-box chat-popup" id="'+ id +'" value="chatbox">';
    element = element + '<div class="popup-head">';
    element = element + '<div class="popup-head-left">'+ id +'</div>';
    element = element + '<div class="popup-head-right"><a href="javascript:close_popup(\''+ id +'\');">&#10005;</a></div>';
    element = element + '<div style="clear: both"></div></div><div class="popup-messages" value ="popup-messages"><div id="content"></div>\n' +
        '<div id="messagearea" value ="messagearea">\n' +
        '    <span id="status">Connecting...</span>\n' +
        '    <input type="text" id="input" value="Default" onKeyPress="javascript:doit_onkeypress( event, \''+ id +'\');" />\n' +
        '</div></div></div>';

    document.getElementsByTagName("body")[0].innerHTML = document.getElementsByTagName("body")[0].innerHTML + chatPopup;

    popups.unshift(id);

    calculate_popups();

}

//this is used to close a popup
function close_popup(id)
{
    for(var iii = 0; iii < popups.length; iii++)
    {
        if(id == popups[iii])
        {
            Array.remove(popups, iii);

            document.getElementById(id).style.display = "none";

            calculate_popups();

            return;
        }
    }
}

//displays the popups. Displays based on the maximum number of popups that can be displayed on the current viewport width
function display_popups()
{
    var right = 220;

    var iii = 0;
    for(iii; iii < total_popups; iii++)
    {
        if(popups[iii] != undefined)
        {
            var element = document.getElementById(popups[iii]);
            element.style.right = right + "px";
            right = right + 320;
            element.style.display = "block";
        }
    }

    for(var jjj = iii; jjj < popups.length; jjj++)
    {
        var element = document.getElementById(popups[jjj]);
        element.style.display = "none";
    }
}

function calculate_popups()
{
    var width = window.innerWidth;
    if(width < 540)
    {
        total_popups = 0;
    }
    else
    {
        width = width - 200;
        //320 is width of a single popup box
        total_popups = parseInt(width/320);
    }

    display_popups();

}



function generate_sidebar(){

    console.log("in side bar",myList.length);

    if(myList.length==0){
        document.getElementById("sidebar").innerHTML = document.getElementById("sidebar").innerHTML + '<span> No users </span>'
        return;
    }
    var i = 0;
    var element = '<div>';
    for(i; i<myList.length ;i++){
        element = element + '<div class="sidebar-name">';
        element = element + '<a href="javascript:register_popup(\''+ myList[i] +'\');"><span>'+myList[i]+'</span></a></div>';

    }
    document.getElementById("sidebar").innerHTML =  element+'<div>';
}
//recalculate when window is loaded and also when window is resized.
window.addEventListener("resize", calculate_popups);
//window.addEventListener("load", open_connection);
console.log("below event listener document ready",connection);
window.addEventListener("load", calculate_popups);
window.addEventListener("load", generate_sidebar);
//window.myList = myList;




        function doit_onkeypress(event, id, element ){
            console.log("in doit", value);
            $(function () {

            if (event.keyCode == 13 || event.which == 13){
                var message = new Object();
                message.Userid = userid;
                message.Username = name;
                message.Receiverid = id;
                var time = new Date();
            //    var timeStr = time.toLocaleTimeString();
                message.Time = time;
                message.Message = value;
                console.log("imessage is",JSON.stringify(message));
                connection.send(JSON.stringify(message));
                $(this).val('');
                if(msg.trim().length != 0){
                    var chatbox = $(this).parents().parents().parents().attr("rel") ;
                    $('<div class="msg-right">'+msg+'</div>').insertBefore('[rel="'+chatbox+'"] .msg_push');
                    $('.msg_body').scrollTop($('.msg_body')[0].scrollHeight);
                }
            }

            });
            return;


        }

        //creates markup for a new popup. Adds the id to popups array.
       // var message = new Object();


$(document).on('keypress', 'textarea' , function(e) {
    if (e.keyCode == 13 ) {
        var msg = $(this).val();
        $(this).val('');
        if(msg.trim().length != 0){
            var chatboxId = $(this).parents().parents().parents().attr("rel") ;
            var message = new Object();
            message.Userid = userid;
            message.Username = name;
            message.Receiverid = chatboxId;
            var time = new Date();
            message.Message = msg;
            message.Time = time;
            message.Type = "Text";
            console.log("imessage is",JSON.stringify(message));
            connection.send(JSON.stringify(message));
            $('<div class="msg-right">'+ userid + '@ ' + (time.getHours() < 10 ? '0'
                + time.getHours() : time.getHours()) + ':'
                + (time.getMinutes() < 10
                    ? '0' + time.getMinutes() : time.getMinutes())
                + ': ' + msg +'</div>').insertBefore('[rel="'+chatboxId+'"] .msg_push');
            $('.msg_body').scrollTop($('.msg_body')[0].scrollHeight);
        }
    }
});

        //calculate the total number of popups suitable and then populate the toatal_popups variable.
$(function () {

   // window.addEventListener("load", open_connection);
    var log = $('#log');

        console.log("inside document ready",connection);
        var ws = "ws://localhost:8000/ws/";
        var url = ws.concat(userid);
console.log(url);
    connection = new WebSocket(url);
    alert("connection");
    console.log("below event listener document ready",connection);
        connection.onerror = function (error) {
            // just in there were some problems with connection...
            log.html($('<p>', {
                text: 'Sorry, but there\'s some problem with your '
                    + 'connection or the server is down.'
            }));
        };
        // most important part - incoming messages
        connection.onmessage = function (message) {
            // try to parse JSON message. Because we know that the server
            // always returns JSON this should work without any problem but
            // we should make sure that the massage is not chunked or
            // otherwise damaged.
            console.log("data: "+ typeof message.data);
            console.log("data: message "+ message.data);
            try {
                var json = JSON.parse(message.data);
            } catch (e) {
                console.log('Invalid JSON: ', message.data);
                return;
            }
            // NOTE: if you're not sure about the JSON structure
            // check the server source code above
            // first response from the server with user's color
            if (json.Type === 'OnlineUsers') {
              var users  = json.Users;
               // window.myList = users.keys;
                myList = new Array();
                console.log(myList);
               // myList = users.keys;
                for (var p in users) {
                    console.log(p);
                    myList.push(p);
                }
                generate_sidebar();
                // from now user can start sending messages
            }
            else if (json.Type === 'Text') { // it's a single message

                // let the user write another message
                //input.removeAttr('disabled');
                addMessage(json.UserId, json.Message, new Date(json.Time));
            } else {
                console.log('Hmm..., I\'ve never seen JSON like this:', json);
            }
        };

    function addMessage(author, message, dt) {


        var element = document.getElementById(author);
        if(element === null){
            register_popup(author);
        }
        $('<div class="msg-left">'+ author + '@ ' + (dt.getHours() < 10 ? '0'
            + dt.getHours() : dt.getHours()) + ':'
            + (dt.getMinutes() < 10
                ? '0' + dt.getMinutes() : dt.getMinutes())
            + ': ' + message +'</div>').insertBefore('[rel="'+ author +'"] .msg_push');
        $('.msg_body').scrollTop($('.msg_body')[0].scrollHeight);


    }
        });

