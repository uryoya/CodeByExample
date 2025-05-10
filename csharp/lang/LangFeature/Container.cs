namespace LangFeature;

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