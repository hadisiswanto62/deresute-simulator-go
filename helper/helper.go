package helper

// Scale scales from min to max (max-min is divided into `steps` parts evenly) (in uint16)
func Scale(min, max uint16, steps, currentStep int8) uint16 {
	x := min + (max-min)/(uint16(steps)-1)*uint16(currentStep)
	return x
}
