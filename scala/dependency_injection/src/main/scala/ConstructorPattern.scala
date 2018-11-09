/* コンストラクタパターンを使ったDIのサンプル。
 *
 * ClientはServiceを実装したオブジェクトをコンストラクタから注入されることで
 * Serviceのメソッドを使うことができる。
 */
object ConstructorPattern extends App {
  trait Service {
    def greeting(s: String): String
  }

  class ServiceImpl extends Service {
    def greeting(s: String): String = s"Hello ${s}!!"
  }

  class ServiceMock extends Service {
    def greeting(s: String): String = s"Hello ${s}! I am mock object!"
  }

  class Client(service: Service) {
    def hello(s: String): Unit = println(service.greeting(s))
  }

  val c = new Client(new ServiceImpl)
  c.hello("scala")

  val cm = new Client(new ServiceMock)
  cm.hello("scala")
}
