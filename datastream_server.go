package goquic

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"os"
	"strings"
)

// implement IncomingDataStreamCreator for Server
type SpdyServerSession struct {
	server        *QuicSpdyServer
	sessionFnChan chan func()
}

func (s *SpdyServerSession) CreateIncomingDynamicStream(quicServerStream *QuicServerStream, streamId uint32) DataStreamProcessor {
	stream := &SimpleServerStream{
		streamId:         streamId,
		server:           s.server,
		buffer:           new(bytes.Buffer),
		sessionFnChan:    s.sessionFnChan,
		quicServerStream: quicServerStream,
	}
	return stream
}

// implement DataStreamProcessor for Server
type SimpleServerStream struct {
	closed           bool
	streamId         uint32 // Just for logging purpose
	header           http.Header
	peerAddress      string
	buffer           *bytes.Buffer
	server           *QuicSpdyServer
	quicServerStream *QuicServerStream
	sessionFnChan    chan func()
	closeNotifyChan  chan bool
}

func (stream *SimpleServerStream) OnInitialHeadersComplete(header http.Header, peerAddress string) {
	stream.header = header
	stream.peerAddress = peerAddress
}

func (stream *SimpleServerStream) OnTrailingHeadersComplete(header http.Header) {
	// Should not be called
}

func (stream *SimpleServerStream) OnDataAvailable(data []byte, isClosed bool) {
	stream.buffer.Write(data)
	if isClosed {
		stream.ProcessRequest()
	}
}

func (stream *SimpleServerStream) OnClose() {
	if stream.closeNotifyChan != nil && !stream.closed {
		stream.closeNotifyChan <- true
	}
	stream.closed = true
}

func (stream *SimpleServerStream) ProcessRequest() {
	header := stream.header
	req := new(http.Request)
	req.Method = header.Get(":method")
	req.RequestURI = header.Get(":path")
	req.Proto = header.Get(":version")
	req.Header = header
	req.Host = header.Get(":host")
	req.RemoteAddr = stream.peerAddress
	rawPath := header.Get(":path")
	
	//################## BEGIN: WRITING PATH TO A FILE ##################\\
	f, err := os.OpenFile("metricas.txt", os.O_APPEND|os.O_WRONLY,0600)
	defer f.Close()
	if err != nil {
		fmt.Println(" Error opening File! ", err)
		return
	}
	if (strings.Contains(rawPath, ".m4s")) {
		
		//FORMATO: 320x240
		if (strings.Contains(rawPath, "45226")) {
			n2, err := f.WriteString(";;;;;;;;;;;;;;")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n2)
			n3, err := f.WriteString("320x240")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n3)
			n4, err := f.WriteString("\n")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n4)
			f.Sync()
		}
		
		//FORMATO: 480x360
		if (strings.Contains(rawPath, "177437")) {
			n2, err := f.WriteString(";;;;;;;;;;;;;;")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n2)
			n3, err := f.WriteString("480x360")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n3)
			n4, err := f.WriteString("\n")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n4)
		}
		
		//FORMATO: 854x480
		if (strings.Contains(rawPath, "509091")) {
			n2, err := f.WriteString(";;;;;;;;;;;;;;")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n2)
			n3, err := f.WriteString("854x480")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n3)
			n4, err := f.WriteString("\n")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n4)
		}
		
		//FORMATO: 1280x720
		if (strings.Contains(rawPath, "782553")) {
			n2, err := f.WriteString(";;;;;;;;;;;;;;")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n2)
			n3, err := f.WriteString("1280x720")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n3)
			n4, err := f.WriteString("\n")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n4)
		}
		
		//FORMATO: 1920x1080
		if (strings.Contains(rawPath, "2087347")) {
			n2, err := f.WriteString(";;;;;;;;;;;;;;")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n2)
			n3, err := f.WriteString("1920x1080")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n3)
			n4, err := f.WriteString("\n")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n4)
		}
		
		//Formato nao tratado
		if !((strings.Contains(rawPath, "45226")) || (strings.Contains(rawPath, "177437")) || (strings.Contains(rawPath, "509091")) || (strings.Contains(rawPath, "782553")) || (strings.Contains(rawPath, "2087347"))) {
			n2, err := f.WriteString(";;;;;;;;;;;;;;")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n2)
			n3, err := f.WriteString("Formato nao tratado em datastream_server.go")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n3)
			n4, err := f.WriteString("\n")
			if err != nil {
				fmt.Println(" Error writing File! ", err)
				return
			}
			fmt.Printf("wrote %d bytes\n", n4)
		}
	}
																																	 
	//################## END: WRITING PATH TO A FILE ##################\\

	url, err := url.ParseRequestURI(rawPath)
	if err != nil {
		fmt.Println(" Error! ", err)
		return
		// TODO(serialx): Send error message
	}

	url.Scheme = header.Get(":scheme")
	url.Host = header.Get(":host")
	req.URL = url
	// TODO(serialx): To buffered async read
	req.Body = ioutil.NopCloser(stream.buffer)

	// Remove SPDY headers
	for k, _ := range header {
		if len(k) > 0 && k[0] == ':' {
			header.Del(k)
		}
	}

	go func() {
		w := &spdyResponseWriter{
			serverStream: stream.quicServerStream,
			spdyStream:   stream,
			header:       make(http.Header),
		}
		sw := &spdyResponseBufferedWriter{
			serverStream: stream.quicServerStream,
			spdyStream:   stream,
		}
		w.w = bufio.NewWriter(sw)

		if stream.server.Handler != nil {
			stream.server.Handler.ServeHTTP(w, req)
		} else {
			http.DefaultServeMux.ServeHTTP(w, req)
		}

		err := w.w.Flush()
		if err != nil {
			fmt.Println("?????", err)
			return // XXX
		}

		stream.sessionFnChan <- func() {
			if stream.closed {
				return
			}

			trailers := make(http.Header)
			for key, v := range w.header {
				if _, ok := w.headerSent[key]; !ok {
					trailers[key] = v
				}
			}
			// XXX: How about header appending or deleting? is it available?
			if len(trailers) > 0 {
				stream.quicServerStream.WriteTrailers(trailers)
			} else {
				stream.quicServerStream.WriteOrBufferData(nil, true)
			}
		}
	}()
}

