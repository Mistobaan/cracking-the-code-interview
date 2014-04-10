
def main() = {
    var i = 0
    while ( i < args.length){
        println("" + i + " " + args(i))
        i += 1
    }
}


// Chapter 2

def echoargs() = args.foreach( a => println(a) )

def echoargsV2() = args.foreach(println)

def for_eacho_args() = for ( e <- args) println(e)


// Chapter 3

// types with one operator, apply, equals
// 
class S() {
    var v = 0

    def update(value: Int, other: Int){
        v = other * value
    }

    def apply(value: Int): S = {
        v *= value
        this
    }

    def to(value: Int) = {
        new Array(value)
    }
}

def PlayWithS() = {
    val s = new S
    println(s(10).v)
    s(10) = 20

    println(s to 10)

    println(s.v)
}

PlayWithS()

// If the operator ends with a colon then is right applied
def ListPrepending() = {

    val first = "a" :: "b" :: Nil
    val second = "Second" :: Nil

    val third = first ::: second

    println(first)
    println(second)
    println(third)
}

ListPrepending()


def ListPractice(){
   val l = List("Cool", "tools", "rule")
   val thrill = "Will" :: "fill" :: "until" :: Nil

   thrill(2)
   thrill.count( s => s.length == 4)
   thrill.tail

}


def TuplePractice() = {

    val tuple1 = (0, "ciao")

    println( tuple1._1)
    println( tuple1._2 )

}

TuplePractice()

import scala.collection.mutable.Set

def TestMutableSet() = {
    var immutable =  Set("Hello", "World", "Hello")
    immutable += "Hello"
    println(immutable)

    val mutableSet = Set("Hello", "World", "Hello")
    mutableSet += "Hello"
    println(mutableSet)
}

TestMutableSet()


def TestStrangeObject() = {
    println( (1 -> "Ciao")._2 )
}

TestStrangeObject()




