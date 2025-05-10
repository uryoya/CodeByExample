namespace LangFeature.クラス;

public class Container
{
    public int Capacity1 = 10; // 規定値
    public int Capacity2; // コンストラクタによる指定

    public Container(int capacity2) => Capacity2 = capacity2;
}

public class Container2(int capacity)
{
    public int Capacity = capacity; // プライマリコンストラクタによる宣言
}

public class Customer()
{
}

/**
 * オブジェクト初期化子
 *
 * フィールド宣言時に `required` を使うとインスタンス作成時にプロパティ名
 * を明示して初期化することができる。
 */
public class Person
{
    public required string LastName { get; set; }
    public required string FirstName { get; set; }

    public string FullName()
    {
        return $"{LastName} {FirstName}";
    }
}
