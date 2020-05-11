package main

import (
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

func PrintS3UserIdentity(eu events.S3UserIdentity) {
	MyPrintf("PrincipalID", eu.PrincipalID)
}

func PrintS3RequestParameters(er events.S3RequestParameters) {
	MyPrintf("SourceIPAddress", er.SourceIPAddress)
}

func PrintResponseElements(er map[string]string) {
	for k, v := range er {
		MyPrintf(k, v)
	}
}

func PrintS3Bucket(eb events.S3Bucket) {
	MyPrintf("Name", eb.Name) // bucket name
	PrintS3UserIdentity(eb.OwnerIdentity)
	MyPrintf("Arn", eb.Arn)
}

func PrintS3Object(eb events.S3Object) {
	MyPrintf("Key", eb.Key) // path to image
	MyPrintf("Size", strconv.FormatInt(eb.Size, 10))
	MyPrintf("URLDecodedKey", eb.URLDecodedKey)
	MyPrintf("VersionID", eb.VersionID)
	MyPrintf("ETag", eb.ETag)
	MyPrintf("Sequencer", eb.Sequencer)
}

func PrintS3Entity(ee events.S3Entity) {
	MyPrintf("SchemaVersion", ee.SchemaVersion)
	MyPrintf("ConfigurationID", ee.ConfigurationID)
	PrintS3Bucket(ee.Bucket)
	PrintS3Object(ee.Object)
}

func PrintS3EventRecord(er events.S3EventRecord) {
	MyPrintf("EventVersion", er.EventVersion)
	MyPrintf("EventSource", er.EventSource)
	MyPrintf("AWSRegion", er.AWSRegion)
	MyPrintf("EventTime", er.EventTime.Format(time.RFC822))
	MyPrintf("EventName", er.EventName)
	MyPrintf("PrincipalID", "")
	PrintS3UserIdentity(er.PrincipalID)
	MyPrintf("RequestParameters", "")
	PrintS3RequestParameters(er.RequestParameters)
	MyPrintf("ResponseElements", "")
	PrintResponseElements(er.ResponseElements)
	MyPrintf("S3", "")
	PrintS3Entity(er.S3)
}
