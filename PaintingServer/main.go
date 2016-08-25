package main
import (
	"fmt"
	"log"
	"net/http"
	"os"
        "image"
        "image/color"
        "image/gif"
        "math"
        "math/rand"
)
func main() {
        //Support Dual Stack CFY && Legacy
        var port string = os.Getenv("PORT")
        if port == "" {
		port = "8800"
	}
        fmt.Fprintf(os.Stdout, "APP Port is %s \n", port)
	http.HandleFunc("/", drawsomething)
        http.HandleFunc("/help", helper)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func drawsomething(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "image/gif")
        var palette = []color.Color{color.White, color.Black}
        const (
                cycles  = 5     
                res     = 0.001
                size    = 100
                nframes = 120
                delay   = 2 
        )
        freq := rand.Float64() * 3.0
        anim := gif.GIF{LoopCount: nframes}
        phase := 0.0
        for i := 0; i < nframes; i++ {
                rect := image.Rect(0, 0, 2*size+1, 2*size+1)
                img := image.NewPaletted(rect, palette)
                for t := 0.0; t < cycles*2*math.Pi; t += res {
                        x := math.Sin(t)
                        y := math.Sin(t*freq + phase)
                        img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
                                1)
                }
                phase += 0.1
                anim.Delay = append(anim.Delay, delay)
                anim.Image = append(anim.Image, img)
        }
        gif.EncodeAll(w, &anim)
}

func helper(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "Usage : /")
}
