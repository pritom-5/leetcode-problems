/** 
* @param {number} a 
* @param {number} b 
* @returns {number}
 * */
function add (a, b) {
	return a + b;
}

/** 
 * @param {One} info
 * @returns {string}
 * */
function createNewInfo (info) {
	return info.firstName + info.secondName + " -- " + info.age.toString();

}

/** @param {string} s */
function validPallindrome (s) {
	const reversed = s.split("").map(item => {
		const c = item.toLowerCase().charCodeAt(0);
		if (c >= 98 && c <= 122) {
			return item
		} 
			return ""
		
	}).toReversed().join("");

	return reversed === s;
}



(function tryLoops () {
	for (let i = 0; i < 9; i+= 3) {
		for (let j = 0; j < 9; j+= 3) {
			console.log(j, i);
		}
		console.log("----------------");
	}
})()





























