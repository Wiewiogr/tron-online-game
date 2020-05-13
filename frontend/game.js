var canvas = document.getElementById('game-window');
var ctx = canvas.getContext('2d');

let socket = new WebSocket('ws://127.0.0.1:8080/ws');
console.log('Attempting Connection...');

var board = [];
var playerId = 1;

let colors = ['#e74c3c', '#0095DD', `#f4d03f`, `#27ae60`];

socket.onmessage = (messageEvent) => {
  message = JSON.parse(messageEvent.data);
  if (message.type == 'newPlayerId') {
    playerId = message.id;
  } else if (message.type == 'playersTrace') {
    board = message.board;
  }
};

document.addEventListener('keydown', keyDownHandler, false);
function keyDownHandler(e) {
  if (e.key == 'ArrowRight') {
    socket.send(JSON.stringify({ id: playerId, key: 'Right' }));
  } else if (e.key == 'ArrowLeft') {
    socket.send(JSON.stringify({ id: playerId, key: 'Left' }));
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
    let position = board[id].position;

    ctx.beginPath();
    ctx.strokeStyle = colors[id % 4];
    ctx.lineWidth = 5
    for (const [i, trace] of board[id].traces.entries()) {
      if (i == 0) {
        ctx.moveTo(trace.x, trace.y);
      } else {
        ctx.lineTo(trace.x, trace.y);
      }
    }
    ctx.lineTo(position.x, position.y)
    ctx.stroke();
    ctx.closePath();

    ctx.beginPath();
    
    ctx.arc(position.x, position.y, 5, 0, Math.PI * 2);
    ctx.fillStyle = colors[id % 4];
    ctx.fill();
    ctx.closePath();
  }
}
