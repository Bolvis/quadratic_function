#include <iostream>
#include <cmath>

float * kwadratowa(float a, float b, float c);
float getNumber();

int main(){
    float a,b,c;
    
    do {
        if (a == 0) std::cout<<"a can not be 0"<<std::endl;
        std::cout<<"Set a ->" ;
        a = getNumber();
    } while (a == 0);
    
    std::cout<<"Set b ->" ;
    b = getNumber();
    std::cout<<"Set c ->" ;
    c = getNumber();

    float * r = kwadratowa(a,b,c);
    if(r[3]){
        std::cout<<"x1 = "<<r[1]<<std::endl;
        std::cout<<"x2 = "<<r[2]<<std::endl;
    }
    std::cout<<"Δ = "<<r[0]<<std::endl;
    
    return 0;
}

float * kwadratowa(float a, float b, float c){
    static float r[4];
    float delta = powf(b,2) + (-4*a*c);
    r[0] = delta;
    if(delta >= 0){
        r[1] = (-b - sqrtf(delta)) / (2*a);
        r[2] = (-b + sqrtf(delta)) / (2*a);
        r[3] = 1;
    }
    return r;
}

float getNumber(){
    float num;
    std::cin>>num;
    while (std::cin.fail())
    {
        std::cin.clear();
        std::cin.ignore();
        std::cout << "Not a valid number. Please reenter ->";
        std::cin >> num;
    }
    std::cout<<"You've set number to "<<num<<std::endl;
    return num;
}
