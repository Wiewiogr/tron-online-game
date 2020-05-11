var canvas = document.getElementById('game-window');
var ctx = canvas.getContext('2d');

let socket = new WebSocket('ws://127.0.0.1:8080/ws');
console.log('Attempting Connection...');

var board = []
var playerId = 1;

socket.onmessage = (messageEvent) => {
  message = JSON.parse(messageEvent.data)
  if(message.type == "newPlayerId") {
    playerId = message.id
  } else if (message.type == "playersPosition") {
    board = message.board
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
  for (id in board) {
    let position = board[id]
    ctx.beginPath();
    ctx.arc(position.x, position.y, 5, 0, Math.PI * 2);
    ctx.fillStyle = '#0095DD';
    ctx.fill();
    ctx.closePath();  
  } 
}

