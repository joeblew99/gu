// Package core.js provides javascript functions which provide similar functionalities
// to allow patching provided virtual DOM and query events and dom nodes as needed.

function PatchDOM(){

}

function GetEvent(ev){
  var c = ev.constructor

  switch c{
    case UIEvent:
		return events.NewBaseEvent(&events.AnimationEvent{
			Core:          ev,
			AnimationName: ev.Get("animationName").String(),
			ElapsedTime:   ev.Get("elapsedTime").Float(),
		}, handle)
    case KeyboardEvent:
  }
}

// fromBlob transform the providded Object blob into a byte slice.
function fromBlob(o) {
	if o == null || o == undefined {
		return
	}

  var data = null

	fileReader := new FileReader()
	fileReader.onloadend = function(){
		data = new Uint8Array(fileReader.result)
	}))

	fileReader.readAsArrayBuffer(o)

	return data
}

// fromFile transform the providded Object blob into a byte slice.
function fromFile(o) []byte {
	if o == null || o == undefined {
		return
	}

  var data = null

	fileReader := new FileReader()
	fileReader.onloadend = function(){
		data = new Uint8Array(fileReader.result)
	}))

	fileReader.readAsArrayBuffer(o)

	return data
}

// toInputSourceCapability returns the events.InputDeviceCapabilities from the object.
function toInputSourceCapability(o ) {
	if o == null || o == undefined {
		return
	}

	return {
		FiresTouchEvent: o.firesTouchEvent,
	}
}

// toMotionData returns a motionData object from the Object.
function toMotionData(o)  {
	var md = {X:0.0, Y: 0.0, Z: 0.0}

	if o == null || o == undefined {
		return md
	}

	md.X = o.x
	md.Y = o.y
	md.Z = o.z
	return md
}

// toRotationData returns a RotationData object from the Object.
function toRotationData(o) events.RotationData {
	var md events.RotationData
	if o == null || o == undefined {
		return md
	}
	md.Alpha = o.Get("alpha").Float()
	md.Beta = o.Get("beta").Float()
	md.Gamma = o.Get("gamma").Float()
	return md
}

// toMediaStream returns a events.MediaStream object.
function toMediaStream(o) events.MediaStream {
	var stream events.MediaStream
	if o == null || o == undefined {
		return stream
	}

	stream.Active = o.Get("active").Bool()
	stream.Ended = o.Get("ended").Bool()
	stream.ID = o.Get("id").String()

	audioTracks := o.Get("getAudioTracks")
	if audioTracks != null && audioTracks != undefined {
		for i := 0; i < audioTracks.Length(); i++ {
			track := audioTracks.Index(i)
			settings := track.Call("getSettings", null)

			stream.Audios = append(stream.Audios, events.MediaStreamTrack{
				Core:       track,
				Enabled:    track.Get("enabled").Bool(),
				ID:         track.Get("id").String(),
				Kind:       track.Get("kind").String(),
				Label:      track.Get("label").String(),
				Muted:      track.Get("muted").Bool(),
				ReadyState: track.Get("readyState").Bool(),
				Remote:     track.Get("remote").Bool(),
				AudioSettings: &events.MediaAudioTrackSettings{
					ChannelCount:     settings.Get("channelCount").Int(),
					EchoCancellation: settings.Get("echoCancellation").Bool(),
					Latency:          settings.Get("latency").Float(),
					SampleRate:       settings.Get("sampleRate").Int64(),
					SampleSize:       settings.Get("sampleSize").Int64(),
					Volume:           settings.Get("volume").Float(),
					MediaTrackSettings: events.MediaTrackSettings{
						DeviceID: settings.Get("deviceId").String(),
						GroupID:  settings.Get("groupId").String(),
					},
				},
			})
		}
	}

	videosTracks := o.Get("getVideoTracks")
	if videosTracks != null && videosTracks != undefined {
		for i := 0; i < videosTracks.Length(); i++ {
			track := videosTracks.Index(i)
			settings := track.Call("getSettings", null)

			stream.Videos = append(stream.Videos, events.MediaStreamTrack{
				Core:       track,
				Enabled:    track.Get("enabled").Bool(),
				ID:         track.Get("id").String(),
				Kind:       track.Get("kind").String(),
				Label:      track.Get("label").String(),
				Muted:      track.Get("muted").Bool(),
				ReadyState: track.Get("readyState").Bool(),
				Remote:     track.Get("remote").Bool(),
				VideoSettings: &events.MediaVideoTrackSettings{
					AspectRatio: settings.Get("aspectRation").Float(),
					FrameRate:   settings.Get("frameRate").Float(),
					Height:      settings.Get("height").Int64(),
					Width:       settings.Get("width").Int64(),
					FacingMode:  settings.Get("facingMode").String(),
					MediaTrackSettings: events.MediaTrackSettings{
						DeviceID: settings.Get("deviceId").String(),
						GroupID:  settings.Get("groupId").String(),
					},
				},
			})
		}
	}

	return stream
}

