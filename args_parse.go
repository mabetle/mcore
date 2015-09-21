package mcore

type Args []string

func NewArgsFromString(args string)(Args){
	v := String(args).Split(" ")
	return Args(v)
}

func NewArgs(args []string)Args{
	return Args(args)
}

func (a Args)IsHasFlag(flag string)bool{
	return String(flag).IsInArray(a)
}

func (a Args)ParseString(flag string)(r string){

	return
}

func (a Args)ParseInt(flag string)(r int){

	return
}

func (a Args)NArgs()(int){
	return len(a)
}

func (a Args)NFlags()(r int){

	return
}

func (a Args)VArgs(index int)string{
	return a[index]
}

