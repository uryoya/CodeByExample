using FluentAssertions;
using LangFeature.レコード;

namespace LangFeature.Tests;

/**
 * レコード
 *
 * - https://learn.microsoft.com/ja-jp/dotnet/csharp/fundamentals/types/records
 */
public sealed class レコード
{
    [Fact]
    public void インスタンス化()
    {
        var person = new Person("安吾", "坂口");
        var person2 = new Person2
        {
            FirstName = "公房",
            LastName = "安部",
        };

        person.FirstName.Should().Be("安吾");
        person.LastName.Should().Be("坂口");

        person2.FirstName.Should().Be("公房");
        person2.LastName.Should().Be("安部");
    }

    [Fact]
    public void 等価性()
    {
        var person1 = new Person("安吾", "坂口");
        var person2 = new Person("安吾", "坂口");

        (person1 == person2).Should().BeTrue();
        ReferenceEquals(person1, person2).Should().BeFalse();
    }

    [Fact]
    public void コピー()
    {
        var ranpo = new Person("乱歩", "江戸川");

        var conan = ranpo with { FirstName = "コナン" };
        var ranpo2 = ranpo with { };

        conan.Should().Be(new Person("コナン", "江戸川"));
        (conan == ranpo).Should().BeFalse();
        ranpo2.Should().Be(ranpo);
        (ranpo2 == ranpo).Should().BeTrue();
    }
}
