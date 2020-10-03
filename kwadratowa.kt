import kotlin.math.pow
import kotlin.math.sqrt

var positiveDelta = false

fun main(args: Array<String>){

    print("set a ->")
    var x = readLine()
    val a = x?.toFloat()
    if (0f == a) print("u dumb af")
    print("set b ->")
    x = readLine()
    val b = x?.toFloat()
    print("set c ->")
    x = readLine()
    val c = x?.toFloat()

    val r = if (a != null && b != null && c != null){
        println("quadratic function for a = $a b = $b c = $c")
        kwadratowa(a,b,c)
    }else{
        println("quadratic function for a = 1 b = 1 c = 1")
        kwadratowa(1f,1f,1f)
    }

    if(positiveDelta){
        println("x1 = ${r[1]}")
        println("x2 = ${r[2]}")
    }
    println("Δ = ${r[0]}")
}

fun kwadratowa(a: Float,b: Float,c: Float): FloatArray{
    val r = FloatArray(3)
    val delta = b.pow(2) - (4*a*c)
    r[0] = delta
    if (delta >= 0){
        r[1] = -1 * b - sqrt(delta)
        r[2] = -1 * b + sqrt(delta)
        positiveDelta = true
    }
    return r;
}