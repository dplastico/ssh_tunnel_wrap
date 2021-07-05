# ssh_tunnel_wrap

In order to practice some GO LANG and working on a project with a customer , I created this small tool, I'm not a programmer and should be use it wit precaution, but I think it could be useful to automate  the process on POCs

**ssh_tun** is a go script that creates a .exe files that can run the ssh tunnel command from [Check Point Connect for remote users](https://www.checkpoint.com/harmony/connect-sase/clientless-remote-access/) tunnel applications. It just automate the process of typing/copying the command from the website to the command line each time you want to start the tunnel app. With a few modifications you can use it to create linux binary also.

**How to use:**

Install go : [HERE, GO](https://golang.org/doc/install)

**Clone the repo**

Change the port to be redirected and the connection information of the variables in the script **ssh_tun.go** and name the key generated on you connect tenant as cp.pem (you can change the name also if you prefer, just adjust the command on the script)

```golang
	var conn string
	conn = "binario_test@tunnel.us-2.checkpoint.security" //CHANGE connection information
	var port string
	port = "7777" //CHANGE  port to be redirected (localhost)
```

**Create a binary:**
Execute the following:

> \>set GOBIN=<C:\path to the cloned repo \>

then:
> \>go build ssh_tun.go

This will generated a .exe file that you can then execute on the same directory of the cp.pem file. As an example,
you can run it in a terminal just executing:

> \>ssh_tun.exe

You can run it on background:

> \>start /b ssh_tun.exe

Or you can execute it when windows starts (I gues this can be done by AD directives also)

add it to:
> C:\Users\{Username}\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup 
or type in run (win+r):
>shell:startup 
