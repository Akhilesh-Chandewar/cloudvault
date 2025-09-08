package main

import (
	"log"
	"sync"

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
	Options  FileServerOptions
	peerLock sync.Mutex
	peers    map[string]p2p.Peer
	Storage  *Storage
	quitch   chan struct{}
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
		peers:   make(map[string]p2p.Peer),
	}
}

func (fs *FileServer) Stop() {
	close(fs.quitch)
}

func (fs *FileServer) OnPeer(p p2p.Peer) error {
	fs.peerLock.Lock()
	defer fs.peerLock.Unlock()
	fs.peers[p.RemoteAddr().String()] = p
	log.Println("New peer connected: ", p.RemoteAddr().String())
	return nil
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
