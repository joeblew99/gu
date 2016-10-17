package trees

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/influx6/gu/events"
)

// FinalizeHandle defines a function type which has the root and item concerned.
type FinalizeHandle func(root, item Markup)

// HTML defines an interface with a single method that returns a html string of
// its content.
type HTML interface {
	HTML() string
	EHTML() template.HTML
}

// Routes defines an interface which exposes the internal Routing struct used
// by a Markup.

// Markup provide a basic specification type of how a element resolves
// its content.
type Markup interface {
	HTML
	Children
	Identity
	Morphers
	Appliable
	Removable
	Routes
	Properties
	Reconcilable
}

// Element represent a concrete implementation of a element node.
type Element struct {
	*Routing
	removed         bool
	autoclose       bool
	allowEvents     bool
	allowChildren   bool
	allowStyles     bool
	allowAttributes bool

	uid         string
	hash        string
	tagname     string
	textContent string

	events         []Event
	children       []Markup
	styles         []Property
	attrs          []Property
	morphers       []Morpher
	finalizers     []FinalizeHandle
	onceFinalizers []FinalizeHandle
	eventManager   events.EventManagers
}

// NewText returns a new Text instance element
func NewText(txt string) *Element {
	em := NewElement("text", false)
	em.allowChildren = false
	em.allowAttributes = false
	em.allowStyles = false
	em.allowEvents = false
	em.textContent = txt
	return em
}

// NewElement returns a new element instance giving the specificed name
func NewElement(tag string, hasNoEndingTag bool) *Element {
	return &Element{
		Routing:         NewRouting("*", &RemoveMorpher{}),
		allowChildren:   true,
		allowStyles:     true,
		allowAttributes: true,
		allowEvents:     true,
		uid:             RandString(8),
		hash:            RandString(10),
		tagname:         strings.ToLower(strings.TrimSpace(tag)),
		autoclose:       hasNoEndingTag,
	}
}

// EHTML returns the html string wrapped by a template.HTML type to avoid getting
// escaped by go templates. The returned html is rendered using the default
// SimpleElementWriter and represents the DOM of the giving element.
func (e *Element) EHTML() template.HTML {
	return template.HTML(SimpleElementWriter.Print(e))
}

// HTML returns the html string representing the DOM of the giving element.
//The returned html is rendered using the default SimpleElementWriter.
func (e *Element) HTML() string {
	return SimpleElementWriter.Print(e)
}

// AutoClosed returns true/false if this element uses a </> or a <></> tag convention
func (e *Element) AutoClosed() bool {
	return e.autoclose
}

// Empty resets the elements children list as 0 length
func (e *Element) Empty() {
	e.children = e.children[:0]
}

//==============================================================================

// Routes defines an interface which allows the access and setting of the
// routes which affects the giving markup.
type Routes interface {
	Router() *Routing
	UseRouter(*Routing)
}

// Router returns the internal router used by the element.
func (e *Element) Router() *Routing {
	return e.Routing
}

// UseRouter changes the internal router to be used by this element alone.
// Connecting all internal children routers to the supplied and discarding
// its previous router.
func (e *Element) UseRouter(rm *Routing) {
	old := e.Routing
	old.Flush()

	for index, morpher := range e.morphers {
		if morpher == old {
			e.morphers = append(e.morphers[:index], e.morphers[index+1:]...)
			break
		}
	}

	e.Routing = rm
	for _, child := range e.children {
		e.Routing.Register(child.Router())
	}
}

//==============================================================================

// Properties interface defines a type that has Attributes
type Properties interface {
	AddAttribute(Property)
	Attributes() []Property

	AddStyle(Property)
	Styles() []Property

	AddEvent(Event)
	Events() []Event
}

// AddEvent adds an event into the event list for this element.
func (e *Element) AddEvent(ev Event) {
	e.events = append(e.events, ev)
}

// Events return the elements events
func (e *Element) Events() []Event {
	return e.events
}

// Styles return the internal style list of the element
func (e *Element) Styles() []Property {
	return e.styles
}

