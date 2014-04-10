

class Rational(n: Int, d: Int){
    require(d!=0, "Argument must be positive")
    val numerator: Int = n
    val denom: Int = d

    def this(n:Int) = this(n, 1)

    override def toString() = n + "/" + d
    def +(other: Rational): Rational = {
        new Rational(numerator + other.numerator , d + other.denom)
    }

    def *(other:Rational)() = new Rational(numerator * other.numerator, denom * other.denom)
}

println(new Rational(10, 100) + new Rational(100, 30))


println(new Rational(5))


implicit def intToRational(x:Int) = new Rational(x)

println(5 * new Rational(5) )