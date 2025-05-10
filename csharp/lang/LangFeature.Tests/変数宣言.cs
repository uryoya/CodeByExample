using FluentAssertions;

namespace LangFeature.Tests;

/**
 * 変数
 *
 * - https://learn.microsoft.com/ja-jp/dotnet/csharp/language-reference/language-specification/variables
 * - https://learn.microsoft.com/ja-jp/dotnet/csharp/language-reference/statements/declarations
 */
public sealed class 変数宣言
{
    // https://learn.microsoft.com/ja-jp/dotnet/csharp/language-reference/statements/declarations
    [Fact]
    public void ローカル変数()
    {
        // 変数宣言と初期化を別にできる
        string greeting;
        int a, b, c;
        List<double> xs;
        // 宣言時に初期化することもできる
        double pi = 3.14;

        // greeting.Should().Be("Hello");
        // ^^^^^^^^ 未割り当ての屁数は使用できない

        greeting = "Hello";
        a = 1;
        b = 2;
        c = 3;
        xs = new List<double>();
        xs.Add(pi);

        greeting.Should().Be("Hello");
        a.Should().Be(1);
        b.Should().Be(2);
        c.Should().Be(3);
        xs.Should().BeEquivalentTo([3.14]);
    }

    // https://learn.microsoft.com/ja-jp/dotnet/csharp/language-reference/statements/declarations
    [Fact]
    public void ローカル定数()
    {
        const string Greeting = "こんにちは";
        const double MinLimit = -10.0, MaxLimit = 10.0;

        Greeting.Should().Be("こんにちは");
        MinLimit.Should().Be(-10.0);
        MaxLimit.Should().Be(10.0);
    }

    // https://learn.microsoft.com/ja-jp/dotnet/csharp/language-reference/statements/declarations#implicitly-typed-local-variables
    [Fact]
    public void 型推論()
    {
        var greeting = "Hello";
        var a = 32;
        var xs = new List<double>();

        greeting.Should().Be("Hello");
        a.Should().Be(32);
        xs.Should().BeOfType(typeof(List<double>));
    }

    [Fact]
    public void 参照変数()
    {
        // TODO: https://learn.microsoft.com/ja-jp/dotnet/csharp/language-reference/statements/declarations#reference-variables
    }

    [Fact]
    public void scoped_ref()
    {
        // TODO: https://learn.microsoft.com/ja-jp/dotnet/csharp/language-reference/statements/declarations#scoped-ref
    }
}
