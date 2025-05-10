namespace LangFeature.レコード;

public record Person(string FirstName, string LastName);

public record Person2
{
    public required string FirstName { get; init; }
    public required string LastName { get; init; }
}
