Front End:
Aaron Sudyam
Evan Robinson

Back End:
Daniel Moraes
Andrew Jackson

## User Stories

Backend User Stories:

#1:
As a user, I expect that my login data will be stored securely and that my password will be hashed.

Task: Hash all passwords before storing in database and securely handle login information.

#2:

As a college student, I want to create a profile log in, and I expect my login to be stored by the site and it will remember me.

Requirements/Task:
-Store users uniquely based on primary key(i.e no duplicates)
-Store username, email and password in DB to be accessed when needed

#3:
As a college student/user, I expect that the website will be able to communicate between it's own front end and back end.

Task: Connect front-end and back-end. Store user data from fron-tend into back-end

## Issues
Backend:
#1 Create Initial Backend storage of users:
Utilize MySQL and Golang to create a 

#2 Push users to database:
With SQL server setup next step was to actually create tables and send users into the DB.

#3 Query user data from database:
Once tables -with test user data- were created next step was to work on querying the DB.

#4 Connect SQL database with frontend:
Once base functionality was created to send and recieve user data from the DB, next step was to connect it to the front-end and parse data back and forth.
## Completed
Backend:
All backend issue were completed, complications arose with all of them however.

#1:
It was initially a bit difficult to get SQL and Golang to interact properly, it was some issue with the package interacting with the local host. Eventually, the server was established using an online server hosting platform.
#2 
It actually turned out to be quite a hassle getting user data properly uploaded into the database, wherein when a string was inputed to the table it would simply read as '0'.Varchar was not working when attempting to input strings. Eventually a workaround was found wherein upon creation of the table variables that would contain strings were declared as "text" variables but within GOLANG code were declared as "string" variables. 
#3
Querying user data went faster than uploading but we did run into issues with querying properly and recieving usernames properly.
#4
We ran into some issues referencing the database properly but eventually got it up.

## Why
