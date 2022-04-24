function onFetchStringClicked(sender) {
	fetchString().then(value => updateOutput(value));
}

function onFetchIntClicked(sender) {
	fetchInt().then(value => updateOutput(value));
}

function onFetchArrayClicked(sender) {
	fetchArray().then(value => updateOutput(value));
}

function onFetchTableClicked(sender) {
	fetchTable().then(value => updateOutput(value));
}

function onFetchComplexClicked(sender) {
	fetchComplex().then(value => updateOutput(value));
}

function onGetComplexClicked(sender) {
	invokeGetComplex()
}

function updateOutput(value) {
	document.getElementById("output").innerText = JSON.stringify(value);
}

function getComplex() {
	return {
		foo: 1228,
		bar: "hello",
		baz: [1, 2, {
			hoge: 123,
			piyo: 456,
		}],
	}
}
