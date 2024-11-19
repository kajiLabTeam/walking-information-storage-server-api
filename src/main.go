package main

import (
	"fmt"
	"walking-information-storage-server-api/minio"
	"walking-information-storage-server-api/repository"
)

func main() {
	repository.GetPedestrian()
	minio.Getbuckets()
	buckets,err := minio.Getbuckets()
	if err != nil {
		fmt.Println(err)
	}
	for _, bucket := range buckets.Buckets {
		fmt.Println(*bucket.Name)
	}

	minio.ListObject("test")
	objects,err := minio.ListObject("test")
	if err != nil {
		fmt.Println(err)
	}	
	for _, object := range objects.Contents {
		fmt.Println(*object.Key)
	}

	minio.GetObject("test","acc.csv")
	object,err := minio.GetObject("test","acc.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(object)
}
