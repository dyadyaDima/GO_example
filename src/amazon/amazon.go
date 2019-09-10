package amazon

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func UploadByPresignedUrl(PresignedUrl string, fileForUpload string, mimetype string) (bool, string) {

    videoFile, err := os.Open(fileForUpload)
    if err != nil {
	return true, "Can't open uploaded file:"+fileForUpload
    }
    defer videoFile.Close()

    // create a new buffer base on file size
    fileInfo, _ := videoFile.Stat()
    var fileSize int64 = fileInfo.Size()
    buffer := make([]byte, fileSize)

    // read file content to buffer
    videoFile.Read(buffer)
    fileBytes := bytes.NewReader(buffer) // convert to io.ReadSeeker type

    req, err := http.NewRequest("PUT", PresignedUrl, fileBytes)
    if err != nil {
	return true, "Can't execute PUT request to AWS!"
    }

    req.Header.Set("Content-Type", mimetype)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
	return true, "Can't execute request!"
    }
    defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)
    fmt.Println(string(body))

    return false, ""
}