var httpmod = require('http');
var url = require('url');
var uridispatcher = require('httpdispatcher');

var time = new Date().getTime();

//Parse CF VCAP_SERVICES
var vcap_services = JSON.parse(process.env.VCAP_SERVICES);

//Parse CF VCAP_APPLICATION
var vcap_application = JSON.parse(process.env.VCAP_APPLICATION);

uridispatcher.onGet("/authen", function(req, res) {
  var reqData = url.parse(req.url, true).query;

  //On bloque le Cross-Origin Request
  res.setHeader('Content-type', 'text/plain');
  res.setHeader('Access-Control-Allow-Origin', '*');
  
  var resp = 'Authent OK pour  '+reqData['login']; 
 
  console.log(time+"[ GET /authen ] Reponse : "+res);
  res.end(resp);
});

uridispatcher.onPost("/authen", function(req, res) {
});

//Write each request and response to STDOUT (to test CF logging aggregation feature).
//Example to STDERR ?
var server = httpmod.createServer(function (req, res) {
  uridispatcher.dispatch(req, res);
  console.log(time+"[ RECEIVING ] : "+req.url);
});

//Support CF & Legacy 
server.listen(process.env.PORT || 8000);
