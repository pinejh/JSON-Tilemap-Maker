const remote = require('electron').remote;
const main = remote.require('./main.js');

var wInput = document.getElementById('width');
var hInput = document.getElementById('height');

var cancel = document.createElement('button');
cancel.textContent = 'Cancel';
cancel.addEventListener('click', () => {
  remote.getCurrentWindow().close();
}, false);
document.body.appendChild(cancel);

var create = document.createElement('button');
create.textContent = 'Create';
create.addEventListener('click', () => {
  var w = wInput.value;
  var h = hInput.value;
  if (w >= 1 && w <= 100 && h >=1 && h <= 100) {
    main.CreateTilemap(w, h);
    remote.getCurrentWindow().close();
  } else {
    alert('Width and Height must be between 1 and 100.');
  }
}, false);
document.body.appendChild(create);
