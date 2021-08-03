package iobox

type Signal struct {
	Message string
	Value   uint32
}

type IOBox struct {
	Outputs []chan<- Signal
	Inputs  []<-chan Signal
}

func New() *IOBox {
	return &IOBox{}
}

func (iob *IOBox) ConnectOutputTo(inputIOBox *IOBox) {
	ch := make(chan Signal)
	iob.Outputs = append(iob.Outputs, ch)
	inputIOBox.Inputs = append(inputIOBox.Inputs, ch)
}

func (iob *IOBox) Send(signal Signal) {
	for index := range iob.Outputs {
		iob.Outputs[index] <- signal
	}
}

func (iob *IOBox) Receive(callback func(signal Signal)) {
	for index := range iob.Inputs {
		go func(index int) {
			for {
				signal := <-iob.Inputs[index]
				callback(signal)
			}
		}(index)
	}
}
