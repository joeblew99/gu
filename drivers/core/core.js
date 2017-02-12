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
			AnimationName: ev.animationName,
			ElapsedTime:   ev.elapsedTime,
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

	fileReader = new FileReader()
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

	fileReader = new FileReader()
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
	md.Alpha = o.alpha
	md.Beta = o.beta
	md.Gamma = o.gamma
	return md
}

// toMediaStream returns a events.MediaStream object.
function toMediaStream(o) events.MediaStream {
	var stream events.MediaStream
	if o == null || o == undefined {
		return stream
	}

	stream.Active = o.active
	stream.Ended = o.ended
	stream.ID = o.id

	audioTracks = o.getAudioTracks()
	if audioTracks != null && audioTracks != undefined {
		for i = 0; i < audioTracks.length; i++ {
			track = audioTracks.Index(i)
			settings = track.getSettings()

			stream.Audios = append(stream.Audios, events.MediaStreamTrack{
				Core:       track,
				Enabled:    track.enabled,
				ID:         track.id,
				Kind:       track.kind,
				Label:      track.label,
				Muted:      track.muted,
				ReadyState: track.readyState,
				Remote:     track.remote,
				AudioSettings: &events.MediaAudioTrackSettings{
					ChannelCount:     settings.channelCount.Int(),
					EchoCancellation: settings.echoCancellation,
					Latency:          settings.latency,
					SampleRate:       settings.sampleRate.Int64(),
					SampleSize:       settings.sampleSize.Int64(),
					Volume:           settings.volume,
					MediaTrackSettings: events.MediaTrackSettings{
						DeviceID: settings.deviceId,
						GroupID:  settings.groupId,
					},
				},
			})
		}
	}

	videosTracks = o.getVideoTracks()
	if videosTracks != null && videosTracks != undefined {
		for i = 0; i < videosTracks.length; i++ {
			track = videosTracks.Index(i)
			settings = track.getSettings()

			stream.Videos = append(stream.Videos, events.MediaStreamTrack{
				Core:       track,
				Enabled:    track.enabled,
				ID:         track.id,
				Kind:       track.kind,
				Label:      track.label,
				Muted:      track.muted,
				ReadyState: track.readyState,
				Remote:     track.remote,
				VideoSettings: &events.MediaVideoTrackSettings{
					AspectRatio: settings.aspectRation,
					FrameRate:   settings.frameRate,
					Height:      settings.height.Int64(),
					Width:       settings.width.Int64(),
					FacingMode:  settings.facingMode,
					MediaTrackSettings: events.MediaTrackSettings{
						DeviceID: settings.deviceId,
						GroupID:  settings.groupId,
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

	for i = 0; i < o.length; i++ {
		ev = o.Index(i)
		touches = append(touches, events.Touch{
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

	pad.DisplayID = o.displayId
	pad.ID = o.id
	pad.Index = o.index.Int()
	pad.Mapping = o.mapping
	pad.Connected = o.connected
	pad.Timestamp = o.timestamp

	axes = o.axes
	if axes != null && axes != undefined {
		for i = 0; i < axes.length; i++ {
			pad.Axes = append(pad.Axes, axes.Index(i))
		}
	}

	buttons = o.buttons
	if buttons != null && buttons != undefined {
		for i = 0; i < buttons.length; i++ {
			button = buttons.Index(i)
			pad.Buttons = append(pad.Buttons, events.Button{
				Value:   button.value,
				Pressed: button.pressed,
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
	dt.DropEffect = o.dropEffect
	dt.EffectAllowed = o.effectAllowed

	types = o.types
	if types != null && types != undefined {
		dt.Types = shell.ObjectToStringList(types)
	}

	var dItems []events.DataTransferItem

	items = o.items
	if items != null && items != undefined {
		for i = 0; i < items.length; i++ {
			item = items.Index(i)
			dItems = append(dItems, events.DataTransferItem{
				Name: item.name,
				Size: item.size.Int(),
				Data: fromFile(item),
			})
		}
	}

	var dFiles []events.DataTransferItem

	files = o.files
	if files != null && files != undefined {
		for i = 0; i < files.length; i++ {
			item = files.Index(i)
			dFiles = append(dFiles, events.DataTransferItem{
				Name: item.name,
				Size: item.size.Int(),
				Data: fromFile(item),
			})
		}
	}

	dt.Items = events.DataTransferItemList{Items: dItems}
	dt.Files = dFiles
	return dt
}
