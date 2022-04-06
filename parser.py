from enum import Enum, auto
import sys
from typing import List, Optional

class File:
    def __init__(self, filename: str) -> None:
        self.filename = filename
    
    def text(self):
        with open(self.filename) as file:
            return file.read()

class Position:
    def __init__(self, row: int, col: int, filename: str) -> None:
        self.row = row
        self.col = col
        self.filename = filename
    
    def copy(self):
        return Position(self.row, self.col, self.filename)

    def next(self, linefeed: bool):
        if linefeed:
            self.row += 1
            self.col = 1
        else:
            self.col += 1

    def __repr__(self) -> str:
        return f'{self.filename}:{self.row}:{self.col}:'

def fail(msg: str, pos: Optional[Position]):
    print(f'FAILED: {msg}')
    if pos: print(f'    at {pos}')
    exit(1)

class TT(Enum):
    EOF = auto()
    IDENTIFIER = auto()
    KEYWORD = auto()
    INT = auto()
    LPAREN = auto()
    RPAREN = auto()
    LBRACKET = auto()
    RBRACKET = auto()
    LBRACE = auto()
    RBRACE = auto()
    ADD_OP = auto()
    SUB_OP = auto()
    MUL_OP = auto()
    DIV_OP = auto()
    MOD_OP = auto()
    EXP_OP = auto()
    ASGN_OP = auto()
    COLON = auto()
    COMMA = auto()

class Token:
    def __init__(self, type: TT, value: str, fp: Position) -> None:
        self.type = type
        self.value = value
        self.fp = fp
        self.first_on_line = False

    def __repr__(self) -> str:
        return f"[{str(self.type)[3:]}:'{self.value}']"

KEYWORDS: List[str] = [
    'if',
    'else',
    'for',
    'while',
    'break',
    'fn',
    'return',
    'end',
    'let',
    'u8',
    'u16',
    'u32',
    'u64',
    'i8',
    'i16',
    'i32',
    'i64',
    'char',
    'usize',
    'uptr',
]

class Lexer:
    def tokenize(self, file: File) -> List[Token]:
        text = file.text()
        if len(text) <= 0:
            raise fail('no text to be tokenized', None)
        self.text = text
        self.pos = 0
        self.c = text[0]
        self.done = False
        self.fp = Position(1, 1, file.filename)
        tokens: List[Token] = []
        while not self.done:
            if self.c in '\n\t ':
                self.next()
            elif self.c in 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_':
                tokens.append(self.make_name());
            elif self.c in '123456789':
                tokens.append(self.make_int())
            elif self.c == '(':
                tokens.append(Token(TT.LPAREN, self.c, self.fp.copy()))
                self.next()
            elif self.c == ')':
                tokens.append(Token(TT.RPAREN, self.c, self.fp.copy()))
                self.next()
            elif self.c == '[':
                tokens.append(Token(TT.LBRACKET, self.c, self.fp.copy()))
                self.next()
            elif self.c == ']':
                tokens.append(Token(TT.RBRACKET, self.c, self.fp.copy()))
                self.next()
            elif self.c == '{':
                tokens.append(Token(TT.LBRACE, self.c, self.fp.copy()))
                self.next()
            elif self.c == '}':
                tokens.append(Token(TT.RBRACE, self.c, self.fp.copy()))
                self.next()
            elif self.c == '+':
                tokens.append(Token(TT.ADD_OP, self.c, self.fp.copy()))
                self.next()
            elif self.c == '-':
                tokens.append(Token(TT.SUB_OP, self.c, self.fp.copy()))
                self.next()
            elif self.c == '*':
                tokens.append(self.make_mul_or_exp_op())
            elif self.c == '/':
                tokens.append(Token(TT.DIV_OP, self.c, self.fp.copy()))
                self.next()
            elif self.c == '%':
                tokens.append(Token(TT.MOD_OP, self.c, self.fp.copy()))
                self.next()
            elif self.c == '=':
                tokens.append(Token(TT.ASGN_OP, self.c, self.fp.copy()))
                self.next()
            elif self.c == ':':
                tokens.append(Token(TT.COLON, self.c, self.fp.copy()))
                self.next()
            elif self.c == ',':
                tokens.append(Token(TT.COMMA, self.c, self.fp.copy()))
                self.next()
            else:
                raise fail(f"unexpected character '{self.c}'", self.fp.copy())
        tokens.append(Token(TT.EOF, '\0', self.fp.copy()))
        return tokens

    def make_name(self) -> Token:
        value = self.c
        self.next()
        while not self.done and self.c in 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_':
            value += self.c
            self.next()
        if value in KEYWORDS:
            return Token(TT.KEYWORD, value, self.fp.copy())
        else:
            return Token(TT.IDENTIFIER, value, self.fp.copy())
    
    def make_int(self) -> Token:
        value = self.c
        self.next()
        while not self.done and self.c in '1234567890':
            value += self.c
            self.next()
        return Token(TT.INT, value, self.fp.copy())

    def make_mul_or_exp_op(self) -> Token:
        value = self.c
        self.next()
        if self.c == '*':
            value += self.c
            self.next()
            return Token(TT.EXP_OP, value, self.fp.copy())
        else:
            return Token(TT.MUL_OP, value, self.fp.copy())

    def next(self):
        self.pos += 1
        if self.pos < len(self.text):
            self.c = self.text[self.pos]
        else:
            self.done = True
            self.c = '\0'
        self.fp.next(self.c == '\n')


