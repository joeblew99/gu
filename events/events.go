package events

import (
	"github.com/gu-io/gu/notifications/mque"
)

// DeltaMode is an indication of the units of measurement for a delta value.
type DeltaMode uint64

// KeyCode represents a system and implementation dependent numerical
// code identifying the unmodified value of the pressed key.
type KeyCode uint8

// KeyLocation represents the location of the key on the keyboard or
// other input device.
type KeyLocation uint8

// Keyboard keys.
const (
	KeyBackspace          KeyCode = 8
	KeyTab                KeyCode = 9
	KeyEnter              KeyCode = 13
	KeyShift              KeyCode = 16
	KeyCtrl               KeyCode = 17
	KeyAlt                KeyCode = 18
	KeyPauseBreak         KeyCode = 19
	KeyCapsLock           KeyCode = 20
	KeyEsc                KeyCode = 27
	KeySpace              KeyCode = 32
	KeyPageUp             KeyCode = 33
	KeyPageDown           KeyCode = 34
	KeyEnd                KeyCode = 35
	KeyHome               KeyCode = 36
	KeyLeft               KeyCode = 37
	KeyUp                 KeyCode = 38
	KeyRight              KeyCode = 39
	KeyDown               KeyCode = 40
	KeyPrintScreen        KeyCode = 44
	KeyInsert             KeyCode = 45
	KeyDelete             KeyCode = 46
	Key0                  KeyCode = 48
	Key1                  KeyCode = 49
	Key2                  KeyCode = 50
	Key3                  KeyCode = 51
	Key4                  KeyCode = 52
	Key5                  KeyCode = 53
	Key6                  KeyCode = 54
	Key7                  KeyCode = 55
	Key8                  KeyCode = 56
	Key9                  KeyCode = 57
	KeyA                  KeyCode = 65
	KeyB                  KeyCode = 66
	KeyC                  KeyCode = 67
	KeyD                  KeyCode = 68
	KeyE                  KeyCode = 69
	KeyF                  KeyCode = 70
	KeyG                  KeyCode = 71
	KeyH                  KeyCode = 72
	KeyI                  KeyCode = 73
	KeyJ                  KeyCode = 74
	KeyK                  KeyCode = 75
	KeyL                  KeyCode = 76
	KeyM                  KeyCode = 77
	KeyN                  KeyCode = 78
	KeyO                  KeyCode = 79
	KeyP                  KeyCode = 80
	KeyQ                  KeyCode = 81
	KeyR                  KeyCode = 82
	KeyS                  KeyCode = 83
	KeyT                  KeyCode = 84
	KeyU                  KeyCode = 85
	KeyV                  KeyCode = 86
	KeyW                  KeyCode = 87
	KeyX                  KeyCode = 88
	KeyY                  KeyCode = 88
	KeyZ                  KeyCode = 90
	KeyMeta               KeyCode = 91
	KeyMenu               KeyCode = 93
	KeyNumPad0            KeyCode = 96
	KeyNumPad1            KeyCode = 97
	KeyNumPad2            KeyCode = 98
	KeyNumPad3            KeyCode = 99
	KeyNumPad4            KeyCode = 100
	KeyNumPad5            KeyCode = 101
	KeyNumPad6            KeyCode = 102
	KeyNumPad7            KeyCode = 103
	KeyNumPad8            KeyCode = 104
	KeyNumPad9            KeyCode = 105
	KeyNumPadMult         KeyCode = 106
	KeyNumPadPlus         KeyCode = 107
	KeyNumPadMin          KeyCode = 109
	KeyNumPadDot          KeyCode = 110
	KeyNumPadDecimal      KeyCode = 111
	KeyF1                 KeyCode = 112
	KeyF2                 KeyCode = 113
	KeyF3                 KeyCode = 114
	KeyF4                 KeyCode = 115
	KeyF5                 KeyCode = 116
	KeyF6                 KeyCode = 117
	KeyF7                 KeyCode = 118
	KeyF8                 KeyCode = 119
	KeyF9                 KeyCode = 120
	KeyF10                KeyCode = 121
	KeyF11                KeyCode = 122
	KeyF12                KeyCode = 123
	KeyNumLock            KeyCode = 144
	KeyMute               KeyCode = 173
	KeyVolumeDown         KeyCode = 174
	KeyVolumeUp           KeyCode = 175
	KeySemicolon          KeyCode = 186
	KeyEqual              KeyCode = 187
	KeyComa               KeyCode = 188
	KeyDash               KeyCode = 189
	KeyDot                KeyCode = 190
	KeySlash              KeyCode = 191
	KeyBackquote          KeyCode = 192
	KeySquareBracketLeft  KeyCode = 219
	KeyBackslash          KeyCode = 220
	KeySquareBracketRight KeyCode = 221
	KeyQuote              KeyCode = 222
)

