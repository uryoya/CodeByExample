/**
 * ドワンゴ研修資料
 * https://dwango.github.io/scala_text/trait.html
 *
 * トレイト
 */

// トレイトの基本
{
  trait TraitA

  trait TraitB

  class ClassA

  class ClassB

  // コンパイルできる
  class ClassC extends ClassA with TraitA with TraitB

  // コンパイルできない
  // class ClassD extends ClassA with ClassB

  // 直接インスタンス化できない
  // val a = new TraitA
}



// 菱形継承問題
{
  trait TraitA {
    def greet(): Unit
  }

  trait TraitB extends TraitA {
    def greet(): Unit = println("Good morning!")
  }

  trait TraitC extends TraitA {
    def greet(): Unit = println("Good evening!")
  }

  // コンパイルできない
  // class ClassA extends TraitA with TraitB with TraitC

  // ClassA で greet を override する
  class ClassA extends TraitB with TraitC {
    override def greet(): Unit = println("How are you?")
  }

  // superに肩を指定してメソッドを呼び出すことでTraitBやTraitCのメソッドを
  // 指定して使うこともできる
  class ClassB extends TraitB with TraitC {
    override def greet(): Unit = super[TraitB].greet()
  }

  (new ClassA).greet()

  (new ClassB).greet()
}

// 線形化 (linearization)
{
  trait TraitA {
    def greet(): Unit
  }

  trait TraitB extends TraitA {
    override def greet(): Unit = println("Good morning!")
  }

  trait TraitC extends TraitA {
    override def greet(): Unit = println("Good evening!")
  }

  // トレイトがミックスインされた順番をトレイトの継承順番と見做す
  class ClassA extends TraitB with TraitC

  // 後からミックスインしたTraitCが優先される
  (new ClassA).greet()

  class ClassB extends TraitC with TraitB
  (new ClassB).greet()
}

// 自分型
