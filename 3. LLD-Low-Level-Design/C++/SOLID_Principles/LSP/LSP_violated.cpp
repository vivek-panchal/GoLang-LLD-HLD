#include<iostream>
#include <bits/stdc++.h>

using namespace std;

class Account{
    public:
        virtual void withdraw(double amount) = 0;
        virtual void deposit(double amount) = 0;
};
class SavingAccount : public Account{
    private:
    double balance;

    public:
    SavingAccount(){
        balance = 0;
    }
    void deposit(double amount){
        balance += amount;
        cout << "Deposited " << amount << " to Saving Account. New balance: " << balance << endl;
    }

    void withdraw(double amount){
        if(amount > balance){
            cout << "Insufficient funds in Saving Account." << endl;
        } else {
            balance -= amount;
            cout << "Withdrew " << amount << " from Saving Account. New balance: " << balance << endl;
        }
    }
};

class CurrentAccount : public Account {
    private:
    double balance;

    public:
    CurrentAccount(){
        balance = 0;
    }
    void deposit(double amount){
        balance += amount;
        cout << "Deposited " << amount << " to Current Account. New balance: " << balance << endl;
    }
    void withdraw(double amount){
        if(amount > balance){
            cout << "Insufficient funds in Current Account." << endl;
        } else {
            balance -= amount;
            cout << "Withdrew " << amount << " from Current Account. New balance: " << balance << endl;
        }
    }
};

class FixedDepositAccount : public Account {
    private:
    double balance;

    public:
    FixedDepositAccount(){
        balance = 0;
    }

    void deposit(double amount){
        balance += amount;
        cout << "Deposited " << amount << " to Fixed Deposit Account. New balance: " << balance << endl;
    }

    void withdraw(double amount){
        cout << "Withdrawals are not allowed from Fixed Deposit Account before maturity." << endl;
    }

};

class BankClient {
    private:
    vector<Account*> accounts;

    public:
    BankClient(vector<Account*> accounts) {
        this->accounts = accounts;
    }
    void performTransactions() {
        for(Account* account : accounts) {
            account->deposit(1000);

           try{
            account->withdraw(500);
           } catch(...) {
               cout << "An error occurred during withdrawal." << endl;
           }
        }
    }
};

int main(){
    vector<Account*> accounts;
    accounts.push_back(new SavingAccount());
    accounts.push_back(new CurrentAccount());
    accounts.push_back(new FixedDepositAccount());

    BankClient client(accounts);
    client.performTransactions();


    return 0;
}