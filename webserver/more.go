package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey there</h1>")
	fmt.Fprintf(w, "<p>Sunt do laboris mollit do ad id sunt cillum esse tempor in in consectetur. Est ipsum laborum laborum amet ut duis. Ad voluptate officia elit in mollit laboris anim incididunt incididunt. Reprehenderit adipisicing excepteur pariatur dolore in velit do ea amet consectetur nisi et. Tempor exercitation anim adipisicing cillum ipsum eu ut veniam veniam magna. Exercitation excepteur occaecat non eu adipisicing. In amet enim occaecat dolore nostrud veniam nulla nostrud esse minim irure dolor culpa ad.</p>")
	fmt.Fprintf(w, "<p>Go is fast<p>")
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>VARIABLES</strong>")

	// print on multiple lines
	fmt.Fprintf(w, `<p>First line</p>
		<p>Line 2</p>`)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
