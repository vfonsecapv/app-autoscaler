'use strict';

var fs = require('fs');
var path = require('path');
var supertest = require("supertest");
var should = require('should');

var settings = JSON.parse(
  fs.readFileSync(path.join(__dirname, '../../config/settings.json'), 'utf8'));
var auth = new Buffer(settings.username + ":" + settings.password).toString('base64');


describe('getCatalog RESTful API', function() {
  var server;
  beforeEach(function() {
    delete require.cache[require.resolve('../../lib/index.js')];
    server = require(path.join(__dirname, '../../lib/index.js'));
  });

  afterEach(function(done) {
    server.close(done);
  });

  it("should return catalog json", function(done) {
    var catalog = JSON.parse(fs.readFileSync(path.join(__dirname, '../../config/catalog.json'), 'utf8'));

    supertest(server)
      .get("/v2/catalog")
      .set("Authorization", "Basic " + auth)
      .expect(200)
      .expect("Content-type", /json/)
      .end(function(err, res) {
        res.status.should.equal(200);
        JSON.stringify(res.body).should.equal(JSON.stringify(catalog));
        done();
      });
  });

});