import 'dart:io';
import 'dart:math';

void main() {
  print('Hello, World!');

  /* 変数宣言 */
  // 変数と型推論
  int num = 0;
  num = 1;
  var num_var = 0;
  num_var = 1;
  print({"num": num, "num_var": num_var});

  // 定数
  final num_final = 0; // finalは実行時に値を決定する
  // num_final = 0; // エラー
  const num_const = 0; // constはコンパイル時に値を決定する
  print({"num_final": num_final, "num_const": num_const});

  // 数値リテラル
  // int
  final x = 1;
  final hex = 0xFF;
  final exponent = 1.42e5;
  print({"x": x, "hex": hex, "exponent": exponent});
  // double
  final y = 1.1;
  final exponents = 1.42e5;
  print({"y": y, "exponents": exponents});

  // 文字列リテラル
  var str1 = 'Very famous guitarist';
  var str2 = "Very famous guitarist";
  print({"str1": str1, "str2": str2});

  final name = 'taro';
  print('Hello, $name');
  print('Hello, ${name.toUpperCase()}'); // 式の結果を文字列に埋め込む場合は${}で囲む

  // コレクションのリテラル
  // List
  var list = [1, 2, 3];
  print(list);
  // Set
  var set = {'猿', '雉', '犬'};
  print(set);
  // Map
  var map = {'猿': 'サル', '雉': 'キジ', '犬': 'イヌ'};
  print(map);
  final setOrMap = {};
  print(setOrMap.runtimeType); // Mapとして推論される

  /* Null安全 */
  int? nullableNum; // nullで初期化される
  nullableNum = nullOrZero();
  if (nullableNum != null) {
    print(nullableNum);
  }
  // null認識演算子
  String? nullStr;
  print(nullStr?.length); // null
  // 非null表明演算子
  String? nonNullStr = "Hello";
  print(nonNullStr!.length);
  // 非null昇格
  if (nonNullStr != null) {
    print(nonNullStr.length);
  }
  // ??演算子
  String? nullStr1 = 'some string';
  String? nullStr2;
  print(nullStr1 ?? 'default string'); // some string
  print(nullStr2 ?? 'default string'); // default string

  /* 関数 */
  String greet(String name) {
    return 'Hello, $name';
  }

  print(greet('taro'));

  // 引数
  // 省略可能引数
  void makeColor(int red, int green, int blue, [int? alpha]) {
    print({"red": red, "green": green, "blue": blue, "alpha": alpha});
  }

  makeColor(0xff, 0x00, 0x3e); // alphaを省略して呼び出し
  makeColor(0xff, 0x00, 0x3e, 0xff); // alphaを与えて呼び出し

  // 省略可能引数にはdefault値を与えることもできる
  void makeColor1(int red, int green, int blue, [int alpha = 0xff]) {
    print({"red": red, "green": green, "blue": blue, "alpha": alpha});
  }

  makeColor1(0xff, 0x00, 0x3e); // alphaを省略して呼び出し

  // 名前付き引数
  void makeColor2(int red, int green, int blue, {int? alpha}) {
    print({"red": red, "green": green, "blue": blue, "alpha": alpha});
  }

  makeColor2(0xff, 0x00, 0x3e, alpha: 0xff); // 名前を指定して呼び出し

  // 名前付き引数（required）
  void makeColor3(
      {required int red, required int green, required int blue, int? alpha}) {
    print({"red": red, "green": green, "blue": blue, "alpha": alpha});
  }

  makeColor3(red: 0xff, green: 0x00, blue: 0x4e); // 名前を指定して呼び出し
  makeColor3(green: 0xff, red: 0x00, blue: 0x4e); // 順番を入れ替えることもできる

  // 関数の省略記法
  // アロー構文
  int double(int x) => x * 2;
  print(double(2));

  // 第一級関数
  final int Function(int) f = double; // 関数doubleを変数fに代入
  print(f(2));

  // ラムダ式
  final int Function(int) g = (int x) {
    int i = 100;
    return x * x + i;
  };
  print(g(3));

  /* クラス */
}

int? nullOrZero() => null;

/* クラス */
// コンストラクタを記述しない場合は引数なしのデフォルトコンストラクタが暗黙的に定義される
class Point {
  int x = 0;
  int y = 0;
}

// コンストラクタ
class Point1 {
  int x;
  int y;
  Point1(int xPosition, int yPosition)
      : x = xPosition,
        y = yPosition;
}

// コンストラクタの糖衣構文
// コンストラクタ引数を直接初期値として与える記述には糖衣構文が用意されている
class Point2 {
  int x;
  int y;
  Point2(this.x, this.y);
}

