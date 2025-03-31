Here’s a step-by-step guide to uninstall and reinstall PostgreSQL on your Fedora system for your Go CRUD project:

### **Uninstall PostgreSQL**
1. **Remove PostgreSQL Packages:**
   Run the following command to remove PostgreSQL and its server:
   ```bash
   sudo dnf remove postgresql postgresql-server
   ```
   Confirm the removal when prompted by typing `y`.

2. **Clean Up Data Directory:**
   Ensure the PostgreSQL data directory is empty or deleted:
   ```bash
   sudo rm -rf /var/lib/pgsql/data
   ```

3. **Remove Residual Files:**
   If any configuration files or residual files remain, you can clean them up:
   ```bash
   sudo rm -rf /etc/postgresql*
   ```

### **Reinstall PostgreSQL**
1. **Update System Packages:**
   Ensure your system is up-to-date:
   ```bash
   sudo dnf update -y
   ```

2. **Install PostgreSQL and Server:**
   Reinstall PostgreSQL using the following command:
   ```bash
   sudo dnf install postgresql postgresql-server
   ```

3. **Initialize the Database:**
   Set up the initial database structure:
   ```bash
   sudo postgresql-setup --initdb
   ```

4. **Enable and Start PostgreSQL Service:**
   Enable PostgreSQL to start on boot and start the service:
   ```bash
   sudo systemctl enable postgresql
   sudo systemctl start postgresql
   ```

5. **Verify Installation:**
   Check if PostgreSQL is running correctly:
   ```bash
   sudo systemctl status postgresql
   ```

### **Post-Installation Configuration**
1. Switch to the `postgres` user:
   ```bash
   sudo -i -u postgres
   ```

2. Create a new database and user for your Go project:
   ```bash
   createuser --interactive
   createdb go_crud_db
   ```

3. Set a password for the `postgres` user:
   ```bash
   psql -c "ALTER USER postgres WITH PASSWORD 'yourpassword';"
   ```

This process ensures a clean uninstall and reinstall, preparing PostgreSQL for your Go CRUD project setup[1][3][4].

```
[dxlord@haru] go-postgres-crud$ psql -U postgres -W
Password: 
psql (16.8)
Type "help" for help.

postgres=# CREATE DATABASE go_crud_db;
CREATE DATABASE
postgres=# \l

 go_crud_db | postgres | UTF8     | libc            | en_US.UTF-8 | en_US.UTF-8 |            |          
 | 

postgres=# \q
[dxlord@haru] go-postgres-crud$ go run main.go
✅ Connected to PostgreSQL with GORM!
[dxlord@haru] go-postgres-crud$ 
```