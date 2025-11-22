#include <iostream>
#include <vector>

using namespace std;

class Product{
    public:
    string name;
    double price;
    Product(string name , double price) {
        this->name = name;
        this->price = price;
    }
};
class ShoppingCart{
    private:
    vector<Product*> products;
    
    public:
    void addProduct(Product* product) {
        products.push_back(product);
    }
    const vector<Product*>& getProducts() const {
        return products;
    }

    double calculateTotal() {
        double total = 0;
        for (const auto& product : products) {
            total += product->price;
        }
        return total;
    }

    void printReceipt() {
        cout << "Receipt:" << endl;
        for (const auto& product : products) {
            cout << product->name << " - Rs " << product->price << endl;
        }
        cout << "Total: Rs " << calculateTotal() << endl;
    }

    void saveToFile(const string& filename) {
        // Simulating file saving
        cout << "Saving receipt to " << filename << endl;
    }
};

int main(){

    ShoppingCart* cart = new ShoppingCart();
    cart->addProduct(new Product("Book", 500));
    cart->addProduct(new Product("Pen", 50));
    cart->printReceipt();
    cart->saveToFile("receipt.txt");

    cout << "SRP Violated Example in C++" << endl;
    return 0;
}