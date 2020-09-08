# MP0
#Daniel Flores and Nicholas Hillis



----
# To Run
To run this program you first need to type the following commands into terminal


```bash
go run process_a.go
``` 

This will open the tcp server and will proceed to listen for the client

Here is where you activate the tcp client:

```bash
go run process_b.go
```

Then it will prompt you to enter the following data (I'm using sample text entries)
```bash
To: floresd@bc.edu
From: hillisn@bc.edu
Title: Homework?
Content: Today our homework is due. Do you want to work on it?
```
Also it should be noted that we put in a print statement to show the user that the client is exiting, so it will also say
```bash
TCP client exiting...
```
This means the message has been sent and we received the ACK from the server. Therefore we can look over to process b and see the output:
```text
To: floresd@bc.edu

From: hillisn@bc.edu

Title: Homework?

Content: Today our homework is due. Do you want to work on it?

(This line is where the current time will show up)
```

----
### Processes

Our code runs entirely on the two previous processes with a being the client and b being the server.

We set up multiple readers to get the information of the email line by line and then sort it so it prints in the desired format.

### Sources

The following sites were used to help create our code:
```text
https://www.youtube.com/watch?v=9NTJF-IqEfw&t=639s

https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/

https://github.com/Dariusrussellkish/Example-MP-Solution

https://tutorialedge.net/golang/reading-console-input-golang/
```