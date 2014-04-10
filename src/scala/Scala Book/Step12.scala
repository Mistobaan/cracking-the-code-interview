import scala.io.Source

if (args.length > 0){
    for (line <- Source.fromFile(args(0)).getLines() ) {
        println(line.length + " " + line)
    }


    def widthOfLenght(s: String): s.length.toString.length

    val list = Source.fromFile(args(0)).getLines().toList()

    val longestLine = list.reduceLeft( (a, b) => if (a.length > b.length) a else b)

    val maxWidth = widthOfLenght(longestLine)

    

}
else
    Console.err.println("Please enter filename")


