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

const t = "a man is, lying on the floor";
validPallindrome(t);
