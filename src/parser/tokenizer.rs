use std::fmt;

#[derive(PartialEq, Clone)]
pub enum TokenType {
    Add,
    Sub,
    Mul,
    Div,
    Exp,
    LParen,
    RParen,
    Int(u32),
    Char(char),
    Identifier(String),
    Keyword(String),
    Whitespace,
}

#[derive(Clone)]
pub struct Token {
    pub token_type: TokenType,
    pub next: Option<Box<Token>>,
}

fn tokenize_char(c: char) -> Token {
    match c {
        '+' => Token {
            token_type: TokenType::Add,
            next: None,
        },
        '-' => Token {
            token_type: TokenType::Sub,
            next: None,
        },
        '*' => Token {
            token_type: TokenType::Mul,
            next: None,
        },
        '/' => Token {
            token_type: TokenType::Div,
            next: None,
        },
        '^' => Token {
            token_type: TokenType::Exp,
            next: None,
        },
        '(' => Token {
            token_type: TokenType::LParen,
            next: None,
        },
        ')' => Token {
            token_type: TokenType::RParen,
            next: None,
        },
        '0'..='9' => Token {
            token_type: TokenType::Int((c as u32) - 48),
            next: None,
        },
        ' ' | '\n' => Token {
            token_type: TokenType::Whitespace,
            next: None,
        },
        'a'..='z' | 'A'..='Z' => Token {
            token_type: TokenType::Char(c),
            next: None,
        },
        _ => panic!("invalid character {}", c),
    }
}

pub fn tokenize_string(s: String) -> Token {
    s.chars()
        .map(|c| tokenize_char(c))
        .filter(|t| t.token_type != TokenType::Whitespace)
        .rev()
        .reduce(|acc, t| Token {
            next: Some(Box::new(acc)),
            token_type: t.token_type,
        })
        .unwrap_or(Token {
            token_type: TokenType::Whitespace,
            next: None,
        })
}

impl fmt::Display for Token {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", self.token_type,)
    }
}

impl fmt::Display for TokenType {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            f,
            "{}",
            match self {
                TokenType::Add => "add".to_string(),
                TokenType::Sub => "sub".to_string(),
                TokenType::Mul => "mul".to_string(),
                TokenType::Div => "div".to_string(),
                TokenType::Exp => "exp".to_string(),
                TokenType::LParen => "l_paren".to_string(),
                TokenType::RParen => "r_paren".to_string(),
                TokenType::Int(x) => format!("int({x})"),
                TokenType::Char(c) => format!("char({c})"),
                TokenType::Identifier(s) => format!("id({s})"),
                TokenType::Keyword(s) => format!("keyword({s})"),
                TokenType::Whitespace => "whitespace".to_string(),
            }
        )
    }
}
