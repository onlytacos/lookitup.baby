package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("HONKS ONLINE ON PORT %v", port)
	
	http.Handle("/kylie", http.RedirectHandler("https://twitter.com/kyliebytes/status/1335335986583142402", 302))
	http.Handle("/ian", http.RedirectHandler("https://twitter.com/Dixie3Flatline/status/1334701146703679488", 302))
	http.Handle("/pop", http.RedirectHandler("https://www.youtube.com/watch?v=EDUy3Y_w9Tk", 302))
	http.Handle("/cloud", http.RedirectHandler("https://twitter.com/IanColdwater/status/1327350827078406144", 302))
	http.Handle("/kubecon", http.RedirectHandler("https://www.youtube.com/watch?v=ZY4kTDCsJ44", 302))
	http.Handle("/betty", http.RedirectHandler("https://twitter.com/BettyJunod/status/1109186194959552513", 302))
	http.Handle("/defcon", http.RedirectHandler("https://twitter.com/fox0x01/status/1344775712511819778", 302))
	http.Handle("/", http.RedirectHandler("https://twitter.com/IanColdwater/status/1292895288546545666", 302))


	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
