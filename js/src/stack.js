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

	/** 
	 * @param {number[]} temps 
	 * @returns {number[]}
	 * */
	nextTemperatures (temps) {
		const res = new Array(temps.length).fill(0);

		/** @type {{temp: number, index: number}[]}*/
		const stack = [];

		for (let i = 0; i < temps.length; i++) {
			const curr_temp = temps[i];

			while (stack.length > 0 && 
						curr_temp > stack[stack.length - 1].temp) {

				const temp_temperature =  stack.pop();

				if (temp_temperature) {
					const {index} = temp_temperature;

					res[index] = i - index;
				}
			}

			stack.push({temp: curr_temp, index: i});
		}

		// loop through every temps
		// while loop where stack not empty and temp > last stack item
		// update res array
		// add temp to the stack

		 return res;
	}
}

const stack = new Stack();
export { stack };















