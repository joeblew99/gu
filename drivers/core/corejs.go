// Package core contains javascript sources which are exported into specific
// drivers webview which exposes similar functionality needed to interoperate with
// the platform.

// Document is auto-generate and should not be modified by hand.

//go:generate go run generate.go

package core

// JavascriptDriverCore contains the javascript code to be injected into a webview
// to provide similar functionality desired to have it working with Gu.
var JavascriptDriverCore = `// Package core.js provides javascript functions which provide similar functionalities
// to allow patching provided virtual DOM and query events and dom nodes as needed.

// unwanted defines specific object functions we dont want passed in during
// collection of object property names.
var unwanted = {"constructor": true,"toString": true}

// PatchDOM patches the provided elements into the target from the current DOM.
// It crawls a live version of the DOM, removing, replacing and adding node
// changes as needed, until the dom resembles it's shadow/fragmentDOM.
function PatchDOM(fragmentDOM, liveDOM, replace){
	if(!live.hasChildNodes()){
		live.appendChild(fragmentDOM)
		return
	}

	var shadowNodes = fragmentDOM.childNodes
	var liveNodes = liveDOM.childNodes

	for(var index = 0; index < shadowNodes.length; index++){
		var node = shadowNodes[index]

		if(node.constructor === Text){
			if(isEmptyTextNode(node)){
				live.appendChild(node)
				continue
			}


			if(index < liveNodes.length){
				liveNode = liveNodes[index]
				liveDOM.insertBefore(liveNode, node)
				continue
			}

			live.appendChild(node)
			continue
		}

		var nodeTagName = node.tagName
		var nodeId = node.getAttribute("id")
		// var nodeClass = node.getAttribute("class")
		var nodeUID = node.getAttribute("uid")
		var nodeAttr = node.attributes
		var nodeKids = node.childNodes
		var nodeSel = nodeTagName+"[uid="+nodeUID+"]"
		var nodeRemoved = node.hasAttribute("NodeRemoved")
		var nodeHash = node.getAttribute("hash")

		if(!nodeId && nodeUID && nodeHash){
			addIfNoEqual(live, node)
			continue
		}

		if(!nodeUID && !nodeHash){
			if(nodeId){
				var found = live.querySelectorAll("#"+nodeId)
				if(!found.length){
					live.appendChild(node)
					continue
				}

				live.replaceNode(found, node)
				continue
			}
		}

		var allTargets = live.querySelectorAll(nodeSel)
		if(!allTargets.length){
			live.appendChild(node)
			continue
		}

		for(var jindex = 0; jindex < allTargets.length; jindex++){
			var curTarget = allTargets[jindex]

			if(nodeRemoved){
				live.remove(curTarget)
				continue
			}

			if(replace){
				live.replaceNode(curTarget, node)
				continue
			}

			var liveHash = curTarget.getAttribute("hash")

			if(liveHash === nodeHash){
				continue
			}


			if(!curTarget.childNodes.length){
				live.replaceNode(curTarget, node)
				continue
			}

			removeAllTextNodes(curTarget)

			for(var key in nodeAttr){
				var attr = nodeAttr[key]
				curTarget.setAttribute(attr.nodeName, attr.nodeValue)
			}

			curTargetChilds = curTarget.childNodes
			if(!curTargetChilds.length){
				curTarget.innerHTML = ""
				curTarget.appendChild.apply(curTarget, nodeKids)
				continue
			}

			PatchDOM(node, curTarget, replace)
		}
	}
}

// addIfNoEqual adds a giving node into the target if its not found to match any
// child nodes of the target and if one is found then that is replaced with the
// provided new node.
function addIfNoEqual(target, node){
	for(var i = 0; i < list.length; i++){
		var against = target[i]
		if(against.IsEqualNode(node)){
			target.replaceNode(against, node)
			return
		}
	}

	target.appendChild(node)
}

// isEmptyTextNode returns true/false if the node is an empty text node.
function isEmptyTextNode(node){
	if(node.nodeType !== 3){
		return false
	}

	return node.textContent === ""
}

// removeAllTextNodes removes all residing textnodes in the provided node.
function removeAllTextNodes(parent){
	var list = parent.childNodes

	for(var i = 0; i < list.length; i++){
		var node = list[i]
		if(node.nodeType === 3){
			parent.removeChild(node)
		}
	}
}

// createDOMFragment creates a DocumentFragment from the provided HTML.
function createDOMFragment(elemString){
	var div  = document.createElement("div")
	div.innerHTML = elemString

	var fragment = document.createDocumentFragment()
	fragment.appendChild.Apply(fragment, div.childNodes)

	return div
}

// GetEvent returns the event as a object which can be jsonified and
// sent over the pipeline.
function GetEvent(ev){
	var eventObj

	var c = ev.constructor
	switch (c) {
		case MutationRecord:
			eventObj = DeepClone(ev,{
				Functions: false,
			})

			var added = map(eventObj.addedNodes, function(elem){
				return StringifyHTML(elem)
			})

			var removed = map(eventObj.removedNodes, function(elem){
				return StringifyHTML(elem)
			})

			var presib = StringifyHTML(eventObj.preSibling)
			var nextsib = StringifyHTML(eventObj.nextSibling)

			eventObj.AddedNodes = added
			eventObj.RemovedNodes = removed
			eventObj.PreSibling = presib
			eventObj.NextSibling = nextsib

		case MediaStream:
			eventObj = toMediaStream(ev)

		default:
			eventObj = DeepClone(ev,{
				Functions: false,
			})
	}

	return eventObj
}

// Type returns the type of the native constructor of the passed in object.
function Type(item){
	if(item !== undefined && item != null) {
		return (item.constructor.toString().match(/function (.*)\(/)[1])
	}
}

// filters out the giving items not matching the provided function.
function filter(item, fn){
	var filtered = []

	for(key in item){
		if(fn(item[key], key, item)){
			filtered.push(item[key])
		}
	}

	return filtered
}

// map maps new values throug the provided function skipping null returns.
function map(item, fn){
	var mapped = []

	for(key in item){
		var res = fn(item[key], key, item)
		if(res){
			mapped.push(item[key])
		}
	}

	return mapped
}

// reverse returns the list reversed in order.
function reverse(list){
	var reversed = []

	for(var i = list.length - 1; i > 0; i--){
		reversed.push(list[i])
	}

	return reversed
}

// mapFlatten maps new values throug the provided function skipping null returns.
function mapFlatten(item, fn){
	var mapped = []

	for(key in item){
		var res = fn(item[key], key, item)
		if(!res){
			continue
		}

		switch(Type(res)){
			case "Array":
				 Array.prototype.push.apply(mapped, res)
				 break
			default:
				mapped.push(item[key])
		}
	}

	return mapped
}

// capitalize returns a capitalized string.
function capitalize(val){
	if(val !== ""){
		var newVal = [val[0].toUpperCase()]
		newVal.push(val.substring(1))
		return newVal.join('')
	}

	return val
}

// isUpperCase returns true/false if the string is in uppercase.
function isUpperCase(val){
	return val.toUpperCase() === val
}

// ConstructorKeys returns the constructor keys for the giving object.
function Keys(item){
	// If we can use the getOwnPropertyNames function in ES5 then use this has
	// inherited properties are desired as well.
	if("getOwnPropertyNames" in Object){
		return Object.getOwnPropertyNames(item)
	}

	// If we can use the Object.keys function in ES5 then use this has
	// we can manage with the provided set.
	if("keys" in Object){
		return Object.keys(item)
	}

	var keys = []

	for(var key in item){
		keys.push(key)
	}

	// Check if keys are empty and if not, probably declared object
	// returned.
	if(keys.length){
		return keys
	}

	// Attempt using the __proto__ object if we can copy. We are probably back in
	// Old JS land.
	if(item.__proto__){
		for(var key in item.__proto__){
			keys.push(key)
		}

		return keys
	}

	// Attempt using the protototype object if we can copy. We are probably back in
	// Old JS land.
	if(item.prototype){
		for(var key in item.prototype){
			keys.push(key)
		}

		return keys
	}

	// Digress to access prototype from constructor and
	// using the protototype object if we can copy. We are probably back in
	// Old JS land.
	if(item.constructor.prototype){
		for(var key in item.constructor.prototype){
			keys.push(key)
		}

		return keys
	}

	return keys
}

// ConstructorDeepClone returns the clone of the items prototype.
function ConstructorDeepClone(item){
	if(item.__proto__){
		return DeepClone(item.prototype)
	}

	if(item.prototype){
		return DeepClone(item.prototype)
	}

	if(item.constructor.prototype){
		return DeepClone(item.prototype)
	}
}

// exceptObjects are objects which we dont want uncloned but kept intact.
// Also, these elements will lead to massive cyclic issues, keep them intact and deal
// in another approach.
var exceptObjects = {HTMLElement: true, NodeList: true, HTMLDocument: true, Node: true, Document: true}

// defaultOptions defines a set of optional values allowed when cloning objects.
var defaultOptions = {Functions: true, LastTree: [] }

// DeepClone clones all internal properties of the provided object, re-creating
// internal key-value pairs accessible to the object even in prototype inheritance.
// Functions are not runned except for custom types which are checked accordingly.
function DeepClone(item, options){
	if(item === undefined || item == null){
		return item
	}

	c = item.constructor

	if(!options){
		options = defaultOptions
	}

	switch(c){
		case Function:
			if(options.AllowFunctions){
				return item
			}

			return

		case Number:
			return item

		case HTMLElement:
			return item

		case Node:
			return item

		case Element:
			return item

		case Boolean:
			return item

		case String:
			return item

		case Blob:
			return fromBlob(item)

		case File:
			return fromFile(item)

		case Uint8Array:
			var newArray = new Uint8Array
			for(var index in item){
				newArray.push(item[index])
			}

			return newArray

		case Float64Array:
			var newArray = new Float32Array
			for(var index in item){
				newArray.push(item[index])
			}

			return newArray

		case Float32Array:
			var newArray = new Float32Array
			for(var index in item){
				newArray.push(item[index])
			}

			return newArray

		case TouchList:
			return toTouches(item)

		case MediaStream:
			return toMediaStream(item)

		case Gamepad:
			return toGamepad(item)

		case DataTransfer:
			return toDataTransfer(item)

		case Array:
			var newArray = []
			for(var index in item){
				indexElem = item[index]
				newArray[index] = DeepClone(indexElem, options)
			}

			return newArray

		default:
			var newObj = {}
			var roots = GetRoots(item)

			// If the element is a child of this givng root in the exceptObjects
			// then return it has is, because we need it intact and unchanged.
			for(var root in roots){
				var base = roots[root]
				if(exceptObjects[Type(base)]){
					return item
				}
			}

			// If we are passed the previous tree, then check if we have
			// someone in that root as well.
			// if(options.LastTree){
			//   for(var root in options.LastTree){
			//     var base = roots[root]
			//     if(exceptObjects[Type(base)]){
			//       return item
			//     }
			//   }
			// }

			var rootProtos = reverse(filter(roots, function(val){
				return Type(val) != "Object"
			}))

			// Are we dealing with a empty object without parent, then we are probably
			// dealing with a declared map/hash that points directly to the Object constructor.
			if(rootProtos.length === 0){
				rootProtos.push(item)
			}

			// Run through all parent constructs and pull keys, we want
			// to have all inherited properties as well.
			var keys = mapFlatten(rootProtos, function(root){
				return filter(Keys(root), function(val){

					// If functions are not allowed and we have on here, then skip.
					if(!options.AllowFunctions && Type(item[val]) === "Function"){
						return false
					}

					var allowed = !unwanted[val]
					var isNotConstant = !isUpperCase(val)

					return allowed && isNotConstant
				})
			})

			for(var index in keys){
				var key = keys[index]
				newObj[capitalize(key)] = DeepClone(item[key], {
					Functions: options.Functions,
					// LastTree: roots,
				})
			}

			return newObj
	}
}

// StringifyHTML returns the html of the element and it's content.
function StringifyHTML(elem, deep) {
	var div = document.createElement("div")
	div.appendChild(elem.cloneNode(deep))
	return div.innerHTML
}

// GetRoots retrieves all root properties for which the element runs down.
function GetRoots(o) {
	var roots = []
	var found = {}

	var proto = o.constructor.prototype

	while(true){
		if(proto == undefined || proto == null) {
			break
		}

		if(found[proto]){
			break
		}

		roots.push(proto)
		found[proto] = true

		if("__proto__" in proto){
			proto = proto.__proto__
		}
	}

	return roots
}

// fromBlob transform the providded Object blob into a byte slice.
function fromBlob(o) {
	if(o == null || o == undefined){
		return
	}

	var data = null

	fileReader = new FileReader()
	fileReader.onloadend = function(){
		data = new Uint8Array(fileReader.result)
	}

	fileReader.readAsArrayBuffer(o)

	return data
}

// fromFile transform the providded Object blob into a byte slice.
function fromFile(o) {
	if(o == null || o == undefined){
		return
	}

	var data = null

	fileReader = new FileReader()
	fileReader.onloadend = function(){
		data = new Uint8Array(fileReader.result)
	}

	fileReader.readAsArrayBuffer(o)

	return data
}

// toInputSourceCapability returns the events.InputDeviceCapabilities from the object.
function toInputSourceCapability(o) {
	if(o == null || o == undefined){
		return
	}

	return {
		FiresTouchEvent: o.firesTouchEvent,
	}
}

// toMotionData returns a motionData object from the Object.
function toMotionData(o)  {
	var md = {X:0.0, Y: 0.0, Z: 0.0}

	if(o == null || o == undefined){
		return md
	}

	md.X = o.x
	md.Y = o.y
	md.Z = o.z
	return md
}

// toRotationData returns a RotationData object from the Object.
function toRotationData(o)  {
	if(o == null || o == undefined){
		return
	}

	md.Alpha = o.alpha
	md.Beta = o.beta
	md.Gamma = o.gamma
	return md
}

// toMediaStream returns a events.MediaStream object.
function toMediaStream(o) {
	if(o == null || o == undefined){
		return
	}

	stream.Active = o.active
	stream.Ended = o.ended
	stream.ID = o.id
	stream.Audios = []
	stream.Videos = []

	var audioTracks = o.getAudioTracks()
	if(audioTracks != null && audioTracks != undefined){
		for(i = 0; i < audioTracks.length; i++ ){
			var track = audioTracks[i]
			var settings = track.getSettings()

			stream.Audios.push({
				Enabled:    track.enabled,
				ID:         track.id,
				Kind:       track.kind,
				Label:      track.label,
				Muted:      track.muted,
				ReadyState: track.readyState,
				Remote:     track.remote,
				AudioSettings: {
					ChannelCount:     settings.channelCount.Int(),
					EchoCancellation: settings.echoCancellation,
					Latency:          settings.latency,
					SampleRate:       settings.sampleRate.Int64(),
					SampleSize:       settings.sampleSize.Int64(),
					Volume:           settings.volume,
					MediaTrackSettings: {
						DeviceID: settings.deviceId,
						GroupID:  settings.groupId,
					},
				},
			})
		}
	}

	var videosTracks = o.getVideoTracks()
	if(videosTracks != null && videosTracks != undefined){
		for( i = 0; i < videosTracks.length; i++) {
			var track = videosTracks[i]
			var settings = track.getSettings()

			stream.Videos.push({
				Enabled:    track.enabled,
				ID:         track.id,
				Kind:       track.kind,
				Label:      track.label,
				Muted:      track.muted,
				ReadyState: track.readyState,
				Remote:     track.remote,
				VideoSettings: {
					AspectRatio: settings.aspectRation,
					FrameRate:   settings.frameRate,
					Height:      settings.height.Int64(),
					Width:       settings.width.Int64(),
					FacingMode:  settings.facingMode,
					MediaTrackSettings: {
						DeviceID: settings.deviceId,
						GroupID:  settings.groupId,
					},
				},
			})
		}
	}

	return stream
}

function toTouches(o) {
	if(o == null || o == undefined){
		return
	}

	var touches = []

	for(i = 0; i < o.length; i++){
		var ev = o.item(i)
		touches.push({
			ClientX:    ev.clientX,
			ClientY:    ev.clientY,
			OffsetX:    ev.offsetX,
			OffsetY:    ev.offsetY,
			PageX:      ev.pageX,
			PageY:      ev.pageY,
			ScreenX:    ev.screenX,
			ScreenY:    ev.screenY,
			Identifier: ev.identifier,
		})

	}

	return touches
}

// toGamepad returns a Gamepad struct from the js object.
function toGamepad(o) {
	var pad = {}

	if(o == null || o == undefined){
		return pad
	}

	pad.DisplayID = o.displayId
	pad.ID = o.id
	pad.Index = o.index.Int()
	pad.Mapping = o.mapping
	pad.Connected = o.connected
	pad.Timestamp = o.timestamp
	pad.Axes = []
	pad.Buttons = []

	var axes = o.axes
	if(axes != null && axes != undefined) {
		for(i = 0; i < axes.length; i++ ){
			pad.Axes.push(axes[i])
		}
	}

	var buttons = o.buttons
	if(buttons != null && buttons != undefined) {
		for(i = 0; i < buttons.length; i++){
			button = buttons[i]
			pad.Buttons.push({
				Value:   button.value,
				Pressed: button.pressed,
			})
		}
	}

	return pad
}

// toDataTransfer returns a transfer object from the Object.
function toDataTransfer(o) {
	if(o == null || o == undefined){
		return
	}

	var dt = {}
	dt.DropEffect = o.dropEffect
	dt.EffectAllowed = o.effectAllowed
	df.Types = o.types
	df.Items = []

	var items = o.items
	if(items != null && items != undefined){
		for(i = 0; i < items.length; i++ ){
			item = items.DataTransferItem(i)
			dItems.push({
				Name: item.name,
				Size: item.size.Int(),
				Data: fromFile(item),
			})
		}
	}

	var dFiles = []

	files = o.files
	if(files != null && files != undefined){
		for(i = 0; i < files.length; i++){
			item = files[i]
			dFiles.push({
				Name: item.name,
				Size: item.size.Int(),
				Data: fromFile(item),
			})
		}
	}

	dt.Items = {Items: dItems}
	dt.Files = dFiles
	return dt
}
`
