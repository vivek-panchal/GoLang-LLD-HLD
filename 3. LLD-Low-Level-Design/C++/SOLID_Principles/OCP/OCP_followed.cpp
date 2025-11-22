#include<iostream>
#include <bits/stdc++.h>

using namespace std;

class Product{
    public:
    string name;
    double price;
    Product(string name , double price){
        this->name = name;
        this->price = price;
    }
};

class ShoppingCart{
    private:
    vector<Product*> products;

    public:

    void addProduct( Product* product){
        products.push_back(product);
    }

    const vector<Product*>& getProducts() {
        return products;
    }

    double calculateTotal(){
        double total = 0.0;
        for(const auto product : products){
            total += product->price;
        }
        return total;
    }
};

class ShoppingCartPrinter{
    private:
    ShoppingCart* cart;
    public:
    ShoppingCartPrinter(ShoppingCart* cart){
        this->cart = cart;
    }
    void printReceipt(){
        cout << "Receipt:" << endl;
        for(const auto product : cart->getProducts()){
            cout << product->name << " - $" << product->price << endl;
        }
        cout << "Total: $" << cart->calculateTotal() << endl;
    }
};

class Persistence{
    private:
    ShoppingCart* cart;
    public:
    virtual void save(ShoppingCart* cart) = 0;
    
};

class SQLPersistence : public Persistence{
    public:
    void save(ShoppingCart* cart) override{
        // Simulate saving to SQL database
        cout << "Saving shopping cart to SQL database..." << endl;
    }
};

class NoSQLPersistence : public Persistence{
    public:
    void save(ShoppingCart* cart) override{
        // Simulate saving to NoSQL database
        cout << "Saving shopping cart to NoSQL database..." << endl;
    }
};

class FilePersistence : public Persistence{
    public:
    void save(ShoppingCart* cart) override{
        // Simulate saving to a file
        cout << "Saving shopping cart to a file..." << endl;
    }
};

int main(){
    ShoppingCart* cart = new ShoppingCart();
    cart->addProduct(new Product("Laptop", 999.99));
    cart->addProduct(new Product("Mouse", 49.99));
    cart->addProduct(new Product("Keyboard", 79.99));
    ShoppingCartPrinter* printer = new ShoppingCartPrinter(cart);
    printer->printReceipt();
    Persistence* sqlPersistence = new SQLPersistence();
    sqlPersistence->save(cart);
    Persistence* nosqlPersistence = new NoSQLPersistence();
    nosqlPersistence->save(cart);
    Persistence* filePersistence = new FilePersistence();
    filePersistence->save(cart);
    return 0;
}
