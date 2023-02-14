VIDEO LINKS: https://drive.google.com/drive/folders/1MfYMeo8X4RtdYbrz1etKz4Lu9D6LQRwh?usp=sharing

Front End:
Aaron Sudyam
Evan Robinson

Back End:
Daniel Moraes
Andrew Jackson

## User Stories

### Backend User Stories:

#1:
As a college student, I want to create a profile log in, and I expect my login to be stored by the site and it will remember me.

Tasks for Sprint 1:
- Establish communication between the front-end and backend to allow for authentication
- Store users uniquely based on primary key(i.e no duplicates)
- Store username, email and password in DB to be accessed when needed

Further task for Sprint 2: 
- Enable session cookies so users can stay logged in.

#2:
As a college student/user, I expect to be able to see items that people are willing to trade on the on the site. 

Task for Sprint 1: 
- Connect front-end and back-end. Store user data from fron-tend into back-end. Configure databases and tables in SQL and establish Golang-SQL communication.

Further task for Sprint 2:
- Load in items available for trade on a dashboard with a title and picture. Have items be selectable to view additional information upon clicking on them.

#3:
As a user, I expect that my login data will be stored securely and that my password will be hashed.

Task for Sprint 1: 
- Hash all passwords before storing in database and securely handle login information.



### Frontend User Stories:

#1
As a user, I expect to be able to sign up and login into my account.

Task: Create a login page for users to access their account.

#2
As a college student, I want to trade my toaster for someone's blender, so that I can make protein shakes instead of bagels for breakfast.

Task: Create a UI for users to swap their items for other items from different users.

#3
As a user, I want to see all the available items available to trade.

Task: Present users with an initial landing page showcasing items in a grid.

#4
As a user, I want to be able to filter items so I can find the items that interest me.

Task: Add a filter UI component with labels that users can select and deselect to filter by.

## Issues
### Backend:
#1 Create Initial Backend storage of users:
Utilize MySQL and Golang to create a database with users stored in a table. Must be able to communicate with database from Golang.

#2 Push users to database:
With SQL server setup next step was to actually create tables and send users into the DB.

#3 Query user data from database:
Once tables -with test user data- were created the next step was to work on querying the DB.

#4 Connect SQL database with frontend:
Once base functionality was created to send and receive user data from the DB, next step was to connect it to the front-end and parse data back and forth.

#5 Test signup and login functionality
Ensure that users can signup and login to the site if and only if they input proper credentials

#6 Setup REST API
Enable communication with the front end for the following pages:
- Login
- Signup
- Home
Create a backend API doc detailing front/backend communication.

#7 Setup product database as nested tables within user database
- Create item tables for each user and generate sample products


### Frontend:
#1 Login page
Sending user input from the frontend to the backend to verify authentication.

#2 Swap functionality
Need to set up a mock backend to switch items in the database.

#3 Landing page
Making sure the items are presented cleanly and inline.

## Completed
### Backend:
#1 Complete Initial Backend Storage of Users

#2 Push Users to Database

#3 Query User Data from Database

#5 Test Signup and Login Functionality

#6 Setup REST API (For Login, Signup, Home

### Frontend:
The official login page is almost complete, however it just needs to be connected to the backend.
The swap page and landing page need to be revamped UI wise, but everything else works.
We still need to add the filter. Have not had the time to implement so far.

## Not Completed and Why?
### Backend:
#4 Connect SQL database with frontend

We have not yet merged the newest front-end designs with the backend. However, test pages we have created are able to communicate with the backend server which is connected and communicating with our SQL database.

#7 Setup Product Database as Nested Tables Within User Database

This was more of a stretch goal for this Sprint. We did not have time to achieve this task this time around, but now that we have a better grasp on creating SQL queries and creating tables in SQL this should not be a difficult task for early on in Sprint 2.


### Frontend
#2 

It actually turned out to be quite a hassle getting user data properly uploaded into the database, wherein when a string was inputed to the table it would simply read as '0'.Varchar was not working when attempting to input strings. Eventually a workaround was found wherein upon creation of the table variables that would contain strings were declared as "text" variables but within GOLANG code were declared as "string" variables. 
#3

Querying user data went faster than uploading but we did run into issues with querying properly and recieving usernames properly.
#4
We ran into some issues referencing the database properly but eventually got it up.

## Difficulties Faced + Lessons Learned
### Backend
#1:

It was initially a bit difficult to get SQL and Golang to interact properly, it was some issue with the package interacting with the local host. Eventually, the server was established using an online server hosting platform. Also, there is a lot of outdated documentation for golang along with many functions being deprecated which caused a bit of confusion.

#2 

It actually turned out to be quite a hassle getting user data properly uploaded into the database, wherein when a string was imputed to the table it would simply read as '0'.Varchar was not working when attempting to input strings. Eventually a workaround was found wherein upon creation of the table variables that would contain strings were declared as "text" variables but within Golang code were declared as "string" variables. 


#4

Angular's http client was being very difficult to work with. We found that the client would prefix any url with its own file path within the src directory including the front end port, making it impossible to make calls to the backend as these could not be edited. After a while we found that we could construct an http interceptor to intercept all http calls from angular and fix the server url to them. 