// コンストラクタ本体が実行される前に非null許容型のフィールドは初期化する必要がある
// class Point3 {
//   int x;
//   int y;
//   Point3(int xPosition, int yPosition) {
//     x = xPosition;
//     // hello_world.dart:158:7: Error: Field 'x' should be initialized because its type 'int' doesn't allow null.
//     // int x;
//     //     ^
//     y = yPosition;
//     // hello_world.dart:159:7: Error: Field 'y' should be initialized because its type 'int' doesn't allow null.
//     //   int y;
//     //       ^
//   }
// }

// getter & setter
class Rectangle {
  Rectangle(this.width, this.height);

  final int width;
  final int height;

  // getを使ってgetterを定義
  int get area => width * height;
}

// constantコンストラクタ
// constantコンストラクタで定義すると、クラスインスタンスがコンパイル時定数として扱われる。
// インスタンス変数はすべてfinalである必要がある。
class ImmutablePoint {
  const ImmutablePoint(this.x, this.y);
  final int x;
  final int y;
}

// constantコンストラクタで作成したインスタンスはconst変数に代入した場合、常に同じインスタンスを参照する。
// singletonとは異なる。
const immutablePoint = ImmutablePoint(0, 0);

// 名前付きコンストラクタ
// 名前を持った複数のコンストラクタを定義することができる。
class Point4 {
  const Point4(this.x, this.y);
  const Point4.origin()
      : x = 0,
        y = 0; // 名前付きコンストラクタ

  final int x;
  final int y;
}

// factoryコンストラクタ
class UserData {
  static final Map<int, UserData> _cache = {};

  UserData();

  factory UserData.fromCache(int userId) {
    // キャッシュを探す
    final cache = _cache[userId];
    if (cache != null) {
      return cache;
    }

    // キャッシュがなければ新規作成
    final newInstance = UserData();
    _cache[userId] = newInstance;
    return newInstance;
  }
}

// クラスの継承
class Animal {
  String greet() => 'hello';
}

class Dog extends Animal {}

// オーバーライドする場合はoverrideキーワードをつけることが推奨されている
class Cat extends Animal {
  @override
  String greet() => 'meow';
}

var sample1 = () {
  final dog = Dog();
  print(dog.greet()); // hello
  final cat = Cat();
  print(cat.greet()); // meow
};

// 親クラスのコンストラクタ
// 親クラスのコンストラクタに引数がある場合は、明示的に子クラスのコンストラクタで親クラスのコンストラクタを呼び出す必要がある
class Animal1 {
  Animal1(this.name);
  final String name;
}

class Dog1 extends Animal1 {
  Dog1(String name) : super(name);
}

class Cat1 extends Animal1 {
  // 親クラスのコンストラクタに引数をそのまま渡す場合の糖衣構文がある
  Cat1(super.name);
}

var sample2 = () {
  final dog = Dog1('dog');
  print(dog.name); // dog
  final cat = Cat1('cat');
  print(cat.name); // cat
};

// 暗黙のインターフェース
// Dartでは明示的なインターフェースがない代わりにクラスは暗黙のうちにインターフェースを実装する
// implementsでは全てのメンバ、メソッドを実装する必要があるのがextendsとの違い
class Person implements Animal {
  @override
  String greet() => 'hello';
}

// mixin
mixin class Hose {
  void run() {
    print('run');
  }
}

mixin class Bird {
  void fly() {
    print('fly');
  }
}

class Pegasus with Hose, Bird {}

var sample3 = () {
  final pegasus = Pegasus();
  pegasus.run(); // run
  pegasus.fly(); // fly
};

// キャスト演算子
// as
class Vehicle {
  void run() {
    print('running');
  }
}

class Car extends Vehicle {
  @override
  void run() {
    print('ブーン');
  }
}

class M1M2Abrams extends Vehicle {
  void shoot() {
    print('boom!!!');
  }
}

var sample4 = () {
  final vehicle1 = M1M2Abrams();
  (vehicle1 as M1M2Abrams).shoot(); // boom!!!
  final vehicle2 = Car();
  (vehicle2 as M1M2Abrams).shoot(); // エラー
};

// is
// TypeScriptのような型の絞り込みができる
var sample5 = () {
  Vehicle vehicle1 = M1M2Abrams();
  if (vehicle1 is M1M2Abrams) {
    vehicle1.shoot(); // boom!!!
  }
};

/* 非同期処理 */
void syncDoSomething() {
  const path = 'some/file/path.txt';
  Future<String> content = File(path).readAsString();
  content.then((content) {
    print(content);
  });
}

Future<void> asyncDoSomething() async {
  const path = 'some/file/path.txt';
  final content = await File(path).readAsString();
  print(content);
}
