const electron = require('electron');
const {app, BrowserWindow, Menu, Accelerator} = electron;

app.setName('JSON Tilemap Maker');

const template = [
  {
    label: 'File',
    submenu: [
      {
        label: 'New Tilemap...',
        accelerator: 'CmdOrCtrl+N',
        click: function (menuItem, browserWindow, event) {
          exports.newTilemap();
        }
      },
      {
        label: 'Open Tilemap...',
        accelerator: 'CmdOrCtrl+O',
        click: function (menuItem, browserWindow, event) {
          console.log(menuItem.label)
        }
      },
      {
        label: 'Load Spritesheet...',
        accelerator: 'CmdOrCtrl+L',
        click: function (menuItem, browserWindow, event) {
          console.log(menuItem.label)
        }
      },
      {type: 'separator'},
      {
        label: 'Save',
        accelerator: 'CmdOrCtrl+S',
        click: function (menuItem, browserWindow, event) {
          console.log(menuItem.label)
        }
      },
      {
        label: 'Save As...',
        accelerator: 'Shift+CmdOrCtrl+S',
        click: function (menuItem, browserWindow, event) {
          console.log(menuItem.label)
        }
      }
    ]
  },
  {
    role: 'editMenu'
  },
  {
    role: 'windowMenu'
  }
];

if (process.platform === 'darwin') {
  template.unshift({
    label: app.getName(),
    submenu: [
      {role: 'quit'}
    ]
  })
}

const menu = Menu.buildFromTemplate(template);

app.on('ready', function () {
  let win = new BrowserWindow({width: 800, height: 600, resizable: false});
  win.loadURL(`file://${__dirname}/index.html`);
  win.webContents.openDevTools();
  Menu.setApplicationMenu(menu);
});

exports.newTilemap = function () {
  let win = new BrowserWindow({width: 300, height: 200, resizable: false});
  win.loadURL(`file://${__dirname}/newtilemap.html`);
}
exports.CreateTilemap = function (width, height) {
  console.log(width, height);
}
