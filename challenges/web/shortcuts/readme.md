This exercises continues to expand on Golang knowledge. These "shortcuts" are mini Go programs that run a simple command and print out the result, which is read by the main.go file and printed out to the user. All a user has to do is create a Golang script that prints out the flag from /home/jeff/flag.txt using exec.Command from Golang. 

After a user uploads the Go script, the list is updated but the list which contains the name of the file with the .go extension. If a user clicks on it, the server will not execute it because it appends ".go" to the name of the file from the list. Meaning, the user must upload a second file called "flag" without any extension. Once both are uploaded the user can easily click on the flag link and retrieve the flag.