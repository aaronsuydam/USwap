# Frontend Work Completed

Creation of more unit tests.
Functioning login page for registered users by verifying credentials to backend.
Utilized JWTs to keep track of user logged in state.
Updated site wide styles and theming.
Created service to store user credentials in session storage.
Created AuthGuard to validate if a user has access to certain routes using the new auth system.

# Frontend Unit Tests

Components: Test all components and root app to ensure they are created successfully. Test LoginPage for proper login state and storage of a user.
Signup Service: Service is successfully posting user object to backend and receiving it back.
Auth Service: Service is successfully logging a user in and assigning a JWT.
Cypress Test: Created e2e unit test to test post functionality of signup service and auth service.
Form: Validate user input into signup form fields.

# Backend Changes and Documentation

Sprint 3 changes:
Added a bunch of new routes for creating and gettings users, items and swap requests, along with unit tests. Added functionality to accept or reject swap rejects, which swaps the owner of items in the items database. Implemented real uuids instead of counting rows for ids in all tables. Significantly cleaned up handler and database code to handle all queries in the database package instead of in the routes thmselves. Added tests for everything, including  tests cross-referencing and making edits between tables.

More info on our documentation and changes is below: 

Server Our test server runs on Golang’s Gorilla Mux on port 4201. It implements CORS for cross origin resource sharing. REST API All calls to the backend will be made via https calls with json body encoding.
NOTE: For site urls, please leave the base url off of any backend calls. Angular prefixes its own base url depending on folder setup which causes discrepancies in http calls. All calls have been fixed to go through a network interceptor which prefixes the correct server url automatically.

## Pages Login (login) Calls:
Post - Provide args username and password.
Assigns a user a JWT in a cookie.

## Signup (signup) Calls:
Post - Provide args username and password

## Home (home) Home/dashboard page. Holds links to login and will hold an intro dashboard. Calls
None right now. Future Calls:
Get - Receive item postings

## Item (item)

## Profile (/profile)

## Backend Documentation
Server Our test server runs on Golang’s Gorilla Mux on port 4201. 
It implements CORS for cross origin resource sharing.
## REST API All calls to the backend will be made via https calls with json body encoding.
NOTE: For site urls, please leave the base url off of any backend calls. Angular prefixes its own base url depending on folder setup which causes discrepancies in http calls. All calls have been fixed to go through a network interceptor which prefixes the correct server url automatically.
## Pages Login (login) Calls:
Post - Provide args username and password.
Assigns a user a JWT in a cookie.
## Signup (signup) Calls:
Post - Provide args username and password
Home (home) Home/dashboard page. Holds links to login and will hold an intro dashboard. 
Calls:  
None right now. 
##### Future Calls:
Get - Receive item postings
Item (item) 
Profile (profile)  
## Database Users Table
User objects containing a base of user_id, user_name, user_email and user_password  
  
The above framework allows us to quickly store a user based on their basic login information  
  
Eventually, there will be a table of items contained within each User row, a nested table will allow for easy access to user   items and easy item management Password Storage  


Passwords are encrypted with Bcrypt and hashes are stored in the database.  


## Database Items Table
Item objects containing a base of item_id, item_name, item_description, user_id and image_path  


The above framework allows us to quickly store an item on its basic information  


This table can be easily cross referenced with the users table to find a specific users items  


## Database Swap Table
Swap objects containing a base of swap_id, sender_id, sender_item_id	receiver_id, receiver_item_id  


The above framework allows us to quickly store a swap request on its basic information  


This table can be easily cross referenced with the users table and the items table to initiate a swap request to find a specific users items  


Note that this table only stores the request, not the swap itself


Database Unit Tests:
TestDatabaseConfig - self explanatory  


TestUserTableCreation- self explanatory  


TestGetUser - tests getting a user row from the table based on user_id  


TestGetUserItem - this test gets a singular item based on item id  


TestGetUserItems - this test gets all the items that a user has, based on user_id  


TestGetSwapRequest- this test will get swap request based  on swap id  
