# Status S3Client 
POC to replace MinIO mc client with our custom implementation with minimal functionality. This new client use AWS SDK V2.

# Getting Started
To start contributing you need to install `Go Lang` version >= 1.15 as this is prerequisite for `AWS SDK V2`. 
We are using `Go Lang` version 1.19 so try to use this version if possible. For local development we
recommend to have `docker` and `docker-compose` installed on your computer.

## Install Go libraries
You can install used libraries manually by executing:
```sh
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/credentials
go get github.com/aws/aws-sdk-go-v2/service/s3
go get github.com/aws/smithy-go
```

Or you can install all of them by executing:
```sh 
go build ./...
```
or
```sh
go run ./...
```

## Install Docker
You can install `docker` and `docker-compose` by executing:
```sh
brew install docker docker-compose
```

> On M1, M2 generation mac's you should consider to install [colima](https://github.com/abiosoft/colima) container run time.
> For x86_64 mac's you have more options to choice from. So take your own decision and install the one you like the most.

## Run docker with MinIO locally
We are using `MinIO` release `RELEASE.2021-04-22T15-44-28Z` as thi is the last know version with Apache License Version 2.0.
> !!! Be careful here as all newer version use GNU Affero General Public License (AGPL).

To start your local instance and run it in background:
```sh
docker-compose -f minio-standalone-compose.yaml up -d
```

To stop your local instance and remove volume:
```sh
docker-compose -f minio-standalone-compose.yaml down -v
```

# Build and Test
TBD
