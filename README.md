# backupdatabase

Backup your MySQL / MariaDB database to .sql file using Go Language.

We put the database dump information in a textfile, separated with "#" symbol. This file consists of :
- MySQL dump file location
- Database's host (IP Address or domain)
- Database's port
- Database's name
- Database's username
- Database's password

We use the command-like function to execute the dumping file. The common command function to dump the database is more like this 
`mysqldump -h[Host] -P[port] -u[Username] -p[Password] [database_name] --[options]`
While in this code, we use the `os/exec` library to execute above command. Then we add some options such as :
- `--events`, to include the Event Scheduler.
- `--routines`, to include the Routines Function or Procedure.
- `--triggers`, to include the Triggers.
- `--single-transaction`, to dump the entire database without locking it.