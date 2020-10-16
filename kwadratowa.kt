import kotlin.math.pow
import kotlin.math.sqrt

fun main() {
    val a = readNumber("A")
    val b = readNumber("B")
    val c = readNumber("C")

    if(a == 0f) return println("a == 0")

    val delta = b.pow(2) - (4 * a * c)
    if(delta > 0) {
        val x1 = (-b - sqrt(delta)) / (2 * a)
        val x2 = (-b + sqrt(delta)) / (2 * a)
        println("X1: $x1, X2: $x2")
    }
    else if(delta == 0f) println("X: ${-b / (2 * a)}")
    else println("Delta < 0")
}

fun readNumber(name: String): Float = print("Podaj $name: ").run {
    readLine()?.toFloatOrNull() ?: readNumber(name)
}
