package trees

var (
	dot        = byte('.')
	coma       = byte(',')
	hash       = byte('#')
	space      = byte(' ')
	bracket    = byte('[')
	endbracket = byte(']')
)

// Query defines a package level variable for access the query interface
// which handles running css queries on markup structures.
var Query queryCtrl

type queryCtrl struct{}

// Selector defines a structure which defines the requirements for a given
// matching to be processed.
type Selector struct {
	Tag      string
	Id       string
	Attr     string
	Classes  []string
	Children []*Selector
}

// ParseSelector returns the giving selector parsed out into its individual
// sections.
func (queryCtrl) ParseSelector(sel string) []*Selector {
	var sels []*Selector
	items := []byte(sel)
	itemsLen := len(items)

	var index int
	var block []byte
	var doChildren bool
	var seenSpace bool

	sels = append(sels, &Selector{})

parseLoop:
	for index < itemsLen {
		item := items[index]

		if seenSpace && item == space {
			index++
			continue
		}

		csel := sels[len(sels)-1]

		switch item {
		case space:
			if len(block) != 0 {
				if doChildren {
					child := csel.Children[len(csel.Children)-1]
					child.Tag = string(block)
					block = nil
				} else {
					csel.Tag = string(block)
					block = nil
				}
			}

			seenSpace = true
			index++
			continue parseLoop

		case coma:
			if len(block) != 0 {
				if doChildren {
					child := csel.Children[len(csel.Children)-1]
					child.Tag = string(block)
					block = nil
				} else {
					csel.Tag = string(block)
					block = nil
				}
			}

			seenSpace = false
			doChildren = false

			newSel := &Selector{}
			sels = append(sels, newSel)
			csel = newSel
			index++
			continue parseLoop

		case dot:
			var child *Selector
			if len(block) != 0 {
				if doChildren {
					child = csel.Children[len(csel.Children)-1]
					child.Tag = string(block)
					block = nil
				} else {
					csel.Tag = string(block)
					block = nil
				}
			}

			{
				var blk []byte

				for {
					index++
					if index >= itemsLen {
						if len(blk) != 0 {
							if doChildren {
								child.Classes = append(child.Classes, string(blk))
							} else {
								csel.Classes = append(csel.Classes, string(blk))
							}

							blk = nil
						}

						break
					}

					item = items[index]

					if item == dot {
						if doChildren {
							child.Classes = append(child.Classes, string(blk))
						} else {
							csel.Classes = append(csel.Classes, string(blk))
						}

						blk = nil
						continue
					}

					if item == space || item == coma {
						if doChildren {
							child.Classes = append(child.Classes, string(blk))
						} else {
							csel.Classes = append(csel.Classes, string(blk))
						}

						blk = nil
						continue parseLoop
					}

					blk = append(blk, item)
				}
			}

		case hash:
			var child *Selector
			if len(block) != 0 {
				if doChildren {
					child = csel.Children[len(csel.Children)-1]
					child.Tag = string(block)
					block = nil
				} else {
					csel.Tag = string(block)
					block = nil
				}
			}

			{
				var blk []byte
				for {
					index++

					if index >= itemsLen {
						if doChildren {
							child.Id = string(blk)
						} else {
							csel.Id = string(blk)
						}

						blk = nil
						break
					}

					item = items[index]
					if item == dot || item == space || item == coma {
						index--

						if doChildren {
							child.Id = string(blk)
						} else {
							csel.Id = string(blk)
						}

						blk = nil
						continue parseLoop
					}

					blk = append(blk, item)
				}
			}

		case bracket:
			var child *Selector
			if len(block) != 0 {
				if doChildren {
					child = csel.Children[len(csel.Children)-1]
					child.Tag = string(block)
					block = nil
				} else {
					csel.Tag = string(block)
					block = nil
				}
			}

			{
				var blk []byte

				for {
					index++
					if index >= itemsLen {
						if len(blk) != 0 {
							if doChildren {
								child.Attr = string(blk)
							} else {
								csel.Attr = string(blk)
							}
						}
						break
					}

					item = items[index]
					if item == endbracket {
						blk = append(blk, item)

						if doChildren {
							child.Attr = string(blk)
						} else {
							csel.Attr = string(blk)
						}

						continue parseLoop
					}

					blk = append(blk, item)
				}
			}

		default:
			block = append(block, item)
		}

		index++
	}

	return sels
}

// Query returns the first element matching the giving selector.
func (queryCtrl) Query(root *Markup, sel string) *Markup {
	var found *Markup

	return found
}

// QueryAll returns all elements matching the giving selector.
func (queryCtrl) QueryAll(root *Markup, sel string) []*Markup {
	var found []*Markup

	return found
}

func (queryCtrl) classFor(root *Markup, class string) *Markup {
	var found *Markup

	return found
}

func (queryCtrl) idFor(root *Markup, id string) *Markup {
	var found *Markup

	return found
}

func (queryCtrl) contentContainsFor(root *Markup, attr string, content string) *Markup {
	var found *Markup

	return found
}

func (queryCtrl) contentMatchFor(root *Markup, attr string, content string) *Markup {
	var found *Markup

	return found
}
