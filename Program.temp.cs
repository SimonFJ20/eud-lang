using System;
using System.Collections.Generic;
using System.Linq;

enum TokenType {
    INT,
    FLOAT,
    CHAR,
    STRING,
    IDENTIFIER,
    KEYWORD,

    ADD_OP,
    SUB_OP,
    MUL_OP,
    DIV_OP,
    MOD_OP,
    EXP_OP,
    LNOT_OP,
    BNOT_OP,
    INCR_OP,
    DECR_OP,
    MBAC_OP,
    COMB_OP,

    LPAREN,
    RPAREN,
    LBRACE,
    RBRACE,
    LBRACKET,
    RBRACKET,

    NEWLINE,
    EOF,
}

class Token {
    private TokenType type;
    private string text = "";

    public string[] keywords = {
        "if",
        "then",
        "else",
        "while",
        "for",
        "do",
        "continue",
        "break",
        "match",
        "func",
        "return",
        "end",
        "i16",
        "i32",
        "i64",
        "u8",
        "u16",
        "u32",
        "u64",
        "f32",
        "f64",
    };

    public Token(TokenType type, string text)
    {
        this.type = type;
        this.text = text;
    }

    public Token(TokenType type)
    {
        this.type = type;
    }
}

class Lexer {
    private int pos = 0;
    private string text = "";
    private char c = '\0';
    private bool done = false;

    public Token[] Tokenize(string text)
    {
        pos = 0;
        this.text = text;
        done = false;
        Next();

        LinkedList<Token> tokens = new LinkedList<Token>();

        for (; !done; Next())
        {
            if (IsNumber(c))
                tokens.Add(IntOrFloat());
            else if (IsLetter(c))
                tokens.Add(IdentifierOrKeyword());
            else {
                switch (c) {
                    case ' ':
                    case '\n':
                    case '\r':
                    case '\t':
                        Next();
                        break;
                    case '\'':
                        tokens.Add(CharLiteral());
                        break;
                    case '"':
                        tokens.Add(StringLiteral());
                        break;
                    
                }
            }
        }
        tokens.Add(new Token(TokenType.EOF));
        return tokens.ToArray();
    }

    void IsLetter(char c) {
        return c >= 'a' && c <= '<' && c >= 'A' && c <= 'Z';
    }

    void IsNumber(char c) {
        return c >= '0' && c <= '9';
    }

    void Next()
    {
        pos++;
        if (pos < text.Length)
            c = text[pos];
        else
        {
            c = '\0';
            done = true;
        }
    }
}

interface Node
{

}

class Parser
{
    private int pos = 0;
    private Token[] tokens;
    private Token tok = new Token(TokenType.EOF);
    private bool done = false;

    public Node Parse(Token[] tokens)
    {
        pos = 0;
        this.tokens = tokens;
        done = false;
        Next();

        for (; !done; Next())
        {
        }
        
        throw new NotImplementedException();

    }

    void Next()
    {
        pos++;
        if (pos < tokens.Length)
            tok = tokens[pos];
        else
        {
            tok = new Token(TokenType.EOF);
            done = true;
        }
    }

}

public class Program
{
    public static void Main(string[] args)
    {
        Console.Out.WriteLine("Legit fuck c#");
    }
}
