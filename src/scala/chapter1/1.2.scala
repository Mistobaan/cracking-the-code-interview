
import scala.runtime.RichChar

def reverseInPlace(s: String): String = {
   val reverse = new Array[RichChar](s.length)   
   for (i <- (0 to (s.length >> 1))) {
     reverse(i) = s(s.length -i -1)
     reverse(s.length -i -1) = s(i)
   }
   return reverse.mkString
}

def reverseLeft(s: String): String = s.foldLeft("") ( (a,b) => 
    b + a
)

def reverseRight(s: String): String = s.foldRight("") ( (a,b) => 
    b + a
)


def time[R](iterations:Int, block: => R) = {
    val t0 = System.nanoTime()
    for ( i <- 0 to iterations){
       block    // call-by-name
    }
    val t1 = System.nanoTime()
    println("Elapsed time: " + (t1 - t0) + "ns")
}



time(1000, {
    "Hello\u0041".reverse
})

time(1000, {
    reverseRight("Hello\u0041")
})

time(1000, {
    reverseInPlace("Hello\u0041")
})

time(1000, {
    reverseLeft("Hello\u0041")
})
