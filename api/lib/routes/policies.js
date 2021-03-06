var models = require('../models')();
var express = require('express');
var router = express.Router();
var logger = require('../log/logger');
var HttpStatus = require('http-status-codes');

router.put('/:app_id/policy',function(req, res) {
  logger.info('Policy creation request received for app id : ' + req.params.app_id);
  models.policy_json.create({
    policy_json: req.body,
    app_id: req.params.app_id
  }).then(function(result) {
    var result = {
      'success': true,
      'error': null,
      'result': result
    };
    res.status(HttpStatus.OK).send(result);
  }).catch(function(error) {
    logger.error ('Policy creation failed for app id : ' + req.params.app_id, error);
    var error = {
      'success': false,
      'error': error,
      'result': null
    };
    res.status(HttpStatus.BAD_REQUEST).send(error)
  });
});

module.exports = router;
