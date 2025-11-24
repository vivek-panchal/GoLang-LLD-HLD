#include<iostream>
using namespace std;

class Database {
public:
    virtual void save(string data) = 0;
};  

class MySQLDatabase : public Database {
public:
    void save(string data) override {
        cout << "Saving data to MySQL Database: " << data << endl;
    }
};  

class MongoDBDatabase : public Database {
public:
    void save(string data) override {
        cout << "Saving data to MongoDB Database: " << data << endl;
    }
};

class UserService { 
private:
    Database* db;
public:
    UserService(Database* database) {
        db = database;
    }
    
    void storeUser(string userData) {
        db->save(userData);
    }
};

int main(){
    MySQLDatabase mysql;
    MongoDBDatabase mongodb;
    UserService userServiceMySQL(&mysql);
    UserService userServiceMongoDB(&mongodb);
    userServiceMySQL.storeUser("User1 Data");
    userServiceMongoDB.storeUser("User2 Data");
    return 0;
}