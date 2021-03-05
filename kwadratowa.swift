import Foundation

var positiveDelta = false

print("Set a -> ",terminator: "")
let a = Float(readLine()!) ?? 0

print("You've set number to \(a)\nSet b -> ",terminator: "")
let b = Float(readLine()!) ?? 0

print("You've set number to \(b)\nSet c -> ",terminator: "")
let c = Float(readLine()!) ?? 0
print("You've set number to \(a)")

let r = kwadratowe(a,b,c)
if(positiveDelta){
    print("x1 = \(r[1])")
    print("x2 = \(r[2])")
}
print("Δ = \(r[0])")

func kwadratowe(_ a: Float,_ b: Float,_ c: Float) -> [Float]{
    if(a == 0){
        print("Bad arguments")
        return []
    }else{
    var r: [Float] = Array(repeating: 0, count: 3)
    let delta = pow(b,2) + (-4*a*c)
    r[0] = delta
    if (delta >= 0){
        r[1] = (-1 * b - delta.squareRoot())/(2*a)
        r[2] = (-1 * b + delta.squareRoot())/(2*a) 
        positiveDelta = true
    }else {positiveDelta = false}
    return r
    }   
}