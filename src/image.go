package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/disintegration/imaging"
)

// base file name
const basefilename = "%s%s.%s"

// get random string
const rsLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type ResizedInfo struct {
	FileName        string    `json:"FileName"`
	ResizedFilePath string    `json:"ResizedFilePath"`
	TypeOfFile      string    `json:"TypeOfFile"`
	SizeX           int       `json:"SizeX"`
	SizeY           int       `json:"SizeY"`
	CreatedAt       time.Time `json:"CreateAt"`
}

func (r ResizedInfo) JsonStr() string {
	json, err := json.Marshal(r)
	if err != nil {
		MyPrintErr(err)
		return ""
	}
	return string(json)
}

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = rsLetters[rand.Intn(len(rsLetters))]
	}
	return string(b)
}

// Calculate resize parcentage
func ResizePercentage(sx, sy, rex, rey int) float64 {
	px := float64(rex) / float64(sx)
	py := float64(rey) / float64(sy)

	if px <= py {
		return px
	}
	return py
}

// CreateNewImage
func CreateNewImage(srcImg image.Image, dsizex, dsizey int, t imaging.Format) (string, error) {
	MyPrintf("dsizex", strconv.Itoa(dsizex))
	MyPrintf("dsizey", strconv.Itoa(dsizey))

	tmppath := "/tmp/tmp-image"
	switch t {
	case imaging.JPEG:
		tmppath += ".jpg"
	case imaging.GIF:
		tmppath += ".gif"
	case imaging.PNG:
		tmppath += ".png"
	}

	imgDst := imaging.Resize(srcImg, dsizex, dsizey, imaging.Lanczos)
	if err := imaging.Save(imgDst, tmppath); err != nil {
		MyPrintErr(err)
		return "", err
	}

	return tmppath, nil
}

func PutFileToS3(s3svc *s3.S3, bucket, path, fname string, tmppath string) (string, error) {

	f, err := os.Open(tmppath)
	defer f.Close()
	if err != nil {
		MyPrintErr(err)
		return "", err
	}

	fi, _ := f.Stat()
	size := fi.Size()
	buffer := make([]byte, size)

	f.Read(buffer)

	fbytes := bytes.NewReader(buffer)
	fType := http.DetectContentType(buffer)
	MyPrintf("DetectContetnType", fType)

	upath := path + "/" + fname

	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(upath),
		Body:          fbytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fType),
	}
	resp, err := s3svc.PutObject(params)
	if err != nil {
		MyPrintErr(err)
	} else {
		fmt.Println(resp)
	}
	return upath, err
}

func CreateNewFileName(t imaging.Format, prefix string) string {
	rand.Seed(time.Now().UnixNano())
	switch t {
	case imaging.JPEG:
		return (fmt.Sprintf(basefilename, prefix, RandString(16), "jpg"))
	case imaging.GIF:
		return (fmt.Sprintf(basefilename, prefix, RandString(16), "gif"))
	case imaging.PNG:
		return (fmt.Sprintf(basefilename, prefix, RandString(16), "png"))
	}
	return ""
}

func ResizeImage(bucket, srcpath, dstpath, prefix string, rex, rey int) (ResizedInfo, error) {
	out := ResizedInfo{"", "", "", 0, 0, time.Now()}
	sess := session.Must(session.NewSession())

	s3svc := s3.New(sess)

	srcObject, err := s3svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(srcpath),
	})

	if err != nil {
		MyPrintErr(err)
		return out, err
	}
	defer srcObject.Body.Close()

	// Decode Object into image.Image
	img, err := imaging.Decode((srcObject.Body), imaging.AutoOrientation(true))
	if err != nil {
		MyPrintErr(err)
		return out, err
	}
	// Gussing format from filename
	t, err := imaging.FormatFromFilename(srcpath)
	if err != nil {
		MyPrintErr(err)
		return out, err
	}
	MyPrintf("TypeOfImage", t.String())

	// get size of a target image
	rct := img.Bounds()
	sx := rct.Dx()
	sy := rct.Dy()
	MyPrintf("Width", strconv.Itoa(rct.Dx()))
	MyPrintf("Height", strconv.Itoa(rct.Dy()))

	p := ResizePercentage(sx, sy, rex, rey)
	out.SizeX = int(float64(sx) * p)
	out.SizeY = int(float64(sy) * p)

	tmppath, err := CreateNewImage(img, out.SizeX, out.SizeY, t)
	if err != nil {
		return out, err
	}

	out.FileName = CreateNewFileName(t, prefix)
	out.ResizedFilePath, err = PutFileToS3(s3svc, bucket, dstpath, out.FileName, tmppath)
	out.CreatedAt = time.Now()
	out.TypeOfFile = t.String()

	return out, err
}
