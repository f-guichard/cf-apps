package com.orange.fab;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestOperations;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

@RestController
public class HystrixController {

	@Value("${microservice.version}")
	private String microservice_version;
	
	@Value("${microservice.name}")
	private String microservice_name;

    @RequestMapping(method=RequestMethod.GET, value="/help")
    public ResponseEntity<String> help()
    {
    	HttpHeaders httpHeaders = new HttpHeaders();
        httpHeaders.set(microservice_name, microservice_version);
    	return new ResponseEntity<String>("Spring OK", httpHeaders, HttpStatus.OK);
    }
}
