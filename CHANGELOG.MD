# [0.0.13]
ADDED:
 - client.ChangeAgent to allow for setting of the smartsheet-change-agent header

# [0.0.12]
ADDED:
 - etColumns to return all columns in a sheet
 - pdateColumn to update a column in a sheet
 - column_resonse, and column_get structs
 - webhook callback and event structs
 - Put_Call function for httpd put calls
 - TestGetColumns test function

UPDATED:
 - column stuct to include more fields
 - sheet struct to include more fields
 - getSheet function to match the new request structure
 - request logging
 - rev GoLang from 1.17 -> 1.23
 - GetSheet test to use columns not owner id

# [0.0.11]
ADDED:
 - added smartsheet error handling

# [0.0.10]
Updated:
 - Moved some structs around to more aproprate files
 - Change data handling from strings to byte slices in root call function

Added:
 - discussion create and gets based on row or sheet
 - comment create and gets
 - smartsheet_request struct for smartsheet post/get requests
 - inlcudes struct for calls that can optionally include other objects
 - Post_Call function for httpd post calls
 - additional tests for comments and discussions

# [0.0.9]
ADD:
 - added new User and Comment structs
 - added Call and Get_Call functions for httpd call re-use
 - switch getReport to use Call and Get_Call functions

# [0.0.8]
ADD:
 - added new report_column struct for the report struct
 - added virtual_column_id to cell struct for reports

# [0.0.5]
ADD:
 - Some logging to the GetReport function

# [0.0.4]
ADD:
 - added COLUMNS to report struct

# [0.0.3]
FIX:
- fixed Image id to string
- fixed cell value to interface

# [0.0.2]

ADD:

- added GetReport call to smartsheet

- added alot of structs

- added some tests

- split sheet calls and report calls into seperate files

# [0.0.1]
inital commit
