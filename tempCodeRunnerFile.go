result, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, bucket := range result.Buckets {
		fmt.Printf("%s\n", aws.StringValue(bucket.Name))
	}