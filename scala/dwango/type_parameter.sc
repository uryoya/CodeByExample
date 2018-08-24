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
// 共変 (convariant)
{

}
