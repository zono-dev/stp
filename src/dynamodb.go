package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/zono-dev/stplib"
)

// WriteImgInfo write the data of a image which was uploaded in dtable.
func WriteImgInfo(dtable string, imginfo *stplib.ImgInfo) error {
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
	table := db.Table(dtable)

	err := table.Put(imginfo).Run()

	if err != nil {
		MyPrintErr(err)
		return err
	}
	return nil
}
