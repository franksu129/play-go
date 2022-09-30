const EVENT_MESSAGE = 'message';
const EVENT_OTHER = 'other';
const LEFT = 'left';
const RIGHT = 'right';

const userPhotos = [
  'https://cdn-icons-png.flaticon.com/512/7927/7927405.png',
  'https://cdn-icons-png.flaticon.com/512/7927/7927404.png',
  'https://cdn-icons-png.flaticon.com/512/7927/7927417.png',
  // "https://www.flaticon.com/svg/static/icons/svg/3408/3408720.svg"
];

var chatroom = document.getElementsByClassName('msger-chat');

function getEventMessage(msg) {
  var msg = `<div class="msg-left">${msg}</div>`;
  return msg;
}

function getMessage(name, img, side, text) {
  const d = new Date();
  //   Simple solution for small apps
  var msg = `
    <div class="msg ${side}-msg">
      <div class="msg-img" style="background-image: url(${img})"></div>

      <div class="msg-bubble">
        <div class="msg-info">
          <div class="msg-info-name">${name}</div>
          <div class="msg-info-time">${d.getFullYear()}/${d.getMonth()}/${d.getDay()} ${d.getHours()}:${d.getMinutes()}</div>
        </div>

        <div class="msg-text">${text}</div>
      </div>
    </div>
  `;
  return msg;
}

function insertMsg(msg, domObj) {
  domObj.insertAdjacentHTML('beforeend', msg);
  domObj.scrollTop += 500;
}

function getRandomNum(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

export default {
  user: {
    name: '',
    poto: '',
  },
  ws: null,
  init: function (url, name) {
    this.user.name = name;
    this.user.poto = userPhotos[getRandomNum(0, userPhotos.length)];

    this.ws = new WebSocket(url);
    this.ws.onmessage = function (e) {
      var m = JSON.parse(e.data);
      var msg = '';
      switch (m.event) {
        case EVENT_MESSAGE:
          if (m.name == this.user.name) {
            msg = getMessage(m.name, m.photo, RIGHT, m.content);
          } else {
            msg = getMessage(m.name, m.photo, LEFT, m.content);
          }
          break;
        case EVENT_OTHER:
          if (m.name != this.user.name) {
            msg = getEventMessage(m.name + ' ' + m.content);
          } else {
            msg = getEventMessage('您已' + m.content);
          }
          break;
      }
      insertMsg(msg, chatroom[0]);
    }.bind(this);

    this.ws.onclose = (e) => {
      console.log(e);
    };
  },
  send: function (msg) {
    this.ws.send(
      JSON.stringify({
        event: 'message',
        photo: this.user.poto,
        name: this.user.name,
        content: msg,
      })
    );
  },
};
