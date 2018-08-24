/**
 * 型パラメータ
 * https://dwango.github.io/scala_text/type-parameter.html
 *
 * 文法
 * ```
 * class <クラス名>[<型パラメータ1>, <型パラメータ2>, ...](<クラス引数>) {
 *   (<フィールド定義>|<メソッド定義>)
 * }
 * ```
 */

// 基本
{
  class Cell[A](var value: A) {
    def put(newValue: A): Unit = {
      value = newValue
    }

    def get(): A = value
  }

  val cell = new Cell[Int](1)
  cell.put(2)
  cell.get() // Int = 2
  // cell.put("Something") // error: type mismatch;

  // 使用例: 2値を返したい時
  class Pair[A, B](val a: A, val b: B) {
    override def toString(): String = "(" + a + "," + b + ")"
  }

  def devide(m: Int, n: Int): Pair[Int, Int] = new Pair[Int, Int](m / n, m % n)
  val result = devide(7, 3)
  println(result) // Pair[Int, Int] = (2, 1)
  // memo: 同様のケースはよくあるので、TupleがScalaには用意されている
}

// 変位指定 (variance)
//  https://dwango.github.io/scala_text/type-parameter.html#%E5%A4%89%E4%BD%8D%E6%8C%87%E5%AE%9A%EF%BC%88variance%EF%BC%89

// 共変 (convariant)
//
// 何も指定しない型パラメータは非変 (invariant) になる。
// 非変とは型パラメータを持つクラスG、型パラメータA BがあったときA=Bの時のみ
//   val : G[A] = G[B]
// という代入が許される性質。
//
// 共変とは、型パラメータを持つクラスG、型パラメータA Bがあったとき、
// AがBを継承している (B >: A) ときにのみ
// という代入が許される性質。
//
// 文法:
// ```
// class G[+A]  // Aは共変
// ```
{
  class Pair[+A, +B](val a: A, val b: B) {
    override def toString(): String = "(" + a + "," + b + ")"
  }

  val pair: Pair[AnyRef, AnyRef] = new Pair[String, String]("foo", "bar")
  println(pair)
}
