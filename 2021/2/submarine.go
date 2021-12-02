package day2

type SubmarinePosition struct {
	horizontal int
	depth int
	aim int
}

func (position SubmarinePosition) GoForward(amount int) SubmarinePosition {
	return SubmarinePosition{
		horizontal: position.horizontal + amount,
		depth: position.depth + (amount * position.aim),
		aim: position.aim,
	}
}

func (position SubmarinePosition) GoDown(amount int) SubmarinePosition {
	return SubmarinePosition{
		horizontal: position.horizontal,
		depth: position.depth,
		aim: position.aim + amount,
	}
}

func (position SubmarinePosition) GoUp(amount int) SubmarinePosition {
	return SubmarinePosition{
		horizontal: position.horizontal,
		depth: position.depth,
		aim: position.aim - amount,
	}
}

func (position SubmarinePosition) Value() int {
	return position.horizontal * position.depth
}
