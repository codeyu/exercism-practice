using System;
using System.Linq;
public static class Bob
{
    public static string Hey(string statement)
    {
        var sure = "Sure.";
        var yell = "Whoa, chill out!";
        var ignore = "Fine. Be that way!";
        var other = "Whatever.";
        
        if(string.IsNullOrWhiteSpace(statement))
        {
            return ignore;
        }
        else if(IsYell(statement) && hasLetter(statement))
        {
            return yell;
        }
        else if(statement.Trim().EndsWith("?"))
        {
            return sure;
        }
        return other;
    }

    private static bool IsYell(string str) => str.Where(x=>Char.IsLetter(x)).All(x=>Char.IsUpper(x));
    private static bool hasLetter(string str) => str.Any(x=>Char.IsLetter(x));
}