var path = require('path');
var url = require('url');
var express = require('express');

var time = new Date().getTime();
var httpserver = express();
httpserver.use(express.static(path.join(__dirname, 'static')));

//Parse CF VCAP_SERVICES
var vcap_services = JSON.parse(process.env.VCAP_SERVICES);

//Parse CF VCAP_APPLICATION
var vcap_application = JSON.parse(process.env.VCAP_APPLICATION);

httpserver.get("/", function(req, res) {
  //Old school
  var reqData = url.parse(req.url, true).query;

  //On bloque le Cross-Origin Request
  res.setHeader('Content-type', 'text/html');
  res.setHeader('Access-Control-Allow-Origin', '*');
  
  var resp = 'Authent OK pour  '+reqData['login']; 
   
  res.sendfile('./static/index.html');
  console.log(time+"[ GET / ] Reponse : "+res);
});

httpserver.get("/services", function(req, res) {
   res.send(vcap_services);
});

httpserver.post("/application", function(req, res) {
   res.send(vcap_application);
});

//Write each request and response to STDOUT (to test CF logging aggregation feature).
//Example to STDERR ?

//Support CF & Legacy
httpserver.listen(process.env.PORT || 8000);
