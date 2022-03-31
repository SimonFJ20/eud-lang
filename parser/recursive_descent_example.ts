
/*
Grammar

expression      :   addition

addition        :   addition "+" subtraction
                |   subtraction

subtraction     :   subtraction "-" multiplication
                |   multiplication

multiplication  :   multiplication "*" division
                |   division

division        :   division "/" exponentation
                |   exponentation

exponentation   :   value "**" exponentation
                |   value

value           :   int
                |   "(" expression ")"

*/

enum TokenType {
    INT,
    ADD_OP,
    SUB_OP,
    MUL_OP,
    DIV_OP,
    EXP_OP,
    LPAREN,
    RPAREN,
    EOF,
}

type Token = {
    type: TokenType,
    value: string,
}

interface BaseNode {}

interface ExpressionNode extends BaseNode {}

class AdditionNode implements ExpressionNode {
    constructor (public left: ExpressionNode, public right: ExpressionNode) {}
}

class SubtractionNode implements ExpressionNode {
    constructor (public left: ExpressionNode, public right: ExpressionNode) {}
}

class MultiplicationNode implements ExpressionNode {
    constructor (public left: ExpressionNode, public right: ExpressionNode) {}
}

class DivisionNode implements ExpressionNode {
    constructor (public left: ExpressionNode, public right: ExpressionNode) {}
}

class ExponentationNode implements ExpressionNode {
    constructor (public left: ExpressionNode, public right: ExpressionNode) {}
}

class IntLiteral implements ExpressionNode {
    constructor (public token: Token) {}
}

class Parser {
    private tokens: Token[];
    private pos: number;
    private tok: Token;
    private done: boolean;
    

    public parse(tokens: Token[]): BaseNode {
        this.tokens = tokens;
        this.pos = 0;
        this.tok = tokens[0];
        this.done = false;

        return this.makeExpression();
    }

    private makeExpression(): ExpressionNode {
        return this.makeAddition();
    }

    private makeAddition(): ExpressionNode {
        const left = this.makeAddition();
        if (this.tok.type === TokenType.ADD_OP) {
            this.next();
            const right = this.makeSubtraction();
            return new AdditionNode(left, right);
        } else {
            return this.makeSubtraction();
        }
    }

    private makeSubtraction(): ExpressionNode {
        const left = this.makeSubtraction();
        if (this.tok.type === TokenType.SUB_OP) {
            this.next();
            const right = this.makeMultiplication();
            return new SubtractionNode(left, right);
        } else {
            return this.makeMultiplication();
        }
    }

    private makeMultiplication(): ExpressionNode {
        const left = this.makeMultiplication();
        if (this.tok.type === TokenType.MUL_OP) {
            this.next();
            const right = this.makeDivision();
            return new MultiplicationNode(left, right);
        } else {
            return this.makeDivision();
        }
    }

    private makeDivision(): ExpressionNode {
        const left = this.makeDivision();
        if (this.tok.type === TokenType.DIV_OP) {
            this.next();
            const right = this.makeExponentation();
            return new DivisionNode(left, right);
        } else {
            return this.makeExponentation();
        }
    }

    private makeExponentation(): ExpressionNode {
        const left = this.makeValue();
        if (this.tok.type === TokenType.EXP_OP) {
            this.next();
            const right = this.makeExponentation();
            return new ExponentationNode(left, right);
        } else {
            return left;
        }
    }

    private makeValue(): ExpressionNode {
        const token = this.tok;
        this.next()
        if (token.type === TokenType.INT) {
            return new IntLiteral(this.tok);
        } else if (token.type === TokenType.LPAREN) {
            const expression = this.makeExpression();
            if (this.tok.type !== TokenType.RPAREN)
                throw new Error('fuck you')
            return expression;
        } else {
            throw new Error('fuck you')
        }
    }

    private next() {
        this.pos++;
        if (this.pos >= this.tokens.length) {
            this.done = true;
            this.tok = {type: TokenType.EOF, value: ''};
        } else {
            this.tok = this.tokens[this.pos]
        }
    }

}
