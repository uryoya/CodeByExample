using Xunit;
using LangFeature;  // テスト対象の名前空間
using FluentAssertions;

namespace LangFeature.Tests
{
    public class SampleTests
    {
        [Fact]
        public void 足し算が正しく動く()
        {
            // Arrange
            var calc = new Calculator();  // 仮にLangFeatureにあるクラス

            // Act
            var result = calc.Add(2, 3);

            // Assert
            Assert.Equal(5, result);

            result.Should().Be(5);
        }
    }
}

class Calculator
{
    public int Add(int a, int b)
    {
        return a + b;
    }
}
