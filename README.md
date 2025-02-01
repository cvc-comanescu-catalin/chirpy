# chirpy
bootdev go server

go mod init to create a Go module for your project
go mod init github.com/cvc-comanescu-catalin/chirpy

go build -o out && ./out

Steps
Create a new http.ServeMux
Create a new http.Server struct.
Use the new "ServeMux" as the server's handler
Set the .Addr field to ":8080"
Use the server's ListenAndServe method to start the server
Build and run your server (e.g. go build -o out && ./out)
Open http://localhost:8080 in your browser. You should see a 404 error because we haven't connected any handler logic yet. Don't worry, that's what is expected for the tests to pass for now.



sudo apt update
sudo apt install postgresql postgresql-contrib

psql --version
sudo passwd postgres
sudo service postgresql start

sudo -u postgres psql
CREATE DATABASE chirpy;
\c chirpy
ALTER USER postgres PASSWORD 'postgres';
SELECT version();


protocol://username:password@host:port/database
postgres://postgres:postgres@localhost:5432/chirpy
psql "postgres://wagslane:@localhost:5432/chirpy"
psql postgres://postgres:postgres@localhost:5432/chirpy


cd into the sql/schema directory and run:
goose postgres <connection_string> up
goose postgres postgres://postgres:postgres@localhost:5432/chirpy up
goose postgres postgres://postgres:postgres@localhost:5432/chirpy down

psql chirpy
\dt



https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html
https://db-engines.com/en/ranking