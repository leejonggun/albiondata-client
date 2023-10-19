del rsrc_windows_*
del albiondata-client.exe
del albiondata-client.*.bak
del .albiondata-client.*.old

del albiondata-client-amd64-installer.exe

go install github.com/tc-hib/go-winres@latest

go-winres make

set GOOS=windows
set GOARCH=amd64
go build -ldflags "-s -w -X main.version=2.0" -o albiondata-client.exe -v -x albiondata-client.go

go-winres patch albiondata-client.exe

dir albiondata-client*

copy albiondata-client.exe albiondata-client.exe.copy
"C:\Program Files (x86)\GnuWin32\bin\gzip.exe" albiondata-client.exe
move albiondata-client.exe.gz customize-windows-amd64.exe.gz
move albiondata-client.exe.copy albiondata-client.exe
