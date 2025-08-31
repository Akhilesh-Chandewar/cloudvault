package main

import "github.com/Akhilesh-Chandewar/cloudvault/p2p"

type FileServerOptions struct {
	StorageRoot   string
	PathTransform PathTransform
	Transport     p2p.Transport
	TransportOpts p2p.TCPTransportOpts
}

type FileServer struct {
	Options FileServerOptions
	Storage *Storage
}

func NewFileServer(opts FileServerOptions) *FileServer {
	storageOptions := StorageOptions{
		Root:          opts.StorageRoot,
		PathTransform: opts.PathTransform,
	}

	return &FileServer{
		Options: opts,
		Storage: NewStorage(storageOptions),
	}
}

func (fs *FileServer) Start() error {
	if err := fs.Options.Transport.ListenAndAccept(); err != nil {
		return err
	}
	return nil
}
