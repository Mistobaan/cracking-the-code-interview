
val hex = 0x5

println(hex)

val hex2 = 0x00FF

println(hex2)


val double1 = 3e5d

println(double1)

// unicode

val uni1 = '\u0041'

println(uni1)


val B\u0041\u0044 = 1

val pythonStyleString = """

    Multiline whooo

    """

println(pythonStyleString)

val pipeTrick = """|0
    |A
    |B
    |C
    """.stripMargin

println(pipeTrick)


// Symbols

val mySymbol = 'favoriteAlbum

println(mySymbol)


val s = "Hello World"

println(s indexOf 'o')

// prefix operators

println((2.0).unary_-)


def pepper () = { println("pepper"); true}
def salt () = { println("salt"); false}

pepper() && salt()

salt() && pepper()
