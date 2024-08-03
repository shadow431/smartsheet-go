
to create a new client
```go
client := smartsheet.NewClient()
``` 

the client will get the access token from the environment variable SMARTSHEET_ACCESS_TOKEN

to run tests the following environment variables are needed

```bash
SMARTSHEET_ACCESS_TOKEN=your_access_token

#TEST IDS
REPORT_ID=
ROW_ID=
SHEET_ID=
COMMENT_ID=
DISCUSSION_ID=

```