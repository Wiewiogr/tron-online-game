var canvas = document.getElementById('game-window');
var ctx = canvas.getContext('2d');

let socket = new WebSocket('ws://127.0.0.1:8080/ws');
console.log('Attempting Connection...');

var players = []

socket.onmessage = (message) => {
  players = JSON.parse(message.data)
  console.log(parsedMessage);
};

var id = 1;
document.addEventListener('keydown', keyDownHandler, false);
function keyDownHandler(e) {
  if (e.key == 'ArrowRight') {
    socket.send(JSON.stringify({id:id, key:"Right"}))
  } else if (e.key == 'ArrowLeft') {
    socket.send(JSON.stringify({id:id, key:"Left"}))
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

