# my_go_cronjob_template
A template for Go Cronjob application by Jackk-Doe.<br>

- Create basic cronjob application.<br>
- Use github.com/go-co-op/gocron package to trigger cronjob. <br>
- Create a custom logger display format in JSON.<br>
- Separate Development, UAT and Production mode. (DEV, UAT, BUILD, TEST & PROD)<br>


## Initialise go mod and install necessary go packages (fiber, dotenv and others) 
```
go mod init <YOUR_GO_MOD_NAME>
go mod tidy
```

## After creating go module
Replace all `github.com/Jackk-Doe/my_go_cronjob_template` with <YOUR_GO_MOD_NAME> you named in `go mod init` command.<br>

### Files to be changed :
- cmd/cronjob/main.go
- configs/dotenv/dotenv.go
- pkg/jobs/test_job.go

## Create .env file, use .env.example file as a reference.


## Run command
`$ go run cmd/cronjob/main.go`

Or debug via CompileDaemon with: <br>

`$ CompileDaemon -build="go build -o ./build/cronjob cmd/cronjob/main.go" -command="./build/cronjob"`


## Example of my custom log format in JSON

INFO level logging :
```
{
	"level": "info",
	"msg": "STARTING_CRON_JOBS_APP...",
	"time": "2021-08-25T09:00:00+07:00",
	"id": "f7b3b3b4-0b3b-4b3b-8b3b-0b3b3b3b3b3b",
	"datas": {}
}

{
   "level":"info",
   "msg":"APP_INFO",
   "id":"-",
   "datas":{
      "APP_BUILD_AT":"",
      "APP_NAME":"SIC CRON Job app - DEV - LOCAL",
      "APP_RUN_AT":"2025-06-10 14:33:42",
      "APP_VERSION":"v0.0.1",
      "MODE":"dev"
   }
}
```

ERROR level logging :
```
{
   "level":"error",
   "msg":"load .env file error in DEV mode",
   "id":"-",
   "error":"open .env: no such file or directory"
   "datas": {}
}
```

