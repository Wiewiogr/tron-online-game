var canvas = document.getElementById('game-window');
var ctx = canvas.getContext('2d');

let socket = new WebSocket('ws://127.0.0.1:8080/ws');
console.log('Attempting Connection...');

var players = []
var playerId = 1;

socket.onmessage = (messageEvent) => {
  message = JSON.parse(messageEvent.data)
  if(message.type == "newPlayerId") {
    playerId = message.id
  } else if (message.type == "playersPosition") {
    players = message.players
  }
};


document.addEventListener('keydown', keyDownHandler, false);
function keyDownHandler(e) {
  if (e.key == 'ArrowRight') {
    socket.send(JSON.stringify({id:playerId, key:"Right"}))
  } else if (e.key == 'ArrowLeft') {
    socket.send(JSON.stringify({id:playerId, key:"Left"}))
  }
}

function tick() {
  draw();
  requestAnimationFrame(tick);
}
tick();

function draw() {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  for (i = 0; i < players.length; i++) {
    let player = players[i]
    ctx.beginPath();
    ctx.arc(player.x, player.y, 5, 0, Math.PI * 2);
    ctx.fillStyle = '#0095DD';
    ctx.fill();
    ctx.closePath();  
  } 
}

