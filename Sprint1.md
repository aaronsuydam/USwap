Front End:
Aaron Sudyam
Evan Robinson

Back End:
Daniel Moraes
Andrew Jackson

## User Stories

### Backend User Stories:

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
Utilize MySQL and Golang to create a 

#2 Push users to database:
With SQL server setup next step was to actually create tables and send users into the DB.

#3 Query user data from database:
Once tables -with test user data- were created next step was to work on querying the DB.

#4 Connect SQL database with frontend:
Once base functionality was created to send and recieve user data from the DB, next step was to connect it to the front-end and parse data back and forth.

### Frontend:
#1 Login page
Sending user input from the frontend to the backend to verify authentication.

#2 Swap functionality
Need to set up a mock backend to switch items in the database.

#3 Landing page
Making sure the items are presented cleanly and inline.

## Completed
### Backend:
All backend issue were completed, complications arose with all of them however.

#1:
It was initially a bit difficult to get SQL and Golang to interact properly, it was some issue with the package interacting with the local host. Eventually, the server was established using an online server hosting platform.
#2 
It actually turned out to be quite a hassle getting user data properly uploaded into the database, wherein when a string was inputed to the table it would simply read as '0'.Varchar was not working when attempting to input strings. Eventually a workaround was found wherein upon creation of the table variables that would contain strings were declared as "text" variables but within GOLANG code were declared as "string" variables. 
#3
Querying user data went faster than uploading but we did run into issues with querying properly and recieving usernames properly.
#4
We ran into some issues referencing the database properly but eventually got it up.

### Frontend:
The login page is almost complete, however it just needs to be connected to the backend
The swap page and landing page need to be revamped UI wise, but everything else works.
We still need to add the filter. Have not had the time to implement so far.