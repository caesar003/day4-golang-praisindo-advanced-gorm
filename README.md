# Go GORM 

## Prerequisites

 Ensure you have postgresql installed and running, if you're using another database, you might need to tweak a few settings yourself.

### Install postgresql (if not already installed)

```shell
sudo apt update && sudo apt install postgresql postgresql-contrib
```
### Connect to PostgreSQL as the `postgres` user:

```shell
sudo -u postgres psql
```
### Create a new user and database

```sql
-- Create user with a strong, unique password
CREATE USER superadmin WITH PASSWORD 'supersecretpassword1234';

-- Create database with <your_name> as the owner
CREATE DATABASE go_playground OWNER superadmin; 

-- Exit PostgreSQL prompt
\q
```

### Verify the New User and Database

```shell
psql -h localhost -p 5432 -U superadmin -d go_playground
```

You will be prompted to enter the password you set earlier. Upon successful login, you should be able to interact with your new database.

## Project Setup

### Clone the repository


```shell
git clone https://github.com/caesar003/day4-golang-praisindo-advanced-gorm.git ~/go-gorm
```

### Configure database credentials

1. Navigate to the project directory:

```shell
cd ~/go-gorm
```

2. Copy the sample credentials file.

```shell
cp credentials-copy.json credentials.json 
```

3. Open `credentials.json` in your preffered text editor and replace the placeholder values with your actual database credentials.

### Install the dependencies

Ensure all dependencies are installed by running:

```shell
go mod tidy
```

### Run the the development server
Start the server using the following command

```shell
go run main.go
```

