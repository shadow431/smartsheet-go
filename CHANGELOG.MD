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
