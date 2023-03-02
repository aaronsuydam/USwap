Sprint 2 file
# Back end functionality/documentation:


Server: Our test server runs on Golang’s Gorilla Mux on port 4201. It implements CORS for cross origin resource sharing. REST API All calls to the backend will be made via https calls with json body encoding.

NOTE: For site urls, please leave the base url off of any backend calls. Angular prefixes its own base url depending on folder setup which causes discrepancies in http calls. All calls have been fixed to go through a network interceptor which prefixes the correct server url automatically. 


Pages Login (/login) 
Calls: 
  PUT: Enter user login information in json format. Includes "username" and "password"


Signup (/signup) 
Calls: 
  PUT: Enter signup information in json format. Includes "username", "email", and "password."



  Post - Provide args username and password
  
Home (/, /home) Home/dashboard page. Holds links to login and will hold an intro dashboard. Calls

None right now. Future Calls:
Get - Receive item postings
Item (/item)

Profile (/profile)

      

## Database Tables:

Table functionality:
Utilizing a dual table functionality consisting of a User Table and a UserItems table.
User Table stores: user_id,user_name,user_email and user_password  

UserItems Table stores: row_num, item_name,item_description and user_id

The dual table functionality is accomplished using a process similar to a foreign key constraint.
Each user_id item in the User table is unique and is the primary key.
Thus, we can pass values into the UserItems table with the user_id as a column, allowing us to cross reference the tables.
By cross referencing the User table with the UserItems table we can then find all the items of a particular user.


