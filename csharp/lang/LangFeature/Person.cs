namespace LangFeature;

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