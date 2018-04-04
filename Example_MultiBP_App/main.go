package main
import (
	"fmt"
	"log"
	"net/http"
	"os"
        "os/exec"
        "io/ioutil"
)
//Idea : python requests to get token, java RestTemplate for api exchanges and go for api controller + IPC Sync
func main() {
        //Dual Stack support (CFY && Legacy)
        var port string = os.Getenv("PORT")
        if port == "" {
		port = "8800"
	}
        fmt.Fprintf(os.Stdout, "APP Port is %s \n", port)
	go func(val int8) {
        	log.Println("VCAP_APPLICATION imported : "+os.Getenv("VCAP_APPLICATION"))
        }(127)
	/* Get it done right ....... */
	go func(val int8) {
                log.Println("Launch Spring Boot App")
                path, err := exec.LookPath("java")
                if err != nil {
                  log.Println("no java in PATH")
                  return
                }
                cmd := exec.Command(path, "-jar","java/target/springapp.jar")
                stdoutStderr, err := cmd.CombinedOutput()
                if err != nil {
                  log.Println("Erreur : ", err)
                }
                log.Println(string(stdoutStderr[:]))
        }(126)
        http.HandleFunc("/", helper)
        http.HandleFunc("/python", python)
        http.HandleFunc("/java", java)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func helper(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "To Do")
}

//Slide with go func {} (+chan to minimize socket usage) 
func python(w http.ResponseWriter, r *http.Request) {
        log.Println("Calling /python")
        path, err := exec.LookPath("python")
	if err != nil {
	  log.Println("no python in PATH")
          fmt.Fprintf(w, "no python in PATH")
          return
	}
        cmd := exec.Command(path, "python/hello.py")
        stdoutStderr, err := cmd.CombinedOutput()
        if err != nil { 
          log.Println("Erreur : ", err)
        } 
        fmt.Fprintf(w, string(stdoutStderr[:]))
}

func java(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling /java")
        //REquest spring boot app on 127.0.0.1:23001/help (defined via hardcoded server.port)
        resp, err := http.Get("http://127.0.0.1:23001/help")
        if err != nil {
          log.Println("Erreur : ", err)
        }
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
          log.Println("Erreur : ", err)
        }
        bodyString := string(body)
        fmt.Fprintf(w, bodyString)
}
