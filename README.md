# What's stp/stph?

stp/stph helps you make system of delivering and management image files on AWS for your Hugo site.
stp work with the Lambda, DynamoDB, S3, and CloudFront and it is set up by terraform.

stph ( = stp helper ) work as local web server for management images and display Hugo shortcodes.

Motivation why I made stp/stph is just for learning Go and AWS.

## Note

I knew [Cloudinary](https://cloudinary.com/) after I finished to make stp/stph...:cold_sweat:

Anyway, at this time, I have plans to add some features to stp/stph for enjoying Go and AWS.


# Features: v0.0.1

- Automatically resizing images when the images are uploaded into your S3 bucket.
- Managiment your images in AWS S3 bucket with `stph`.

Of couse, all AWS resources are managed by terraform so you don't need to create/destroy them in your AWS console.

# Setup stp

1. Install Go and terraform (Also I recommend to use [tfenv](https://github.com/tfutils/tfenv.) for management your terraform version with `.terraform-version` file.)
2. Setup your aws cli account and install your credentials into your `$HOME/.aws`.
3. Git clone `stp` in your workdir.
4. `cd stp/src`
4. `make build`
5. `cd stp/tf` and `terraform init` and `terraform apply`.
6. Wait about 5-10 minitues for setup all AWS resources. Especially setting up CloudFront needs about 5 or more minitues.
7. After apply complete, now you can use stp/stph.

Try to upload your jpg, png or gif file(s) into your S3 bucket's `org` directory.

A resized image will be made your S3 buckets's `resize` directory.

The original and resized images can be access from web browser with CloudFront's URL like `https://abcdefghijklmn.cloudfront.net/resize/resized-xxxxxxxxx.jpg`

Check your aws console to get CloudFront Domain and DynamoDB item like below.

![CloudFrontDomain](https://d3i0o7y01oiqpa.cloudfront.net/org/stp_cloudfront.png)

If you upload an image into your S3 bucket's `org`, an item will be write down to your DynamoDB table.
![DynamoDB](https://d3i0o7y01oiqpa.cloudfront.net/org/stp_dynamodb.png)

# Setup stph

1. Git clone `stph` repository.
2. `go build` to build `stph`.
3. Write your DynamoDB table name, S3 bucket name, CloudFront domain to your `configs/settings.yaml`. ( Copy `sample_setting.yaml` to `settings.yaml` )

Execute stph and access `http://localhost:8080` from web browser. If it works, web page shows your images in your S3 bucket like below.

![stph](https://d3i0o7y01oiqpa.cloudfront.net/org/stph_toppage.JPG)

`stph` will help you manage your image(delete and copy strings of Hugo shortcode).

If you want to use shortcode, see the [how-to-use-stp-shortcodes.md in stph repository](https://github.com/zono-dev/stph/blob/master/shortcodes/how-to-use-stp-shortcodes.md).
