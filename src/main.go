package main

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/zono-dev/stplib"
)

// SetImageInfo sets ResizedInfo and orgpath in ImgInfo
func SetImageInfo(r ResizedInfo, orgpath string) *stplib.ImgInfo {
	i := stplib.ImgInfo{}
	i.FileName = r.FileName
	i.CreatedAt = r.CreatedAt
	i.OrgPath = orgpath
	i.ResizedFilePath = r.ResizedFilePath
	i.FileType = r.TypeOfFile
	i.SizeX = r.SizeX
	i.SizeY = r.SizeY
	return &i
}

// CheckEvent is handler of the Lambda.
// It a S3 PUT event data and resize the image in the S3 bucket,
// then put new image file that is resized in S3 path.
// Finally it write the image info in DynamoDB table.
func CheckEvent(event events.S3Event) {
	env := Setup()
	for _, i := range event.Records {
		if env.Vlog {
			PrintS3EventRecord(i)
		} else {
			MyPrintf("EventTime", i.EventTime.Format(time.RFC822))
			MyPrintf("Name", i.S3.Bucket.Name)
			MyPrintf("Key", i.S3.Object.Key)
			MyPrintf("Size", strconv.FormatInt(i.S3.Object.Size, 10))
		}
		imginfo, err := ResizeImage(i.S3.Bucket.Name, i.S3.Object.Key, "resize", env.Prefix, env.ResizeX, env.ResizeY)
		if err != nil {
			MyPrintErr(err)
		} else {
			MyPrintf("ResizedFilePath", imginfo.JsonStr())
		}

		// Write info in DynamoDB Table
		i := SetImageInfo(imginfo, i.S3.Object.Key)
		err = WriteImgInfo(env.Dtable, i)
		if err != nil {
			MyPrintErr(err)
		}
	}
}

// HandleRequest is the starting point of Lambda process.
func HandleRequest(ctx context.Context, event events.S3Event) (string, error) {
	CheckEvent(event)
	return "Complete.", nil
}

// main starts Lambda.
func main() {
	lambda.Start(HandleRequest)
}
