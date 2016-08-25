var httpmod = require('http');
var url = require('url');
var uridispatcher = require('httpdispatcher');

var time = new Date();

//Parse CF VCAP_SERVICES
var vcap_services = JSON.parse(process.env.VCAP_SERVICES);

//Parse CF VCAP_APPLICATION
var vcap_application = JSON.parse(process.env.VCAP_APPLICATION);

uridispatcher.onGet("/xxx", function(req, res) {
});

uridispatcher.onPost("/yyy", function(req, res) {
});

//Support CF & Legacy 
server.listen(process.env.PORT || 8000);
