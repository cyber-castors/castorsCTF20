This exercise works by validating the cookie from a client's browser against a predefined number: 3123213. The trick with the cookie is that with each request the server incrementes the cookie's value by 3, also the server grabas the cookies value and subtracts 35 when comparing it with the predefined number. Hence, the user must manipulate their cookie to 3123248. 

The second part of the exercise consists of injecting an SQL database to achieve data leakage. Once the user identifies the SQL vulnerability they must enumerate the database using queries such as "union select all 1,2,3,4..." where the numbers represent the tables columns. Again, once, the user validates the amout of columns they can leak data using queries such as "union select all 1,2,3,column_name from information_schema.columns where table_name ='users'". After identifying colum and table names, users can easily query the Users table, "union select all 1,2,3,username,password from Users;"

After leaking the admin's password they must crack the hash retrieved from the database and navigate to /dealer and login to retrieve the flag. The is found in rockyou.txt.


Extract Data from DB:
id=1 union all select 1,2,3,4 table_name from infomation_schema.tables
id=1 union all select 1,2,3,4 column_name from information_schema.colums where table_name='users'
id=1 union all select 1,2,3,username, password from users