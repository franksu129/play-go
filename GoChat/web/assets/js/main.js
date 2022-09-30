import chatSocket from './chatScocket.js';

var text = document.getElementById('msg');
var send = document.getElementById('send');
var PERSON_NAME = 'Guest' + Math.floor(Math.random() * 1000);

let person = prompt('Please enter your name:', PERSON_NAME);
if (person != '') {
  PERSON_NAME = person;
}

var url = 'ws://' + window.location.host + '/ws?id=' + PERSON_NAME;
chatSocket.init(url, PERSON_NAME);

send.onclick = function (e) {
  handleMessageEvent();
};

text.onkeydown = function (e) {
  if (e.keyCode === 13 && text.value !== '') {
    handleMessageEvent();
  }
};

function handleMessageEvent() {
  if (text.value !== '') {
    chatSocket.send(text.value);
    text.value = '';
  }
}