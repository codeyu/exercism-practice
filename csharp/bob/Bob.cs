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
        bool isYell(string str)
        {
            var hasUpper = false;
            foreach(var c in str)
            {
                if(Char.IsLetter(c))
                {
                    if(Char.IsUpper(c))
                    {
                        hasUpper = true;
                        continue;
                    }
                    return false;
                }
            }
            return hasUpper;
        }
        if(string.IsNullOrWhiteSpace(statement))
        {
            return ignore;
        }
        else if(isYell(statement))
        {
            return yell;
        }
        else if(statement.Trim().EndsWith("?"))
        {
            return sure;
        }
        return other;
    }
}