// AddStyle adds a property to the style property list.
func (e *Element) AddStyle(p Property) {
	e.styles = append(e.styles, p)
}

// Attributes return the internal attribute list of the element
func (e *Element) Attributes() []Property {
	return e.attrs
}

// AddAttribute adds a property to the attribute property list.
func (e *Element) AddAttribute(p Property) {
	e.attrs = append(e.attrs, p)
}

//==============================================================================

// Eventers provide an interface type for elements able to register and load
// event managers.
type Eventers interface {
	LoadEvents()
	UseEventManager(events.EventManagers) bool
}

// UseEventManager adds a eventmanager into the markup and if not available before automatically registers
// the events with it,once an event manager is registered to it,it will and can not be changed
func (e *Element) UseEventManager(man events.EventManagers) bool {
	if e.eventManager != nil {
		man.AttachManager(e.eventManager)
		return false
	}

	e.eventManager = man
	e.LoadEvents()
	return true
}

// LoadEvents loads up the events registered by this and by its children into each respective
// available events managers
func (e *Element) LoadEvents() {
	if e.eventManager != nil {
		e.eventManager.DisconnectRemoved()

		for _, ev := range e.events {
			if es, _ := e.eventManager.NewEventMeta(ev.Meta()); es != nil {
				es.Q(ev.Fire)
			}
		}

	}

	//load up the children events also
	for _, em := range e.children {
		if ems, ok := em.(Eventers); ok {
			if !ems.UseEventManager(e.eventManager) {
				ems.LoadEvents()
			}
		}
	}
}

//==============================================================================

// MarkupState defines a markup which reveals internal state of a Markup.
type MarkupState interface {
	AllowStyles() bool
	AllowAttributes() bool
	AllowChildren() bool
	AllowEvents() bool
}

// AllowEvents returns true/false if the element allows events.
func (e *Element) AllowEvents() bool {
	return e.allowEvents
}

// AllowChildren returns true/false if the element allows attributes.
func (e *Element) AllowChildren() bool {
	return e.allowChildren
}

// AllowAttributes returns true/false if the element allows attributes.
func (e *Element) AllowAttributes() bool {
	return e.allowAttributes
}

// AllowStyles returns true/false if the element allows styles.
func (e *Element) AllowStyles() bool {
	return e.allowStyles
}

//==============================================================================

// Identity defines an interface for identifiable structures.
type Identity interface {
	Name() string
	UID() string
	Hash() string
	EventID() string
}

// EventID returns the selector used for tagging events for a markup.
func (e *Element) EventID() string {
	return fmt.Sprintf("%s[uid='%s']", strings.ToLower(e.Name()), e.UID())
}

// Name returns the tag name of the element
func (e *Element) Name() string {
	return e.tagname
}

// UID returns the current uid of the Element
func (e *Element) UID() string {
	return e.uid
}

// Hash returns the current hash of the Element
func (e *Element) Hash() string {
	return e.hash
}

//==============================================================================

// Morphers exposes a method to allow adding morphers.
type Morphers interface {
	AddMorpher(...Morpher)
	ApplyMorphers() Markup
}

// AddMorpher adds the provided morphers into the elements lists.
func (e *Element) AddMorpher(m ...Morpher) {
	e.morphers = append(e.morphers, m...)
}

// ApplyMorphers calls all elemental morphers sequentially applying them to the
// element and passes the result as the input of the next morpher unless. If
// any morpher returns nil, then the element is reused again until all morphers
// are called.
func (e *Element) ApplyMorphers() Markup {
	var base Markup

	for index, child := range e.children {
		e.children[index] = child.ApplyMorphers()
	}

	for _, morpher := range e.morphers {
		if base == nil {
			base = morpher.Morph(e)
			continue
		}

		base = morpher.Morph(base)
	}

	return base
}

//==============================================================================

// TextMarkup defines a interface for text based markup.
type TextMarkup interface {
	TextContent() string
}

// TextContent returns the elements text value if its a text
// type else an empty string.
func (e *Element) TextContent() string {
	return e.textContent
}

