rmdir /s /q "./build"
mkdir build
go build -C ./backend -o ../build -ldflags -H=windowsgui -tags release
yarn --cwd ./frontend build
