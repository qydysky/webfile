package main
import (
	"net/http"
    "strconv"
	xpart "github.com/qydysky/part"
	wt "github.com/qydysky/other/wt"
)

func main() {
	web :=  http.NewServeMux()

	wt := wt.WT()
	
	wt.WebRoot = "./html/"
	wt.SavePath = "save/"

	wt.Web(web)

	webAddr := "0.0.0.0"
	// xpart.Port().Set("wt",8089)//use random port by remove this line

	server := &http.Server{
		Addr:         webAddr+":"+strconv.Itoa(xpart.Port().Get("wt")),
		Handler:      web,
	}

	xpart.Logf().I("start:",server.Addr)
	xpart.Logf().I("open",server.Addr+"/wt/ to upload")
	server.ListenAndServe()

}