package pageHandler

import (
	"io/ioutil"
	"net/http"
	"ticketBackend/sessionHandler"
)

// localhost:.../saveTicket
// Speichert den Text aus den Textareas in mail, subject, text
func SaveTicketHandler(response http.ResponseWriter, request *http.Request) {
	if sessionHandler.IsUserLoggedIn(request) {
		mail := request.FormValue("ticketMail")
		subject := request.FormValue("ticketSubject")
		text := request.FormValue("ticketText")

		// Testdatei für die Eingabe
		inputTest := string("Mail\n" + mail + "\n\nSubject\n" + subject + "\n\nText\n" + text)
		ioutil.WriteFile("./assets/TicketTest", []byte(inputTest), 0600)

		// Zurück zu der Ticketseite
		http.Redirect(response, request, "/ticket", http.StatusFound)

	} else {
		http.Redirect(response, request, "/", 302)
	}
}