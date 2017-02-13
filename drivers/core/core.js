// Package core.js provides javascript functions which provide similar functionalities
// to allow patching provided virtual DOM and query events and dom nodes as needed.

var unwanted = {"constructor": true,"toString": true}

// PatchDOM patches the provided elements into the target from the current DOM.
function PatchDOM(fragmentDOM, liveDOM, replace){

}

// GetEvent returns the event as a object which can be jsonified and
// sent over the pipeline.
function GetEvent(ev){
  var eventObj = DeepClone(ev)

  var c = ev.constructor
  switch (c) {
    case MutationRecord:
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
  if(Object.getOwnPropertyNames){
    return Object.getOwnPropertyNames(item)
  }

  // If we can use the Object.keys function in ES5 then use this has
  // we can manage with the provided set.
  if(Object.keys){
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

// DeepClone clones all internal properties of the provided object, re-creating
// internal key-value pairs accessible to the object even in prototype inheritance.
// Functions are not runned except for custom types which are checked accordingly.
function DeepClone(item){
  c = item.constructor

  switch(c){
    case String:
      return item

    case Number:
      return item

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

    case Array:
      var newArray = []
      for(var index in item){
        indexElem = item[index]
        newArray[index] = DeepClone(indexElem)
      }

      return newArray

    default:
      var newObj = {}
      var rootProtos = reverse(filter(GetRoots(item), function(val){
        return Type(val) != "Object"
      }))

      // Run through all parent constructs and pull keys, we want
      // to have all inherited properties as well.
      var keys = mapFlatten(rootProtos, function(root){
        return filter(Keys(root), function(val){
          var allowed = !unwanted[val]
          var isNotConstant = !isUpperCase(val)
          return allowed && isNotConstant
        })
      })

      for(var index in keys){
        var key = keys[index]
        newObj[capitalize(key)] = item[key]
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
		var ev = o[i]
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
			item = items[i]
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
