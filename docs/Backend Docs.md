﻿Backend Documentation


Server
Our test server runs on Golang’s Gorilla Mux on port 4201. It implements CORS for cross origin resource sharing. 
REST API
All calls to the backend will be made via https calls with json body encoding. 


NOTE: For site urls, please leave the base url off of any backend calls. Angular prefixes its own base url depending on folder setup which causes discrepancies in http calls. All calls have been fixed to go through a network interceptor which prefixes the correct server url automatically.
Pages
Login (/login)
Calls:
* Post - Provide args username and password.


Signup (/signup)
Calls:
* Post - Provide args username and password


Home (/, /home)
Home/dashboard page. Holds links to login and will hold an intro dashboard.
Calls
* None right now. 
Future Calls:
* Get - Receive item postings


Item (/item)
Get(item ID)
- Returns itemID, itemName, itemDescription, userID, and imagePath from Items table


Profile (/profile)


Database Info

Users Table
* The above framework allows us to quickly store a user based on their basic login information
* Passwords are encrypted with Bcrypt and hashes are stored in the database.
Contains:
- user_id
- user_name
- user_email
- user_password


Items Table
Contains:
- item_id
- item_name
- item_description
- user_id
- image_path

Swap Table
Contains:
- swap_id
- sender_id
- sender_item_id
- receiver_id
- receiver_item_id
