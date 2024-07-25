package main

import (
	"log"
	"net/http"
)

func raiz (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pagina inicial"))
}

func home (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, mundo!"))
}	

func users (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Endpoint de usuários"))
}

func main() {
		// http é um pacote do Go que permite fazer requisições HTTP - é um protocolo de comunicaçao - base da comunicacao web

		//cliente: faz a requisicao - servidor - recebe a requisicao e devolve a resposta

		// GET, POST, PUT, DELETE, HEAD, OPTIONS, TRACE, CONNECT (metodos HTTP)

		// Request: o que eu quero fazer
		// Response: o que foi feito

		// Request: URL, método, cabeçalho, body
		// Response: status, cabeçalho, body

		//Rotas: caminhos que a requisição pode fazer:
		      //URI- identificador do recurso
		      //Método - GET, POST, PUT, DELETE
	
	// como fazer uma requisição HTTP em Go?
	// 1. Criar um cliente
	// 2. Fazer a requisição
	// 3. Tratar a resposta
		
	//como criar um servidor HTTP em Go?
	// 1. Definir as rotas
	// 2. Iniciar o servidor
	// 3. Tratar as requisições
	// 4. Enviar as respostas
	
	
	http.HandleFunc("/", raiz)
	http.HandleFunc("/home", home)
	http.HandleFunc("/users", users)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}
