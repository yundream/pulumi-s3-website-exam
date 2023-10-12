package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a bucket and expose a website index document
		f, err := NewS3Folder(ctx, "s3-website-bucket", "www", &FolderArgs{})
		if err != nil {
			fmt.Println(">> ", err.Error())
			return err
		}
		ctx.Export("bucketName", f.bucketName)
		ctx.Export("websiteUrl", f.websiteUrl)
		return nil
	})
}
