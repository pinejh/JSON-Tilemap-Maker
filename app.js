var fs = require('fs');
var jsonfile = require('jsonfile');

var file = '/tmp/data.json';

jsonfile.writeFile(file, {}, function (err) {
  console.error(err);
});

console.log(fs);
console.log(jsonfile);
