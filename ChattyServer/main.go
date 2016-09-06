package main
import (
	"fmt"
	"log"
	"net/http"
	"os"
)
func main() {
        //Dual Stack support (CFY && Legacy)
        var port string = os.Getenv("PORT")
        if port == "" {
		port = "8800"
	}
        fmt.Fprintf(os.Stdout, "APP Port is %s \n", port)
	http.HandleFunc("/", telleverything)
        http.HandleFunc("/help", info)
        http.HandleFunc("/info", info)
        http.HandleFunc("/headers", headers)
        http.HandleFunc("/queries", queries)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func info(w http.ResponseWriter, r *http.Request) {
       w.Header().Set("JOX-TYPE", "INFO")
       fmt.Fprintf(w, "Usage : \n/ => Tout avoir\n/headers => Obtenir tous les headers de la query\n/queries => Obtenir tous les params de la query\n")
} 
func headers(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Nom header => Valeur header\n")
        for k, v := range r.Header {
                fmt.Fprintf(w, "%q => %q\n", k, v)
        }
}
func queries(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
        fmt.Fprintf(w, "Host = %q\n", r.Host)
        fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
        if err := r.ParseForm(); err != nil {
                log.Print(err)
        }
        fmt.Fprintf(w, "Nom param => Valeur param\n")
        for k, v := range r.Form {
                fmt.Fprintf(w, "%q => %q\n", k, v)
        }
}
func telleverything(w http.ResponseWriter, r *http.Request) {
        headers(w,r)
        queries(w, r)
}
    
