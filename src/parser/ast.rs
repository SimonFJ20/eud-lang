use crate::parser::tokenizer::{Token, TokenType};
use std::fmt;

pub enum StatementType {
    Declaration,
    FuncDef,
    Expression,
}

#[derive(Clone, Copy, Debug)]
pub enum ExpressionType {
    VarAccess,
    VarAssign,
    Add,
    Sub,
    Mul,
    Div,
    Exp,
    Int,
    FuncCall,
}

pub trait Statement {
    fn statement_type(&self) -> StatementType;
}

pub trait Expression {
    fn expression_type(&self) -> ExpressionType;
}

pub trait ExpressionStatement: Expression + Statement + fmt::Display {}

struct IntLiteral {
    tok: Token,
}

impl Expression for IntLiteral {
    fn expression_type(&self) -> ExpressionType {
        ExpressionType::Int
    }
}

impl Statement for IntLiteral {
    fn statement_type(&self) -> StatementType {
        StatementType::Expression
    }
}
impl ExpressionStatement for IntLiteral {}

impl fmt::Display for IntLiteral {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", self.tok)
    }
}

struct LeftRightExpression {
    expression_type: ExpressionType,
    left: Box<dyn ExpressionStatement>,
    right: Box<dyn ExpressionStatement>,
}

impl Expression for LeftRightExpression {
    fn expression_type(&self) -> ExpressionType {
        self.expression_type
    }
}
impl Statement for LeftRightExpression {
    fn statement_type(&self) -> StatementType {
        StatementType::Expression
    }
}
impl ExpressionStatement for LeftRightExpression {}

impl fmt::Display for LeftRightExpression {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            f,
            "{:?}({}, {})",
            self.expression_type(),
            self.left,
            self.right
        )
    }
}

pub struct Parser {
    tok: Option<Token>,
}

impl Parser {
    pub fn from(t: Token) -> Self {
        Self { tok: Some(t) }
    }
    pub fn parse(&mut self) -> Box<dyn ExpressionStatement> {
        self.make_expression()
    }
    fn make_expression(&mut self) -> Box<dyn ExpressionStatement> {
        return self.make_addition();
    }
    fn make_addition(&mut self) -> Box<dyn ExpressionStatement> {
        let left = self.make_subtraction();
        if self.tok.is_none() {
            return left;
        }
        if self.tok.as_ref().unwrap().token_type == TokenType::Add {
            self.next();
            let right = self.make_addition();
            return Box::new(LeftRightExpression {
                expression_type: ExpressionType::Add,
                left: left,
                right: right,
            });
        } else {
            return left;
        }
    }
    fn make_subtraction(&mut self) -> Box<dyn ExpressionStatement> {
        let left = self.make_multiplication();
        if self.tok.is_none() {
            return left;
        }
        if self.tok.as_ref().unwrap().token_type == TokenType::Sub {
            self.next();
            let right = self.make_subtraction();
            return Box::new(LeftRightExpression {
                expression_type: ExpressionType::Sub,
                left: left,
                right: right,
            });
        } else {
            return left;
        }
    }
    fn make_multiplication(&mut self) -> Box<dyn ExpressionStatement> {
        let left = self.make_division();
        if self.tok.is_none() {
            return left;
        }
        if self.tok.as_ref().unwrap().token_type == TokenType::Mul {
            self.next();
            let right = self.make_multiplication();
            return Box::new(LeftRightExpression {
                expression_type: ExpressionType::Mul,
                left: left,
                right: right,
            });
        } else {
            return left;
        }
    }
    fn make_division(&mut self) -> Box<dyn ExpressionStatement> {
        let left = self.make_exponentation();
        if self.tok.is_none() {
            return left;
        }
        if self.tok.as_ref().unwrap().token_type == TokenType::Mul {
            self.next();
            let right = self.make_division();
            return Box::new(LeftRightExpression {
                expression_type: ExpressionType::Div,
                left: left,
                right: right,
            });
        } else {
            return left;
        }
    }
    fn make_exponentation(&mut self) -> Box<dyn ExpressionStatement> {
        let left = self.make_value();
        if self.tok.is_none() {
            return left;
        }
        if self.tok.as_ref().unwrap().token_type == TokenType::Exp {
            self.next();
            let right = self.make_exponentation();
            return Box::new(LeftRightExpression {
                expression_type: ExpressionType::Exp,
                left: left,
                right: right,
            });
        } else {
            return left;
        }
    }
    fn make_value(&mut self) -> Box<dyn ExpressionStatement> {
        let token = self.tok.clone().unwrap();
        self.next();

        if token.token_type == TokenType::LParen {
            let expr = self.make_expression();
            if self.tok.as_ref().unwrap().token_type != TokenType::RParen {
                panic!("unexpected: token_type != RParen")
            }
            self.next();
            expr
        } else if match self.tok.as_ref().unwrap().token_type {
            TokenType::Int(_) => true,
            _ => false,
        } {
            Box::new(IntLiteral {
                tok: self.tok.clone().unwrap(),
            })
        } else if match token.token_type {
            TokenType::Int(_) => true,
            _ => false,
        } {
            Box::new(IntLiteral { tok: token })
        } else {
            self.make_expression()
        }
    }
    fn next(&mut self) {
        let next = self.tok.clone().unwrap().next;
        if next.is_none() {
            // finished parsing
        } else {
            self.tok = Some(*(next.unwrap()));
        }
    }
}
