# net-cat
## About Project
This project consists on recreating the `NetCat` in a `Server-Client` Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

NetCat, nc system command, is a command-line utility that reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open `TCP` connections, send UDP packages, listen on arbitrary `TCP` and UDP ports and many more.

To see more information about NetCat inspect the manual `man nc`.
## Usage
```
$ go run . $port 
```
### Example
- `1st` Terminal
```bash
$ go run cmd/server/main.go || make server
Listening on the port :8989
2022/08/30 19:16:23 Listening the connections on the server 127.0.0.1:8989
2022/08/30 19:16:32 new client has connected: 127.0.0.1:51534
2022/08/30 19:16:32 new client has connected: 127.0.0.1:51536
2022/08/30 19:26:17 127.0.0.1:51534 was disconnected from the server
2022/08/30 19:26:17 127.0.0.1:51536 was disconnected from the server
^Csignal: interrupt
$ 
```
- `2nd` Terminal
```bash
$ go run cmd/client/main.go || nc localhost 8989 || make client
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Aidana
[2021-10-27 19:14:10][Aidana]:Hello 
Sultan has joined our chat...
[2021-10-27 19:14:43][Sultan]:Hello
[2021-10-27 19:14:43][Aidana]:How Are you?
[2021-10-27 19:15:31][Sultan]:I am fine, and you?
[2021-10-27 19:15:51][Aidana]:Good
Sultan has left our chat...
[2021-10-27 19:16:10][Aidana]:^C
$ 
```
- `3rd` Terminal
```bash
$go run cmd/client/main.go || nc localhost 8989 || make client
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Sultan
[2021-10-27 19:14:20][Aidana]:Hello
[2021-10-27 19:14:32][Sultan]:Hello
[2021-10-27 19:15:04][Aidana]:How Are you?
[2021-10-27 19:15:04][Sultan]:I am fine, and you?
[2021-10-27 19:15:31][Aidana]:Good
[2021-10-27 19:16:03][Sultan]:^C
$ 
```
## Authors
- [AidanaErbolatova]