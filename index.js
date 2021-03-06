'use strict';
// make `.jsx` file requirable by node
require('node-jsx').install();
var express = require('express');
var bodyParser = require('body-parser')
var renderer = require('react-engine');
var app = express();
app.use(bodyParser.json())

// create the view engine with `react-engine`
var engine = renderer.server.create();
// set the engine
app.engine('.jsx', engine);
// set the view directory
app.set('views', __dirname + '/public/views');
// set js as the view engine
app.set('view engine', 'jsx');
// finally, set the custom view
app.set('view', renderer.expressView);

//expose public folder as static assets
app.use(express.static(__dirname + '/public'));


app.post('/render', function(req, res) {
  var variables = req.body;
  res.render('index', variables);
})


var server = app.listen(3000, function() {
  var host = server.address().address;
  var port = server.address().port;
  console.log('Example app listening at http://%s:%s', host, port);
});
