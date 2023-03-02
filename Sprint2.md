## Backend Work Completed - Work accomplished this sprint:  

Creation and population of userItems table
Linking of Items table with Users table.
Fixed blocking issues revolving around CORs middleware with implementation of own middleware.
Unit tests for database, utils, and handlers.
Refactoring and organization of code for better project structure


## Backend Unit Tests
Database: Added unit tests for database configuration, database connection, table creation, and pinging the database after setup.
Hashing: Added unit tests to ensure success of utils hashing functionality and password entry.
Handling: Added mockups of route testing. Need to look into test mocking for Go as the database may not be started during these tests. 

## Backend Documentation:
### Pages
All pages require json formatted data

Login (/login) 
Calls: 
POST: Submit credentials to log a user in. 
Args - “username”, “password”


Signup (/signup)  

Calls:  
POST:: Submit credentials to create a new account. 
Args - "username", "email", and "password."
  
Home (/, /home) Home/dashboard page. Holds links to login and will hold an intro dashboard.  
Calls:  

None right now.   
Future Calls:  

Get - Receive item postings
Item (/item)

Profile (/profile)
None right now.  
Future Calls:  

Get - Requests profile info and items belonging to user

## Database Tables:

Table functionality:
Utilizing a dual table functionality consisting of a User Table and a UserItems table.
User Table stores: user_id,user_name,user_email and user_password  

UserItems Table stores: row_num, item_name,item_description and user_id

The dual table functionality is accomplished using a process similar to a foreign key constraint.
Each user_id item in the User table is unique and is the primary key.
Thus, we can pass values into the UserItems table with the user_id as a column, allowing us to cross reference the tables.
By cross referencing the User table with the UserItems table we can then find all the items of a particular user.
