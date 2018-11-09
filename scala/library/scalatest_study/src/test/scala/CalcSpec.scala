import org.scalatest._
import org.scalatest.concurrent.TimeLimits
import org.scalatest.time.SpanSugar._

class CalcSpec extends FlatSpec with DiagrammedAssertions with TimeLimits {
  //https://dwango.github.io/scala_text/test.html
  val calc = new Calc

  "sum関数" should "整数の配列を取得し、それらを足し合わせた整数を返すことができる" in {
    assert(calc.sum(Seq(1, 2, 3)) === 6)
    //    assert(calc.sum(Seq(1, 2, 3)) === 7) // test will be failed
    assert(calc.sum(Seq(0)) === 0)
    assert(calc.sum(Seq(-1, 1)) === 0)
    assert(calc.sum(Seq()) === 0)
  }

  it should "Intの最大を上回った際にはオーバーフローする" in {
    assert(calc.sum(Seq(Integer.MAX_VALUE, 1)) === Integer.MIN_VALUE)
  }

  "isPrime関数" should "その値が素数であるかどうかのブール値を返す" in {
    assert(calc.isPrime(0) === false)
    assert(calc.isPrime(-1) === false)
    assert(calc.isPrime(2))
    assert(calc.isPrime(17))
  }
  it should "100万以下の値の素数判定を１秒以内で処理できる" in {
    failAfter(1000 millis) {
      assert(calc.isPrime(9999991))
    }
  }
}
