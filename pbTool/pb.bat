set CURR_DIR=%cd%

set OUTDIR=%cd%

protoc.exe --plugin=protoc-gen-go=protoc-gen-go.exe --go_out %OUTDIR% --proto_path "." *.proto
protoc.exe -I=%CURR_DIR% --csharp_out=%OUTDIR% %CURR_DIR%/*.proto
protoc.exe -I=%CURR_DIR% --java_out=%OUTDIR% %CURR_DIR%/*.proto

pause