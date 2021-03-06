var express = require('express');
var bodyParser = require('body-parser');
var policies = require('./lib/routes/policies');
var logger = require('./lib/log/logger');
var app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use('/v1/apps',policies);

require('./lib/models')(function (error) {
  if (!error) {
    var server = app.listen(process.env.PORT || 3002, function() {
	  logger.info('Autoscaler API server listening on port ' + server.address().port);
    });  
  } else {
    logger.error ('Autoscaler API server failed to start with error', error);
  }
});

module.exports = app;
