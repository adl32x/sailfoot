package driver

const JsClickWithText = `
	var rootElement = document.body;

	var filter = {
        acceptNode: function(node){
            if(node.nodeType === document.TEXT_NODE && node.nodeValue.includes(text)){
                 return NodeFilter.FILTER_ACCEPT;
            }
            return NodeFilter.FILTER_REJECT;
        }
    }
    var nodes = [];
    var walker = document.createTreeWalker(document.body, NodeFilter.SHOW_TEXT, filter, false);
    while(walker.nextNode()){
       nodes.push(walker.currentNode.parentNode);
    }
    if (nodes.length > 0) {
        nodes[0].click();
        return 1;
    }
    return 0;
	`

const JsHasText = `
		var rootElement = document.body;

		var filter = {
	        acceptNode: function(node){
	            if(node.nodeType === document.TEXT_NODE && node.nodeValue.includes(text)){
	                 return NodeFilter.FILTER_ACCEPT;
	            }
	            return NodeFilter.FILTER_REJECT;
	        }
	    }
	    var nodes = [];
	    var walker = document.createTreeWalker(document.body, NodeFilter.SHOW_TEXT, filter, false);
	    while(walker.nextNode()){
	       nodes.push(walker.currentNode.parentNode);
	    }
	    if (nodes.length > 0) {
	        return 1;
	    }
	    return 0;
		`

const JsClickClosest = `
		var rootElement = document.querySelector(text);
		if (rootElement == null) { return 0 }

		var rect = rootElement.getBoundingClientRect()

		var elements = document.querySelectorAll(text2)

		if (elements.length == 0) { return 0 }

		var elementToClick = 0
		var distance = Number.MAX_VALUE
		for (var i = 0; i < elements.length; i++) {
			var rect2 = elements[i].getBoundingClientRect();
			var thisDistance = Math.hypot(rect2.x - rect.x, rect2.y - rect.y);
			if (thisDistance < distance) {
				elementToClick = i;
				distance = thisDistance;
			}
		}

		elements[elementToClick].click()
		return 1;
		`

const JsInputEmpty = `
		var elements = document.querySelectorAll('input:not([type=hidden]):not([type=button])')

		if (elements.length == 0) { return 0 }

		for (var i = 0; i < elements.length; i++) {
			var element = elements[i];
			if (element.hidden || element.value != "") { continue; }
			element.setAttribute('data-sailfoot-empty', '');
			return 1
		}

		return 0;
		`
const JsInputEmptyReset = `document.querySelector('input[data-sailfoot-empty]').removeAttribute('data-sailfoot-empty');`

const JsActiveElement = `
		document.activeElement.setAttribute('data-sailfoot-active-element', '');
		return 1;
		`

const JsActiveElementReset = `document.querySelector('[data-sailfoot-empty]').removeAttribute('data-sailfoot-active-element');`
