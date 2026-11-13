package main

import (
	"fmt"      // Scrittura del messaggio nella risposta HTTP
	"log"      // Messaggi di log e gestione dell’errore del server
	"net/http" // Server HTTP e gestione delle richieste
	"os"       // Accesso alle variabili d’ambiente
)

func main() {
	// Legge la variabile d’ambiente una sola volta, all’avvio.
	// Nel container il valore arriva da ENV oppure da docker run -e.
	messaggio := os.Getenv("MESSAGGIO")

	// Usa questo valore se la variabile è assente oppure vuota.
	if messaggio == "" {
		messaggio = "Laboratorio Docker con Go"
	}

	// Registra una funzione sul router HTTP predefinito.
	// "/" intercetta anche i percorsi senza un handler più specifico.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w permette di costruire la risposta al client.
		// r contiene la richiesta ricevuta: metodo, URL, header, ecc.
		// La funzione accede a messaggio definito nel main: è una closure.
		// Fprintln scrive il messaggio nel corpo della risposta,
		// aggiungendo un newline; lo stato HTTP predefinito è 200 OK.
		fmt.Fprintln(w, "Server Go:", messaggio)
//		fmt.Fprintln(w, messaggio)
	})

	// Il logger standard scrive su stderr: visibile con docker logs.
	log.Println("Server in ascolto sulla porta 8000")

	// ":8000" ascolta sulla porta 8000 senza limitarsi al localhost.
	// nil seleziona il router predefinito usato da http.HandleFunc.
	// ListenAndServe resta in esecuzione e gestisce le richieste;
	// se restituisce un errore, log.Fatal lo stampa e termina il processo.
	log.Fatal(http.ListenAndServe(":8000", nil))
}