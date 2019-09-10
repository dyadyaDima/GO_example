package media

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"../amazon"
	"../ffmpeg"
	"../helpers"
)

type Media struct {
    ID				string `json:"id" validate:"required"`
    Name			string `json:"name"`
    Mimetype			string `json:"mimetype"`
    Size			string `json:"size"`
    File			http.File `json:"file"`
    PresignedUrl320		string `json:"presigned_url_320"`
    PresignedUrlThumb320	string `json:"presigned_url_thumb_320"`
    PresignedUrl1280		string `json:"presigned_url_1280"`
}

var mediaFile Media

func Main(w http.ResponseWriter, r *http.Request) {

    mediaFile.ID = r.FormValue("id")
    mediaFile.Name = r.FormValue("name")
    mediaFile.Mimetype = r.FormValue("mimetype")
    mediaFile.Size = r.FormValue("size")
    mediaFile.PresignedUrl320 = r.FormValue("presigned_url_320")
    mediaFile.PresignedUrlThumb320 = r.FormValue("presigned_url_thumb_320")
    mediaFile.PresignedUrl1280 = r.FormValue("presigned_url_1280")

    // FormFile returns the first file for the given key `File`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("file")
    if err != nil {
		helpers.GetError(w, "Error Retrieving the File:"+handler.Filename)
		return
    }
    defer file.Close()

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    tempFile, err := ioutil.TempFile("/tmp", "*_"+mediaFile.Name)
    if err != nil {
		helpers.GetError(w, "Can't create temporary file")
		return
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
		helpers.GetError(w, "Can't read uploaded file")
		return
    }

    // write this byte array to our temporary file
    tempFile.Write(fileBytes)

    // SET ffmpeg file name
    ffmpegFilePathArray := strings.Split(tempFile.Name(), "/")
    ffmpegFileName := ffmpegFilePathArray[len(ffmpegFilePathArray)-1]
    ffmpegFileNameArray := strings.Split(ffmpegFileName, ".")
    ffmpegFileNameWithoutExtension := ffmpegFileNameArray[len(ffmpegFileNameArray)-2]

    // BEGIN 320: resize & crop video to 320, upload to aws s3
    ffmpegFileLink320 := ffmpeg.VideoPreview320(tempFile.Name(), ffmpegFileNameWithoutExtension, w)
    uploadErrorVideo320, uploadErrorVideo320msg := amazon.UploadByPresignedUrl(mediaFile.PresignedUrl320, ffmpegFileLink320, "video/mp4")
    if uploadErrorVideo320 {
		helpers.GetError(w, uploadErrorVideo320msg)
		return
    }
    // END 320

    // BEGIN THUMB 320: resize & crop video to 320, upload to aws s3
    ffmpegFileLinkThumb320 := ffmpeg.VideoPreviewThumbnail320(ffmpegFileLink320, ffmpegFileNameWithoutExtension, w)
    uploadErrorThumb320, uploadErrorThumb320msg := amazon.UploadByPresignedUrl(mediaFile.PresignedUrlThumb320, ffmpegFileLinkThumb320, "image/jpeg")
    if uploadErrorThumb320 {
		helpers.GetError(w, uploadErrorThumb320msg)
		return
    }
    // END THUMB 320

    // BEGIN 1280: resize & crop video to 1280, upload to aws s3
	ffmpegFileLink1280 := ffmpeg.VideoMain1280(tempFile.Name(), ffmpegFileNameWithoutExtension, w)
    uploadErrorVideo1280, uploadErrorVideo1280msg := amazon.UploadByPresignedUrl(mediaFile.PresignedUrl1280, ffmpegFileLink1280, "video/mp4")
    if uploadErrorVideo1280 {
		helpers.GetError(w, uploadErrorVideo1280msg)
		return
    }
    // END 1280

    errSize, errSizeMsg, size := helpers.GetFileSize("/tmp/ffmpeg_1280_"+ffmpegFileNameWithoutExtension+".mp4", w)
    if errSize {
		helpers.GetError(w, errSizeMsg)
		return
    }

    result := map[string]interface{}{
		"Error": "",
		"Result": map[string]interface{}{
			mediaFile.ID: map[string]interface{}{
			"name": mediaFile.Name,
			"mimetype": "video/mp4",
			"size": size,
			},
	},
	}

    // REMOVE FILES
    helpers.RemoveFile(tempFile.Name(), w)
    helpers.RemoveFile(ffmpegFileLink320, w)
    helpers.RemoveFile(ffmpegFileLinkThumb320, w)
    helpers.RemoveFile(ffmpegFileLink1280, w)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}