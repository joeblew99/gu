package gopherjs

import (
	"errors"
	"strings"
	"sync"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gu-io/gu/eventx"
	"github.com/gu-io/gu/notifications/mque"
	"github.com/gu-io/gu/shell"
)

// GetEventByType returns the appropriate event from the provided js.Object, by
// introspectign the `type` field of the event. This is fallback for when constructor
// returns `Event`.
func GetEventByType(ev *js.Object, handle mque.End) *eventx.BaseEvent {
	if ev == nil || ev == js.Undefined {
		return nil
	}

	c := prepareEventName(ev.Get("type").String())

	switch c {
	case strings.ToLower("Animation"):
		return eventx.NewBaseEvent(&eventx.AnimationEvent{
			Core:          ev,
			AnimationName: ev.Get("animationName").String(),
			ElapsedTime:   ev.Get("elapsedTime").Float(),
		}, handle)
	case strings.ToLower("AudioProcessing"):
		return eventx.NewBaseEvent(&eventx.AudioProcessingEvent{
			Core:         ev,
			PlaybackTime: ev.Get("playbackTime").Float(),
		}, handle)
	case strings.ToLower("BeforeInput"):
		return eventx.NewBaseEvent(&eventx.BeforeInputEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("BeforeUnload"):
		return eventx.NewBaseEvent(&eventx.BeforeUnloadEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("Blob"):
		return eventx.NewBaseEvent(&eventx.BlobEvent{
			Core: ev,
			Data: fromBlob(ev.Get("data")),
		}, handle)
	case strings.ToLower("Change"):
		return eventx.NewBaseEvent(&eventx.ChangeEvent{
			Core:  ev,
			Value: ev.Get("target").Get("value").String(),
		}, handle)
	case strings.ToLower("Clipboard"):
		return eventx.NewBaseEvent(&eventx.ClipboardEvent{
			Core: ev,
			Data: toDataTransfer(ev.Get("clipboardData")),
		}, handle)
	case strings.ToLower("Close"):
		return eventx.NewBaseEvent(&eventx.CloseEvent{
			Core:     ev,
			Code:     ev.Get("code").Int(),
			Reason:   ev.Get("reason").String(),
			WasClean: ev.Get("wasClean").Bool(),
		}, handle)
	case strings.ToLower("Composition"):
		return eventx.NewBaseEvent(&eventx.CompositionEvent{
			Core:   ev,
			Text:   ev.Get("text").String(),
			Data:   ev.Get("data").String(),
			Locale: ev.Get("locale").String(),
		}, handle)
	case strings.ToLower("CSSFontFaceLoad"):
		return eventx.NewBaseEvent(&eventx.CSSFontFaceLoadEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("Custom"):
		return eventx.NewBaseEvent(&eventx.CustomEvent{
			Core:   ev,
			Detail: ev.Get("detail").Interface(),
		}, handle)
	case strings.ToLower("DeviceLight"):
		return eventx.NewBaseEvent(&eventx.DeviceLightEvent{
			Core:  ev,
			Value: ev.Get("value").Float(),
		}, handle)
	case strings.ToLower("DeviceMotion"):
		return eventx.NewBaseEvent(&eventx.DeviceMotionEvent{
			Core:                         ev,
			Interval:                     ev.Get("interval").Float(),
			Acceleration:                 toMotionData(ev.Get("acceleration")),
			AccelerationIncludingGravity: toMotionData(ev.Get("accelerationIncludingGravity")),
			RotationRate:                 toRotationData(ev.Get("rotationRate")),
		}, handle)
	case strings.ToLower("DeviceOrientation"):
		return eventx.NewBaseEvent(&eventx.DeviceOrientationEvent{
			Core:     ev,
			Absolute: ev.Get("absolute").Bool(),
			Alpha:    ev.Get("alpha").Float(),
			Beta:     ev.Get("beta").Float(),
			Gamma:    ev.Get("gamma").Float(),
		}, handle)
	case strings.ToLower("DeviceProximity"):
		return eventx.NewBaseEvent(&eventx.DeviceProximityEvent{
			Core:  ev,
			Max:   ev.Get("max").Float(),
			Min:   ev.Get("min").Float(),
			Value: ev.Get("value").Float(),
		}, handle)
	case strings.ToLower("DOMTransaction"):
		return eventx.NewBaseEvent(&eventx.DOMTransactionEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("DragStart"):
		return eventx.NewBaseEvent(&eventx.DragStartEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
	case strings.ToLower("DragExit"):
		return eventx.NewBaseEvent(&eventx.DragExitEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
	case strings.ToLower("DragEnd"):
		return eventx.NewBaseEvent(&eventx.DragEndEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
	case strings.ToLower("DragEnter"):
		return eventx.NewBaseEvent(&eventx.DragEnterEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
	case strings.ToLower("DragLeave"):
		return eventx.NewBaseEvent(&eventx.DragLeaveEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
	case strings.ToLower("DragOver"):
		return eventx.NewBaseEvent(&eventx.DragOverEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
	case strings.ToLower("Drop"):
		return eventx.NewBaseEvent(&eventx.DropEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
	case strings.ToLower("Drag"):
		return eventx.NewBaseEvent(&eventx.DragEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
	case strings.ToLower("EditingBeforeInput"):
		return eventx.NewBaseEvent(&eventx.EditingBeforeInputEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("Error"):
		var err error
		if evx := ev.Get("error"); evx != nil && evx != js.Undefined {
			err = errors.New(ev.Get("error").String())
		}

		return eventx.NewBaseEvent(&eventx.ErrorEvent{
			Core:       ev,
			Message:    ev.Get("message").String(),
			Filename:   ev.Get("filename").String(),
			LineNumber: ev.Get("lineno").Int(),
			ColNumber:  ev.Get("colno").Int(),
			Error:      err,
		}, handle)
	case strings.ToLower("Focus"):
		return eventx.NewBaseEvent(&eventx.FocusEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("Gamepad"):
		return eventx.NewBaseEvent(&eventx.GamepadEvent{
			Core:    ev,
			Gamepad: toGamepad(ev.Get("gamepad")),
		}, handle)
	case strings.ToLower("HashChange"):
		return eventx.NewBaseEvent(&eventx.HashChangeEvent{
			Core: ev,
			Old:  ev.Get("oldURL").String(),
			New:  ev.Get("newURL").String(),
		}, handle)
	case strings.ToLower("IDBVersionChange"):
		return eventx.NewBaseEvent(&eventx.IDBVersionChangeEvent{
			Core:       ev,
			OldVersion: ev.Get("oldVersion").Int64(),
			NewVersion: ev.Get("newVersion").Int64(),
		}, handle)
	case strings.ToLower("Keyboard"):
		return eventx.NewBaseEvent(&eventx.KeyboardEvent{
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
			KeyLocation:   eventx.KeyLocation(ev.Get("KeyLocation").Uint64()),
			KeyCode:       eventx.KeyCode(ev.Get("keyCode").Uint64()),
		}, handle)
	case strings.ToLower("MediaStream"):
		return eventx.NewBaseEvent(&eventx.MediaStreamEvent{
			Core:   ev,
			Stream: toMediaStream(ev.Get("stream")),
		}, handle)
	case strings.ToLower("Message"):
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

		return eventx.NewBaseEvent(&eventx.MessageEvent{
			Core:   ev,
			Data:   data,
			Origin: ev.Get("origin").String(),
			Source: ev.Get("source").String(),
			Port:   ev.Get("port").Int(),
		}, handle)
	case strings.ToLower("Mouse"):
		return eventx.NewBaseEvent(&eventx.MouseEvent{
			UIEvent: &eventx.UIEvent{
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
	case strings.ToLower("Mutation"):
		return eventx.NewBaseEvent(&eventx.MutationEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("OfflineAudioCompletion"):
		return eventx.NewBaseEvent(&eventx.OfflineAudioCompletionEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("PageTransition"):
		return eventx.NewBaseEvent(&eventx.PageTransitionEvent{
			Core:      ev,
			Persisted: ev.Get("persisted").Bool(),
		}, handle)
	case strings.ToLower("Pointer"):
		return eventx.NewBaseEvent(&eventx.PointerEvent{
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
	case strings.ToLower("Fetch"):
		var req shell.WebRequest

		if ro := ev.Get("request"); ro != nil && ro != js.Undefined {
			req, _ = shell.ObjectToWebRequest(ro)
		}

		return eventx.NewBaseEvent(&eventx.FetchEvent{
			Core:     ev,
			IsReload: ev.Get("isReload").Bool(),
			ClientID: ev.Get("clientId").String(),
			Request:  req,
		}, handle)
	case strings.ToLower("PopState"):
		return eventx.NewBaseEvent(&eventx.PopStateEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("Progress"):
		return eventx.NewBaseEvent(&eventx.ProgressEvent{
			Core:             ev,
			LengthComputable: ev.Get("lengthComputable").Bool(),
			Loaded:           ev.Get("loaded").Uint64(),
			Total:            ev.Get("total").Int(),
		}, handle)
	case strings.ToLower("Related"):
		return eventx.NewBaseEvent(&eventx.RelatedEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("RTCPeerConnectionIce"):
		return eventx.NewBaseEvent(&eventx.RTCPeerConnectionIceEvent{
			Core:      ev,
			Candidate: ev.Get("candidate").String(),
		}, handle)
	case strings.ToLower("Sensor"):
		return eventx.NewBaseEvent(&eventx.SensorEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("Storage"):
		return eventx.NewBaseEvent(&eventx.StorageEvent{
			Core:        ev,
			Key:         ev.Get("key").String(),
			NewValue:    ev.Get("newValue").String(),
			OldValue:    ev.Get("oldValue").String(),
			URL:         ev.Get("url").String(),
			StorageArea: ev.Get("storageArea").Interface(),
		}, handle)
	case strings.ToLower("SVG"):
		return eventx.NewBaseEvent(&eventx.SVGEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("SVGZoom"):
		return eventx.NewBaseEvent(&eventx.SVGZoomEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("Time"):
		return eventx.NewBaseEvent(&eventx.TimeEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("Touch"):
		return eventx.NewBaseEvent(&eventx.TouchEvent{
			Core:          ev,
			AltKey:        ev.Get("altKey").Bool(),
			CtrlKey:       ev.Get("ctrlKey").Bool(),
			MetaKey:       ev.Get("metaKey").Bool(),
			ShiftKey:      ev.Get("shiftKey").Bool(),
			TargetTouches: toTouches(ev.Get("touches")),
			Touches:       toTouches(ev.Get("targetTouches")),
		}, handle)
	case strings.ToLower("Track"):
		return eventx.NewBaseEvent(&eventx.TrackEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("Transition"):
		return eventx.NewBaseEvent(&eventx.TransitionEvent{
			Core:          ev,
			ElapsedTime:   ev.Get("elapsedTime").Float(),
			PropertyName:  ev.Get("propertyName").String(),
			PseudoElement: ev.Get("pseudoElement").String(),
		}, handle)
	case strings.ToLower("UI"):
		return eventx.NewBaseEvent(&eventx.UIEvent{
			Core:               ev,
			IsChar:             ev.Get("isChar").Bool(),
			LayerX:             ev.Get("layerX").Float(),
			LayerY:             ev.Get("layerY").Float(),
			PageX:              ev.Get("pageX").Float(),
			PageY:              ev.Get("pageY").Float(),
			Detail:             ev.Get("detail").Int(),
			SourceCapabilities: toInputSourceCapability(ev.Get("sourceCapabilities")),
		}, handle)
	case strings.ToLower("UserProximity"):
		return eventx.NewBaseEvent(&eventx.UserProximityEvent{
			Core: ev,
		}, handle)
	case strings.ToLower("Wheel"):
		return eventx.NewBaseEvent(&eventx.WheelEvent{
			Core:      ev,
			DeltaX:    ev.Get("deltaX").Float(),
			DeltaY:    ev.Get("deltaX").Float(),
			DeltaZ:    ev.Get("deltaX").Float(),
			DeltaMode: eventx.DeltaMode(ev.Get("deltaMode").Uint64()),
		}, handle)
	}

	return eventx.NewBaseEvent(ev, handle)
}

// GetEvent returns the appropriate event from the provided structures.
func GetEvent(ev *js.Object, handle mque.End) *eventx.BaseEvent {
	if ev == nil || ev == js.Undefined {
		return nil
	}

	c := ev.Get("constructor")

	// If we recieve a EventObject type whoes constructor is the base Event
	// type then we have to use the `type` field.
	if c == js.Global.Get("Event") {
		return GetEventByType(ev, handle)
	}

	switch c {
	case js.Global.Get("AnimationEvent"):
		return eventx.NewBaseEvent(&eventx.AnimationEvent{
			Core:          ev,
			AnimationName: ev.Get("animationName").String(),
			ElapsedTime:   ev.Get("elapsedTime").Float(),
		}, handle)
	case js.Global.Get("AudioProcessingEvent"):
		return eventx.NewBaseEvent(&eventx.AudioProcessingEvent{
			Core:         ev,
			PlaybackTime: ev.Get("playbackTime").Float(),
		}, handle)
	case js.Global.Get("BeforeInputEvent"):
		return eventx.NewBaseEvent(&eventx.BeforeInputEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("BeforeUnloadEvent"):
		return eventx.NewBaseEvent(&eventx.BeforeUnloadEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("BlobEvent"):
		return eventx.NewBaseEvent(&eventx.BlobEvent{
			Core: ev,
			Data: fromBlob(ev.Get("data")),
		}, handle)
	case js.Global.Get("ChangeEvent"):
		return eventx.NewBaseEvent(&eventx.ChangeEvent{
			Core:  ev,
			Value: ev.Get("target").Get("value").String(),
		}, handle)
	case js.Global.Get("ClipboardEvent"):
		return eventx.NewBaseEvent(&eventx.ClipboardEvent{
			Core: ev,
			Data: toDataTransfer(ev.Get("clipboardData")),
		}, handle)
	case js.Global.Get("CloseEvent"):
		return eventx.NewBaseEvent(&eventx.CloseEvent{
			Core:     ev,
			Code:     ev.Get("code").Int(),
			Reason:   ev.Get("reason").String(),
			WasClean: ev.Get("wasClean").Bool(),
		}, handle)
	case js.Global.Get("CompositionEvent"):
		return eventx.NewBaseEvent(&eventx.CompositionEvent{
			Core:   ev,
			Text:   ev.Get("text").String(),
			Data:   ev.Get("data").String(),
			Locale: ev.Get("locale").String(),
		}, handle)
	case js.Global.Get("CSSFontFaceLoadEvent"):
		return eventx.NewBaseEvent(&eventx.CSSFontFaceLoadEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("CustomEvent"):
		return eventx.NewBaseEvent(&eventx.CustomEvent{
			Core:   ev,
			Detail: ev.Get("detail").Interface(),
		}, handle)
	case js.Global.Get("DeviceLightEvent"):
		return eventx.NewBaseEvent(&eventx.DeviceLightEvent{
			Core:  ev,
			Value: ev.Get("value").Float(),
		}, handle)
	case js.Global.Get("DeviceMotionEvent"):
		return eventx.NewBaseEvent(&eventx.DeviceMotionEvent{
			Core:                         ev,
			Interval:                     ev.Get("interval").Float(),
			Acceleration:                 toMotionData(ev.Get("acceleration")),
			AccelerationIncludingGravity: toMotionData(ev.Get("accelerationIncludingGravity")),
			RotationRate:                 toRotationData(ev.Get("rotationRate")),
		}, handle)
	case js.Global.Get("DeviceOrientationEvent"):
		return eventx.NewBaseEvent(&eventx.DeviceOrientationEvent{
			Core:     ev,
			Absolute: ev.Get("absolute").Bool(),
			Alpha:    ev.Get("alpha").Float(),
			Beta:     ev.Get("beta").Float(),
			Gamma:    ev.Get("gamma").Float(),
		}, handle)
	case js.Global.Get("DeviceProximityEvent"):
		return eventx.NewBaseEvent(&eventx.DeviceProximityEvent{
			Core:  ev,
			Max:   ev.Get("max").Float(),
			Min:   ev.Get("min").Float(),
			Value: ev.Get("value").Float(),
		}, handle)
	case js.Global.Get("DOMTransactionEvent"):
		return eventx.NewBaseEvent(&eventx.DOMTransactionEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DragStartEvent"):
		return eventx.NewBaseEvent(&eventx.DragStartEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
		return eventx.NewBaseEvent(&eventx.DragExitEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
		return eventx.NewBaseEvent(&eventx.DragEndEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
		return eventx.NewBaseEvent(&eventx.DragEnterEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
		return eventx.NewBaseEvent(&eventx.DragLeaveEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
		return eventx.NewBaseEvent(&eventx.DragOverEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
		return eventx.NewBaseEvent(&eventx.DropEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
		return eventx.NewBaseEvent(&eventx.DragEvent{
			MouseEvent: &eventx.MouseEvent{
				UIEvent: &eventx.UIEvent{
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
		return eventx.NewBaseEvent(&eventx.EditingBeforeInputEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("ErrorEvent"):
		var err error
		if evx := ev.Get("error"); evx != nil && evx != js.Undefined {
			err = errors.New(ev.Get("error").String())
		}

		return eventx.NewBaseEvent(&eventx.ErrorEvent{
			Core:       ev,
			Message:    ev.Get("message").String(),
			Filename:   ev.Get("filename").String(),
			LineNumber: ev.Get("lineno").Int(),
			ColNumber:  ev.Get("colno").Int(),
			Error:      err,
		}, handle)
	case js.Global.Get("FocusEvent"):
		return eventx.NewBaseEvent(&eventx.FocusEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("GamepadEvent"):
		return eventx.NewBaseEvent(&eventx.GamepadEvent{
			Core:    ev,
			Gamepad: toGamepad(ev.Get("gamepad")),
		}, handle)
	case js.Global.Get("HashChangeEvent"):
		return eventx.NewBaseEvent(&eventx.HashChangeEvent{
			Core: ev,
			Old:  ev.Get("oldURL").String(),
			New:  ev.Get("newURL").String(),
		}, handle)
	case js.Global.Get("IDBVersionChangeEvent"):
		return eventx.NewBaseEvent(&eventx.IDBVersionChangeEvent{
			Core:       ev,
			OldVersion: ev.Get("oldVersion").Int64(),
			NewVersion: ev.Get("newVersion").Int64(),
		}, handle)
	case js.Global.Get("KeyboardEvent"):
		return eventx.NewBaseEvent(&eventx.KeyboardEvent{
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
			KeyLocation:   eventx.KeyLocation(ev.Get("KeyLocation").Uint64()),
			KeyCode:       eventx.KeyCode(ev.Get("keyCode").Uint64()),
		}, handle)
	case js.Global.Get("MediaStreamEvent"):
		return eventx.NewBaseEvent(&eventx.MediaStreamEvent{
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

		return eventx.NewBaseEvent(&eventx.MessageEvent{
			Core:   ev,
			Data:   data,
			Origin: ev.Get("origin").String(),
			Source: ev.Get("source").String(),
			Port:   ev.Get("port").Int(),
		}, handle)
	case js.Global.Get("MouseEvent"):
		return eventx.NewBaseEvent(&eventx.MouseEvent{
			UIEvent: &eventx.UIEvent{
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
		return eventx.NewBaseEvent(&eventx.MutationEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("OfflineAudioCompletionEvent"):
		return eventx.NewBaseEvent(&eventx.OfflineAudioCompletionEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("PageTransitionEvent"):
		return eventx.NewBaseEvent(&eventx.PageTransitionEvent{
			Core:      ev,
			Persisted: ev.Get("persisted").Bool(),
		}, handle)
	case js.Global.Get("PointerEvent"):
		return eventx.NewBaseEvent(&eventx.PointerEvent{
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

		return eventx.NewBaseEvent(&eventx.FetchEvent{
			Core:     ev,
			IsReload: ev.Get("isReload").Bool(),
			ClientID: ev.Get("clientId").String(),
			Request:  req,
		}, handle)
	case js.Global.Get("PopStateEvent"):
		return eventx.NewBaseEvent(&eventx.PopStateEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("ProgressEvent"):
		return eventx.NewBaseEvent(&eventx.ProgressEvent{
			Core:             ev,
			LengthComputable: ev.Get("lengthComputable").Bool(),
			Loaded:           ev.Get("loaded").Uint64(),
			Total:            ev.Get("total").Int(),
		}, handle)
	case js.Global.Get("RelatedEvent"):
		return eventx.NewBaseEvent(&eventx.RelatedEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("RTCPeerConnectionIceEvent"):
		return eventx.NewBaseEvent(&eventx.RTCPeerConnectionIceEvent{
			Core:      ev,
			Candidate: ev.Get("candidate").String(),
		}, handle)
	case js.Global.Get("SensorEvent"):
		return eventx.NewBaseEvent(&eventx.SensorEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("StorageEvent"):
		return eventx.NewBaseEvent(&eventx.StorageEvent{
			Core:        ev,
			Key:         ev.Get("key").String(),
			NewValue:    ev.Get("newValue").String(),
			OldValue:    ev.Get("oldValue").String(),
			URL:         ev.Get("url").String(),
			StorageArea: ev.Get("storageArea").Interface(),
		}, handle)
	case js.Global.Get("SVGEvent"):
		return eventx.NewBaseEvent(&eventx.SVGEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("SVGZoomEvent"):
		return eventx.NewBaseEvent(&eventx.SVGZoomEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("TimeEvent"):
		return eventx.NewBaseEvent(&eventx.TimeEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("TouchEvent"):
		return eventx.NewBaseEvent(&eventx.TouchEvent{
			Core:          ev,
			AltKey:        ev.Get("altKey").Bool(),
			CtrlKey:       ev.Get("ctrlKey").Bool(),
			MetaKey:       ev.Get("metaKey").Bool(),
			ShiftKey:      ev.Get("shiftKey").Bool(),
			TargetTouches: toTouches(ev.Get("touches")),
			Touches:       toTouches(ev.Get("targetTouches")),
		}, handle)
	case js.Global.Get("TrackEvent"):
		return eventx.NewBaseEvent(&eventx.TrackEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("TransitionEvent"):
		return eventx.NewBaseEvent(&eventx.TransitionEvent{
			Core:          ev,
			ElapsedTime:   ev.Get("elapsedTime").Float(),
			PropertyName:  ev.Get("propertyName").String(),
			PseudoElement: ev.Get("pseudoElement").String(),
		}, handle)
	case js.Global.Get("UIEvent"):
		return eventx.NewBaseEvent(&eventx.UIEvent{
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
		return eventx.NewBaseEvent(&eventx.UserProximityEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("WheelEvent"):
		return eventx.NewBaseEvent(&eventx.WheelEvent{
			Core:      ev,
			DeltaX:    ev.Get("deltaX").Float(),
			DeltaY:    ev.Get("deltaX").Float(),
			DeltaZ:    ev.Get("deltaX").Float(),
			DeltaMode: eventx.DeltaMode(ev.Get("deltaMode").Uint64()),
		}, handle)
	}

	return eventx.NewBaseEvent(ev, handle)
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

// toInputSourceCapability returns the eventx.InputDeviceCapabilities from the js.object.
func toInputSourceCapability(o *js.Object) *eventx.InputDeviceCapabilities {
	if o == nil || o == js.Undefined {
		return nil
	}

	return &eventx.InputDeviceCapabilities{
		FiresTouchEvent: o.Get("firesTouchEvent").Bool(),
	}
}

// toMotionData returns a motionData object from the js.Object.
func toMotionData(o *js.Object) eventx.MotionData {
	var md eventx.MotionData

	if o == nil || o == js.Undefined {
		return md
	}

	md.X = o.Get("x").Float()
	md.Y = o.Get("y").Float()
	md.Z = o.Get("z").Float()
	return md
}

// toRotationData returns a RotationData object from the js.Object.
func toRotationData(o *js.Object) eventx.RotationData {
	var md eventx.RotationData
	if o == nil || o == js.Undefined {
		return md
	}
	md.Alpha = o.Get("alpha").Float()
	md.Beta = o.Get("beta").Float()
	md.Gamma = o.Get("gamma").Float()
	return md
}

// toMediaStream returns a eventx.MediaStream object.
func toMediaStream(o *js.Object) eventx.MediaStream {
	var stream eventx.MediaStream
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

			stream.Audios = append(stream.Audios, eventx.MediaStreamTrack{
				Core:       track,
				Enabled:    track.Get("enabled").Bool(),
				ID:         track.Get("id").String(),
				Kind:       track.Get("kind").String(),
				Label:      track.Get("label").String(),
				Muted:      track.Get("muted").Bool(),
				ReadyState: track.Get("readyState").Bool(),
				Remote:     track.Get("remote").Bool(),
				AudioSettings: &eventx.MediaAudioTrackSettings{
					ChannelCount:     settings.Get("channelCount").Int(),
					EchoCancellation: settings.Get("echoCancellation").Bool(),
					Latency:          settings.Get("latency").Float(),
					SampleRate:       settings.Get("sampleRate").Int64(),
					SampleSize:       settings.Get("sampleSize").Int64(),
					Volume:           settings.Get("volume").Float(),
					MediaTrackSettings: eventx.MediaTrackSettings{
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

			stream.Videos = append(stream.Videos, eventx.MediaStreamTrack{
				Core:       track,
				Enabled:    track.Get("enabled").Bool(),
				ID:         track.Get("id").String(),
				Kind:       track.Get("kind").String(),
				Label:      track.Get("label").String(),
				Muted:      track.Get("muted").Bool(),
				ReadyState: track.Get("readyState").Bool(),
				Remote:     track.Get("remote").Bool(),
				VideoSettings: &eventx.MediaVideoTrackSettings{
					AspectRatio: settings.Get("aspectRation").Float(),
					FrameRate:   settings.Get("frameRate").Float(),
					Height:      settings.Get("height").Int64(),
					Width:       settings.Get("width").Int64(),
					FacingMode:  settings.Get("facingMode").String(),
					MediaTrackSettings: eventx.MediaTrackSettings{
						DeviceID: settings.Get("deviceId").String(),
						GroupID:  settings.Get("groupId").String(),
					},
				},
			})
		}
	}

	return stream
}

func toTouches(o *js.Object) eventx.TouchList {
	var th eventx.TouchList

	if o == nil || o == js.Undefined {
		return th
	}

	var touches []eventx.Touch

	for i := 0; i < o.Length(); i++ {
		ev := o.Index(i)
		touches = append(touches, eventx.Touch{
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
func toGamepad(o *js.Object) eventx.Gamepad {
	var pad eventx.Gamepad
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
			pad.Buttons = append(pad.Buttons, eventx.Button{
				Value:   button.Get("value").Float(),
				Pressed: button.Get("pressed").Bool(),
			})
		}
	}

	return pad
}

// toDataTransfer returns a transfer object from the js.Object.
func toDataTransfer(o *js.Object) eventx.DataTransfer {
	var dt eventx.DataTransfer
	if o == nil || o == js.Undefined {
		return dt
	}
	dt.DropEffect = o.Get("dropEffect").String()
	dt.EffectAllowed = o.Get("effectAllowed").String()

	types := o.Get("types")
	if types != nil && types != js.Undefined {
		dt.Types = shell.ObjectToStringList(types)
	}

	var dItems []eventx.DataTransferItem

	items := o.Get("items")
	if items != nil && items != js.Undefined {
		for i := 0; i < items.Length(); i++ {
			item := items.Index(i)
			dItems = append(dItems, eventx.DataTransferItem{
				Name: item.Get("name").String(),
				Size: item.Get("size").Int(),
				Data: fromFile(item),
			})
		}
	}

	var dFiles []eventx.DataTransferItem

	files := o.Get("files")
	if files != nil && files != js.Undefined {
		for i := 0; i < files.Length(); i++ {
			item := files.Index(i)
			dFiles = append(dFiles, eventx.DataTransferItem{
				Name: item.Get("name").String(),
				Size: item.Get("size").Int(),
				Data: fromFile(item),
			})
		}
	}

	dt.Items = eventx.DataTransferItemList{Items: dItems}
	dt.Files = dFiles
	return dt
}

// prepareEventName lowers the event name into lowercase and removes all non string
// characters.
func prepareEventName(name string) string {
	name = strings.ToLower(name)
	name = strings.Replace(name, "-", "", -1)
	name = strings.Replace(name, "_", "", -1)
	return name
}
