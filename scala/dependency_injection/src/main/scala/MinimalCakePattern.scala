/* Minimal Cake パターンを使ったDIのサンプル。
 *
 * 参考: https://qiita.com/pab_tech/items/1c0bdbc8a61949891f1f
 */
// Serviceの依存関係を記述
trait UsesService {
  val service: Service
}

trait Service {
  def greeting(s: String): String
}

// Serviceに対する実装
trait MixInService extends UsesService {
  val service: Service = ServiceImpl
}

object ServiceImpl extends Service {
  def greeting(s: String): String = s"Hello ${s}!!"
}

// Serviceに対する実装 (モックアップ)
trait MixInMockService extends UsesService {
  val service: Service = ServiceMock
}

object ServiceMock extends Service {
  def greeting(s: String): String = s"Hello ${s}! I am mock object!"
}

// Client
trait Client extends UsesService {
  def hello(s: String): Unit = println(service.greeting(s))
}

// Clientに対して依存オブジェクトの注入
object Client extends Client with MixInService

// Clientに対して依存オブジェクト(モックアップ)の注入
object TestClient extends Client with MixInMockService

// main
object MinimalCakePattern extends App {
  val c = Client

  c.hello("scala")

  val cm = TestClient

  cm.hello("scala")
}
