package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 0, "The port to test")
	flag.Parse()

	if *port == 0 {
		fmt.Printf("La porta specificata (%d) non è valida. Usa -port <numero> per impostarne una.\n", *port)
		os.Exit(1)
	}

	fmt.Printf(`Avvio il server sulla porta %d. Gli altri devono connettersi a http://<il tuo ip pubblico>:%d
Assicurati di aver aperto la porta che stai testando nel tuo modem
Puoi anche controllare che il mio programma funzioni andando su http://127.0.0.1:%d, se non vedi nulla c'è qualcosa che non va :)
`, *port, *port, *port)

	err := http.ListenAndServe(":"+fmt.Sprint(*port), http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		ip := GetIP(r)

		rw.WriteHeader(200)
		_, err := rw.Write([]byte(fmt.Sprintf(`
<head>
<meta charset="UTF-8" />
<meta http-equiv="X-UA-Compatible" content="IE=edge" />
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<title>Port Tester</title>
<style>
body{
font-family: system-ui, -apple-system, Segoe UI, Roboto, Ubuntu, Cantarell,
	Noto Sans, sans-serif, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif,
	'Apple Color Emoji', 'Segoe UI Emoji';
}
main{
	display: grid;
	place-content: center;
	text-align: center;
}
</style>
</head>
<body>
<main>
<h1>Funziona!</h1>
<p>Il tuo ip pubblico è %s</p>
</main>
</body>
`, ip)))

		if err != nil {
			panic(err)
		}

	}))

	if err != nil {
		panic(err)
	}

}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
