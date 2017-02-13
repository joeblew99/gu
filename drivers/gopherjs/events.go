package gopherjs

import (
	"errors"
	"sync"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gu-io/gu/events"
	"github.com/gu-io/gu/notifications/mque"
	"github.com/gu-io/gu/shell"
)

// GetEvent returns the appropriate event from the provided structures.
func GetEvent(ev *js.Object, handle mque.End) *events.BaseEvent {
	if ev == nil || ev == js.Undefined {
		return nil
	}

	c := ev.Get("constructor")

	switch c {
	case js.Global.Get("AnimationEvent"):
		return events.NewBaseEvent(&events.AnimationEvent{
			Core:          ev,
			AnimationName: ev.Get("animationName").String(),
			ElapsedTime:   ev.Get("elapsedTime").Float(),
		}, handle)
	case js.Global.Get("AudioProcessingEvent"):
		return events.NewBaseEvent(&events.AudioProcessingEvent{
			Core:         ev,
			PlaybackTime: ev.Get("playbackTime").Float(),
		}, handle)
	case js.Global.Get("BeforeInputEvent"):
		return events.NewBaseEvent(&events.BeforeInputEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("BeforeUnloadEvent"):
		return events.NewBaseEvent(&events.BeforeUnloadEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("BlobEvent"):
		return events.NewBaseEvent(&events.BlobEvent{
			Core: ev,
			Data: fromBlob(ev.Get("data")),
		}, handle)
	case js.Global.Get("ChangeEvent"):
		return events.NewBaseEvent(&events.ChangeEvent{
			Core:  ev,
			Value: ev.Get("target").Get("value").String(),
		}, handle)
	case js.Global.Get("ClipboardEvent"):
		return events.NewBaseEvent(&events.ClipboardEvent{
			Core: ev,
			Data: toDataTransfer(ev.Get("clipboardData")),
		}, handle)
	case js.Global.Get("CloseEvent"):
		return events.NewBaseEvent(&events.CloseEvent{
			Core:     ev,
			Code:     ev.Get("code").Int(),
			Reason:   ev.Get("reason").String(),
			WasClean: ev.Get("wasClean").Bool(),
		}, handle)
	case js.Global.Get("CompositionEvent"):
		return events.NewBaseEvent(&events.CompositionEvent{
			Core:   ev,
			Text:   ev.Get("text").String(),
			Data:   ev.Get("data").String(),
			Locale: ev.Get("locale").String(),
		}, handle)
	case js.Global.Get("CSSFontFaceLoadEvent"):
		return events.NewBaseEvent(&events.CSSFontFaceLoadEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("CustomEvent"):
		return events.NewBaseEvent(&events.CustomEvent{
			Core:   ev,
			Detail: ev.Get("detail").Interface(),
		}, handle)
	case js.Global.Get("DeviceLightEvent"):
		return events.NewBaseEvent(&events.DeviceLightEvent{
			Core:  ev,
			Value: ev.Get("value").Float(),
		}, handle)
	case js.Global.Get("DeviceMotionEvent"):
		return events.NewBaseEvent(&events.DeviceMotionEvent{
			Core:                         ev,
			Interval:                     ev.Get("interval").Float(),
			Acceleration:                 toMotionData(ev.Get("acceleration")),
			AccelerationIncludingGravity: toMotionData(ev.Get("accelerationIncludingGravity")),
			RotationRate:                 toRotationData(ev.Get("rotationRate")),
		}, handle)
	case js.Global.Get("DeviceOrientationEvent"):
		return events.NewBaseEvent(&events.DeviceOrientationEvent{
			Core:     ev,
			Absolute: ev.Get("absolute").Bool(),
			Alpha:    ev.Get("alpha").Float(),
			Beta:     ev.Get("beta").Float(),
			Gamma:    ev.Get("gamma").Float(),
		}, handle)
	case js.Global.Get("DeviceProximityEvent"):
		return events.NewBaseEvent(&events.DeviceProximityEvent{
			Core:  ev,
			Max:   ev.Get("max").Float(),
			Min:   ev.Get("min").Float(),
			Value: ev.Get("value").Float(),
		}, handle)
	case js.Global.Get("DOMTransactionEvent"):
		return events.NewBaseEvent(&events.DOMTransactionEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DragStartEvent"):
		return events.NewBaseEvent(&events.DragStartEvent{
			MouseEvent: &events.MouseEvent{
				UIEvent: &events.UIEvent{
					Core:               ev,
					IsChar:             ev.Get("isChar").Bool(),
					LayerX:             ev.Get("layerX").Float(),
					LayerY:             ev.Get("layerY").Float(),
					PageX:              ev.Get("pageX").Float(),
					PageY:              ev.Get("pageY").Float(),
					Detail:             ev.Get("detail").Int(),
					SourceCapabilities: toInputSourceCapability(ev.Get("sourceCapabilities")),
				},
				ClientX:  ev.Get("clientX").Float(),
				ClientY:  ev.Get("clientY").Float(),
				OffsetX:  ev.Get("offsetX").Float(),
				OffsetY:  ev.Get("offsetY").Float(),
				PageX:    ev.Get("pageX").Float(),
				PageY:    ev.Get("pageY").Float(),
				ScreenX:  ev.Get("screenX").Float(),
				ScreenY:  ev.Get("screenY").Float(),
				MovemenX: ev.Get("movementX").Float(),
				MovemenY: ev.Get("movementY").Float(),
				Button:   ev.Get("button").Int(),
				Detail:   ev.Get("clientX").Int(),
				AltKey:   ev.Get("altKey").Bool(),
				CtrlKey:  ev.Get("ctrlKey").Bool(),
				MetaKey:  ev.Get("metaKey").Bool(),
				ShiftKey: ev.Get("shiftKey").Bool(),
			},
			DataTransfer: toDataTransfer(ev.Get("dataTransfer")),
			Core:         ev,
		}, handle)
	case js.Global.Get("DragExitEvent"):
		return events.NewBaseEvent(&events.DragExitEvent{
			MouseEvent: &events.MouseEvent{
				UIEvent: &events.UIEvent{
					Core:               ev,
					IsChar:             ev.Get("isChar").Bool(),
					LayerX:             ev.Get("layerX").Float(),
					LayerY:             ev.Get("layerY").Float(),
					PageX:              ev.Get("pageX").Float(),
					PageY:              ev.Get("pageY").Float(),
					Detail:             ev.Get("detail").Int(),
					SourceCapabilities: toInputSourceCapability(ev.Get("sourceCapabilities")),
				},
				ClientX:  ev.Get("clientX").Float(),
				ClientY:  ev.Get("clientY").Float(),
				OffsetX:  ev.Get("offsetX").Float(),
				OffsetY:  ev.Get("offsetY").Float(),
				PageX:    ev.Get("pageX").Float(),
				PageY:    ev.Get("pageY").Float(),
				ScreenX:  ev.Get("screenX").Float(),
				ScreenY:  ev.Get("screenY").Float(),
				MovemenX: ev.Get("movementX").Float(),
				MovemenY: ev.Get("movementY").Float(),
				Button:   ev.Get("button").Int(),
				Detail:   ev.Get("clientX").Int(),
				AltKey:   ev.Get("altKey").Bool(),
				CtrlKey:  ev.Get("ctrlKey").Bool(),
				MetaKey:  ev.Get("metaKey").Bool(),
				ShiftKey: ev.Get("shiftKey").Bool(),
			},
			Core:         ev,
			DataTransfer: toDataTransfer(ev.Get("dataTransfer")),
		}, handle)
	case js.Global.Get("DragEndEvent"):
		return events.NewBaseEvent(&events.DragEndEvent{
			MouseEvent: &events.MouseEvent{
				UIEvent: &events.UIEvent{
					Core:               ev,
					IsChar:             ev.Get("isChar").Bool(),
					LayerX:             ev.Get("layerX").Float(),
					LayerY:             ev.Get("layerY").Float(),
					PageX:              ev.Get("pageX").Float(),
					PageY:              ev.Get("pageY").Float(),
					Detail:             ev.Get("detail").Int(),
					SourceCapabilities: toInputSourceCapability(ev.Get("sourceCapabilities")),
				},
				ClientX:  ev.Get("clientX").Float(),
				ClientY:  ev.Get("clientY").Float(),
				OffsetX:  ev.Get("offsetX").Float(),
				OffsetY:  ev.Get("offsetY").Float(),
				PageX:    ev.Get("pageX").Float(),
				PageY:    ev.Get("pageY").Float(),
				ScreenX:  ev.Get("screenX").Float(),
				ScreenY:  ev.Get("screenY").Float(),
				MovemenX: ev.Get("movementX").Float(),
				MovemenY: ev.Get("movementY").Float(),
				Button:   ev.Get("button").Int(),
				Detail:   ev.Get("clientX").Int(),
				AltKey:   ev.Get("altKey").Bool(),
				CtrlKey:  ev.Get("ctrlKey").Bool(),
				MetaKey:  ev.Get("metaKey").Bool(),
				ShiftKey: ev.Get("shiftKey").Bool(),
			},
			DataTransfer: toDataTransfer(ev.Get("dataTransfer")),
			Core:         ev,
		}, handle)
	case js.Global.Get("DragEnterEvent"):
		return events.NewBaseEvent(&events.DragEnterEvent{
			MouseEvent: &events.MouseEvent{
				UIEvent: &events.UIEvent{
					Core:               ev,
					IsChar:             ev.Get("isChar").Bool(),
					LayerX:             ev.Get("layerX").Float(),
					LayerY:             ev.Get("layerY").Float(),
					PageX:              ev.Get("pageX").Float(),
					PageY:              ev.Get("pageY").Float(),
					Detail:             ev.Get("detail").Int(),
					SourceCapabilities: toInputSourceCapability(ev.Get("sourceCapabilities")),
				},
				ClientX:  ev.Get("clientX").Float(),
				ClientY:  ev.Get("clientY").Float(),
				OffsetX:  ev.Get("offsetX").Float(),
				OffsetY:  ev.Get("offsetY").Float(),
				PageX:    ev.Get("pageX").Float(),
				PageY:    ev.Get("pageY").Float(),
				ScreenX:  ev.Get("screenX").Float(),
				ScreenY:  ev.Get("screenY").Float(),
				MovemenX: ev.Get("movementX").Float(),
				MovemenY: ev.Get("movementY").Float(),
				Button:   ev.Get("button").Int(),
				Detail:   ev.Get("clientX").Int(),
				AltKey:   ev.Get("altKey").Bool(),
				CtrlKey:  ev.Get("ctrlKey").Bool(),
				MetaKey:  ev.Get("metaKey").Bool(),
				ShiftKey: ev.Get("shiftKey").Bool(),
			},
			DataTransfer: toDataTransfer(ev.Get("dataTransfer")),
			Core:         ev,
		}, handle)
	case js.Global.Get("DragLeaveEvent"):
		return events.NewBaseEvent(&events.DragLeaveEvent{
			MouseEvent: &events.MouseEvent{
				UIEvent: &events.UIEvent{
					Core:               ev,
					IsChar:             ev.Get("isChar").Bool(),
					LayerX:             ev.Get("layerX").Float(),
					LayerY:             ev.Get("layerY").Float(),
					PageX:              ev.Get("pageX").Float(),
					PageY:              ev.Get("pageY").Float(),
					Detail:             ev.Get("detail").Int(),
					SourceCapabilities: toInputSourceCapability(ev.Get("sourceCapabilities")),
				},
				ClientX:  ev.Get("clientX").Float(),
				ClientY:  ev.Get("clientY").Float(),
				OffsetX:  ev.Get("offsetX").Float(),
				OffsetY:  ev.Get("offsetY").Float(),
				PageX:    ev.Get("pageX").Float(),
				PageY:    ev.Get("pageY").Float(),
				ScreenX:  ev.Get("screenX").Float(),
				ScreenY:  ev.Get("screenY").Float(),
				MovemenX: ev.Get("movementX").Float(),
				MovemenY: ev.Get("movementY").Float(),
				Button:   ev.Get("button").Int(),
				Detail:   ev.Get("clientX").Int(),
				AltKey:   ev.Get("altKey").Bool(),
				CtrlKey:  ev.Get("ctrlKey").Bool(),
				MetaKey:  ev.Get("metaKey").Bool(),
				ShiftKey: ev.Get("shiftKey").Bool(),
			},
			DataTransfer: toDataTransfer(ev.Get("dataTransfer")),
			Core:         ev,
		}, handle)
	case js.Global.Get("DragOverEvent"):
		return events.NewBaseEvent(&events.DragOverEvent{
			MouseEvent: &events.MouseEvent{
				UIEvent: &events.UIEvent{
					Core:               ev,
					IsChar:             ev.Get("isChar").Bool(),
					LayerX:             ev.Get("layerX").Float(),
					LayerY:             ev.Get("layerY").Float(),
					PageX:              ev.Get("pageX").Float(),
					PageY:              ev.Get("pageY").Float(),
					Detail:             ev.Get("detail").Int(),
					SourceCapabilities: toInputSourceCapability(ev.Get("sourceCapabilities")),
				},
				ClientX:  ev.Get("clientX").Float(),
				ClientY:  ev.Get("clientY").Float(),
				OffsetX:  ev.Get("offsetX").Float(),
				OffsetY:  ev.Get("offsetY").Float(),
				PageX:    ev.Get("pageX").Float(),
				PageY:    ev.Get("pageY").Float(),
				ScreenX:  ev.Get("screenX").Float(),
				ScreenY:  ev.Get("screenY").Float(),
				MovemenX: ev.Get("movementX").Float(),
				MovemenY: ev.Get("movementY").Float(),
				Button:   ev.Get("button").Int(),
				Detail:   ev.Get("clientX").Int(),
				AltKey:   ev.Get("altKey").Bool(),
				CtrlKey:  ev.Get("ctrlKey").Bool(),
				MetaKey:  ev.Get("metaKey").Bool(),
				ShiftKey: ev.Get("shiftKey").Bool(),
			},
			DataTransfer: toDataTransfer(ev.Get("dataTransfer")),
			Core:         ev,
		}, handle)
	case js.Global.Get("DropEvent"):
		return events.NewBaseEvent(&events.DropEvent{
			MouseEvent: &events.MouseEvent{
				UIEvent: &events.UIEvent{
					Core:               ev,
					IsChar:             ev.Get("isChar").Bool(),
					LayerX:             ev.Get("layerX").Float(),
					LayerY:             ev.Get("layerY").Float(),
					PageX:              ev.Get("pageX").Float(),
					PageY:              ev.Get("pageY").Float(),
					Detail:             ev.Get("detail").Int(),
					SourceCapabilities: toInputSourceCapability(ev.Get("sourceCapabilities")),
				},
				ClientX:  ev.Get("clientX").Float(),
				ClientY:  ev.Get("clientY").Float(),
				OffsetX:  ev.Get("offsetX").Float(),
				OffsetY:  ev.Get("offsetY").Float(),
				PageX:    ev.Get("pageX").Float(),
				PageY:    ev.Get("pageY").Float(),
				ScreenX:  ev.Get("screenX").Float(),
				ScreenY:  ev.Get("screenY").Float(),
				MovemenX: ev.Get("movementX").Float(),
				MovemenY: ev.Get("movementY").Float(),
				Button:   ev.Get("button").Int(),
				Detail:   ev.Get("clientX").Int(),
				AltKey:   ev.Get("altKey").Bool(),
				CtrlKey:  ev.Get("ctrlKey").Bool(),
				MetaKey:  ev.Get("metaKey").Bool(),
				ShiftKey: ev.Get("shiftKey").Bool(),
			},
			DataTransfer: toDataTransfer(ev.Get("dataTransfer")),
			Core:         ev,
		}, handle)
	case js.Global.Get("DragEvent"):
		return events.NewBaseEvent(&events.DragEvent{
			MouseEvent: &events.MouseEvent{
				UIEvent: &events.UIEvent{
					Core:               ev,
					IsChar:             ev.Get("isChar").Bool(),
					LayerX:             ev.Get("layerX").Float(),
					LayerY:             ev.Get("layerY").Float(),
					PageX:              ev.Get("pageX").Float(),
					PageY:              ev.Get("pageY").Float(),
					Detail:             ev.Get("detail").Int(),
					SourceCapabilities: toInputSourceCapability(ev.Get("sourceCapabilities")),
				},
				ClientX:  ev.Get("clientX").Float(),
				ClientY:  ev.Get("clientY").Float(),
				OffsetX:  ev.Get("offsetX").Float(),
				OffsetY:  ev.Get("offsetY").Float(),
				PageX:    ev.Get("pageX").Float(),
				PageY:    ev.Get("pageY").Float(),
				ScreenX:  ev.Get("screenX").Float(),
				ScreenY:  ev.Get("screenY").Float(),
				MovemenX: ev.Get("movementX").Float(),
				MovemenY: ev.Get("movementY").Float(),
				Button:   ev.Get("button").Int(),
				Detail:   ev.Get("clientX").Int(),
				AltKey:   ev.Get("altKey").Bool(),
				CtrlKey:  ev.Get("ctrlKey").Bool(),
				MetaKey:  ev.Get("metaKey").Bool(),
				ShiftKey: ev.Get("shiftKey").Bool(),
			},
			DataTransfer: toDataTransfer(ev.Get("dataTransfer")),
			Core:         ev,
		}, handle)
	case js.Global.Get("EditingBeforeInputEvent"):
		return events.NewBaseEvent(&events.EditingBeforeInputEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("ErrorEvent"):
		var err error
		if evx := ev.Get("error"); evx != nil && evx != js.Undefined {
			err = errors.New(ev.Get("error").String())
		}

		return events.NewBaseEvent(&events.ErrorEvent{
			Core:       ev,
			Message:    ev.Get("message").String(),
			Filename:   ev.Get("filename").String(),
			LineNumber: ev.Get("lineno").Int(),
			ColNumber:  ev.Get("colno").Int(),
			Error:      err,
		}, handle)
	case js.Global.Get("FocusEvent"):
		return events.NewBaseEvent(&events.FocusEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("GamepadEvent"):
		return events.NewBaseEvent(&events.GamepadEvent{
			Core:    ev,
			Gamepad: toGamepad(ev.Get("gamepad")),
		}, handle)
	case js.Global.Get("HashChangeEvent"):
		return events.NewBaseEvent(&events.HashChangeEvent{
			Core: ev,
			Old:  ev.Get("oldURL").String(),
			New:  ev.Get("newURL").String(),
		}, handle)
	case js.Global.Get("IDBVersionChangeEvent"):
		return events.NewBaseEvent(&events.IDBVersionChangeEvent{
			Core:       ev,
			OldVersion: ev.Get("oldVersion").Int64(),
			NewVersion: ev.Get("newVersion").Int64(),
		}, handle)
	case js.Global.Get("KeyboardEvent"):
		return events.NewBaseEvent(&events.KeyboardEvent{
			Core:          ev,
			CharCode:      ev.Get("charCode").Int(),
			Key:           ev.Get("key").String(),
			Locale:        ev.Get("locale").String(),
			AltKey:        ev.Get("altKey").Bool(),
			CtrlKey:       ev.Get("ctrlKey").Bool(),
			MetaKey:       ev.Get("metaKey").Bool(),
			ShiftKey:      ev.Get("shiftKey").Bool(),
			Repeat:        ev.Get("repeat").Bool(),
			Location:      ev.Get("location").Int(),
			ModifiedState: ev.Get("getModifierState").Bool(),
			KeyIdentifier: ev.Get("keyIdentifier").String(),
			KeyLocation:   events.KeyLocation(ev.Get("KeyLocation").Uint64()),
			KeyCode:       events.KeyCode(ev.Get("keyCode").Uint64()),
		}, handle)
	case js.Global.Get("MediaStreamEvent"):
		return events.NewBaseEvent(&events.MediaStreamEvent{
			Core:   ev,
			Stream: toMediaStream(ev.Get("stream")),
		}, handle)
	case js.Global.Get("MessageEvent"):
		var data []byte

		if so := ev.Get("data"); so != nil && so != js.Undefined {
			switch so {
			case js.Global.Get("String"):
				data = []byte(so.String())
				break
			case js.Global.Get("Blob"):
				data = fromBlob(so)
				break
			case js.Global.Get("ArrayBuffer"):
				data = fromFile(so)
				break
			}
		}

		return events.NewBaseEvent(&events.MessageEvent{
			Core:   ev,
			Data:   data,
			Origin: ev.Get("origin").String(),
			Source: ev.Get("source").String(),
			Port:   ev.Get("port").Int(),
		}, handle)
	case js.Global.Get("MouseEvent"):
		return events.NewBaseEvent(&events.MouseEvent{
			UIEvent: &events.UIEvent{
				Core: ev,
			},
			ClientX:  ev.Get("clientX").Float(),
			ClientY:  ev.Get("clientY").Float(),
			OffsetX:  ev.Get("offsetX").Float(),
			OffsetY:  ev.Get("offsetY").Float(),
			PageX:    ev.Get("pageX").Float(),
			PageY:    ev.Get("pageY").Float(),
			ScreenX:  ev.Get("screenX").Float(),
			ScreenY:  ev.Get("screenY").Float(),
			MovemenX: ev.Get("movementX").Float(),
			MovemenY: ev.Get("movementY").Float(),
			Button:   ev.Get("button").Int(),
			Detail:   ev.Get("clientX").Int(),
			AltKey:   ev.Get("altKey").Bool(),
			CtrlKey:  ev.Get("ctrlKey").Bool(),
			MetaKey:  ev.Get("metaKey").Bool(),
			ShiftKey: ev.Get("shiftKey").Bool(),
		}, handle)
	case js.Global.Get("MutationEvent"):
		return events.NewBaseEvent(&events.MutationEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("OfflineAudioCompletionEvent"):
		return events.NewBaseEvent(&events.OfflineAudioCompletionEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("PageTransitionEvent"):
		return events.NewBaseEvent(&events.PageTransitionEvent{
			Core:      ev,
			Persisted: ev.Get("persisted").Bool(),
		}, handle)
	case js.Global.Get("PointerEvent"):
		return events.NewBaseEvent(&events.PointerEvent{
			Core:        ev,
			Width:       ev.Get("width").Int(),
			Height:      ev.Get("height").Int(),
			TiltX:       ev.Get("tiltX").Float(),
			TiltY:       ev.Get("tiltY").Float(),
			Pressure:    ev.Get("pressure").Float(),
			IsPrimary:   ev.Get("isPrimary").Bool(),
			PointerID:   ev.Get("pointerId").Int(),
			PointerType: ev.Get("pointerType").String(),
		}, handle)
	case js.Global.Get("FetchEvent"):
		var req shell.WebRequest

		if ro := ev.Get("request"); ro != nil && ro != js.Undefined {
			req, _ = shell.ObjectToWebRequest(ro)
		}

		return events.NewBaseEvent(&events.FetchEvent{
			Core:     ev,
			IsReload: ev.Get("isReload").Bool(),
			ClientID: ev.Get("clientId").String(),
			Request:  req,
		}, handle)
	case js.Global.Get("PopStateEvent"):
		return events.NewBaseEvent(&events.PopStateEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("ProgressEvent"):
		return events.NewBaseEvent(&events.ProgressEvent{
			Core:             ev,
			LengthComputable: ev.Get("lengthComputable").Bool(),
			Loaded:           ev.Get("loaded").Uint64(),
			Total:            ev.Get("total").Int(),
		}, handle)
	case js.Global.Get("RelatedEvent"):
		return events.NewBaseEvent(&events.RelatedEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("RTCPeerConnectionIceEvent"):
		return events.NewBaseEvent(&events.RTCPeerConnectionIceEvent{
			Core:      ev,
			Candidate: ev.Get("candidate").String(),
		}, handle)
	case js.Global.Get("SensorEvent"):
		return events.NewBaseEvent(&events.SensorEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("StorageEvent"):
		return events.NewBaseEvent(&events.StorageEvent{
			Core:        ev,
			Key:         ev.Get("key").String(),
			NewValue:    ev.Get("newValue").String(),
			OldValue:    ev.Get("oldValue").String(),
			URL:         ev.Get("url").String(),
			StorageArea: ev.Get("storageArea").Interface(),
		}, handle)
	case js.Global.Get("SVGEvent"):
		return events.NewBaseEvent(&events.SVGEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("SVGZoomEvent"):
		return events.NewBaseEvent(&events.SVGZoomEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("TimeEvent"):
		return events.NewBaseEvent(&events.TimeEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("TouchEvent"):
		return events.NewBaseEvent(&events.TouchEvent{
			Core:          ev,
			AltKey:        ev.Get("altKey").Bool(),
			CtrlKey:       ev.Get("ctrlKey").Bool(),
			MetaKey:       ev.Get("metaKey").Bool(),
			ShiftKey:      ev.Get("shiftKey").Bool(),
			TargetTouches: toTouches(ev.Get("touches")),
			Touches:       toTouches(ev.Get("targetTouches")),
		}, handle)
	case js.Global.Get("TrackEvent"):
		return events.NewBaseEvent(&events.TrackEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("TransitionEvent"):
		return events.NewBaseEvent(&events.TransitionEvent{
			Core:          ev,
			ElapsedTime:   ev.Get("elapsedTime").Float(),
			PropertyName:  ev.Get("propertyName").String(),
			PseudoElement: ev.Get("pseudoElement").String(),
		}, handle)
	case js.Global.Get("UIEvent"):
		return events.NewBaseEvent(&events.UIEvent{
			Core:               ev,
			IsChar:             ev.Get("isChar").Bool(),
			LayerX:             ev.Get("layerX").Float(),
			LayerY:             ev.Get("layerY").Float(),
			PageX:              ev.Get("pageX").Float(),
			PageY:              ev.Get("pageY").Float(),
			Detail:             ev.Get("detail").Int(),
			SourceCapabilities: toInputSourceCapability(ev.Get("sourceCapabilities")),
		}, handle)
	case js.Global.Get("UserProximityEvent"):
		return events.NewBaseEvent(&events.UserProximityEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("WheelEvent"):
		return events.NewBaseEvent(&events.WheelEvent{
			Core:      ev,
			DeltaX:    ev.Get("deltaX").Float(),
			DeltaY:    ev.Get("deltaX").Float(),
			DeltaZ:    ev.Get("deltaX").Float(),
			DeltaMode: events.DeltaMode(ev.Get("deltaMode").Uint64()),
		}, handle)
	}

	return events.NewBaseEvent(ev, handle)
}

// fromBlob transform the providded js.Object blob into a byte slice.
func fromBlob(o *js.Object) []byte {
	if o == nil || o == js.Undefined {
		return nil
	}

	var buf []byte
	var wg sync.WaitGroup

	fileReader := js.Global.Get("FileReader").New()
	fileReader.Set("onloadend", js.MakeFunc(func(this *js.Object, _ []*js.Object) interface{} {
		defer wg.Done()
		buf = js.Global.Get("Uint8Array").New(this.Get("result")).Interface().([]uint8)
		return nil
	}))

	fileReader.Call("readAsArrayBuffer", o)

	wg.Wait()
	return buf
}

// fromFile transform the providded js.Object blob into a byte slice.
func fromFile(o *js.Object) []byte {
	if o == nil || o == js.Undefined {
		return nil
	}

	var buf []byte
	var wg sync.WaitGroup

	fileReader := js.Global.Get("FileReader").New()
	fileReader.Set("onloadend", js.MakeFunc(func(this *js.Object, _ []*js.Object) interface{} {
		defer wg.Done()
		buf = js.Global.Get("Uint8Array").New(this.Get("result")).Interface().([]uint8)
		return nil
	}))

	fileReader.Call("readAsArrayBuffer", o)
	// fileReader.Call("readAsBinaryString", o)

	wg.Wait()
	return buf
}

// toInputSourceCapability returns the events.InputDeviceCapabilities from the js.object.
func toInputSourceCapability(o *js.Object) *events.InputDeviceCapabilities {
	if o == nil || o == js.Undefined {
		return nil
	}

	return &events.InputDeviceCapabilities{
		FiresTouchEvent: o.Get("firesTouchEvent").Bool(),
	}
}

// toMotionData returns a motionData object from the js.Object.
func toMotionData(o *js.Object) events.MotionData {
	var md events.MotionData

	if o == nil || o == js.Undefined {
		return md
	}

	md.X = o.Get("x").Float()
	md.Y = o.Get("y").Float()
	md.Z = o.Get("z").Float()
	return md
}

// toRotationData returns a RotationData object from the js.Object.
func toRotationData(o *js.Object) events.RotationData {
	var md events.RotationData
	if o == nil || o == js.Undefined {
		return md
	}
	md.Alpha = o.Get("alpha").Float()
	md.Beta = o.Get("beta").Float()
	md.Gamma = o.Get("gamma").Float()
	return md
}

// toMediaStream returns a events.MediaStream object.
func toMediaStream(o *js.Object) events.MediaStream {
	var stream events.MediaStream
	if o == nil || o == js.Undefined {
		return stream
	}

	stream.Active = o.Get("active").Bool()
	stream.Ended = o.Get("ended").Bool()
	stream.ID = o.Get("id").String()

	audioTracks := o.Get("getAudioTracks")
	if audioTracks != nil && audioTracks != js.Undefined {
		for i := 0; i < audioTracks.Length(); i++ {
			track := audioTracks.Index(i)
			settings := track.Call("getSettings", nil)

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
	if videosTracks != nil && videosTracks != js.Undefined {
		for i := 0; i < videosTracks.Length(); i++ {
			track := videosTracks.Index(i)
			settings := track.Call("getSettings", nil)

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

func toTouches(o *js.Object) events.TouchList {
	var th events.TouchList

	if o == nil || o == js.Undefined {
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
func toGamepad(o *js.Object) events.Gamepad {
	var pad events.Gamepad
	if o == nil || o == js.Undefined {
		return pad
	}

	pad.DisplayID = o.Get("displayId").String()
	pad.ID = o.Get("id").String()
	pad.Index = o.Get("index").Int()
	pad.Mapping = o.Get("mapping").String()
	pad.Connected = o.Get("connected").Bool()
	pad.Timestamp = o.Get("timestamp").Float()

	axes := o.Get("axes")
	if axes != nil && axes != js.Undefined {
		for i := 0; i < axes.Length(); i++ {
			pad.Axes = append(pad.Axes, axes.Index(i).Float())
		}
	}

	buttons := o.Get("buttons")
	if buttons != nil && buttons != js.Undefined {
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

// toDataTransfer returns a transfer object from the js.Object.
func toDataTransfer(o *js.Object) events.DataTransfer {
	var dt events.DataTransfer
	if o == nil || o == js.Undefined {
		return dt
	}
	dt.DropEffect = o.Get("dropEffect").String()
	dt.EffectAllowed = o.Get("effectAllowed").String()

	types := o.Get("types")
	if types != nil && types != js.Undefined {
		dt.Types = shell.ObjectToStringList(types)
	}

	var dItems []events.DataTransferItem

	items := o.Get("items")
	if items != nil && items != js.Undefined {
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
	if files != nil && files != js.Undefined {
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
