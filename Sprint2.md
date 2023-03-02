Sprint 2 file
Back end functionality:

Table functionality:
Utilizing a dual table functionality consisting of a User Table and a UserItems table.
User Table stores: user_id,user_name,user_email and user_password
UserItems Table stores: row_num, item_name,item_description and user_id

The dual table functionality is essentially accomplished using a foreign key constraint.
Each user_id item in the User table is unique and is the primary key.
Thus, we can pass values into the UserItems table with the user_id as a column, allowing us to cross reference the tables.
By cross referencing the User table with the UserItems table we can then find all the items of a particular user.