//==============================================================================

// Cleanable defines a interface for structures to self sanitize their contents.
type Cleanable interface {
	Clean()
}

// Clean cleans out all internal markup marked as removable.
func (e *Element) Clean() {
	for n, elm := range e.children {
		if emr, ok := elm.(Removable); ok {
			if emr.Removed() {
				copy(e.children[n:], e.children[n+1:])
				e.children = e.children[:len(e.children)-1]
			} else {
				if em, ok := elm.(Cleanable); ok {
					em.Clean()
				}
			}
		}
	}
}

//==============================================================================

// Removable defines a self removal type structure.
type Removable interface {
	Remove()
	UnRemove()
	Removed() bool
}

// Remove sets the markup as removable and adds a 'NodeRemoved' attribute to it.
func (e *Element) Remove() {
	if !e.Removed() {
		e.attrs = append(e.attrs, &Attribute{"NodeRemoved", ""})
		e.removed = true
	}
}

// UnRemove sets the markup as not to be removable.
func (e *Element) UnRemove() {
	if !e.Removed() {
		return
	}

	e.removed = false

	for index, attr := range e.attrs {
		if name, _ := attr.Render(); name != "NodeRemoved" {
			continue
		}

		e.attrs = append(e.attrs[:index], e.attrs[1+index:]...)
		return
	}
}

// Removed returns true/false if the Element is marked removed
func (e *Element) Removed() bool {
	return !!e.removed
}

//==============================================================================

// SwappableIdentity defines an interface that allows swapping a structures
// identity information.
type SwappableIdentity interface {
	SwapHash(string)
	SwapUID(string)
	UpdateHash()
}

// SwapUID swaps the uid of the internal Element.
func (e *Element) SwapUID(uid string) {
	e.uid = uid
}

// SwapHash swaps the hash of the internal Element.
func (e *Element) SwapHash(hash string) {
	e.hash = hash
}

// UpdateHash updates the Element hash value
func (e *Element) UpdateHash() {
	e.hash = RandString(10)
}

//==============================================================================

// ElementalMarkup defines a markup for elemental structures, which provide the
// concrete structure for dom nodes.
type ElementalMarkup interface {
	Markup
	TextMarkup
	Cleanable

	Clonable
	Eventers

	SwappableIdentity

	AutoClosed() bool

	Empty()
}

// Reconcilable defines the interface of markups that can reconcile their content against another
type Reconcilable interface {
	Reconcile(Markup) bool
}

