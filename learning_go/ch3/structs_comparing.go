package main

func main() {
	// Comparing structs is possible if they are entirely composed of comparable types - slices or map fields = not comparable
	// Go does not allow for comparison between variables that represent structs of different types - but struct conversion CAN occur...
	type firstPerson struct {
		name string
		age  int
	}

	// It would be possible to take an instance of firstPerson, convert to secondPerson and compare the conversion to a secondPerson instance
	type secondPerson struct {
		name string
		age  int
	}

	// It would NOT be possible to covert here as the order of the fields has changed
	type thirdPerson struct {
		age  int
		name string
	}

	// This cannot be converted to as the field for 'name' is different - the fields do not match
	type fourthPerson struct {
		firstName string
		age       int
	}

	// This cannot be converted to as there is an additional field in the struct
	type fifthPerson struct {
		name           string
		age            int
		favouriteColor string
	}
}
