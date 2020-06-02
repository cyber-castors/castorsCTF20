This challenge needs to have a script that auto generates requests and adds users to the DBSessions. Problem is credentials aren't sent securely. An alternative is to use the jwt tokens printed out by the server to try to login from a different computer, using only the jwt tokens generated. 

To do this the server must be up and running at all times, meaning if it happens to fail the creation of all jwt tokens would need to take place again and so would the capture, hence the need for an automated script.

*UPDATE* 
Due to my lack of understanding of Json Web Tokens I couldn't provide a strong and good challege with a specific vulnerability. If this challenge is used it will only consists of cracking the JWT token. 