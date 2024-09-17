package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("HONKS ONLINE ON PORT %v", port)

	http.Handle("/kylie", http.RedirectHandler("https://www.businessinsider.com/author/kylie-robison", 302))
	http.Handle("/ian", http.RedirectHandler("https://twitter.com/Dixie3Flatline/status/1334701146703679488", 302))
	http.Handle("/cloud", http.RedirectHandler("https://twitter.com/IanColdwater/status/1327350827078406144", 302))
	http.Handle("/kubecon", http.RedirectHandler("https://www.youtube.com/watch?v=ZY4kTDCsJ44", 302))
	http.Handle("/betty", http.RedirectHandler("https://twitter.com/BettyJunod/status/1109186194959552513", 302))
	http.Handle("/defcon", http.RedirectHandler("https://twitter.com/fox0x01/status/1344775712511819778", 302))
	http.Handle("/elle", http.RedirectHandler("https://twitter.com/ElleArmageddon/status/1349197020913827840", 302))
	http.Handle("/acab", http.RedirectHandler("https://twitter.com/ElleArmageddon/status/1333995231063130115", 302))
	http.Handle("/k8s-networking", http.RedirectHandler("https://www.oreilly.com/library/view/kubernetes-networking/9781492081647", 302))
	http.Handle("/poggers", http.RedirectHandler("https://twitter.com/Dixie3Flatline/status/1374538402989703170", 302))
	http.Handle("/coq", http.RedirectHandler("https://twitter.com/TaliaRinger/status/1453348584758390786", 302))
	http.Handle("/grady-uml", http.RedirectHandler("https://x.com/grady_booch/status/1697450877676892455", 302))
	http.HandleFunc("/hachyderm/", Hachyderm)
	http.Handle("/", http.RedirectHandler("https://web.archive.org/web/20200810190413/https://twitter.com/IanColdwater/status/1292895288546545666", 302))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func Hachyderm(res http.ResponseWriter, req *http.Request) {
	if strings.Contains(req.UserAgent(), "Twitterbot") {
		log.Print("twitterbot attempting to access hachyderm, nice try elmo")
		res.WriteHeader(http.StatusTeapot)
		return
	}

	p := strings.Replace(req.URL.Path, "/hachyderm", "", 1)

	http.Redirect(res, req, "https://hachyderm.io"+p, http.StatusFound)
}
