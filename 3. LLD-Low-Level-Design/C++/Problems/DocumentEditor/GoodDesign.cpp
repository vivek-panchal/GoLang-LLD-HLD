#include<iostream>
#include <bits/stdc++.h>
using namespace std;

class DocumentElement {
    public:
    virtual string render() = 0; 
};

class TextElement : public DocumentElement {
    private:
    string text;
    
    public:
    TextElement(string text) {
        this->text = text;
    }
    
    string render() override {
        return text;
    }
};

class ImageElement : public DocumentElement {
    private:
    string imagePath;
    
    public:
    ImageElement(string path) {
        this->imagePath = path;
    }
    
    string render() override {
        return "<img src='" + imagePath + "' />";
    }
};

class NewLineElement : public DocumentElement {
    public:
    string render() override {
        return "\n";
    }
};

class TabSpaceElement : public DocumentElement {
    public:
    string render() override {
        return "\t";
    }
};

class Document {
    private:
    vector<DocumentElement*> documentElements;
    
    public:
    void addElement(DocumentElement* element) {
        documentElements.push_back(element);
    }
    
    string render() {
        string result;
        for(auto element : documentElements) {
            result += element->render();
        }
        return result;
    }
};

class Persistence {
    public:
    virtual void save (string data) = 0;
};

class FileStorage : public Persistence {
   public:
   void save ( string data ) override {
       ofstream outFile("document.txt");
       if(outFile){
            outFile << data;
            outFile.close();
            cout<<"Document saved to document.txt"<<endl;
        }
        else {
            cout << "Error opening file for writing." << endl;
        }
    }
};

class DatabaseStorage : public Persistence {
   public:
   void save ( string data ) override {
       // Simulate saving to a database
       cout << "Document saved to database: " << data << endl;
    }
};

class DocumentEditor {
    private:
    Document* document;
    Persistence* storage;
    string renderedDocument;

    public:
    DocumentEditor(Document* doc, Persistence* storage) {
        this->document = doc;
        this->storage = storage;
    }

    void addText(string text) {
        document->addElement(new TextElement(text));
    }

    void addImage(string path) {
        document->addElement(new ImageElement(path));
    }

    void addNewLine() {
        document->addElement(new NewLineElement());
    }

    void addTabSpace() {
        document->addElement(new TabSpaceElement());
    }

    string renderDocument() {
        if(renderedDocument.empty()) {
            renderedDocument = document->render();
        }
        return renderedDocument;
    }

    void saveDocument() {
        storage->save(renderedDocument);
    }
};

int main() {
    Document* doc = new Document();
    Persistence* storage = new FileStorage(); // Change to DatabaseStorage() to save to database
    DocumentEditor* editor = new DocumentEditor(doc, storage);
    
    editor->addText("Hello, World!");
    editor->addNewLine();
    editor->addTabSpace();
    editor->addText("This is a sample document.");
    editor->addNewLine();
    editor->addImage("image.png");
    
    string rendered = editor->renderDocument();
    cout << "Rendered Document:\n" << rendered << endl;
    
    editor->saveDocument();
    
    return 0;
}