package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// TODO: URLParam
	name := r.URL.Query().Get("name")

	if name != "" {
		fmt.Fprintf(w, `<h1>Welcome, %s!</h1>`, name)
	} else {
		fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Contact Page</h1>
	<p>
		To get in touch, email me at 
		<a href=\"mailto:umishra1504@gmail.com\">umishra1504@gmail.com</a>.
	</p>`)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>FAQ Page</h1>
	<p>Some Frequently Asked Questions:</p>
	<br />
	<ul>
		<li>
			<p><b>Q: Is there a free version?</b></p>
			<p>A: Yes! We offer a free trial for 30 days on any paid plans</p>
		</li>
		<li>
			<p><b>Q: What are your support hours?</b></p>
			<p>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</p>
		</li>
		<li>
			<p><b>Q: How do I contact support?</b></p>
			<p>A: Email us - <a href="mailto:umishra1504@gmail.com">umishra1504@gmail.com</p>
		</li>
	</ul>
	`)
}

func main() {
	r := chi.NewRouter()
	// TODO: Use logger middleware
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Server starting on 3000...")
	http.ListenAndServe(":3000", r)
}
