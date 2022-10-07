// reference https://go.dev/doc/articles/wiki/
package main

// import declaration of fmt and os packages from Go standard library
import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/*
Define the Page structure with a title and body (i.e. how page data is
stored in memory).
The body element is a []byte  rather than a string because that is the type
expected by the io libraries we will use
*/
type Page struct {
	Title string
	Body  []byte
}

/*
Create a save method for persistent storage.
Signature = "This is a method named save that takes as its reciever p, a pointer
to Page. It takes no parameters, and returns a value of type error."
Reference https://go.dev/doc/articles/wiki/ for details on the function
definition.
*/
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

/*
define function for loading pages.
Reference https://go.dev/doc/articles/wiki/ for details on the function
definition.
*/
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

/*test procedure
func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
	http.HandlerFunc("/", handler)
	log.Faral(http.ListenAndServe(":8080", nii))
	http.HanleFunc("view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func handler(w http.ResponseWriter, r *https.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
*/

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</txtarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}

func saveHandler(w http) {

}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
