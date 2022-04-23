var rowId = 0;

function init(sender) {
	appendRow(1, "もぐもぐ", 15)
	appendRow(2, "ぱくぱく", 30)
	appendRow(3, "1,2", 45)
}

function getList() {
	return Array.from(document.querySelectorAll("#list tbody tr")).map(rowElement => {
		return {
			id: parseInt(rowElement.querySelector(".id").innerText),
			name: rowElement.querySelector(".name").innerText,
			age: parseInt(rowElement.querySelector(".age").innerText),
		};
	});
}

function appendRow(id, name, age) {
	let rowElement = document.querySelector("#row-origin").content.cloneNode(true);
	rowElement.querySelector("tr").dataset.id = rowId;
	rowElement.querySelector(".id").innerText = id;
	rowElement.querySelector(".name").innerText = name;
	rowElement.querySelector(".age").innerText = age;
	rowElement.querySelector(".remove").dataset.id = rowId;
	document.querySelector("#list tbody").appendChild(rowElement);
	rowId += 1;
}

function removeRow(rowId) {
	document.querySelector('#list tbody tr[data-id="' + rowId + '"]').remove();
}

function onAddClicked(sender) {
	let editorRootElement = document.querySelector("#list tfoot tr");
	let idEditorElement = editorRootElement.querySelector(".id");
	let nameEditorElement = editorRootElement.querySelector(".name");
	let ageEditorElement = editorRootElement.querySelector(".age");

	let id = idEditorElement.value;
	let name = nameEditorElement.value;
	let age = ageEditorElement.value;

	idEditorElement.value = "";
	nameEditorElement.value = "";
	ageEditorElement.value = "";

	appendRow(id, name, age);
}

function onRemoveClicked(sender) {
	removeRow(parseInt(sender.dataset.id));
}

function onSaveClicked(sender) {
	save()
}
