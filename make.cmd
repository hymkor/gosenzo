setlocal
set GOARCH=386
go build -ldflags "-s -w"
for %%I in (*.exe) do upx %%I
endlocal
