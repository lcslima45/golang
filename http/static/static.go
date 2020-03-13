<<<<<<< HEAD


=======
>>>>>>> 694a4309f7f8aa32cce60f1a28400094a1f7aace
package main
import (
"log"
"net/http"
)
func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
<<<<<<< HEAD
}
=======
}
>>>>>>> 694a4309f7f8aa32cce60f1a28400094a1f7aace