function toTouches(o) events.TouchList {
	var th events.TouchList

	if o == null || o == undefined {
		return th
	}

	var touches []events.Touch

	for i := 0; i < o.Length(); i++ {
		ev := o.Index(i)
		touches = append(touches, events.Touch{
			ClientX:    ev.Get("clientX").Float(),
			ClientY:    ev.Get("clientY").Float(),
			OffsetX:    ev.Get("offsetX").Float(),
			OffsetY:    ev.Get("offsetY").Float(),
			PageX:      ev.Get("pageX").Float(),
			PageY:      ev.Get("pageY").Float(),
			ScreenX:    ev.Get("screenX").Float(),
			ScreenY:    ev.Get("screenY").Float(),
			Identifier: ev.Get("identifier").Float(),
		})

	}

	th.Touches = touches
	th.Length = len(touches)
	return th
}

// toGamepad returns a Gamepad struct from the js object.
function toGamepad(o) events.Gamepad {
	var pad events.Gamepad
	if o == null || o == undefined {
		return pad
	}

	pad.DisplayID = o.Get("displayId").String()
	pad.ID = o.Get("id").String()
	pad.Index = o.Get("index").Int()
	pad.Mapping = o.Get("mapping").String()
	pad.Connected = o.Get("connected").Bool()
	pad.Timestamp = o.Get("timestamp").Float()

	axes := o.Get("axes")
	if axes != null && axes != undefined {
		for i := 0; i < axes.Length(); i++ {
			pad.Axes = append(pad.Axes, axes.Index(i).Float())
		}
	}

	buttons := o.Get("buttons")
	if buttons != null && buttons != undefined {
		for i := 0; i < buttons.Length(); i++ {
			button := buttons.Index(i)
			pad.Buttons = append(pad.Buttons, events.Button{
				Value:   button.Get("value").Float(),
				Pressed: button.Get("pressed").Bool(),
			})
		}
	}

	return pad
}

// toDataTransfer returns a transfer object from the Object.
function toDataTransfer(o) events.DataTransfer {
	var dt events.DataTransfer
	if o == null || o == undefined {
		return dt
	}
	dt.DropEffect = o.Get("dropEffect").String()
	dt.EffectAllowed = o.Get("effectAllowed").String()

	types := o.Get("types")
	if types != null && types != undefined {
		dt.Types = shell.ObjectToStringList(types)
	}

	var dItems []events.DataTransferItem

	items := o.Get("items")
	if items != null && items != undefined {
		for i := 0; i < items.Length(); i++ {
			item := items.Index(i)
			dItems = append(dItems, events.DataTransferItem{
				Name: item.Get("name").String(),
				Size: item.Get("size").Int(),
				Data: fromFile(item),
			})
		}
	}

	var dFiles []events.DataTransferItem

	files := o.Get("files")
	if files != null && files != undefined {
		for i := 0; i < files.Length(); i++ {
			item := files.Index(i)
			dFiles = append(dFiles, events.DataTransferItem{
				Name: item.Get("name").String(),
				Size: item.Get("size").Int(),
				Data: fromFile(item),
			})
		}
	}

	dt.Items = events.DataTransferItemList{Items: dItems}
	dt.Files = dFiles
	return dt
}
