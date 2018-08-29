package main

import (
	"net/http"
	"io"
	"os"
	"os/exec"
)

func streaming(w http.ResponseWriter, r *http.Request) {
	io.Copy(w, os.Stdin)
}

func main() {
	http.HandleFunc("/", streaming)


	cmd := exec.Command("xdg-open", "http://localhost:8080")
	cmd.Run()

	http.ListenAndServe(":8080", nil)
}
