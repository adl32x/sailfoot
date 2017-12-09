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
