using FluentAssertions;
using LangFeature;

namespace LangFeature.Tests;

/**
 * クラス
 *
 * - https://learn.microsoft.com/ja-jp/dotnet/csharp/fundamentals/types/classes
 * - https://learn.microsoft.com/ja-jp/dotnet/csharp/fundamentals/object-oriented/
 */
public sealed class クラス
{
    // https://learn.microsoft.com/ja-jp/dotnet/csharp/fundamentals/types/classes#declaring-classes
    // https://learn.microsoft.com/ja-jp/dotnet/csharp/fundamentals/types/classes#creating-objects
    [Fact]
    public void クラスの宣言と作成()
    {
        // 宣言は /LangFeature/Customer.cs
        Customer customer = new Customer();

        customer.Should().BeOfType(typeof(Customer));
    }

    // https://learn.microsoft.com/ja-jp/dotnet/csharp/fundamentals/types/classes#constructors-and-initialization
    [Fact]
    public void コンストラクタと初期化()
    {
        // 宣言は /LangFeature/Container.cs
        var container1 = new Container(20);
        var container2 = new Container2(100); // 実装側はプライマリコンストラクタを使用
        // 宣言は /LangFeature/Person.cs
        var person = new Person
        {
            FirstName = "乱歩",
            LastName = "江戸川"
        };

        container1.Capacity1.Should().Be(10);
        container1.Capacity2.Should().Be(20);
        container2.Capacity.Should().Be(100);
        person.FullName().Should().Be("江戸川 乱歩");
    }

    // https://learn.microsoft.com/ja-jp/dotnet/csharp/fundamentals/types/classes#class-inheritance
    [Fact]
    public void 継承()
    {
        var manager =  new Manager();
        
        manager.Should().BeAssignableTo<Employee>();
    }
}