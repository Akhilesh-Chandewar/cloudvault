package main

import (
	"log"

	"github.com/Akhilesh-Chandewar/cloudvault/p2p"
)

type FileServerOptions struct {
	StorageRoot    string
	PathTransform  PathTransform
	Transport      p2p.Transport
	TransportOpts  p2p.TCPTransportOpts
	BootstrapAddrs []string
}

type FileServer struct {
	Options FileServerOptions
	Storage *Storage
	quitch  chan struct{}
}

func NewFileServer(opts FileServerOptions) *FileServer {
	storageOptions := StorageOptions{
		Root:          opts.StorageRoot,
		PathTransform: opts.PathTransform,
	}

	return &FileServer{
		Options: opts,
		Storage: NewStorage(storageOptions),
		quitch:  make(chan struct{}),
	}
}

func (fs *FileServer) Stop() {
	close(fs.quitch)
}

func (fs *FileServer) loop() {
	defer func() {
		log.Println("FileServer stopped due to user quit action")
		fs.Options.Transport.Close()
	}()

	for {
		select {
		case rpc := <-fs.Options.Transport.Consume():
			log.Printf("Received RPC: %+v\n", rpc)

		case <-fs.quitch:
			log.Println("FileServer shutting down...")
			return
		}
	}
}
func (fs *FileServer) BootstrapNetwork() error {
	for _, addr := range fs.Options.BootstrapAddrs {
		if len(addr) == 0 {
			continue
		}

		go func(addr string) {
			log.Printf("Bootstrapping with %s\n", addr)
			if err := fs.Options.Transport.Dial(addr); err != nil {
				log.Println("dial error: ", err)
			}
		}(addr)
	}

	return nil
}

func (fs *FileServer) Start() error {
	if err := fs.Options.Transport.ListenAndAccept(); err != nil {
		return err
	}
	if err := fs.BootstrapNetwork(); err != nil {
		return err
	}
	fs.loop()
	return nil
}