class Node:
    def __init__(self, fp: Position) -> None:
        self.fp = fp
    
    def __repr__(self) -> str: return f'{type(self).__name__}'

class Statement(Node):
    def __init__(self, fp: Position) -> None:
        super().__init__(fp)

class Expression(Statement):
    def __init__(self, fp: Position) -> None:
        super().__init__(fp)

class Type(Node):
    def __init__(self, token: Token) -> None:
        super().__init__(token.fp)
        self.token = token

    def __repr__(self) -> str: return super().__repr__() + f'({self.token.value})'

class TypedDecl(Node):
    def __init__(self, target: Token, type: Type) -> None:
        super().__init__(target.fp)
        self.target = target
        self.type = type

    def __repr__(self) -> str: return super().__repr__() + f'({self.target.value}, {self.type})'

class FuncDef(Statement):
    def __init__(self, target: Token, type: Type, params: List[TypedDecl], body: List[Statement], fp: Position) -> None:
        super().__init__(fp)
        self.target = target
        self.params = params
        self.type = type
        self.body = body
    
    def __repr__(self) -> str:
        paramstr = ','.join(map(lambda x:x.__repr__(), self.params))
        bodystr = ','.join(map(lambda x:x.__repr__(), self.body))
        return super().__repr__() + f'({self.target.value}, {self.type}, [{paramstr}], [{bodystr}])'

class VarDecl(Statement):
    def __init__(self, target: Token, type: Type, fp: Position) -> None:
        super().__init__(fp)
        self.target = target
        self.type = type
    
    def __repr__(self) -> str: return super().__repr__() + f'({self.target.value}, {self.type})'

class Assign(Expression):
    def __init__(self, target: Token, value: Expression, fp: Position) -> None:
        super().__init__(fp)
        self.target = target
        self.value = value

    def __repr__(self) -> str: return super().__repr__() + f'({self.target.value}, {self.value})'

class BinaryOperation(Expression):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(fp)
        self.left = left
        self.right = right

    def __repr__(self) -> str: return super().__repr__() + f'({self.left}, {self.right})'

