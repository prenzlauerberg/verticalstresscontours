I created this repository to store / possibly enhance and develop the vertical stress contours drawing algorithm/code. 
(Main reason -> CE366 course has a bonus assignment)
Right now it is written in golang, I preferred it due to its cleanness  and efficiency.
I might also add a basic web page to draw the contours...
----
31.05.2024
Simple html page and a js script is added to send the parameters and get the results from the golang based backend
Axios is used for the http request, 
for drawing the contours , the plain html canvas is used, (just the points are drawn)

In the server side, fiber package used to receive request.

---
How to run?

Golang is a dependency since there aren't any prebuilt binaries shared (nor planned)
https://go.dev/

---
then simply after cloning this repository, 
in the root folder 
"go run main.go"
command is enough to start the server,
Note: this will listen on 9924 port.
    Then heading out to the following link

http://127.0.0.1:9924/ce366
