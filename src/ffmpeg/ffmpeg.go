package ffmpeg

import (
	"../helpers"
	"net/http"
	"strconv"
	"strings"
)

func VideoPreview320(fileName string, ffmpegFileNameWithoutExtension string, w http.ResponseWriter) string {

	// SET ffmpeg file real path for 320 video review
	ffmpegFileLink320 := "/tmp/ffmpeg_320_" + ffmpegFileNameWithoutExtension + ".mp4"
	// BEGIN 320: resize & crop video to 320, upload to aws s3
	args320 := []string{
		"-i",
		fileName,
		"-ss",
		"00:00:00",
		"-t",
		"00:00:06",
		"-f",
		"mp4",
		"-an",
		"-maxrate",
		"0.4M",
		"-bufsize",
		"0.4M",
		"-movflags",
		"faststart",
		"-vf",
		"scale='min(min(320,iw), ih)*iw/min(ih,iw)':'min(min(320,iw), ih)*ih/min(ih,iw)'",
		"-y",
		ffmpegFileLink320}

	_, err320 := helpers.RunCMD("ffmpeg", args320, false)

	if err320 != nil {
		helpers.GetError(w, "Can't create 320 video preview")
		return "Error"
	}

	return ffmpegFileLink320

}

func VideoPreviewThumbnail320(fileName string, ffmpegFileNameWithoutExtension string, w http.ResponseWriter) string {

	// SET ffmpeg file real path for 320 video thumbnail
	ffmpegFileLinkThumb320 := "/tmp/ffmpeg_320_thumb_" + ffmpegFileNameWithoutExtension + ".jpg"

	// GET video duration for creating thumbnail from center
	duration := GetVideoDuration(fileName, w)

	// BEGIN THUMB 320: resize & crop video to 320, upload to aws s3
	argsThumb320 := []string{
		"-i",
		fileName,
		"-ss",
		duration,
		"-vframes",
		"1",
		ffmpegFileLinkThumb320}

	_, errThumb320 := helpers.RunCMD("ffmpeg", argsThumb320, false)

	if errThumb320 != nil {
		helpers.GetError(w, "Can't create 320 thumbnail preview")
		return "Error"
	}

	return ffmpegFileLinkThumb320
}

func VideoMain1280(fileName string, ffmpegFileNameWithoutExtension string, w http.ResponseWriter) string {

	// SET ffmpeg file real path for 1280 main video
	ffmpegFileLink1280 := "/tmp/ffmpeg_1280_" + ffmpegFileNameWithoutExtension + ".mp4"

	// BEGIN 1280: resize & crop video to 1280, upload to aws s3
	args1280 := []string{
		"-i",
		fileName,
		"-ss",
		"00:00:00",
		"-t",
		"00:10:00",
		"-vf",
		"scale='min(max(ih,iw), 1280)*iw/max(ih,iw)':'min(max(ih,iw), 1280)*ih/max(ih,iw)'",
		"-f",
		"mp4",

		"-vcodec",
		"h264",
		"-acodec",
		"aac",

		"-acodec",
		"mp3",
		"-ar",
		"44100",
		"-ab",
		"64k",

		"-y",
		ffmpegFileLink1280}

	_, err1280 := helpers.RunCMD("ffmpeg", args1280, false)

	if err1280 != nil {
		helpers.GetError(w, "Can't create 1280 video")
		return "Error"
	}

	return ffmpegFileLink1280
}

func VideoPreview480(fileName string, ffmpegFileNameWithoutExtension string, w http.ResponseWriter) string {

	// SET ffmpeg file real path for 480 video review
	ffmpegFileLink480 := "/tmp/ffmpeg_480_" + ffmpegFileNameWithoutExtension + ".mp4"
	// BEGIN 480: resize & crop video to 480, upload to aws s3
	args480 := []string{
		"-i",
		fileName,
		"-ss",
		"00:00:00",
		"-t",
		"00:00:06",
		"-f",
		"mp4",
		"-an",
		"-maxrate",
		"0.4M",
		"-bufsize",
		"0.4M",
		"-movflags",
		"faststart",
		"-vf",
		"scale='min(max(ih,iw), 480)*iw/max(ih,iw)':'min(max(ih,iw), 480)*ih/max(ih,iw)'",
		"-y",
		ffmpegFileLink480}

	_, err480 := helpers.RunCMD("ffmpeg", args480, false)

	if err480 != nil {
		helpers.GetError(w, "Can't create 480 video preview")
		return "Error"
	}

	return ffmpegFileLink480

}

func VideoPreviewThumbnail480(fileName string, ffmpegFileNameWithoutExtension string, w http.ResponseWriter) string {

	// SET ffmpeg file real path for 480 video thumbnail
	ffmpegFileLinkThumb480 := "/tmp/ffmpeg_480_thumb_" + ffmpegFileNameWithoutExtension + ".jpg"

	// GET video duration for creating thumbnail from center
	duration := GetVideoDuration(fileName, w)

	// BEGIN THUMB 480: resize & crop video to 480, upload to aws s3
	argsThumb480 := []string{
		"-i",
		fileName,
		"-ss",
		duration,
		"-vframes",
		"1",
		ffmpegFileLinkThumb480}

	_, errThumb480 := helpers.RunCMD("ffmpeg", argsThumb480, false)

	if errThumb480 != nil {
		helpers.GetError(w, "Can't create 480 thumbnail preview")
		return "Error"
	}

	return ffmpegFileLinkThumb480
}

func GetVideoDuration(fileName string, w http.ResponseWriter) string {

	argsSizeThumb480 := []string{
		"-v",
		"error",
		"-show_entries",
		"format=duration",
		"-of",
		"default=noprint_wrappers=1:nokey=1",
		fileName}

	sizeThumb, errSizeThumb480 := helpers.RunCMD("ffprobe", argsSizeThumb480, false)
	if errSizeThumb480 != nil {
		helpers.GetError(w, "Can't get duration for 480 thumbnail preview")
		return "Error"
	}

	clearSizeThumb := strings.Replace(sizeThumb, "\n", "", -1)
	floatSize, errFloatSize := strconv.ParseFloat(clearSizeThumb, 64)
	if errFloatSize != nil {
		helpers.GetError(w, "Can't convert duration to float!")
		return "Error"
	}

	newFloatSize := floatSize/2
	duration := "00:00:0"+strconv.Itoa(int(newFloatSize))

	return duration
}