func (stream *SimpleServerStream) closeNotify() <-chan bool {
	if stream.closeNotifyChan == nil {
		stream.closeNotifyChan = make(chan bool, 1)
	}
	return stream.closeNotifyChan
}

type spdyResponseWriter struct {
	serverStream *QuicServerStream
	spdyStream   *SimpleServerStream
	header       http.Header
	headerSent   http.Header
	wroteHeader  bool
	w            *bufio.Writer
}

func (w *spdyResponseWriter) Header() http.Header {
	return w.header
}

func (w *spdyResponseWriter) Write(buffer []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}

	return w.w.Write(buffer)
}

func cloneHeader(h http.Header) http.Header {
	h2 := make(http.Header, len(h))
	for k, vv := range h {
		vv2 := make([]string, len(vv))
		copy(vv2, vv)
		h2[k] = vv2
	}
	return h2
}

func (w *spdyResponseWriter) WriteHeader(statusCode int) {
	if w.wroteHeader {
		return
	}
	copiedHeader := cloneHeader(w.header)
	w.spdyStream.sessionFnChan <- func() {
		copiedHeader.Set(":status", strconv.Itoa(statusCode))
		copiedHeader.Set(":version", "HTTP/1.1")
		if w.spdyStream.closed {
			return
		}
		w.serverStream.WriteHeader(copiedHeader, false)
	}
	w.wroteHeader = true
	w.headerSent = copiedHeader
}

func (w *spdyResponseWriter) CloseNotify() <-chan bool {
	return w.spdyStream.closeNotify()
}

func (w *spdyResponseWriter) Flush() {
	w.w.Flush()
}

type spdyResponseBufferedWriter struct {
	serverStream *QuicServerStream
	spdyStream   *SimpleServerStream
}

func (sw *spdyResponseBufferedWriter) Write(buffer []byte) (int, error) {
	copiedBuf := make([]byte, len(buffer))
	copy(copiedBuf, buffer)

	sw.spdyStream.sessionFnChan <- func() {
		if sw.spdyStream.closed {
			return
		}
		sw.serverStream.WriteOrBufferData(copiedBuf, false)
	}
	return len(buffer), nil
}