// Keyboard locations constants.
const (
	KeyLocationStandard KeyLocation = iota
	KeyLocationLeft
	KeyLocationRight
	KeyLocationNumpad
)

const (
	// DeltaPixel defines the mouse movement state.
	DeltaPixel = 0

	// DeltaLine defines the mouse movement state.
	DeltaLine = 1

	// DeltaPage defines the mouse movement state.
	DeltaPage = 2
)

// UserProximityEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type UserProximityEvent struct {
	Core interface{} `json:"-"`
}

// UIEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type UIEvent struct {
	Core interface{} `json:"-"`
}

// TrackEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type TrackEvent struct {
	Core interface{} `json:"-"`
}

// TouchEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type TouchEvent struct {
	Core interface{} `json:"-"`
}

// SVGZoomEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type SVGZoomEvent struct {
	Core interface{} `json:"-"`
}

// SVGEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type SVGEvent struct {
	Core interface{} `json:"-"`
}

// MessageEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type MessageEvent struct {
	Core interface{} `json:"-"`
	Data string
}

// MediaStreamEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type MediaStreamEvent struct {
	Core interface{} `json:"-"`
}

// IDBVersionChangeEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type IDBVersionChangeEvent struct {
	Core interface{} `json:"-"`
}

// HashChangeEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type HashChangeEvent struct {
	Core interface{} `json:"-"`
}

// GamepadEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type GamepadEvent struct {
	Core interface{} `json:"-"`
}

// AnimationEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type AnimationEvent struct {
	Core interface{} `json:"-"`
}

// AudioProcessingEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type AudioProcessingEvent struct {
	Core interface{} `json:"-"`
}

// BeforeUnloadEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type BeforeUnloadEvent struct {
	Core interface{} `json:"-"`
}

// BeforeInputEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type BeforeInputEvent struct {
	Core interface{} `json:"-"`
}

// BlobEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type BlobEvent struct {
	Core interface{} `json:"-"`
}

// ClipboardEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type ClipboardEvent struct {
	Core interface{} `json:"-"`
}

// CloseEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type CloseEvent struct {
	Core     interface{} `json:"-"`
	Code     int
	Reason   string
	WasClean bool
}

// CompositionEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type CompositionEvent struct {
	Core interface{} `json:"-"`
}

// CSSFontFaceLoadEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type CSSFontFaceLoadEvent struct {
	Core interface{} `json:"-"`
}

// CustomEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type CustomEvent struct {
	Core interface{} `json:"-"`
}

// DeviceLightEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type DeviceLightEvent struct {
	Core interface{} `json:"-"`
}

// DeviceMotionEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type DeviceMotionEvent struct {
	Core interface{} `json:"-"`
}

// DeviceOrientationEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type DeviceOrientationEvent struct {
	Core interface{} `json:"-"`
}

// DeviceProximityEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type DeviceProximityEvent struct {
	Core interface{} `json:"-"`
}

// DOMTransactionEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type DOMTransactionEvent struct {
	Core interface{} `json:"-"`
}

// DragEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type DragEvent struct {
	Core interface{} `json:"-"`
}

// EditingBeforeInputEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type EditingBeforeInputEvent struct {
	Core interface{} `json:"-"`
}

// ErrorEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type ErrorEvent struct {
	Core interface{} `json:"-"`
}

// FocusEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type FocusEvent struct {
	Core interface{} `json:"-"`
}

// MutationEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type MutationEvent struct {
	Core interface{} `json:"-"`
}

// OfflineAudioCompletionEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type OfflineAudioCompletionEvent struct {
	Core interface{} `json:"-"`
}

// PageTransitionEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type PageTransitionEvent struct {
	Core interface{} `json:"-"`
}

// PointerEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type PointerEvent struct {
	Core interface{} `json:"-"`
}

// PopStateEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type PopStateEvent struct {
	Core interface{} `json:"-"`
}

// ProgressEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type ProgressEvent struct {
	Core interface{} `json:"-"`
}

// RelatedEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type RelatedEvent struct {
	Core interface{} `json:"-"`
}

// RTCPeerConnectionIceEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type RTCPeerConnectionIceEvent struct {
	Core interface{} `json:"-"`
}

// SensorEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type SensorEvent struct {
	Core interface{} `json:"-"`
}

// StorageEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type StorageEvent struct {
	Core interface{} `json:"-"`
}

// TimeEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type TimeEvent struct {
	Core interface{} `json:"-"`
}

// TransitionEvent defines a struct to contain the values of a promiximity
// event fired from a giving DOM.
type TransitionEvent struct {
	Core interface{} `json:"-"`
}

// ChangeEvent represents the data passed in a onchange event.
type ChangeEvent struct {
	Core  interface{} `json:"-"`
	Value string      `json:"-"`
}

// MouseEvent represents data fired when interacting
// with a pointing device (such as a mouse).
type MouseEvent struct {
	*UIEvent `json:"uievent"`
	Core     interface{} `json:"-"`
	ClientX  float64
	ClientY  float64
	PageX    float64
	PageY    float64
	ScreenX  float64
	ScreenY  float64
	MovemenX float64
	MovemenY float64
	Button   int
	Detail   int
	AltKey   bool
	CtrlKey  bool
	MetaKey  bool
	ShiftKey bool
}

// WheelEvent represents data fired when a wheel button of a
// pointing device (usually a mouse) is rotated.
type WheelEvent struct {
	Core      interface{} `json:"-"`
	DeltaX    float64
	DeltaY    float64
	DeltaZ    float64
	DeltaMode DeltaMode
}

// KeyboardEvent represents data fired when the keyboard is used.
type KeyboardEvent struct {
	Core          interface{} `json:"-"`
	CharCode      rune
	KeyCode       KeyCode
	Location      KeyLocation
	AltKey        bool
	CtrlKey       bool
	MetaKey       bool
	ShiftKey      bool
	Key           string
	KeyIdentifier string
	Locale        string
	Repeat        bool
	ModifiedState bool
}

// BasicEventMap defines a event type which defines a event type which is not
// supported by this package.
type BasicEventMap map[string]string

// BaseEvent defines the standard event returned to wrap a core event.
type BaseEvent struct {
	Core   interface{} `json:"-"`
	event  interface{}
	handle mque.End
}

// NewBaseEvent returns a new instance of the base evet.
func NewBaseEvent(event interface{}, handle mque.End) *BaseEvent {
	return &BaseEvent{
		event:  event,
		handle: handle,
	}
}

// RemoveEvent calls the finalizer of the giving event.
func (b *BaseEvent) RemoveEvent() {
	if b.handle != nil {
		b.handle.End()
	}
}

// Underlying returns the core event object returned by the
// base.
func (b *BaseEvent) Underlying() interface{} {
	return b.event
}

// GetEvent returns the base event type expected from the provided value.
func GetEvent() *BaseEvent {
	return nil
}
