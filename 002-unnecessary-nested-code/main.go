// Mistake #2: Unnecessary nested code
//
// Too many nested levels make code hard to read. Return early on errors
// and keep the happy path aligned to the left (no else after a return).
package main

func main() {
	mistake()
	fix()
}

func mistake() {
	if true {
		if true {
			if true {
				// do something
			} else {
				return
			}
		} else {
			return
		}
	} else {
		return
	}
}

func fix() {
	if !true {
		return
	}

	if !true {
		return
	}

	if !true {
		return
	}

	// do something
}
