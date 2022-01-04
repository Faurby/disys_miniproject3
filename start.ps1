wt.exe PowerShell go run .\Replica\Replica.go :6001 `; `
split-pane -V PowerShell go run .\Replica\Replica.go :6002 `; `
split-pane -V PowerShell go run .\Replica\Replica.go :6003 `; `
nt PowerShell go run .\Frontend\Frontend.go :5000 `; `
split-pane -H PowerShell go run .\Frontend\Frontend.go :5001 `; `
move-focus left `; `
split-pane -H PowerShell go run .\Client\Client.go`; `
split-pane -V PowerShell go run .\Client\Client.go`; `
move-focus up `; `
split-pane -V PowerShell go run .\Client\Client.go `;