// Reconcile takes a old markup and reconciles its uid and its children with
// these information,it returns a true/false telling the parent if the children
// swapped hashes.
// The reconcilation uses the order in which elements are added, if the order
// and element types are same then the uid are swapped, else it firsts checks the
// element type, but if not the same adds the old one into the new list as removed
// then continues the check. The system takes position of elements in the old and
// new as very important and I cant stress this enough, "Element Positioning" in
// the markup are very important, If a Anchor was the first element in the old
// render and the next pass returns a Div in the position for that Anchor in the
// new render, the old Anchor will be marked as removed and will be removed from
// the dom and ignored by the writers.
// When two elements position are same and their types are the same then a checkup
// process is doing using the elements attributes, this is done to determine if the
// hash value of the new should be swapped with the old. We cant use style properties
// here because they are the most volatile of the set and will periodically be
// either changed and returned to normal values eg display: none to display: block
// and vise-versa, so only attributes are used in the check process.
func (e *Element) Reconcile(m Markup) bool {
	em, ok := m.(ElementalMarkup)
	if !ok {
		return false
	}

	// are we reconciling the proper elements type ? if not skip (i.e different types cant reconcile eachother)]
	// TODO: decide if we should mark the markup as removed in this case as a catchall system
	if e.Name() != em.Name() {
		return false
	}

	em.Clean()

	//since the tagname are the same, swap uids
	// olduid := em.UID()
	e.SwapUID(em.UID())

	//since the tagname are the same and we have swapped uid, to determine who gets or keeps
	// its hash we will check the attributes against each other, but also the hash is dependent on the
	// children also, if the children observered there was a change
	oldHash := em.Hash()

	// if we have a special case for text element then we do things differently
	if e.Name() == "text" {

		//if the contents are equal,keep the prev hash
		if e.TextContent() == em.TextContent() {
			e.SwapHash(oldHash)
			return false
		}
		return true
	}

	newChildren := e.Children()
	oldChildren := em.Children()
	maxSize := len(newChildren)
	oldMaxSize := len(oldChildren)

	attrChanged := EqualAttributes(e, em)
	styleChanged := EqualStyles(e, em)

	// if the element had no children too, swap hash.
	if maxSize < 1 {
		if oldMaxSize > 1 {
			return true
		}

		if !attrChanged || !styleChanged {
			return true
		}

		e.SwapHash(oldHash)
		return false
	}

	var childChanged bool

	for n, och := range oldChildren {
		if maxSize > n {

			nch := newChildren[n]
			if nch.Name() != och.Name() {
				och.Remove()
				e.AddChild(och)
				childChanged = true
				continue
			}

			if nch.Reconcile(och) {
				childChanged = true
			}

			continue
		}

		och.Remove()
		e.AddChild(och)
		childChanged = true
	}

	ReconcileEvents(e, em)

	if e.eventManager != nil {
		e.eventManager.DisconnectRemoved()
	}

	if !childChanged && attrChanged && styleChanged {
		e.SwapHash(oldHash)
		return false
	}

	return true
}

//==============================================================================

// ChildApplier defines a package level handle that exposes addition of
// children Markup into a Element children list.
var ChildApplier childAdd

type childAdd struct{}

// Apply adds the giving markup children into the root.
func (childAdd) Apply(root Markup, child ...Markup) {
	if rs, ok := root.(MarkupState); ok {
		if !rs.AllowChildren() {
			return
		}

		rootEm := root.(*Element)

		for _, ch := range child {
			rootEm.children = append(rootEm.children, ch)
			root.Router().Register(ch.Router())

			if chm, ok := ch.(Eventers); ok {
				chm.UseEventManager(rootEm.eventManager)
			}
		}
	}
}

// Children defines the interface of an element that has children
type Children interface {
	AddChild(...Markup)
	Children() []Markup
}

// AddChild adds a new markup as the children of this element
func (e *Element) AddChild(em ...Markup) {
	ChildApplier.Apply(e, em...)
}

// Children returns the children list for the element
func (e *Element) Children() []Markup {
	return e.children
}

//==============================================================================

// Appliable define the interface specification for applying changes to elements elements in tree
type Appliable interface {
	Apply(Markup)
}

//Apply adds the giving element into the current elements children tree
func (e *Element) Apply(em Markup) {
	em.AddChild(e)
}

//==============================================================================

// Clonable defines an interface for objects that can be cloned
type Clonable interface {
	Clone() Markup
}

// Clone makes a new copy of the markup structure
func (e *Element) Clone() Markup {
	co := NewElement(e.Name(), e.AutoClosed())

	//copy over the textContent
	co.textContent = e.textContent

	//copy over the attribute lockers
	co.allowChildren = e.allowChildren
	co.allowEvents = e.allowEvents
	co.allowAttributes = e.allowAttributes
	co.eventManager = e.eventManager

	if e.Removed() {
		co.Removed()
	}

	//clone the internal styles
	for _, so := range e.styles {
		so.Clone().Apply(co)
	}

	co.allowStyles = e.allowStyles

	//clone the internal attribute
	for _, ao := range e.attrs {
		ao.Clone().Apply(co)
	}

	// co.allowAttributes = e.allowAttributes
	//clone the internal children
	for _, ch := range e.children {
		if cl, ok := ch.(Clonable); ok {
			cl.Clone().Apply(co)
		}
	}

	for _, ch := range e.events {
		ch.Clone().Apply(co)
	}

	co.morphers = append(co.morphers, e.morphers...)

	return co
}
