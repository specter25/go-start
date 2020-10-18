package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("interfaces")
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello.go"))

	// eg2

	myint := IntCounter(0)
	var inc Incrementer = &myint
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	fmt.Println("multiple interafces problem")
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Youtube listeners"))
	wc.Close()

	//convert Writer closure to a buffered writer closure
	//this si possibel because the struct implements this partivular interface
	bwc, ok := wc.(*BufferedWriterCloser)
	if ok {
		fmt.Println(bwc)

	} else {
		fmt.Println("conversion failed")
	}

	// empty interfaces are interfaces with mo methods
	// can be type casted to any other struct interface or even primitive
	var myObj interface{} = NewBufferedWriterCloser()
	if wcc, ok := myObj.(WriterCloser); ok {
		wcc.Write([]byte("Hello Youtube listeners"))
		wcc.Close()
	}

	//type swutches
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("not int ")
	}

	// now when we define a varibale of the interface type as initialize it with a value struct not a  pointerstruct then all the implemented methods should be value methods
	// var wc WriterClosure = BufferedWriterClosure{}
	//but if we initilaize it via a pointer to the struct then the implemented methods can be bith vvalue and refernce methods
	// var wc WriterClosure = &BufferedWriterClosure{}

}

//some naming conventions if a interface with singlr method and the method name is eg write then the interface name should end with the same name +"er" i.e. Writer

type Writer interface {
	// interface don't describe data they describe behaviour
	// thus we won't define variables we will inturn define function definitions
	Write([]byte) (int, error)
}

//now implement the interface
type ConsoleWriter struct {
	// we will implimently implement the interface
}

func (sw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

//depicts that we can implement an interface by any type not necessaryily a struct
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// interface containing interfaces
type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil

}
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