class Add(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class Sub(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class Mul(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class Div(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class Mod(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class Exp(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class FuncCall(Expression):
    def __init__(self, target: Expression, args: List[Expression]) -> None:
        super().__init__(target.fp)
        self.target = target
        self.args = args
    
    def __repr__(self) -> str:
        argstr = ','.join(map(lambda x:x.__repr__(), self.args))
        return super().__repr__() + f'({self.target}, [{argstr}])'

class Int(Expression):
    def __init__(self, token: Token) -> None:
        super().__init__(token.fp)
        self.token = token
    
    def __repr__(self) -> str: return f'{super().__repr__()}({self.token.value})'

class Var(Expression):
    def __init__(self, token: Token) -> None:
        super().__init__(token.fp)
        self.token = token
    
    def __repr__(self) -> str: return f'{super().__repr__()}({self.token.value})'

class Parser:
    def parse(self, tokens: List[Token]) -> List[Statement]:
        if len(tokens) <= 0:
            raise fail('no tokens to be parsed', None)
        self.tokens = tokens
        self.pos = 0
        self.t = tokens[0]
        self.done = False

        ast = self.make_program()
        return ast

    def make_program(self) -> List[Statement]:
        return self.make_statements()

    def make_statements(self) -> List[Statement]:
        statements: List[Statement] = []
        while not self.done and self.t.type not in [TT.RBRACE, TT.EOF]:
            statements.append(self.make_statement())
        return statements

    def make_statement(self) -> Statement:
        if self.t.value == 'func':
            return self.make_func_def()
        elif self.t.value == 'let':
            return self.make_declaration_statement()
        else:
            return self.make_expression()

    def make_func_def(self) -> FuncDef:
        fp = self.t.fp
        self.next()
        if self.t.type != TT.IDENTIFIER:
            fail(f'expected identifier, got {self.t}', self.t.fp)
        target = self.t
        self.next()
        if self.t.type != TT.LPAREN:
            fail(f'expected \'(\', got {self.t}', self.t.fp)
        self.next()
        params: List[TypedDecl] = []
        while not self.done and self.t.type != TT.RPAREN:
            params.append(self.make_typed_declaration())
            if self.t.type == TT.RPAREN:
                break
            elif self.t.type == TT.COMMA:
                self.next()
            else:
              fail(f'expected \',\', got {self.t}', self.t.fp)
        if self.t.type != TT.RPAREN:
            fail(f'expected \')\', got {self.t}', self.t.fp)
        self.next()
        if self.t.type != TT.COLON:
            fail(f'expected \':\', got {self.t}', self.t.fp)
        self.next()
        if self.t.value not in ['u8', 'u16', 'u32', 'u64', 'i8', 'i16', 'i32', 'i64', 'char', 'usize', 'uptr']:
            fail(f'expected keyword, got {self.t}', self.t.fp)
        type = self.make_type()
        if self.t.type != TT.LBRACE:
            fail("expected '{'" + f', got {self.t}', self.t.fp)
        self.next()
        body = self.make_statements()
        if self.t.type != TT.RBRACE:
            fail("expected '}'" + f', got {self.t}', self.t.fp)
        self.next()
        return FuncDef(target, type, params, body, fp)

    def make_typed_declaration(self) -> TypedDecl:
        target = self.t
        self.next()
        if self.t.type != TT.COLON:
            fail(f'expected \':\', got {self.t}', self.t.fp)
        self.next()
        if self.t.value not in ['u8', 'u16', 'u32', 'u64', 'i8', 'i16', 'i32', 'i64', 'char', 'usize', 'uptr']:
            fail(f'expected keyword, got {self.t}', self.t.fp)
        type = self.make_type()
        return TypedDecl(target, type)

    def make_declaration_statement(self) -> VarDecl:
        fp = self.t.fp
        self.next()
        if self.t.type != TT.IDENTIFIER:
            fail(f'expected identifier, got {self.t}', self.t.fp)
        target = self.t
        self.next()
        if self.t.type != TT.COLON:
            fail(f"expected ':', got {self.t}", self.t.fp)
        self.next()
        type = self.make_type()
        return VarDecl(target, type, fp)

    def make_type(self) -> Type:
        if self.t.value not in ['u8', 'u16', 'u32', 'u64', 'i8', 'i16', 'i32', 'i64', 'char', 'usize', 'uptr']:
            fail(f'expected keyword, got {self.t}', self.t.fp)
        token = self.t
        self.next()
        return Type(token)

    def make_expression(self) -> Expression:
        return self.make_assignment()

    def make_assignment(self) -> Expression:
        target = self.t
        self.next()
        if target.type == TT.IDENTIFIER and self.t.type == TT.ASGN_OP:
            self.next()
            value = self.make_expression()
            return Assign(target, value, target.fp)
        else:
            self.previus()
            return self.make_addition()


    def make_addition(self) -> Expression:
        left = self.make_subtraction()
        if self.t.type == TT.ADD_OP:
            self.next()
            right = self.make_addition()
            return Add(left, right, left.fp)
        else:
            return left

    def make_subtraction(self) -> Expression:
        left = self.make_multiplication()
        if self.t.type == TT.SUB_OP:
            self.next()
            right = self.make_subtraction()
            return Sub(left, right, left.fp)
        else:
            return left

    def make_multiplication(self) -> Expression:
        left = self.make_division()
        if self.t.type == TT.MUL_OP:
            self.next()
            right = self.make_multiplication()
            return Mul(left, right, left.fp)
        else:
            return left

    def make_division(self) -> Expression:
        left = self.make_modulus()
        if self.t.type == TT.DIV_OP:
            self.next()
            right = self.make_division()
            return Div(left, right, left.fp)
        else:
            return left

    def make_modulus(self) -> Expression:
        left = self.make_exponentation()
        if self.t.type == TT.MOD_OP:
            self.next()
            right = self.make_modulus()
            return Mod(left, right, left.fp)
        else:
            return left

    def make_exponentation(self) -> Expression:
        left = self.make_func_call()
        if self.t.type == TT.EXP_OP:
            self.next()
            right = self.make_exponentation()
            return Exp(left, right, left.fp)
        else:
            return left

    def make_func_call(self) -> Expression:
        target = self.make_value()
        if self.t.type == TT.LPAREN:
            self.next()
            args: List[Expression] = []
            while not self.done and self.t.type != TT.RPAREN:
                args.append(self.make_expression())
                if self.t.type == TT.RPAREN:
                    break
                elif self.t.type == TT.COMMA:
                    self.next()
                else:
                    fail(f'expected \',\', got {self.t}', self.t.fp)
            if self.t.type != TT.RPAREN:
                fail(f'expected \')\', got {self.t}', self.t.fp)
            self.next()
            return FuncCall(target, args)
        else:
            return target

    def make_value(self) -> Expression:
        token = self.t
        self.next()
        if token.type == TT.INT:
            return Int(token)
        elif token.type == TT.IDENTIFIER:
            return Var(token)
        elif token.type == TT.LPAREN:
            expression = self.make_expression()
            if self.t.type != TT.RPAREN:
                fail('parenthesis not closed', token.fp)
            self.next()
            return expression
        else:
            raise fail(f"unexpected token {token}", token.fp.copy())

    def next(self):
        self.pos += 1
        if self.pos < len(self.tokens):
            self.t = self.tokens[self.pos]
        else:
            self.done = True
            self.t = Token(TT.EOF, '\0', Position(-1, -1, ''))

    def previus(self):
        self.pos -= 1
        self.t = self.tokens[self.pos]

def main():
    if len(sys.argv) <= 1:
        print('python3 parser.py <file>')
    filename = sys.argv[1]
    tokens = Lexer().tokenize(File(filename))
    for i in tokens: print(i)
    ast = Parser().parse(tokens)
    for i in ast: print(i)

if __name__ == '__main__':
    main()
