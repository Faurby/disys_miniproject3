# disys_miniproject3

Replica:6001 dies after 10 seconds, and the auction ends after 25 seconds.

If you have Powershell and Windows Terminal, you can simply right click start.ps1 and "run with Powershell"

If you dont, then you need to run thses commands, each in a new terminal window.

```powershell
go run .\Replica\Replica.go :6001
go run .\Replica\Replica.go :6002
go run .\Replica\Replica.go :6003
go run .\Frontend\Frontend.go
go run .\Client\Client.go
go run .\Client\Client.go
go run .\Client\Client.go
```
