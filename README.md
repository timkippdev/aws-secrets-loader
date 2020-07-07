# AWS Secrets Loader

An application written in Go that will take a JSON file and will create or delete the values from AWS Secrets Manager.

## Prerequisites
If you intend to compile and run the raw code, you must have Go installed.

## Usage

### Running Pre-Compiled Binary (tested on MacOS only)
```
./bin/aws-secrets-loader <optional_flags>
```

### Manual Compile and Run
```
go run main.go <optional_flags>
```

### Compile Into New Binary (tested on MacOS)
```
go build -o bin/aws-secrets-loader main.go
```

## Optional Flags

You can combine any of the following flags in order to customize your application usage.

Name | Description | Usage | Default Value
---- | ----------- | ----- | -------------
delete | flag to delete the secrets instead of create them | --delete | false
file | the path to the file you want to load | --file=data/somefile.json | data/data.json
region | the AWS region to upload/delete secrets in | --region=us-west-1 | us-west-2