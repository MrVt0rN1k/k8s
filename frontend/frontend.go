package main

import (
	"html/template"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

const htmlTemplate = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Frontend Service</title>
		<script>
			function getMessage() {
				fetch('http://localhost:8080/api/message') 
					.then(response => response.json())
					.then(data => {
						document.getElementById('message').innerText = data.text;
					})
					.catch(error => {
						console.error('Error:', error);
						document.getElementById('message').innerText = 'Error fetching message';
					});
			}
		</script>
	</head>
	<body>
		<h1>Welcome to the Frontend Service</h1>
		<button onclick="getMessage()">Get Message from Backend</button>
		<p id="message"></p>	
	</body>
	</html>
	`

func main() {
	http.HandleFunc("/", handleRoot)

	log.Println("Starting frontend server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
}
