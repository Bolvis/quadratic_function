import math

a = input("set a ->")
b = input("set b ->")
c = input("set c ->")

def kwadratowa(a, b, c):
    x1 = 0.0
    x2 = 0.0
    if a == 0:
        return [null,null,null],False
    delta = pow(b,2) - (4*a*c)
    if delta >= 0:
        x1 = (-1 * b - math.sqrt(delta))/ (2 * a)
        x2 = (-1 * b + math.sqrt(delta))/ (2 * a)
        possitiveDelta = True
    else:
        possitiveDelta = False
    return [delta,x1,x2],possitiveDelta

r,possitiveDelta = kwadratowa(float(a),float(b),float(c))

if possitiveDelta == True:
    print("x1 = ",r[1])
    print("x2 = ",r[2])
print("Δ = ",r[0])
