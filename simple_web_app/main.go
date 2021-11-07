package main
import (
"net/http"
"io"
)

/* form to get user input request */
const form = `<html><body><form action="#" method="post" name="bar">
<input type="text" name="in"/>
<input type="submit" value="Submit"/>
</form></html></body>`

/* handle a simple get request */
func SimpleServer(w http.ResponseWriter, request *http.Request) {
  
}

/* handle a form, both the GET which displays the form
and the POST which processes it.*/
func FormServer(w http.ResponseWriter, request *http.Request) {


  switch request.Method {
      /* display the form to the user */

 
      /* handle the form data */
  }
}

func main() {
 
}