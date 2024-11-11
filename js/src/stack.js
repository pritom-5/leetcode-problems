class Stack {
  /**
   * @param {string[]} tokens
   * @return {number}
   * */
  polishNotation(tokens) {
    /** @type {number[]}*/
    const stack = [];
    const ops = {
      "+": true,
      "-": true,
      "*": true,
      "/": true,
    };

    for (let i = 0; i < tokens.length; i++) {
      const tmp = tokens[i];

      if (tmp in ops) {
        const num1 = stack.pop(),
          num2 = stack.pop();

        if (num1 === undefined || num2 === undefined) {
          continue;
        }

        console.log("stack: ", stack);

        switch (tmp) {
          case "+":
            stack.push(num1 + num2);
            break;
          case "-":
            stack.push(num1 - num1);
            break;
          case "*":
            stack.push(num1 * num2);
            break;
          case "/":
            let value = Math.trunc(num2 / num1);
            //if (value === -0) {
            //  value = 0;
            //}
            stack.push(value);
        }
      } else {
        stack.push(Number(tmp));
      }
    }

    return stack[0];
  }
}

const stack = new Stack();
export { stack };
