// Listのパターンマッチ
val len3: List[Int] = List(1, 2, 3)
val len1: List[Int] = List(1)
val len0: List[Int] = List()

for (foo <- Seq(len0, len1, len3)) {
  foo match {
    case first :: second :: _ => println("first: " + first + " second: " + second)
    case one :: Nil => println("one: " + one)
    case Nil => println("0 length list")
    case _ => println("oops!")
  }
}
