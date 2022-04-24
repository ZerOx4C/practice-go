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

function updateOutput(value) {
	document.getElementById("output").innerText = JSON.stringify(value);
}
