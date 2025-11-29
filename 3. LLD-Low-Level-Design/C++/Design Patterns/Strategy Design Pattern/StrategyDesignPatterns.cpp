/*
Strategy Design Pattern :- Define a family of algorithms, put them into separate classes so that they can be cahanged at runtime.
*/

#include <iostream>
using namespace std;

// Strategy Interface for walking behavior
class WalkableRobot {
    public:
        virtual void walk() = 0;
        virtual ~WalkableRobot() {}
};

// Concrete Strategy for Walk
class NormalWalk : public WalkableRobot {
    public:
        void walk() override {
            cout << "Walking normally." << endl;
        }
};

// Concrete Strategy no walk
class NoWalk : public WalkableRobot {
    public:
        void walk() override {
            cout << "I cannot walk." << endl;
        }
};

// Strategy Interface for talk behavior
class TalkableRobot {
    public:
        virtual void talk() = 0;
        virtual ~TalkableRobot() {}
};

// Concrete Strategy for Talk
class NormalTalk : public TalkableRobot {
    public:
        void talk() override {
            cout << "Talking normally." << endl;
        }
};

// Concrete Strategy no talk
class NoTalk : public TalkableRobot {
    public:
        void talk() override {
            cout << "I cannot talk." << endl;
        }
};

// Strategy Interface for fly behavior
class FlyableRobot {
    public:
        virtual void fly() = 0;
        virtual ~FlyableRobot() {}
};

// Concrete Strategy for Fly
class NormalFly : public FlyableRobot {
    public:
        void fly() override {
            cout << "Flying normally." << endl;
        }
};

// Concrete Strategy no fly
class NoFly : public FlyableRobot {
    public:
        void fly() override {
            cout << "I cannot fly." << endl;
        }
};

// Robot base class

class Robot {
    protected:
        WalkableRobot* walkBehavior;
        TalkableRobot* talkBehavior;
        FlyableRobot* flyBehavior;

    public:
        Robot(WalkableRobot* w, TalkableRobot* t, FlyableRobot* f){
            this->walkBehavior = w;
            this->talkBehavior = t;
            this->flyBehavior = f;
        }
           
        void walk() {
            walkBehavior->walk();
        }

        void talk() {
            talkBehavior->talk();
        }

        void fly() {
            flyBehavior->fly();
        }

       virtual void projection() = 0;
};

// Concrete Robot class

class CompanionRobot : public Robot {
    public:
        CompanionRobot(WalkableRobot* w, TalkableRobot* t, FlyableRobot* f) : Robot(w, t, f) {}

        void projection() override {
            cout << "I am a Companion Robot." << endl;
        }
};

class WorkerRobot : public Robot {
    public:
        WorkerRobot(WalkableRobot* w, TalkableRobot* t, FlyableRobot* f) : Robot(w, t, f) {}

        void projection() override {
            cout << "I am a Worker Robot." << endl;
        }
};

int main() {
    Robot* robot1 = new CompanionRobot(new NormalWalk(), new NormalTalk(), new NoFly());
    robot1->projection();
    robot1->walk();
    robot1->talk();
    robot1->fly();
     
    cout << "------------------------" << endl;
    Robot* robot2 = new WorkerRobot(new NoWalk(), new NoTalk(), new NormalFly());
    robot2->projection();
    robot2->walk();
    robot2->talk(); 
    robot2->fly();


    return 0;
}