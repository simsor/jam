// Local storage routines

const SHEET_PREFIX = "sheet_";

function checkLS() {
    if (typeof localStorage == "undefined") {
        console.log("Your browser doesn't support LocalStorage :(");
        return false;
    }
    return true;
}

function getSheetsList() {
    if (!checkLS()) return [];

    let l = localStorage.getItem("sheets");
    if (l == null) {
        return [];
    }

    return JSON.parse(l);
}

function setSheetsList(sheets) {
    if (!checkLS()) return;
    
    localStorage.setItem("sheets", JSON.stringify(sheets));
}

function updateSheet(name, content) {
    if (!checkLS()) return;    

    let sheets = getSheetsList();

    if (sheets.indexOf(name) == -1) {
        sheets.push(name);
        setSheetsList(sheets);
    }

    localStorage.setItem(SHEET_PREFIX + name, content);
}

function deleteSheet(name) {
    if (!checkLS()) return;

    let sheets = getSheetsList();
    let idx = sheets.indexOf(name);

    if (idx == -1) {
        return;
    }

    sheets.splice(idx, 1);
    setSheetsList(sheets);
    localStorage.removeItem(SHEET_PREFIX + name);
}

function getSheet(name) {
    if (!checkLS()) return "";

    let s = localStorage.getItem(SHEET_PREFIX + name);
    return s;
}
