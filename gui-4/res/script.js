function getValueOf(id) {
	let element = document.getElementById(id);
	if (!element) return null;
	return element.value;
}
