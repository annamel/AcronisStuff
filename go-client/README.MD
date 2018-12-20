# Go REST API Server for Acronis

This is API for Cloud-Based Microservice for for Components of Distributed System Metrics and Logs Aggregation with Storage Them in Cloud Infrastructure

## Overview
This client is set to test server which is set to help different users automate their process of working with information gained from metrics and logs files.
It can help you to get statistic information from your files and see where your processes might have run into trouble.

- API version: 1.1
- Build date: 2018-12-11T19:37:11

### Directory tree
>clientAcronis
>>client
>>>logs files
>>
>>>metrics files
>>
>>>all_api.go
>>
>>>logs_api.go
>>
>>>metrics_api.go
>>
>>tests
>>>appstat_test.go
>>
>>>client_test.go
### Running the server
To run the client, follow these simple steps, beforehand uncomment all necessary stuff:

```
go run client.go
```

To run tests, follow these simple steps:

```
go test tests/appstat_test.go
```