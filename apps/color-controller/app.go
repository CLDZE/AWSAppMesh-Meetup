package main
 
import (
	"os"
	"fmt"
        "time"
	"net/http"
        "github.com/aws/aws-xray-sdk-go/xray"
        
	// "github.com/golangci/golangci-lint/pkg/exitcodes"
)

type Style struct {
  data  string `json:"style"`
}

func main() {

    xray.Configure(xray.Config{
      LogLevel: "warn",
      LogFormat: "[%Level] [%Time] %Msg%n",
    })
        
	c := os.Getenv("COLOR")
	if len(c) == 0{
		os.Setenv("COLOR", "lightblue")//"#F1A94E") //Blue 44B3C2 and Yellow F1A94E 
	}  

	http.Handle("/", xray.Handler(xray.NewFixedSegmentNamer("color-" + os.Getenv("COLOR")), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
               
               ////print request Headers
               // for name, values := range r.Header {
               //   for _, value := range values {
               //     fmt.Println(name, value)
               //   }
               // }
               enableCors(&w)
               fmt.Fprint(w,os.Getenv("COLOR"))
			   fmt.Println("Got request at "+time.Now().Format(time.ANSIC))
	})))


	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html> DASHBOARD Requested: %s\n </html>", r.URL.Path)
	})

	http.HandleFunc("/die", func(w http.ResponseWriter, r *http.Request) {
		die();
	})
        fmt.Println("Running. Listening on 8080")
	http.ListenAndServe(":8080", nil)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func die() {
	os.Exit(3)
}
