package helpers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func RunCMD(path string, args []string, debug bool) (out string, err error) {
	cmd := exec.Command(path, args...)

	var b []byte
	b, _ = cmd.CombinedOutput()
	out = string(b)

	if debug {
		fmt.Println(strings.Join(cmd.Args[:], " "))

		if err != nil {
			fmt.Println("RunCMD ERROR")
			fmt.Println(out)
		}
	}

	return
}

func RemoveFile(fileLink string, w http.ResponseWriter) {
	err := os.Remove(fileLink)
	if err != nil {
		GetError(w, "Can't delete file: "+fileLink)
		return
	}
}

func GetFileSize(fileLink string, w http.ResponseWriter) (bool, string, int64) {

	// open file
	file, err := os.Open(fileLink)
	if err != nil {
		return true, "Can't open file:" + fileLink, 0
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return true, "Can't get file size for:" + fileLink, 0
	}

	return false, "", stat.Size()
}

func GetError(w http.ResponseWriter, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	str := `{"Error": "` + err + `","Result": ""}`
	fmt.Fprint(w, str)
	return
}
