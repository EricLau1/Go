package main

import(
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {

	fmt.Println("Listening port 8080")

	http.HandleFunc("/", handler)

	open("http://localhost:8080/")

	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello world!"))

}

func open(url string) error {
    var cmd string
    var args []string

    switch runtime.GOOS {
    case "windows":
        cmd = "cmd"
        args = []string{"/c", "start"}
    case "darwin":
        cmd = "open"
    default: // "linux", "freebsd", "openbsd", "netbsd"
        cmd = "xdg-open"
    }
    args = append(args, url)
    return exec.Command(cmd, args...).Start()
}