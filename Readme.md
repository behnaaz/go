##Commands
go mod init main.go

go run main.go

##Test
curl http://localhost:8080
curl -X POST http://localhost:8080
curl -X PUT http://localhost:8080
curl http://localhost:8080/test/abs


##Install gorilla/mux
go get github.com/gorilla/mux

##Install MySql driver
go get -u github.com/go-sql-driver/mysql

##Install mysql
sudo apt update
sudo apt install mysql-server
sudo mysql_secure_installation

##Install vim-go
git clone https://github.com/fatih/vim-go.git ~/.vim/pack/plugins/start/vim-go

##DB setup
CREATE USER 'newuser'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON * . * TO 'newuser'@'localhost';
FLUSH PRIVILEGES;
